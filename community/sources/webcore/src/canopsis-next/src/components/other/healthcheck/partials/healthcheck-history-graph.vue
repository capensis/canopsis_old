<template>
  <div
    class="healthcheck-history"
    @touchend.stop=""
  >
    <healthcheck-history-filters
      :interval="query.interval"
      :event-types="query.eventTypes"
      :connector-names="query.connectorNames"
      :deleted-before="deletedBefore"
      class="ml-4 mb-4"
      @update:interval="updateInterval"
      @update:event-types="updateQueryField('eventTypes', $event)"
      @update:connector-names="updateQueryField('connectorNames', $event)"
    />
    <div class="position-relative">
      <c-progress-overlay :pending="pending" />

      <limited-time-line-chart
        ref="chart"
        :datasets="datasets"
        :options="historyChartOptions"
        :dark="$system.dark"
        class="text--primary"
      >
        <template #actions="{ chart }">
          <chart-export-actions
            :chart="chart"
            class="mt-4"
            hide-csv
            @export:png="exportChart"
          />
          <div class="healthcheck-history__zoom">
            <v-btn
              fab
              small
              @click="zoomIn(chart)"
            >
              <v-icon>add</v-icon>
            </v-btn>
            <v-btn
              fab
              small
              @click="zoomOut(chart)"
            >
              <v-icon>remove</v-icon>
            </v-btn>
          </div>
        </template>
      </limited-time-line-chart>
    </div>
  </div>
</template>

<script>
import { debounce } from 'lodash';

import {
  DATETIME_FORMATS,
  HEALTHCHECK_HISTORY_GRAPH_RECEIVED_FACTOR,
  MESSAGE_STATS_INTERVALS,
  QUICK_RANGES,
  TIME_UNITS,
} from '@/constants';
import { COLORS, HEALTHCHECK_HISTORY_FILENAME_PREFIX } from '@/config';

import {
  convertDateToStartOfDayTimestamp,
  convertDateToStartOfUnitTimestamp,
  convertDateToString,
  getNowTimestamp,
} from '@/helpers/date/date';
import { convertStartDateIntervalToTimestamp, convertStopDateIntervalToTimestamp } from '@/helpers/date/date-intervals';
import { convertUnit } from '@/helpers/date/duration';
import { colorToRgba } from '@/helpers/color';
import { saveFile } from '@/helpers/file/files';
import { isMetricsQueryChanged } from '@/helpers/entities/metric/query';

import { entitiesMessageRateStatsMixin } from '@/mixins/entities/message-rate-stats';
import { localQueryMixin } from '@/mixins/query/query';

import HealthcheckHistoryFilters from '@/components/other/healthcheck/partials/healthcheck-history-filters.vue';
import ChartExportActions from '@/components/common/chart/chart-export-actions.vue';

const LimitedTimeLineChart = () => import(/* webpackChunkName: "Charts" */ '@/components/common/chart/limited-time-line-chart.vue');

