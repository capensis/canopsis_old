import i18n from '@/i18n';
import omit from 'lodash/omit';
import set from 'lodash/set';

import request from '@/services/request';
import { API_ROUTES } from '@/config';

export default {
  namespaced: true,
  actions: {
    async fetchItemValuesWithoutStore({ dispatch }, { params }) {
      try {
        const data = await request.post(`${API_ROUTES.stats}/${params.stat.stat.value}`, { ...omit(params, ['stat']), parameters: params.stat.parameters });

        return data.values;
      } catch (err) {
        await dispatch('popup/add', { type: 'error', text: i18n.t('errors.default') }, { root: true });

        return [];
      }
    },

    async fetchListWithoutStore({ dispatch }, { params, aggregate }) {
      try {
        if (aggregate) {
          Object.keys(params.stats).forEach(stat => set(params.stats[stat], 'aggregate', aggregate));
        }

        const data = await request.post(API_ROUTES.stats, { ...params });

        return data;
      } catch (err) {
        await dispatch('popup/add', { type: 'error', text: i18n.t('errors.default') }, { root: true });

        return [];
      }
    },
  },
};
