import qs from 'qs';

import i18n from '@/i18n';
import { API_ROUTES } from '@/config';
import { ENTITIES_TYPES } from '@/constants';
import { userSchema } from '@/store/schemas';
import request from '@/services/request';

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
  },
  getters: {
    items: (state, getters, rootState, rootGetters) =>
      rootGetters['entities/getList'](ENTITIES_TYPES.user, state.allIds),

    meta: state => state.meta,
    pending: state => state.pending,
  },
  mutations: {
    [types.FETCH_LIST](state) {
      state.pending = true;
    },
    [types.FETCH_LIST_COMPLETED](state, { allIds, meta }) {
      state.pending = false;
      state.allIds = allIds;
      state.meta = meta;
    },
    [types.FETCH_LIST_FAILED](state) {
      state.pending = false;
    },
  },
  actions: {
    async create({ dispatch }, { data }) {
      try {
        await request.post(API_ROUTES.createUser, qs.stringify({ user: JSON.stringify(data) }), {
          headers: { 'content-type': 'application/x-www-form-urlencoded' },
        });
        await dispatch('popup/add', { type: 'success', text: i18n.t('success.default') }, { root: true });
      } catch (err) {
        await dispatch('popup/add', { type: 'error', text: i18n.t('errors.default') }, { root: true });
      }
    },

    async remove({ dispatch }, { id }) {
      try {
        await request.post(`${API_ROUTES.deleteUser}/${id}`);
        await dispatch('popup/add', { type: 'success', text: i18n.t('success.default') }, { root: true });
      } catch (err) {
        await dispatch('popup/add', { type: 'error', text: i18n.t('errors.default') }, { root: true });
      }
    },

    async fetchList({ commit, dispatch }, { params } = {}) {
      try {
        commit(types.FETCH_LIST);

        const { normalizedData, data } = await dispatch('entities/fetch', {
          route: API_ROUTES.user,
          schema: [userSchema],
          params,
          dataPreparer: d => d.data,
        }, { root: true });

        commit(types.FETCH_LIST_COMPLETED, {
          allIds: normalizedData.result,
          meta: {
            total: data.total,
          },
        });
      } catch (err) {
        commit(types.FETCH_LIST_FAILED);

        await dispatch('popup/add', { type: 'error', text: i18n.t('errors.default') }, { root: true });
      }
    },
  },
};
