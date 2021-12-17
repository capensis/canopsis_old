import {
  MODALS,
  WEATHER_ACK_EVENT_OUTPUT,
  BUSINESS_USER_PERMISSIONS_ACTIONS_MAP, WEATHER_ACTIONS_TYPES, PBEHAVIOR_TYPE_TYPES,
} from '@/constants';

import { authMixin } from '@/mixins/auth';
import { entitiesPbehaviorMixin } from '@/mixins/entities/pbehavior';
import { entitiesPbehaviorTypesMixin } from '@/mixins/entities/pbehavior/types';
import { isActionTypeAvailableForEntity } from '@/helpers/entities/context';

export const widgetActionPanelServiceEntityMixin = {
  mixins: [
    authMixin,
    entitiesPbehaviorTypesMixin,
    entitiesPbehaviorMixin,
  ],
  data() {
    return {
      unavailableEntitiesAction: {},
    };
  },
  computed: {
    /**
     * @return {Object.<string, Function>}
     */
    actionsMethodsMap() {
      return {
        [WEATHER_ACTIONS_TYPES.entityAck]: this.addAckActionToQueue,
        [WEATHER_ACTIONS_TYPES.entityAssocTicket]: this.showCreateDeclareTicketModal,
        [WEATHER_ACTIONS_TYPES.entityValidate]: this.addValidateActionToQueue,
        [WEATHER_ACTIONS_TYPES.entityInvalidate]: this.addInvalidateActionToQueue,
        [WEATHER_ACTIONS_TYPES.entityPause]: this.showCreateServicePauseEventModal,
        [WEATHER_ACTIONS_TYPES.entityPlay]: this.addPlayActionToQueue,
        [WEATHER_ACTIONS_TYPES.entityCancel]: this.showCancelModal,
        [WEATHER_ACTIONS_TYPES.entityComment]: this.showCreateCommentEventModal,
      };
    },
  },
  methods: {
    removeEntityFromUnavailable(entity) {
      this.unavailableEntitiesAction[entity._id] = false;
    },

    addActionToQueue(action) {
      const {
        availableEntities,
        unavailableEntities,
      } = action.entities.reduce((acc, entity) => {
        if (isActionTypeAvailableForEntity(action.actionType, entity)) {
          acc.availableEntities.push(entity);
        } else {
          acc.unavailableEntities.push(entity);
        }

        return acc;
      }, {
        availableEntities: [],
        unavailableEntities: [],
      });

      this.unavailableEntitiesAction = unavailableEntities.reduce((acc, { _id: id }) => {
        acc[id] = true;

        return acc;
      }, {});

      this.$emit('add:action', {
        ...action,
        entities: availableEntities,
      });
    },

    addEntityAction(actionType, entities) {
      const handler = this.actionsMethodsMap[actionType];

      if (handler) {
        handler(entities);
      }
    },

    /**
     * Filter for available entity actions
     *
     * @param {string} type
     * @return {boolean}
     */
    actionsAccessFilterHandler({ type }) {
      const permission = BUSINESS_USER_PERMISSIONS_ACTIONS_MAP.weather[type];

      return permission
        ? this.checkAccess(permission)
        : true;
    },

    addAckActionToQueue(entities) {
      this.addActionToQueue({
        entities,
        actionType: WEATHER_ACTIONS_TYPES.entityAck,
      });
    },

    showCreateDeclareTicketModal(entities) {
      this.$modals.show({
        name: MODALS.textFieldEditor,
        config: {
          title: this.$t('modals.createDeclareTicket.title'),
          field: {
            name: 'ticket',
            label: this.$t('modals.createAssociateTicket.fields.ticket'),
            validationRules: 'required',
          },
          action: (ticket) => {
            this.addActionToQueue({
              entities,
              actionType: WEATHER_ACTIONS_TYPES.entityAssocTicket,
              payload: { ticket },
            });
          },
        },
      });
    },

    showCreateCommentEventModal(entities) {
      this.$modals.show({
        name: MODALS.createCommentEvent,
        config: {
          action: ({ output }) => {
            this.addActionToQueue({
              entities,
              actionType: WEATHER_ACTIONS_TYPES.entityComment,
              payload: { output },
            });
          },
        },
      });
    },

    addValidateActionToQueue(entities) {
      this.addActionToQueue({
        actionType: WEATHER_ACTIONS_TYPES.entityValidate,
        items: entities,
      });
    },

    addInvalidateActionToQueue(entities) {
      this.addActionToQueue({
        entities,
        actionType: WEATHER_ACTIONS_TYPES.entityInvalidate,
        payload: {
          output: WEATHER_ACK_EVENT_OUTPUT.ack,
        },
      });
    },

    showCreateServicePauseEventModal(entities) {
      this.$modals.show({
        name: MODALS.createServicePauseEvent,
        config: {
          action: async (pause) => {
            const defaultPbehaviorTypes = await this.fetchDefaultPbehaviorTypes();

            const pauseType = defaultPbehaviorTypes.find(({ type }) => type === PBEHAVIOR_TYPE_TYPES.pause);

            this.addActionToQueue({
              entities,
              actionType: WEATHER_ACTIONS_TYPES.entityPause,
              payload: {
                comment: pause.comment,
                reason: pause.reason,
                type: pauseType,
              },
            });
          },
        },
      });
    },

    addPlayActionToQueue(entities) {
      this.addActionToQueue({
        entities,
        actionType: WEATHER_ACTIONS_TYPES.entityPlay,
      });
    },

    showCancelModal(entities) {
      this.$modals.show({
        name: MODALS.textFieldEditor,
        config: {
          title: this.$t('common.output'),
          action: (output) => {
            this.addActionToQueue({
              entities,
              actionType: WEATHER_ACTIONS_TYPES.entityCancel,
              payload: { output },
            });
          },
        },
      });
    },
  },
};
