<template>
  <v-radio-group
    v-field="value"
    :name="name"
    :label="label"
    hide-details
    mandatory
  >
    <v-radio
      v-for="type in types"
      :key="type.value"
      :label="type.label"
      :value="type.value"
      color="primary"
    />
  </v-radio-group>
</template>

<script>
import { computed } from 'vue';

import { AVAILABILITY_DISPLAY_PARAMETERS } from '@/constants';

import { useI18n } from '@/hooks/i18n';

export default {
  props: {
    value: {
      type: Number,
      default: AVAILABILITY_DISPLAY_PARAMETERS.uptime,
    },
    name: {
      type: String,
      default: 'parameter',
    },
    label: {
      type: String,
      required: false,
    },
  },
  setup() {
    const { t, tc } = useI18n();

    const types = computed(() => [{
      value: AVAILABILITY_DISPLAY_PARAMETERS.uptime,
      label: tc('common.uptime'),
    }, {
      value: AVAILABILITY_DISPLAY_PARAMETERS.downtime,
      label: t('common.downtime'),
    }]);

    return {
      types,
    };
  },
};
</script>
