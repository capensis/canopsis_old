import { contextSchema } from '@/store/schemas';
import request from '@/services/request';
import { API_ROUTES } from '@/config';
import { ENTITIES_TYPES } from '@/constants';

export const types = {
  FETCH_LIST: 'FETCH_LIST',
  FETCH_LIST_COMPLETED: 'FETCH_LIST_COMPLETED',
  FETCH_LIST_FAILED: 'FETCH_LIST_FAILED',
  FETCH_GENERAL_LIST: 'FETCH_GENERAL_LIST',
  FETCH_GENERAL_LIST_COMPLETED: 'FETCH_GENERAL_LIST_COMPLETED',
  FETCH_GENERAL_LIST_FAILED: 'FETCH_GENERAL_LIST_FAILED',
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
  getters: {
    allIds: state => state.allIds,
    items: (state, getters, rootState, rootGetters) => rootGetters['entities/getList']('context', state.allIds),
    pending: state => state.pending,
    meta: state => state.meta,
    allIdsGeneralList: state => state.allIds,
    itemsGeneralList: (state, getters, rootState, rootGetters) => rootGetters['entities/getList']('context', state.allIdsGeneralList),
    pendingGeneralList: state => state.pendingGeneralList,
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
    [types.FETCH_GENERAL_LIST](state, { params }) {
      state.pendingGeneralList = true;
      state.fetchingParamsGeneralList = params;
    },
    [types.FETCH_GENERAL_LIST_COMPLETED](state, { allIds }) {
      state.allIdsGeneralList = allIds;
      state.pendingGeneralList = false;
    },
    [types.FETCH_GENERAL_LIST_FAILED](state) {
      state.pendingGeneralList = false;
    },
  },
  actions: {
    fetch({ dispatch }, { params } = {}) {
      return dispatch('entities/fetch', {
        route: API_ROUTES.context,
        schema: [contextSchema],
        params,
        dataPreparer: d => d.data,
        isPost: true,
      }, { root: true });
    },

    async fetchList({ commit, dispatch }, { params } = {}) {
      try {
        commit(types.FETCH_LIST, { params });

        const { normalizedData, data } = await dispatch('fetch', { params });

        commit(types.FETCH_LIST_COMPLETED, {
          allIds: normalizedData.result,
          meta: {
            total: data.total,
          },
        });
      } catch (err) {
        console.error(err);
        commit(types.FETCH_LIST_FAILED);
      }
    },

    async fetchGeneralList({ commit, dispatch }, { params } = {}) {
      try {
        commit(types.FETCH_GENERAL_LIST, { params });

        const { normalizedData } = await dispatch('fetch', { params });

        commit(types.FETCH_GENERAL_LIST_COMPLETED, {
          allIds: normalizedData.result,
        });
      } catch (err) {
        console.error(err);
        commit(types.FETCH_GENERAL_LIST_FAILED);
      }
    },

    async remove({ dispatch }, { id } = {}) {
      try {
        await request.delete(API_ROUTES.context, { params: { ids: id } });

        await dispatch('entities/removeFromStore', {
          id,
          type: ENTITIES_TYPES.context,
        }, { root: true });
      } catch (err) {
        console.warn(err);
      }
    },
    async update(context, { entity }) {
      try {
        await request.put(API_ROUTES.context, {
          _type: 'crudcontext',
          entity,
        });
      } catch (err) {
        console.error(err);
      }
    },
  },
};
