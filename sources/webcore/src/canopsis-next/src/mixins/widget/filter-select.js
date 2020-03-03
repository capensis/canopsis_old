import { isEmpty, isBoolean } from 'lodash';

import { FILTER_DEFAULT_VALUES } from '@/constants';

import prepareMainFilterToQueryFilter from '@/helpers/filter';

export default {
  computed: {
    mainFilterCondition() {
      return this.userPreference.widget_preferences.mainFilterCondition || this.widget.parameters.mainFilterCondition;
    },

    mainFilter() {
      const {
        mainFilter: userMainFilter,
        mainFilterUpdatedAt: userMainFilterUpdatedAt = 0,
      } = this.userPreference.widget_preferences;

      const {
        mainFilter: widgetMainFilter,
        mainFilterUpdatedAt: widgetMainFilterUpdatedAt = 0,
      } = this.widget.parameters;

      let mainFilter = userMainFilter;
      if (isEmpty(mainFilter) && !isEmpty(widgetMainFilter) && widgetMainFilterUpdatedAt >= userMainFilterUpdatedAt) {
        mainFilter = widgetMainFilter;
      }

      return mainFilter || null;
    },

    viewFilters() {
      return this.userPreference.widget_preferences.viewFilters || [];
    },

    widgetViewFilters() {
      const { mainFilter, viewFilters } = this.widget.parameters;

      if (!this.hasAccessToListFilter) {
        return mainFilter ? [mainFilter] : [];
      }

      return viewFilters || [];
    },
  },
  methods: {
    updateFieldsInWidgetPreferences(fields = {}) {
      const hasAccessToUserFilter = this.hasAccessToUserFilter || !isBoolean(this.hasAccessToUserFilter);

      if (hasAccessToUserFilter) {
        return this.updateWidgetPreferencesInUserPreference({
          ...this.userPreference.widget_preferences,
          ...fields,
        });
      }

      return Promise.resolve();
    },

    updateFilters(viewFilters, mainFilter = this.mainFilter) {
      this.updateFieldsInWidgetPreferences({ viewFilters, mainFilter });
      this.updateQueryBySelectedFilterAndCondition(mainFilter, this.mainFilterCondition);
    },

    updateSelectedCondition(condition = FILTER_DEFAULT_VALUES.condition) {
      this.updateFieldsInWidgetPreferences({ mainFilterCondition: condition });
      this.updateQueryBySelectedFilterAndCondition(this.mainFilter, condition);
    },

    updateSelectedFilter(filterObject) {
      this.updateFieldsInWidgetPreferences({ mainFilter: filterObject || {}, mainFilterUpdatedAt: Date.now() });
      this.updateQueryBySelectedFilterAndCondition(filterObject, this.mainFilterCondition);
    },

    updateQueryBySelectedFilterAndCondition(filter, condition) {
      this.query = {
        ...this.query,

        filter: prepareMainFilterToQueryFilter(filter, condition),
      };
    },
  },
};
