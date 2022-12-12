import { createNamespacedHelpers } from 'vuex';

import { ALARM_STEP_FIELDS, ENTITY_TEMPLATE_FIELDS, MAX_LIMIT, PBEHAVIOR_INFO_FIELDS } from '@/constants';

const { mapActions: mapServiceActions } = createNamespacedHelpers('service');

export const entityVariablesMixin = {
  data() {
    return {
      infos: [],
    };
  },
  mounted() {
    this.fetchInfos();
  },
  computed: {
    infosSubVariables() {
      return [
        {
          text: this.$t('common.value'),
          value: 'value',
        },
        {
          text: this.$t('common.description'),
          value: 'description',
        },
      ];
    },

    infosVariables() {
      return this.infos.map(({ value }) => ({
        text: value,
        value,
        variables: this.infosSubVariables,
      }));
    },

    alarmStepValueVariable() {
      return {
        text: this.$t('common.value'),
        value: ALARM_STEP_FIELDS.value,
      };
    },

    alarmStepTimestampVariable() {
      return {
        text: this.$t('common.timestamp'),
        value: ALARM_STEP_FIELDS.timestamp,
      };
    },

    alarmStepMessageVariable() {
      return {
        text: this.$t('common.message'),
        value: ALARM_STEP_FIELDS.message,
      };
    },

    alarmStepAuthorVariable() {
      return {
        text: this.$t('common.author'),
        value: ALARM_STEP_FIELDS.author,
      };
    },

    stateVariables() {
      return [this.alarmStepTimestampVariable, this.alarmStepValueVariable];
    },

    statusVariables() {
      return [this.alarmStepValueVariable];
    },

    ticketVariables() {
      return [this.alarmStepValueVariable];
    },

    snoozeVariables() {
      return [
        this.alarmStepTimestampVariable,
        this.alarmStepAuthorVariable,
        this.alarmStepMessageVariable,
      ];
    },

    ackVariables() {
      return [
        this.alarmStepTimestampVariable,
        this.alarmStepAuthorVariable,
        this.alarmStepMessageVariable,
      ];
    },

    pbehaviorInfoVariables() {
      return [
        {
          text: this.$t('pbehavior.pbehaviorType'),
          value: PBEHAVIOR_INFO_FIELDS.typeName,
        },
        {
          text: this.$tc('pbehavior.pbehaviorReason'),
          value: PBEHAVIOR_INFO_FIELDS.reason,
        },
        {
          text: this.$t('pbehavior.pbehaviorName'),
          value: PBEHAVIOR_INFO_FIELDS.name,
        },
      ];
    },

    variables() {
      return [
        {
          text: this.$t('common.id'),
          value: ENTITY_TEMPLATE_FIELDS.id,
        },
        {
          text: this.$t('common.name'),
          value: ENTITY_TEMPLATE_FIELDS.name,
        },
        {
          text: this.$t('common.infos'),
          value: ENTITY_TEMPLATE_FIELDS.infos,
          variables: this.infosVariables,
        },
        {
          text: this.$t('common.connector'),
          value: ENTITY_TEMPLATE_FIELDS.connector,
        },
        {
          text: this.$t('common.connectorName'),
          value: ENTITY_TEMPLATE_FIELDS.connectorName,
        },
        {
          text: this.$t('common.component'),
          value: ENTITY_TEMPLATE_FIELDS.component,
        },
        {
          text: this.$t('common.resource'),
          value: ENTITY_TEMPLATE_FIELDS.resource,
        },
        {
          text: this.$t('common.state'),
          value: ENTITY_TEMPLATE_FIELDS.state,
          variables: this.stateVariables,
        },
        {
          text: this.$t('common.status'),
          value: ENTITY_TEMPLATE_FIELDS.status,
          variables: this.statusVariables,
        },
        {
          text: this.$t('common.snooze'),
          value: ENTITY_TEMPLATE_FIELDS.snooze,
          variables: this.snoozeVariables,
        },
        {
          text: this.$t('common.ack'),
          value: ENTITY_TEMPLATE_FIELDS.ack,
          variables: this.ackVariables,
        },
        {
          text: this.$t('common.updated'),
          value: ENTITY_TEMPLATE_FIELDS.lastUpdateDate,
        },
        {
          text: this.$t('common.impactLevel'),
          value: ENTITY_TEMPLATE_FIELDS.impactLevel,
        },
        {
          text: this.$t('common.impactState'),
          value: ENTITY_TEMPLATE_FIELDS.impactState,
        },
        {
          text: this.$t('common.category'),
          value: ENTITY_TEMPLATE_FIELDS.categoryName,
        },
        {
          text: this.$t('pbehavior.pbehaviorInfo'),
          value: ENTITY_TEMPLATE_FIELDS.pbehaviorInfo,
          variables: this.pbehaviorInfoVariables,
        },
        {
          text: this.$t('alarm.alarmCreationDate'),
          value: ENTITY_TEMPLATE_FIELDS.alarmCreationDate,
        },
        {
          text: this.$t('common.ticket'),
          value: ENTITY_TEMPLATE_FIELDS.ticket,
          variables: this.ticketVariables,
        },
        {
          text: this.$t('entity.okEvents'),
          value: ENTITY_TEMPLATE_FIELDS.statsOk,
        },
        {
          text: this.$t('entity.koEvents'),
          value: ENTITY_TEMPLATE_FIELDS.statsKo,
        },
        {
          text: this.$t('alarm.alarmDisplayName'),
          value: ENTITY_TEMPLATE_FIELDS.alarmDisplayName,
        },
        {
          text: this.$tc('common.link', 2),
          value: ENTITY_TEMPLATE_FIELDS.links,
        },
      ];
    },
  },
  methods: {
    ...mapServiceActions({ fetchEntityInfosKeysWithoutStore: 'fetchInfosKeysWithoutStore' }),

    async fetchInfos() {
      const { data: infos } = await this.fetchEntityInfosKeysWithoutStore({
        params: { limit: MAX_LIMIT },
      });

      this.infos = infos;
    },
  },
};
