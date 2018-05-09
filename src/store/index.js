import Vue from 'vue';
import Vuex from 'vuex';

import i18nModule from './modules/i18n';
import AlarmsListSettingsModule from './modules/alarms-list-settings';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    i18n: i18nModule,
    AlarmsListSettings: AlarmsListSettingsModule,
  },
});
