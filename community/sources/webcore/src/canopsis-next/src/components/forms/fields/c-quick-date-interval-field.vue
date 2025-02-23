<template>
  <div
    :class="{ 'c-quick-interval--reverse': reverse, 'c-quick-interval--short': short }"
    class="c-quick-interval"
  >
    <template v-if="short">
      <v-menu
        :close-on-content-click="false"
        offset-y
      >
        <template #activator="{ on }">
          <v-text-field
            :value="shortValue"
            :label="$t('common.interval')"
            readonly
            hide-details
            v-on="on"
          />
        </template>
        <v-card width="400px">
          <v-card-text>
            <c-information-block :title="$t('common.interval')">
              <v-divider />
              <v-layout class="mt-2">
                <component
                  :is="withHours ? 'c-date-time-interval-field' : 'c-date-interval-field'"
                  :value="intervalObject"
                  :disabled="disabled"
                  :is-allowed-from-date="isAllowedFromDate"
                  :is-allowed-to-date="isAllowedToDate"
                  :round-hours="withHours"
                  hide-details
                  column
                  @input="updateModel($event)"
                />
                <c-quick-date-interval-type-field
                  :value="range"
                  :ranges="availableQuickRanges"
                  :disabled="disabled"
                  class="c-quick-interval__range ml-4"
                  hide-details
                  return-object
                  @input="updateIntervalRange"
                />
              </v-layout>
            </c-information-block>
          </v-card-text>
        </v-card>
      </v-menu>
    </template>
    <template v-else>
      <c-date-interval-field
        :value="intervalObject"
        :disabled="disabled"
        :is-allowed-from-date="isAllowedFromDate"
        :is-allowed-to-date="isAllowedToDate"
        @input="updateModel($event)"
      />
      <div class="c-quick-interval__range">
        <c-quick-date-interval-type-field
          :value="range"
          :ranges="availableQuickRanges"
          :disabled="disabled"
          hide-details
          return-object
          @input="updateIntervalRange"
        />
      </div>
    </template>
  </div>
</template>

<script>
import { DATETIME_FORMATS, QUICK_RANGES, TIME_UNITS } from '@/constants';

import {
  convertDateToEndOfUnitTimestamp,
  convertDateToString,
  convertDateToTimestamp,
  getNowTimestamp,
  getWeekdayNumber,
} from '@/helpers/date/date';
import {
  convertStartDateIntervalToTimestamp,
  convertStopDateIntervalToTimestamp,
  findQuickRangeValue,
} from '@/helpers/date/date-intervals';

import { formMixin } from '@/mixins/form';

