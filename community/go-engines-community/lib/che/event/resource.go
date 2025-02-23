package event

import (
	"context"
	"errors"
	"fmt"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/contextgraph"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/eventfilter"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/techmetrics"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	libmongo "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
)

type resourceProcessor struct {
	dbClient            libmongo.DbClient
	dbEntityCollection  libmongo.DbCollection
	dbAlarmCollection   libmongo.DbCollection
	contextGraphManager contextgraph.Manager
	eventFilterService  eventfilter.Service
	logger              zerolog.Logger
}

func NewResourceProcessor(
	dbClient libmongo.DbClient,
	contextGraphManager contextgraph.Manager,
	eventFilterService eventfilter.Service,
	logger zerolog.Logger,
) Processor {
	return &resourceProcessor{
		dbClient:            dbClient,
		dbEntityCollection:  dbClient.Collection(libmongo.EntityMongoCollection),
		dbAlarmCollection:   dbClient.Collection(libmongo.AlarmMongoCollection),
		contextGraphManager: contextGraphManager,
		eventFilterService:  eventFilterService,
		logger:              logger,
	}
}

func (p *resourceProcessor) Process(ctx context.Context, event *types.Event) (
	[]types.Entity,
	[]string,
	techmetrics.CheEventMetric,
	error,
) {
	eventMetric := techmetrics.CheEventMetric{
		EventMetric: techmetrics.EventMetric{
			EventType: event.EventType,
		},
	}

	var report contextgraph.Report
	commRegister := libmongo.NewCommandsRegister(p.dbEntityCollection, canopsis.DefaultBulkSize)

	err := p.dbClient.WithTransaction(ctx, func(ctx context.Context) error {
		commRegister.Clear()

		var err error
		report, err = p.contextGraphManager.HandleResource(ctx, event, commRegister)
		if err != nil {
			return fmt.Errorf("cannot handle resource: %w", err)
		}

		return commRegister.Commit(ctx)
	})
	if err != nil {
		return nil, nil, eventMetric, err
	}

	if event.Entity == nil {
		return nil, nil, eventMetric, errors.New("unexpected empty resource")
	}

	eventMetric.EntityType = event.Entity.Type
	eventMetric.IsNewEntity = report.IsNew

	if event.Healthcheck {
		return nil, nil, eventMetric, nil
	}

	// Process event by event filters.
	if event.Entity.Enabled {
		var isInfosUpdated bool
		isInfosUpdated, eventMetric.ExecutedEnrichRules, eventMetric.ExternalRequests, err = p.eventFilterService.ProcessEvent(ctx, event)
		if err != nil {
			return nil, nil, eventMetric, err
		}

		if isInfosUpdated {
			_, err = p.dbEntityCollection.UpdateOne(
				ctx,
				bson.M{"_id": event.Entity.ID},
				bson.M{"$set": bson.M{"infos": event.Entity.Infos}},
			)
			if err != nil {
				return nil, nil, eventMetric, fmt.Errorf("cannot update entities: %w", err)
			}

			eventMetric.IsInfosUpdated = true
			report.CheckResource = true
		}
	}

	// cap = 3 for a full context graph set: resource, component and connector.
	entityIdsToCheck := make([]string, 0, 3)
	entityIdsToUpdateMetrics := make([]string, 0, 3)

	if report.CheckResource {
		entityIdsToCheck = append(entityIdsToCheck, event.Entity.ID)
		entityIdsToUpdateMetrics = append(entityIdsToUpdateMetrics, event.Entity.ID)

		// always add component when we need to check resource, in order
		// to set component infos and check state settings.
		entityIdsToCheck = append(entityIdsToCheck, event.Entity.Component)

		// add component id to metrics update only if it should be checked.
		if report.CheckComponent {
			entityIdsToUpdateMetrics = append(entityIdsToUpdateMetrics, event.Entity.Component)
		}
	}

	if report.CheckConnector {
		entityIdsToCheck = append(entityIdsToCheck, event.Entity.Connector)
		entityIdsToUpdateMetrics = append(entityIdsToUpdateMetrics, event.Entity.Connector)
	}

	// if nothing is changed - leave.
	if len(entityIdsToCheck) == 0 {
		return nil, nil, eventMetric, nil
	}

	// cap = 2 for a potential component and connector counter updates.
	toCountersUpdate := make([]types.Entity, 0, 2)

	err = p.dbClient.WithTransaction(ctx, func(ctx context.Context) error {
		commRegister.Clear()
		toCountersUpdate = toCountersUpdate[:0]

		eventMetric.IsServicesUpdated = false
		eventMetric.IsStateSettingUpdated = false

		var resource types.Entity
		var component types.Entity
		var connector types.Entity

		cursor, err := p.dbEntityCollection.Find(ctx, bson.M{"_id": bson.M{"$in": entityIdsToCheck}})
		if err != nil {
			return err
		}

		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var ent types.Entity

			err = cursor.Decode(&ent)
			if err != nil {
				return err
			}

			if ent.Type == types.EntityTypeResource {
				resource = ent
			} else if ent.Type == types.EntityTypeComponent {
				component = ent
			} else {
				connector = ent
			}
		}

		if resource.ID == "" {
			return errors.New("resource was deleted during event processing")
		}

		// todo: decide if needed
		//if component.ID == "" {
		//	return fmt.Errorf("component was deleted during event processing")
		//}

		// todo: should be called to get fresh services from the db, should be removed when we do something with cache
		err = p.contextGraphManager.LoadServices(ctx)
		if err != nil {
			return fmt.Errorf("cannot refresh services: %w", err)
		}

		p.contextGraphManager.AssignServices(&resource, commRegister)
		eventMetric.IsServicesUpdated = len(event.Entity.ServicesToAdd) > 0 || len(event.Entity.ServicesToRemove) > 0

		if component.ID != "" && report.CheckComponent {
			p.contextGraphManager.AssignServices(&component, commRegister)
			if len(component.ServicesToAdd) > 0 || len(component.ServicesToRemove) > 0 {
				toCountersUpdate = append(toCountersUpdate, component)
			}

			eventMetric.IsStateSettingUpdated, err = p.contextGraphManager.AssignStateSetting(ctx, &component, commRegister)
			if err != nil {
				return fmt.Errorf("cannot assign state settings for a component: %w", err)
			}
		}

		if connector.ID != "" && report.CheckConnector {
			p.contextGraphManager.AssignServices(&connector, commRegister)
			if len(connector.ServicesToAdd) > 0 || len(connector.ServicesToRemove) > 0 {
				toCountersUpdate = append(toCountersUpdate, connector)
			}
		}

		err = p.contextGraphManager.InheritComponentFields(&resource, &component, commRegister)
		if err != nil {
			return fmt.Errorf("cannot inherit component fields: %w", err)
		}

		err = commRegister.Commit(ctx)
		if err != nil {
			return err
		}

		event.Entity = &resource

		return nil
	})
	if err != nil {
		return nil, nil, eventMetric, err
	}

	if eventMetric.IsInfosUpdated {
		go func() {
			err = updateMetaAlarmInfos(ctx, event.Entity.ID, event.Entity.Infos, p.dbAlarmCollection, p.dbEntityCollection)
			if err != nil {
				p.logger.Err(err).Str("entity", event.Entity.ID).Msg("cannot update meta alarm infos")
			}
		}()
	}

	return toCountersUpdate, entityIdsToUpdateMetrics, eventMetric, nil
}
