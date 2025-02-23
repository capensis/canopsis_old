<template>
  <v-layout class="availability-graph position-relative" align-start justify-center>
    <c-progress-overlay :pending="pending" />

    <div class="availability-graph__content">
      <availability-history
        :query="query"
        :availabilities="availabilities"
        :default-show-type="defaultShowType"
        :display-parameter="displayParameter"
        :downloading="downloading"
        @update:sampling="updateQueryField('sampling', $event)"
        @export:csv="exportAvailabilityHistoryAsCsv"
        @export:png="exportAvailabilityHistoryAsPng"
      />
    </div>
  </v-layout>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';
import { isEqual } from 'lodash';

import { AVAILABILITY_DISPLAY_PARAMETERS, DATETIME_FORMATS, SAMPLINGS, TIME_UNITS } from '@/constants';
import { AVAILABILITY_FILENAME_PREFIX } from '@/config';

import { convertDateToString, getDiffBetweenDates } from '@/helpers/date/date';
import { convertMetricsToTimezone } from '@/helpers/entities/metric/list';
import { saveFile } from '@/helpers/file/files';
import { getExportMetricDownloadFileUrl } from '@/helpers/entities/metric/url';
import { openUrlInNewTab } from '@/helpers/url';

import { localQueryMixin } from '@/mixins/query/query';

import AvailabilityHistory from '@/components/other/availability/partials/availability-history.vue';

const { mapActions: mapAvailabilityActions } = createNamespacedHelpers('availability');

export default {
  components: { AvailabilityHistory },
  mixins: [
    localQueryMixin,
  ],
  props: {
    availability: {
      type: Object,
      required: true,
    },
    interval: {
      type: Object,
      required: true,
    },
    defaultShowType: {
      type: Number,
      required: false,
    },
    displayParameter: {
      type: Number,
      default: AVAILABILITY_DISPLAY_PARAMETERS.uptime,
    },
  },
  data() {
    return {
      pending: false,
      downloading: false,
      availabilities: [],
      query: {
        sampling: this.getSamplingByInterval(),
      },
    };
  },
  watch: {
    interval(interval, oldInterval) {
      if (!isEqual(interval, oldInterval)) {
        this.fetchList();
      }
    },
  },
  mounted() {
    this.fetchList();
  },
  methods: {
    ...mapAvailabilityActions({
      fetchAvailabilityHistoryWithoutStore: 'fetchAvailabilityHistoryWithoutStore',
      fetchAvailabilityHistoryExport: 'fetchAvailabilityHistoryExport',
    }),

    getSamplingByInterval() {
      const daysDiff = getDiffBetweenDates(this.interval.to, this.interval.from, TIME_UNITS.day);

      if (daysDiff > 1) {
        return SAMPLINGS.day;
      }

      return SAMPLINGS.hour;
    },

    getQuery() {
      return {
        _id: this.availability.entity._id,
        from: this.interval.from,
        to: this.interval.to,
        sampling: this.query.sampling,
      };
    },

    async fetchList() {
      this.pending = true;

      try {
        const availabilities = await this.fetchAvailabilityHistoryWithoutStore({
          params: this.getQuery(),
        });

        this.availabilities = convertMetricsToTimezone(availabilities);
      } catch (err) {
        console.error(err);
      } finally {
        this.pending = false;
      }
    },

    getFileName() {
      const { from, to } = this.interval;

      const fromTime = convertDateToString(from, DATETIME_FORMATS.short);
      const toTime = convertDateToString(to, DATETIME_FORMATS.short);

      return [
        AVAILABILITY_FILENAME_PREFIX,
        fromTime,
        toTime,
        this.query.sampling,
      ].join('-');
    },

    async exportAvailabilityHistoryAsPng(blob) {
      try {
        await saveFile(blob, this.getFileName());
      } catch (err) {
        console.error(err);

        this.$popups.error({ text: err.message || this.$t('errors.default') });
      }
    },

    async exportAvailabilityHistoryAsCsv() {
      this.downloading = true;

      try {
        const fileData = await this.fetchAvailabilityHistoryExport({
          params: this.getQuery(),
        });

        openUrlInNewTab(getExportMetricDownloadFileUrl(fileData._id));
      } catch (err) {
        this.$popups.error({ text: this.$t('availability.popups.exportCSVFailed') });
      } finally {
        this.downloading = false;
      }
    },
  },
};
</script>

<style lang="scss">
.availability-graph {
  &__content {
    max-width: 900px;
    width: 100%;
  }
}
</style>
