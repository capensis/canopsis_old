import { USER_PERMISSIONS } from '@/constants';

import { authMixin } from '@/mixins/auth';

export const permissionsWidgetsServiceWeatherFilters = {
  mixins: [authMixin],
  computed: {
    hasAccessToListFilters() {
      return this.checkAccess(USER_PERMISSIONS.business.serviceWeather.actions.listFilters);
    },

    hasAccessToEditFilter() {
      return this.checkAccess(USER_PERMISSIONS.business.serviceWeather.actions.editFilter);
    },

    hasAccessToAddFilter() {
      return this.checkAccess(USER_PERMISSIONS.business.serviceWeather.actions.addFilter);
    },

    hasAccessToUserFilter() {
      return this.checkAccess(USER_PERMISSIONS.business.serviceWeather.actions.userFilter);
    },
  },
};
