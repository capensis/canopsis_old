import { API_ROUTES } from '@/config';
import { ENTITIES_TYPES } from '@/constants';

import request from '@/services/request';

import { createEntityModule } from '@/store/plugins/entities';

export default createEntityModule({
  route: API_ROUTES.filters,
  entityType: ENTITIES_TYPES.filter,
  dataPreparer: d => d.data,
  withFetchingParams: true,
  withMeta: true,
}, {
  actions: {
    fetchListWithoutStore(context, { params }) {
      return request.get(API_ROUTES.filters, { params });
    },
  },
});
