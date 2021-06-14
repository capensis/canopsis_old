// mongoadapter contains casbin mongo adapter.
// Adapter loads policy from mongo collection and transform result into casbin format.
// Refactor mongo collection and use casbin mongo adapter after all API is migrated to Go.
package mongoadapter

import (
	"context"
	"git.canopsis.net/canopsis/go-engines/lib/mongo"
	libmodel "git.canopsis.net/canopsis/go-engines/lib/security/model"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"go.mongodb.org/mongo-driver/bson"
)

// NewAdapter creates mongo adapter.
func NewAdapter(db mongo.DbClient) persist.Adapter {
	return &adapter{
		collection: db.Collection(mongo.RightsMongoCollection),
	}
}

// adapter implements casbin adapter interface.
type adapter struct {
	collection mongo.DbCollection
}

// available permissions
const (
	permissionCreate = "create"
	permissionRead   = "read"
	permissionUpdate = "update"
	permissionDelete = "delete"
	permissionCan    = "can"
)

// bitmasks of available permissions
const (
	permissionBitmaskCreate = 8 // 0b1000
	permissionBitmaskRead   = 4 // 0b0100
	permissionBitmaskUpdate = 2 // 0b0010
	permissionBitmaskDelete = 1 // 0b0001
	permissionBitmaskCan    = 1 // 0b0001
)

const (
	CasbinPtypePolicy = "p"
	CasbinPtypeRole   = "g"
)

type objectConfig struct {
	Name         string
	HasCRUDPerms bool
}

// LoadPolicy loads all policy rules from mongo collection.
func (a *adapter) LoadPolicy(model model.Model) (resErr error) {
	objConfByID, err := a.findObjects()

	if err != nil {
		return err
	}

	roleNamesByID, err := a.loadRoles(model, objConfByID)

	if err != nil {
		return err
	}

	err = a.loadSubjects(model, roleNamesByID)

	if err != nil {
		return err
	}

	return nil
}

// SavePolicy isn't implemented.
// Implement it if all API is migrated to Go
// and there is not time to refactor mongo collection.
func (adapter) SavePolicy(model.Model) error {
	panic("implement me")
}

func (adapter) AddPolicy(string, string, []string) error {
	panic("implement me")
}

func (adapter) RemovePolicy(string, string, []string) error {
	panic("implement me")
}

func (adapter) RemoveFilteredPolicy(string, string, int, ...string) error {
	panic("implement me")
}

// findObjects fetches objects from mongo collection and returns map[objectID]objectConfig.
func (a *adapter) findObjects() (_ map[string]objectConfig, resErr error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cursor, err := a.collection.Find(
		ctx,
		bson.M{
			"crecord_type": libmodel.LineTypeObject,
			"crecord_name": bson.M{"$exists": true, "$ne": ""},
		},
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil && resErr == nil {
			resErr = err
		}
	}()

	objConfByID := make(map[string]objectConfig)

	for cursor.Next(ctx) {
		var line libmodel.Rbac
		err = cursor.Decode(&line)

		if err != nil {
			return nil, err
		}

		objConfByID[line.ID] = objectConfig{
			Name:         line.Name,
			HasCRUDPerms: line.ObjectType == libmodel.LineObjectTypeCRUD,
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return objConfByID, nil
}

// loadRoles fetches roles from mongo collection and adds them to casbin policy.
// Method returns map[roleID]roleName.
func (a *adapter) loadRoles(
	model model.Model,
	objConfByID map[string]objectConfig,
) (_ map[string]string, resErr error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cursor, err := a.collection.Find(
		ctx,
		bson.M{
			"crecord_type": libmodel.LineTypeRole,
			"crecord_name": bson.M{"$exists": true, "$ne": ""},
		},
	)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil && resErr == nil {
			resErr = err
		}
	}()

	roleNamesByID := make(map[string]string)
	ptype := CasbinPtypePolicy
	sec := ptype
	permBitmasksByName := map[string]int{
		permissionCreate: permissionBitmaskCreate,
		permissionRead:   permissionBitmaskRead,
		permissionUpdate: permissionBitmaskUpdate,
		permissionDelete: permissionBitmaskDelete,
	}

	for cursor.Next(ctx) {
		var line libmodel.Rbac
		err = cursor.Decode(&line)

		if err != nil {
			return nil, err
		}

		roleName := line.Name

		for objId, permConfig := range line.PermConfigList {
			if objConf, ok := objConfByID[objId]; ok {
				if objConf.HasCRUDPerms {
					for permName, bitmask := range permBitmasksByName {
						if permConfig.Bitmask&bitmask == bitmask {
							model[sec][ptype].Policy = append(
								model[sec][ptype].Policy,
								[]string{roleName, objConf.Name, permName},
							)
						}
					}
				} else if permConfig.Bitmask&permissionBitmaskCan == permissionBitmaskCan {
					model[sec][ptype].Policy = append(
						model[sec][ptype].Policy,
						[]string{roleName, objConf.Name, permissionCan},
					)
				}
			}
		}

		roleNamesByID[line.ID] = roleName
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return roleNamesByID, nil
}

// loadSubjects loads subjects from mongo collection and adds them to casbin policy.
func (a *adapter) loadSubjects(
	model model.Model,
	roleNamesByID map[string]string,
) (resErr error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cursor, err := a.collection.Find(ctx, bson.M{
		"crecord_type": libmodel.LineTypeSubject,
		"role":         bson.M{"$exists": true, "$ne": ""},
	})

	if err != nil {
		return err
	}

	defer func() {
		if err := cursor.Close(ctx); err != nil && resErr == nil {
			resErr = err
		}
	}()

	ptype := CasbinPtypeRole
	sec := ptype

	for cursor.Next(ctx) {
		var line libmodel.Rbac
		err := cursor.Decode(&line)

		if err != nil {
			return err
		}

		subjectID := line.ID

		if roleName, ok := roleNamesByID[line.Role]; ok {
			model[sec][ptype].Policy = append(
				model[sec][ptype].Policy,
				[]string{subjectID, roleName},
			)
		}
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}
