<template>
  <widget-settings
    :submitting="submitting"
    :dirty="hasChanges"
    divider
    @submit="submit"
  >
    <statistics-form
      v-model="form"
      :widget="widget"
      :type="$constants.KPI_RATING_SETTINGS_TYPES.entity"
      :show-filter="hasAccessToFilter"
      :show-interval="hasAccessToInterval"
    />
  </widget-settings>
</template>

<script>
import { SIDE_BARS } from '@/constants';

import { widgetSettingsMixin } from '@/mixins/widget/settings';
import { permissionsWidgetsAlarmStatisticsInterval } from '@/mixins/permissions/widgets/statistics/alarm/interval';
import { permissionsWidgetsAlarmStatisticsFilters } from '@/mixins/permissions/widgets/statistics/alarm/filters';

import WidgetSettings from '../partials/widget-settings.vue';

import StatisticsForm from './form/statistics-form.vue';

export default {
  name: SIDE_BARS.alarmStatisticsSettings,
  components: {
    WidgetSettings,
    StatisticsForm,
  },
  mixins: [
    widgetSettingsMixin,
    permissionsWidgetsAlarmStatisticsInterval,
    permissionsWidgetsAlarmStatisticsFilters,
  ],
};
</script>