export default {
  mixins: [formMixin],
  model: {
    event: 'input',
    prop: 'interval',
  },
  props: {
    interval: {
      type: Object,
      default: () => ({
        from: 0,
        to: 0,
      }),
    },
    accumulatedBefore: {
      type: Number,
      required: false,
    },
    min: {
      type: Number,
      required: false,
    },
    disabled: {
      type: Boolean,
      required: false,
    },
    reverse: {
      type: Boolean,
      default: false,
    },
    quickRanges: {
      type: Array,
      default: () => Object.values(QUICK_RANGES),
    },
    short: {
      type: Boolean,
      default: false,
    },
    withHours: {
      type: Boolean,
      default: false,
    },
    allowFuture: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    shortValue() {
      const range = findQuickRangeValue(this.range.start, this.range.stop, this.availableQuickRanges);

      if (range.value === QUICK_RANGES.custom.value) {
        return `${this.intervalFromString} - ${this.intervalToString}`;
      }

      return this.$t(`quickRanges.types.${range.value}`);
    },

    availableQuickRanges() {
      return this.quickRanges.filter(this.isAllowedQuickRange);
    },

    intervalFromAsTimestamp() {
      return convertStartDateIntervalToTimestamp(this.interval.from);
    },

    intervalToAsTimestamp() {
      return convertStopDateIntervalToTimestamp(this.interval.to);
    },

    intervalFromString() {
      return convertDateToString(
        this.intervalFromAsTimestamp,
        this.withHours ? DATETIME_FORMATS.dateTimePicker : DATETIME_FORMATS.datePicker,
      );
    },

    intervalToString() {
      return convertDateToString(
        this.intervalToAsTimestamp,
        this.withHours ? DATETIME_FORMATS.dateTimePicker : DATETIME_FORMATS.datePicker,
      );
    },

    intervalObject() {
      if (this.withHours) {
        return {
          from: this.intervalFromAsTimestamp,
          to: this.intervalToAsTimestamp,
        };
      }

      return {
        from: this.intervalFromString,
        to: this.intervalToString,
      };
    },

    range() {
      return {
        start: this.interval.from,
        stop: this.interval.to,
      };
    },
  },
  methods: {
    isGreaterMinDate(dateTimestamp) {
      if (this.min) {
        return dateTimestamp >= this.min;
      }

      return true;
    },

    isLessToDate(dateTimestamp) {
      return this.interval.to && dateTimestamp < this.intervalToAsTimestamp;
    },

    isLessNowDate(dateTimestamp) {
      return dateTimestamp <= getNowTimestamp();
    },

    isLessFutureDay(dateTimestamp) {
      return dateTimestamp <= convertDateToEndOfUnitTimestamp(Date.now(), TIME_UNITS.day);
    },

    isGreaterFromDate(dateTimestamp) {
      return dateTimestamp > this.intervalFromAsTimestamp;
    },

    isAllowedAccumulatedFromDate(dateTimestamp) {
      return this.accumulatedBefore > dateTimestamp
        /**
         * NOTE: If the date is before the accumulation date, the data is grouped by week.
         * In this case, we can only select Monday.
         */
        ? getWeekdayNumber(dateTimestamp) === 1
        : true;
    },

    isAllowedAccumulatedToDate(dateTimestamp) {
      return this.accumulatedBefore > dateTimestamp
        /**
         * NOTE: If the date is before the accumulation date, the data is grouped by week.
         * In this case, we can only select Sunday.
         */
        ? getWeekdayNumber(dateTimestamp) === 7
        : true;
    },

    isAllowedFromDate(date) {
      const dateTimestamp = convertDateToTimestamp(date);

      return this.isLessToDate(dateTimestamp)
        && this.isGreaterMinDate(dateTimestamp)
        && this.isAllowedAccumulatedFromDate(dateTimestamp);
    },

    isAllowedToDate(date) {
      const dateTimestamp = convertDateToTimestamp(date);

      return this.isGreaterFromDate(dateTimestamp)
        && this.isLessNowDate(dateTimestamp)
        && this.isAllowedAccumulatedToDate(dateTimestamp);
    },

    isAllowedQuickRange({ start, stop }) {
      if (!start || !stop) {
        return true;
      }

      const startTimestamp = convertStartDateIntervalToTimestamp(start);
      const stopTimestamp = convertStopDateIntervalToTimestamp(stop);

      return this.isGreaterMinDate(startTimestamp)
        && this.isAllowedAccumulatedFromDate(startTimestamp)
        && (this.allowFuture || this.isLessFutureDay(stopTimestamp))
        && this.isAllowedAccumulatedToDate(stopTimestamp);
    },

    updateIntervalRange({ start, stop }) {
      if (start && stop) {
        this.updateModel({
          ...this.interval,
          from: start,
          to: stop,
        });
      }
    },
  },
};
</script>

<style scoped lang="scss">
.c-quick-interval {
  display: inline-flex;
  flex-wrap: wrap;
  column-gap: 24px;

  &__range {
    display: flex;
    max-width: 180px;
  }

  &--reverse {
    flex-direction: row-reverse;
  }

  &--short {
    column-gap: 0;
  }
}
</style>
