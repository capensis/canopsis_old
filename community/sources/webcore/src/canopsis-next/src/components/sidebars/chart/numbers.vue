<template>
  <widget-settings
    :submitting="submitting"
    :dirty="hasChanges"
    divider
    @submit="submit"
  >
    <numbers-widget-form
      v-model="form"
      :with-filters="hasAccessToFilter"
      with-periodic-refresh
      with-preset
    />
  </widget-settings>
</template>

<script>
import { SIDE_BARS } from '@/constants';

import { widgetSettingsMixin } from '@/mixins/widget/settings';
import { permissionsWidgetsNumbersFilters } from '@/mixins/permissions/widgets/chart/numbers/filters';

import WidgetSettings from '../partials/widget-settings.vue';

import NumbersWidgetForm from './form/numbers-widget-form.vue';

export default {
  name: SIDE_BARS.numbersSettings,
  components: {
    WidgetSettings,
    NumbersWidgetForm,
  },
  mixins: [
    widgetSettingsMixin,
    permissionsWidgetsNumbersFilters,
  ],
};
</script>
