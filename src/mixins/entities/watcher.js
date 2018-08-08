import { createNamespacedHelpers } from 'vuex';

const { mapGetters, mapActions } = createNamespacedHelpers('watcher');

/**
 * @mixin
 */
export default {
  computed: {
    ...mapGetters({
      getWatchersListByWidgetId: 'getListByWidgetId',
      getWatchersMetaByWidgetId: 'getMetaByWidgetId',
      getWatchersPendingByWidgetId: 'getPendingByWidgetId',
    }),
    watchers() {
      return this.getWatchersListByWidgetId(this.widget.id);
    },
    watchersMeta() {
      return this.getWatchersMetaByWidgetId(this.widget.id);
    },
    watchersPending() {
      return this.getWatchersPendingByWidgetId(this.widget.id);
    },
  },
  methods: {
    ...mapActions({
      fetchWatcherItem: 'fetchItem',
      fetchWatchersList: 'fetchList',
    }),
  },
};
