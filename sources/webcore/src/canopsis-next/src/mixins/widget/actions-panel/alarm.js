import { get, omit } from 'lodash';

import {
  MODALS,
  ENTITIES_TYPES,
  EVENT_ENTITY_TYPES,
  BUSINESS_USER_RIGHTS_ACTIONS_MAP,
  CRUD_ACTIONS,
  WIDGET_TYPES,
} from '@/constants';

import eventActionsAlarmMixin from '@/mixins/event-actions/alarm';
import entitiesPbehaviorMixin from '@/mixins/entities/pbehavior';

import { convertObjectToTreeview } from '@/helpers/treeview';

import { generateWidgetByType } from '@/helpers/entities';

/**
 * @mixin Mixin for the alarms list actions panel, show modal of the action
 */
export default {
  mixins: [eventActionsAlarmMixin, entitiesPbehaviorMixin],
  methods: {
    createFastAckEvent() {
      let eventData = {};

      if (this.widget.parameters.fastAckOutput && this.widget.parameters.fastAckOutput.enabled) {
        eventData = { output: this.widget.parameters.fastAckOutput.value };
      }

      return this.createEvent(EVENT_ENTITY_TYPES.ack, this.item, eventData);
    },

    async createMassFastAckEvent() {
      let eventData = {};

      if (this.widget.parameters.fastAckOutput && this.widget.parameters.fastAckOutput.enabled) {
        eventData = { output: this.widget.parameters.fastAckOutput.value };
      }
      const ackEventData = this.prepareData(EVENT_ENTITY_TYPES.ack, this.items, eventData);

      await this.createEventAction({
        data: ackEventData,
      });
    },

    showActionModal(name) {
      return () => this.$modals.show({
        name,
        config: this.modalConfig,
      });
    },

    showAckModal() {
      this.$modals.show({
        name: MODALS.createAckEvent,
        config: {
          ...this.modalConfig,
          isNoteRequired: this.widget.parameters.isAckNoteRequired,
        },
      });
    },

    showPbehaviorsListModal() {
      this.$modals.show({
        name: MODALS.pbehaviorList,
        config: {
          ...this.modalConfig,
          pbehaviors: this.item.pbehaviors,
          entityId: this.item.entity._id,
          availableActions: [CRUD_ACTIONS.delete, CRUD_ACTIONS.update],
        },
      });
    },

    showMoreInfosModal() {
      this.$modals.show({
        name: MODALS.moreInfos,
        config: {
          ...this.modalConfig,
          template: this.widget.parameters.moreInfoTemplate,
        },
      });
    },

    showAckRemoveModal() {
      this.$modals.show({
        name: MODALS.createCancelEvent,
        config: {
          ...this.modalConfig,
          title: this.$t('modals.createAckRemove.title'),
          eventType: EVENT_ENTITY_TYPES.ackRemove,
        },
      });
    },

    showVariablesHelperModal() {
      const variables = [];

      const alarmFields = convertObjectToTreeview(omit(this.item, ['entity']), 'alarm');

      variables.push(alarmFields);

      if (this.item.entity) {
        const entityFields = convertObjectToTreeview(this.item.entity, 'entity');
        variables.push(entityFields);
      }

      this.$modals.show({
        name: MODALS.variablesHelp,
        config: {
          ...this.modalConfig,

          variables,
        },
      });
    },

    showAddPbehaviorModal() {
      this.$modals.show({
        name: MODALS.createPbehavior,
        config: {
          pbehavior: {
            filter: {
              _id: { $in: [this.item.d] },
            },
          },
          action: data => this.createPbehavior({
            data,
            parents: [this.item],
            parentsType: ENTITIES_TYPES.alarm,
          }),
        },
      });
    },

    showHistoryModal() {
      const widget = generateWidgetByType(WIDGET_TYPES.alarmList);
      const filter = { $and: [{ 'entity._id': get(this.item, 'entity._id') }] };
      const entityFilter = {
        title: this.item.entity.name,
        filter,
      };

      widget.parameters.mainFilter = entityFilter;
      widget.parameters.viewFilters = [entityFilter];

      this.$modals.show({
        name: MODALS.alarmsList,
        config: {
          widget,
        },
      });
    },

    actionsAccessFilterHandler({ type }) {
      const right = BUSINESS_USER_RIGHTS_ACTIONS_MAP.alarmsList[type];

      if (!right) {
        return true;
      }

      return this.checkAccess(right);
    },
  },
};
