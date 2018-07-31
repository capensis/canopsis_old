import omit from 'lodash/omit';

import request from '@/services/request';
import { API_ROUTES } from '@/config';
import { ENTITIES_TYPES } from '@/constants';
import { userPreferenceSchema } from '@/store/schemas';

export const types = {
  FETCH_LIST: 'FETCH_LIST',
  FETCH_LIST_COMPLETED: 'FETCH_LIST_COMPLETED',
  FETCH_LIST_FAILED: 'FETCH_LIST_FAILED',
  SET_ACTIVE_FILTER: 'SET_ACTIVE_FILTER',
};

export default {
  namespaced: true,
  getters: {
    getItemByWidget: (state, getters, rootState, rootGetters) => (widget) => {
      const id = `${widget.id}_root`;
      const userPreference = rootGetters['entities/getItem'](ENTITIES_TYPES.userPreference, id);

      if (!userPreference) {
        return {
          id,
          _id: id,
          widget_preferences: {},
          crecord_name: 'root', // TODO: username
          widget_id: widget.id,
          widgetXtype: widget.xtype,
          crecord_type: 'userpreferences',
        };
      }

      return userPreference;
    },
  },
  mutations: {
    [types.FETCH_LIST]: state => state.pending = true,
    [types.FETCH_LIST_COMPLETED]: (state) => {
      state.pending = false;
    },
    [types.FETCH_LIST_FAILED]: (state) => {
      state.pending = false;
    },
  },
  actions: {
    /**
     * This action fetches user preferences list
     *
     * @param {function} commit
     * @param {function} dispatch
     * @param {Object} params
     */
    async fetchList({ commit, dispatch }, { params }) {
      try {
        commit(types.FETCH_LIST);

        await dispatch('entities/fetch', {
          route: API_ROUTES.userPreferences,
          schema: [userPreferenceSchema],
          params,
          dataPreparer: d => d.data,
        }, { root: true });

        commit(types.FETCH_LIST_COMPLETED);
      } catch (e) {
        commit(types.FETCH_LIST_FAILED);
        console.warn(e);
      }
    },

    /**
     * This action fetches user preference item by widget id
     *
     * @param {function} dispatch
     * @param {string|number} widgetId
     */
    async fetchItemByWidgetId({ dispatch }, { widgetId }) {
      await dispatch('fetchList', {
        params: {
          limit: 1,
          filter: {
            crecord_name: 'root',
            widget_id: widgetId,
            _id: `${widgetId}_root`, // TODO: change to real user
          },
        },
      });
    },

    /**
     * This action creates user preference
     *
     * @param {function} dispatch
     * @param {Object} userPreference
     */
    async create({ dispatch }, { userPreference }) {
      try {
        const body = omit(userPreference, ['crecord_creation_time', 'crecord_write_time', 'enable']);

        await dispatch('entities/update', {
          route: `${API_ROUTES.userPreferences}`,
          schema: userPreferenceSchema,
          body: JSON.stringify(body),
          dataPreparer: d => d.data[0],
        }, { root: true });
      } catch (err) {
        console.warn(err);
      }
    },

    async setActiveFilter({ commit, getters }, { data, selectedFilter }) {
      try {
        await request.post(API_ROUTES.userPreferences, {
          widget_preferences: {
            user_filters: getters.filters.map(filter => ({
              filter: filter.filter,
              title: filter.title,
            })),
            selected_filter: {
              filter: selectedFilter.filter,
              isActive: true,
              title: selectedFilter.title,
            },
          },
          crecord_name: data.crecord_name,
          widget_id: data.widget_id,
          widgetXtype: data.widgetXtype,
          title: data.title,
          id: data.id,
          _id: data._id,
          crecord_type: data.crecord_type,
        });
        commit(types.SET_ACTIVE_FILTER, selectedFilter);
      } catch (e) {
        console.warn(e);
      }
    },
  },
};
