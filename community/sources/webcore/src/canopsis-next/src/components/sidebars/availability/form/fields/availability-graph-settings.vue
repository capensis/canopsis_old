<template>
  <widget-settings-item :title="$t('settings.availability.graphSettings')">
    <c-enabled-field v-field="form.enabled" />
    <v-expand-transition>
      <div v-if="form.enabled">
        <c-quick-date-interval-type-field
          v-field="form.default_time_range"
          :name="defaultTimeRangeFieldName"
          :ranges="intervalRanges"
          :label="$t('settings.defaultTimeRange')"
        />
        <availability-show-type-radio-field
          v-field="form.default_show_type"
          :label="$t('settings.availability.defaultAvailabilityDisplay')"
          :name="defaultShowTypeFieldName"
        />
      </div>
    </v-expand-transition>
  </widget-settings-item>
</template>

<script>
import { computed } from 'vue';

import { AVAILABILITY_QUICK_RANGES } from '@/constants';

import WidgetSettingsItem from '@/components/sidebars/partials/widget-settings-item.vue';
import AvailabilityShowTypeRadioField from '@/components/other/availability/form/fields/availability-show-type-radio-field.vue';

export default {
  components: { WidgetSettingsItem, AvailabilityShowTypeRadioField },
  model: {
    prop: 'form',
    event: 'input',
  },
  props: {
    form: {
      type: Object,
      required: true,
    },
    name: {
      type: String,
      default: 'availability',
    },
  },
  setup(props) {
    const defaultTimeRangeFieldName = computed(() => `${props.name}.default_time_range`);
    const defaultShowTypeFieldName = computed(() => `${props.name}.show_type`);
    const intervalRanges = computed(() => Object.values(AVAILABILITY_QUICK_RANGES));

    return {
      defaultTimeRangeFieldName,
      defaultShowTypeFieldName,
      intervalRanges,
    };
  },
};
</script>
