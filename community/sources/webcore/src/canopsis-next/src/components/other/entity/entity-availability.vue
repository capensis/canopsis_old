<template>
  <v-layout class="entity-availability position-relative" align-start justify-center>
    <c-progress-overlay :pending="pending" />
    <availability-bar
      v-if="availability"
      :query="query"
      :availability="availability"
      :default-show-type="defaultShowType"
      class="entity-availability__content"
      @update:interval="updateInterval"
    />
  </v-layout>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';

import { QUICK_RANGES } from '@/constants';

import Observer from '@/services/observer';

import { convertQueryIntervalToTimestamp } from '@/helpers/date/date-intervals';

import { localQueryMixin } from '@/mixins/query/query';
import { queryIntervalFilterMixin } from '@/mixins/query/interval';

import AvailabilityBar from '@/components/other/availability/partials/availability-bar.vue';

const { mapActions: mapAvailabilityActions } = createNamespacedHelpers('availability');

export default {
  inject: {
    $periodicRefresh: {
      default() {
        return new Observer();
      },
    },
  },
  components: { AvailabilityBar },
  mixins: [localQueryMixin, queryIntervalFilterMixin],
  props: {
    entity: {
      type: Object,
      required: true,
    },
    defaultTimeRange: {
      type: String,
      default: QUICK_RANGES.today.value,
    },
    defaultShowType: {
      type: Number,
      required: false,
    },
  },
  data() {
    const { start, stop } = QUICK_RANGES[this.defaultTimeRange];

    return {
      pending: false,
      availability: null,
      query: {
        interval: {
          from: start,
          to: stop,
        },
      },
    };
  },
  mounted() {
    this.fetchList();

    this.$periodicRefresh.register(this.fetchList);
  },
  beforeDestroy() {
    this.$periodicRefresh.unregister(this.fetchList);
  },
  methods: {
    ...mapAvailabilityActions({
      fetchAvailabilityWithoutStore: 'fetchAvailabilityWithoutStore',
    }),

    getIntervalQuery() {
      const { interval } = this.query;

      if (!interval) {
        return {};
      }

      return convertQueryIntervalToTimestamp({
        interval,
        timezone: this.$system.timezone,
      });
    },

    getQuery() {
      return {
        _id: this.entity._id,
        ...this.getIntervalQuery(),
      };
    },

    async fetchList() {
      this.pending = true;

      try {
        this.availability = await this.fetchAvailabilityWithoutStore({
          params: this.getQuery(),
        });
      } catch (err) {
        console.error(err);
      } finally {
        this.pending = false;
      }
    },
  },
};
</script>

<style lang="scss">
.entity-availability {
  min-height: 100px;

  &__content {
    max-width: 900px;
    width: 100%;
  }
}
</style>