export default {
  inject: ['$system'],
  components: { ChartExportActions, HealthcheckHistoryFilters, LimitedTimeLineChart },
  mixins: [entitiesMessageRateStatsMixin, localQueryMixin],
  props: {
    maxMessagesPerMinute: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      pending: false,
      messagesStats: [],
      deletedBefore: 0,
      query: {
        eventTypes: [],
        connectorNames: [],
        interval: {
          from: QUICK_RANGES.last30Days.start,
          to: QUICK_RANGES.last30Days.stop,
        },
      },
    };
  },
  computed: {
    interval() {
      return {
        from: convertStartDateIntervalToTimestamp(this.query.interval.from),
        to: convertStopDateIntervalToTimestamp(this.query.interval.to),
      };
    },

    intervalWithSwipeOffset() {
      const offsetForSwipe = Math.round((this.interval.to - this.interval.from) / 2);

      return {
        from: this.interval.from - offsetForSwipe,
        to: this.interval.to + offsetForSwipe,
      };
    },

    sampling() {
      const offsetDays = convertUnit(
        this.intervalWithSwipeOffset.to - this.intervalWithSwipeOffset.from,
        TIME_UNITS.second,
        TIME_UNITS.day,
      );

      return offsetDays > 1
        ? MESSAGE_STATS_INTERVALS.hour
        : MESSAGE_STATS_INTERVALS.minute;
    },

    isMinuteSampling() {
      return this.sampling === MESSAGE_STATS_INTERVALS.minute;
    },

    maxMessagesPerInterval() {
      return this.isMinuteSampling ? this.maxMessagesPerMinute : this.maxMessagesPerMinute * 60;
    },

    suggestedMax() {
      return Math.max.apply(null, this.messagesStats.map(({ rate }) => rate))
        * HEALTHCHECK_HISTORY_GRAPH_RECEIVED_FACTOR;
    },

    datasets() {
      return [{
        backgroundColor: colorToRgba(COLORS.healthcheck.edgeGray, 0.7),
        fill: true,
        data: this.messagesStats.map(({ time, rate }) => ({
          x: time * 1000,
          y: rate,
        })),
      }];
    },

    historyChartOptions() {
      return {
        animation: false,
        scales: {
          x: {
            min: this.interval.from * 1000,
            max: this.interval.to * 1000,
            time: {
              stepSize: 2,
            },
          },
          y: {
            suggestedMax: this.suggestedMax,
          },
        },
        plugins: {
          tooltip: {
            callbacks: {
              label: this.formatTooltipMessage,
            },
          },
          limit: {
            enabled: true,
            scaleID: 'x',
            value: this.maxMessagesPerInterval,
            backgroundColor: COLORS.healthcheck.error,
            borderWidth: 1,
          },
          zoom: {
            limits: {
              x: {
                min: this.deletedBefore * 1000,
                max: convertDateToStartOfUnitTimestamp(getNowTimestamp(), TIME_UNITS.hour) * 1000,
              },
            },
            pan: {
              enabled: true,
              mode: 'x',
              threshold: 5,
              onPanComplete: this.updateChartInterval,
            },
            zoom: {
              mode: 'x',
              wheel: {
                enabled: true,
                speed: 0.1,
              },
              pinch: {
                enabled: true,
              },
              drag: {
                enabled: true,
                modifierKey: 'ctrl',
              },
              onZoom: this.updateChartInterval,
            },
          },
        },
      };
    },
  },
  created() {
    this.debouncedUpdateInterval = debounce(this.updateInterval, 300);
  },
  mounted() {
    this.fetchList();
  },
  methods: {
    formatTooltipMessage(tooltip) {
      const suffix = this.$t(`healthcheck.${this.isMinuteSampling ? 'messagesPerMinute' : 'messagesPerHour'}`);

      return `${tooltip.formattedValue} ${suffix}`;
    },

    customQueryCondition(query, oldQuery) {
      return isMetricsQueryChanged(query, oldQuery, this.deletedBefore);
    },

    async exportChart(chartBlob) {
      try {
        const fromTime = convertDateToString(this.interval.from, DATETIME_FORMATS.long);
        const toTime = convertDateToString(this.interval.to, DATETIME_FORMATS.long);

        await saveFile(chartBlob, `${HEALTHCHECK_HISTORY_FILENAME_PREFIX}${fromTime}-${toTime}`);
      } catch (err) {
        console.error(err);

        this.$popups.error({ text: err.message || this.$t('errors.default') });
      }
    },

    zoomOut(chart) {
      chart.zoom({ x: 0.9 });
    },

    zoomIn(chart) {
      chart.zoom({ x: 1.1 });
    },

    updateChartInterval({ chart }) {
      const { min, max } = chart.scales.x;

      this.debouncedUpdateInterval({
        from: Math.floor(min / 1000),
        to: Math.ceil(max / 1000),
      });
    },

    updateInterval(interval) {
      this.updateQueryField('interval', interval);
    },

    getQuery() {
      const query = {
        ...this.intervalWithSwipeOffset,
        interval: this.sampling,
      };

      if (this.query.eventTypes.length) {
        query.event_types = this.query.eventTypes;
      }

      if (this.query.connectorNames.length) {
        query.connector_names = this.query.connectorNames;
      }

      return query;
    },

    async fetchList() {
      this.pending = true;

      try {
        const { data: messagesStats, meta } = await this.fetchMessageRateStatsWithoutStore({
          params: this.getQuery(),
        });

        this.messagesStats = messagesStats;

        if (meta.deleted_before) {
          this.deletedBefore = convertDateToStartOfDayTimestamp(meta.deleted_before);

          if (this.interval.from < this.deletedBefore) {
            this.updateInterval({
              ...this.query.interval,
              from: this.deletedBefore,
            });
          }
        }
      } catch (err) {
        console.error(err);
      } finally {
        this.pending = false;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.healthcheck-history {
  &__zoom {
    display: flex;
    gap: 8px;
    position: absolute;
    top: 10px;
    right: 10px;
  }
}
</style>
