import request from '@/services/request';
import { API_ROUTES } from '@/config';
import { ENTITIES_TYPES } from '@/constants';
import { watcherOtherSchema } from '@/store/schemas';
import i18n from '@/i18n';
import watchedEntityModule from './watched-entity';

export const types = {
  FETCH_LIST: 'FETCH_LIST',
  FETCH_LIST_COMPLETED: 'FETCH_LIST_COMPLETED',
  FETCH_LIST_FAILED: 'FETCH_LIST_FAILED',
};

export default {
  namespaced: true,
  state: {
    allIds: [],
    meta: {},
    pending: false,
    fetchingParams: {},
    allIdsGeneralList: [],
    pendingGeneralList: false,
  },
  modules: {
    watchedEntity: watchedEntityModule,
  },
  getters: {
    allIds: state => state.allIds,
    items: (state, getters, rootState, rootGetters) => rootGetters['entities/getList'](ENTITIES_TYPES.otherWatcher, state.allIds),
    pending: state => state.pending,
    meta: state => state.meta,
    item: (state, getters, rootState, rootGetters) => id => rootGetters['entities/getItem'](ENTITIES_TYPES.otherWatcher, id),
  },
  mutations: {
    [types.FETCH_LIST](state, { params }) {
      state.pending = true;
      state.fetchingParams = params;
    },
    [types.FETCH_LIST_COMPLETED](state, { allIds, meta }) {
      state.allIds = allIds;
      state.meta = meta;
      state.pending = false;
    },
    [types.FETCH_LIST_FAILED](state) {
      state.pending = false;
    },
  },
  actions: {
    async create(context, params = {}) {
      try {
        await request.post(API_ROUTES.watcher, params);
      } catch (err) {
        console.warn(err);
      }
    },

    async remove({ dispatch }, { id } = {}) {
      try {
        await request.delete(API_ROUTES.watcher, { params: { watcher_id: id } });

        await dispatch('entities/removeFromStore', {
          id,
          type: ENTITIES_TYPES.watcher,
        }, { root: true });
      } catch (err) {
        console.warn(err);
      }
    },

    async fetchList({ dispatch, commit }, { params, filter } = { params: {}, filter: {} }) {
      try {
        commit(types.FETCH_LIST, { params });

        const { normalizedData } = await dispatch('entities/fetch', {
          route: `${API_ROUTES.watchers}/${JSON.stringify(filter)}`,
          schema: [watcherOtherSchema],
          params,
          dataPreparer: d => d,
        }, { root: true });

        commit(types.FETCH_LIST_COMPLETED, { allIds: normalizedData.result, meta: {} });
      } catch (e) {
        commit(types.FETCH_LIST_FAILED);
        console.error(e);
        await dispatch('popup/add', { type: 'error', text: i18n.t('errors.default') }, { root: true });
      }
    },
  },
};
