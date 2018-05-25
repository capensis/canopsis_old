import Vue from 'vue';
import Vuex from 'vuex';

import appModule from './modules/app';
import i18nModule from './modules/i18n';
import modalModule from './modules/modal';
import eventModule from './modules/event';
import mFilterEditorModule from './modules/mfilter-editor';
import entitiesModules from './modules/entities';

import entitiesPlugin from './plugins/entities';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    app: appModule,
    i18n: i18nModule,
    modal: modalModule,
    event: eventModule,
    mFilterEditor: mFilterEditorModule,

    ...entitiesModules,
  },
  plugins: [entitiesPlugin],
});
