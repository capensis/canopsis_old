<template>
  <v-combobox
    v-validate="rules"
    v-field="value"
    :items="eventTypes"
    :label="label ?? $t('common.eventType')"
    :return-object="false"
    :error-messages="errors.collect(name)"
    :name="name"
    v-bind="$attrs"
  >
    <template v-if="$scopedSlots.selection" #selection="props">
      <slot name="selection" v-bind="props" />
    </template>
  </v-combobox>
</template>

<script>
import { computed } from 'vue';

import { EVENT_TYPES } from '@/constants';

export default {
  inject: ['$validator'],
  inheritAttrs: false,
  props: {
    value: {
      type: [String, Array],
      required: true,
    },
    label: {
      type: String,
      required: false,
    },
    required: {
      type: Boolean,
      default: false,
    },
    name: {
      type: String,
      default: 'event_type',
    },
    types: {
      type: Array,
      default: () => [],
    },
  },
  setup(props) {
    const eventTypes = computed(() => (
      props.types?.length
        ? props.types
        : Object.values(EVENT_TYPES).map(value => ({
          value,
          text: value,
        }))));

    const rules = computed(() => ({
      required: props.required,
    }));

    return {
      rules,
      eventTypes,
    };
  },
};
</script>
