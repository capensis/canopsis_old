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
      :type="$constants.KPI_RATING_SETTINGS_TYPES.user"
      :show-filter="hasAccessToFilter"
      :show-interval="hasAccessToInterval"
    />
  </widget-settings>
</template>

<script>
import { SIDE_BARS } from '@/constants';

import { widgetSettingsMixin } from '@/mixins/widget/settings';
import { permissionsWidgetsUserStatisticsInterval } from '@/mixins/permissions/widgets/statistics/user/interval';
import { permissionsWidgetsUserStatisticsFilters } from '@/mixins/permissions/widgets/statistics/user/filters';

import WidgetSettings from '../partials/widget-settings.vue';

import StatisticsForm from './form/statistics-form.vue';

export default {
  name: SIDE_BARS.userStatisticsSettings,
  components: {
    WidgetSettings,
    StatisticsForm,
  },
  mixins: [
    widgetSettingsMixin,
    permissionsWidgetsUserStatisticsInterval,
    permissionsWidgetsUserStatisticsFilters,
  ],
};
</script>
