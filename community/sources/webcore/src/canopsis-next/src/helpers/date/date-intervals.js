import moment from 'moment-timezone';

import {
  DATETIME_FORMATS,
  QUICK_RANGES,
  DATETIME_INTERVAL_TYPES,
  TIME_UNITS,
  SAMPLINGS,
} from '@/constants';

import {
  convertDateToMomentByTimezone,
  convertDateToStartOfDayMoment,
  convertDateToStartOfUnitMoment,
  convertDateToEndOfUnitMoment,
  convertDateToStartOfUnitString,
  convertDateToTimestampByTimezone,
  getLocalTimezone,
  getNowTimestamp,
  subtractUnitFromDate,
} from './date';

/**
 * @typedef {Object} Interval
 * @property {number} from
 * @property {number} to
 */

/**
 * @typedef {Object} IntervalForm
 * @property {number | string} from
 * @property {number | string} to
 */

/**
 * Convert a date interval string to moment date object
 *
 * @param {string} string
 * @param {string} type
 * @returns {Moment}
 */
export const convertStringToDateInterval = (string, type) => {
  const matches = string.match(
    /^(?<initial>now|today)((?<operator>[+--])(?<deltaValue>\d+)(?<deltaUnit>[hdwmMy]{1}))?(\/(?<roundUnit>[hdwmMy]{1}))?$/,
  );

  if (matches && matches.groups) {
    const { initial, operator, deltaValue, deltaUnit, roundUnit } = matches.groups;
    const preparedRoundUnit = roundUnit === TIME_UNITS.week ? 'isoWeek' : roundUnit;

    const isToday = initial === 'today';
    const isStart = type === DATETIME_INTERVAL_TYPES.start;
    let result = moment();

    if (isToday) {
      result = isStart || deltaValue
        ? convertDateToStartOfUnitMoment(getNowTimestamp(), TIME_UNITS.day)
        : convertDateToEndOfUnitMoment(getNowTimestamp(), TIME_UNITS.day);
    }

    if (preparedRoundUnit) {
      if (isStart) {
        result.startOf(preparedRoundUnit);
      } else {
        result.endOf(preparedRoundUnit);
      }
    }

    if (deltaValue && deltaUnit) {
      if (operator === '+') {
        result.add(deltaValue, deltaUnit);
      } else {
        result.subtract(deltaValue, deltaUnit);
      }
    }

    return result;
  }

  throw new Error('Date string pattern not recognized');
};

/**
 * Parse date in every format to moment object
 *
 * @param {LocalDate} date
 * @param {string} type
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit]
 * @return {number | moment.Moment}
 */
export const convertDateIntervalToMoment = (
  date,
  type,
  format = DATETIME_FORMATS.datePicker,
  unit,
) => {
  const momentDate = typeof date === 'number'
    ? moment.unix(date)
    : moment(date, format, true);

  if (!momentDate.isValid()) {
    return convertStringToDateInterval(date, type);
  }

  if (unit) {
    const method = type === DATETIME_INTERVAL_TYPES.start ? 'startOf' : 'endOf';

    return momentDate[method](unit);
  }

  return momentDate;
};
/**
 * Parse date in every format to moment object by timezone
 *
 * @param {LocalDate} date
 * @param {string} type
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit = SAMPLINGS.hour]
 * @param {string} [timezone = getLocalTimezone()]
 * @return {number | moment.Moment}
 */
export const convertDateIntervalToMomentByTimezone = (
  date,
  type,
  format = DATETIME_FORMATS.datePicker,
  unit = SAMPLINGS.hour,
  timezone = getLocalTimezone(),
) => (
  convertDateToMomentByTimezone(
    convertDateIntervalToMoment(date, type, format, unit),
    getLocalTimezone(),
    timezone,
  )
);

/**
 * Convert date interval to timestamp unix system.
 *
 * @param {number} date
 * @param {string} type
 * @param {string} format
 * @param {string} unit
 * @return {number}
 */
export const convertDateIntervalToTimestamp = (date, type, format, unit) => (
  convertDateIntervalToMoment(date, type, format, unit).unix()
);

/**
 * Convert date interval to timestamp unix system.
 *
 * @param {number} date
 * @param {string} type
 * @param {string} format
 * @param {string} unit
 * @param {string} [timezone = getLocalTimezone()]
 * @return {number}
 */
export const convertDateIntervalToTimestampByTimezone = (date, type, format, unit, timezone = getLocalTimezone()) => (
  convertDateIntervalToMomentByTimezone(date, type, format, unit, timezone).unix()
);

/**
 * Convert from value to timestamp
 *
 * @param {LocalDate} date
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit]
 * @return {number}
 */
export const convertStartDateIntervalToTimestamp = (
  date,
  format = DATETIME_FORMATS.datePicker,
  unit,
) => convertDateIntervalToTimestamp(
  date,
  DATETIME_INTERVAL_TYPES.start,
  format,
  unit,
);

/**
 * Convert from value to moment
 *
 * @param {LocalDate} date
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit]
 * @return {Moment}
 */
export const convertStartDateIntervalToMoment = (
  date,
  format = DATETIME_FORMATS.datePicker,
  unit,
) => convertDateIntervalToMoment(
  date,
  DATETIME_INTERVAL_TYPES.start,
  format,
  unit,
);

/**
 * Convert from value to timestamp or moment
 *
 * @param {LocalDate} date
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit = SAMPLINGS.day]
 * @param {string} [timezone = getLocalTimezone()]
 * @return {number}
 */
export const convertStartDateIntervalToTimestampByTimezone = (
  date,
  format = DATETIME_FORMATS.datePicker,
  unit = SAMPLINGS.day,
  timezone = getLocalTimezone(),
) => convertDateIntervalToTimestampByTimezone(
  date,
  DATETIME_INTERVAL_TYPES.start,
  format,
  unit,
  timezone,
);

/**
 * Convert to value to timestamp
 *
 * @param {LocalDate} date
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit]
 * @return {number}
 */
export const convertStopDateIntervalToTimestamp = (
  date,
  format = DATETIME_FORMATS.datePicker,
  unit,
) => convertDateIntervalToTimestamp(
  date,
  DATETIME_INTERVAL_TYPES.stop,
  format,
  unit,
);

/**
 * Convert to value to moment
 *
 * @param {LocalDate} date
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit]
 * @return {Moment}
 */
export const convertStopDateIntervalToMoment = (
  date,
  format = DATETIME_FORMATS.datePicker,
  unit,
) => convertDateIntervalToMoment(
  date,
  DATETIME_INTERVAL_TYPES.stop,
  format,
  unit,
);

/**
 * Convert date interval value to timestamp by timezone
 *
 * @param {LocalDate} date
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [unit = SAMPLINGS.day]
 * @param {string} [timezone = getLocalTimezone()]
 * @return {number}
 */
export const convertStopDateIntervalToTimestampByTimezone = (
  date,
  format = DATETIME_FORMATS.datePicker,
  unit = SAMPLINGS.day,
  timezone = getLocalTimezone(),
) => convertDateIntervalToTimestampByTimezone(
  date,
  DATETIME_INTERVAL_TYPES.stop,
  format,
  unit,
  timezone,
);

/**
 * Prepare date to date object
 *
 * @param {number} date
 * @param {string} type
 * @param {string} [format = DATETIME_FORMATS.dateTimePicker]
 * @param {string} [unit]
 * @returns {Date}
 */
export const convertDateIntervalToDateObject = (
  date,
  type,
  format = DATETIME_FORMATS.dateTimePicker,
  unit,
) => convertDateIntervalToMoment(date, type, format, unit).toDate();

/**
 * Find range object
 *
 * @param {string} start
 * @param {string} stop
 * @param {Object} ranges
 * @param {Object} defaultValue
 * @returns {Object}
 */
export const findQuickRangeValue = (
  start,
  stop,
  ranges = QUICK_RANGES,
  defaultValue = QUICK_RANGES.custom,
) => Object.values(ranges)
  .find(range => start === range.start && stop === range.stop) ?? defaultValue;

/**
 * Get value from quick range period
 *
 * @param {string} value
 * @param {string | number} start
 * @param {string | number} stop
 * @return {Object}
 */
export const getValueFromQuickRange = ({ value, start, stop }) => {
  if (value === QUICK_RANGES.custom.value) {
    return {
      periodUnit: TIME_UNITS.hour,
      periodValue: 1,

      tstart: convertDateToStartOfUnitString(
        subtractUnitFromDate(Date.now(), 1, TIME_UNITS.hour),
        TIME_UNITS.hour,
        DATETIME_FORMATS.dateTimePicker,
      ),

      tstop: convertDateToStartOfUnitString(
        Date.now(),
        TIME_UNITS.hour,
        DATETIME_FORMATS.dateTimePicker,
      ),
    };
  }

  return {
    tstart: start,
    tstop: stop,
  };
};

/**
 * Get diff between start and stop of quick range type
 *
 * @param {string} type
 * @return {number}
 */
export const getDiffBetweenStartAndStopQuickInterval = (type) => {
  const { start, stop } = QUICK_RANGES[type];

  if (!start || !stop) {
    return 0;
  }

  return convertStopDateIntervalToTimestamp(stop) - convertStartDateIntervalToTimestamp(start);
};

/**
 * Get quick range by diff
 *
 * @param {number} diff
 * @param {Array} ranges
 * @return {Object}
 */
export const getQuickRangeByDiffBetweenStartAndStop = (
  diff,
  ranges = Object.values(QUICK_RANGES),
) => ranges.find(range => getDiffBetweenStartAndStopQuickInterval(range.value) === diff) || QUICK_RANGES.custom;

/**
 * Find quick range by interval
 *
 * @param {Object} interval
 * @param {Array} ranges
 * @return {Object}
 */
export const findQuickRangeByInterval = (
  interval,
  ranges = Object.values(QUICK_RANGES),
) => ranges.find(
  range => range.start === interval.from && range.stop === interval.to,
)
  || QUICK_RANGES.custom;

/**
 * Convert interval form to timestamp interval
 *
 * @param {IntervalForm} [interval = {}]
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [timezone = getLocalTimezone()]
 * @returns {Interval}
 */
export const convertMetricIntervalToTimestamp = ({
  interval = {},
  format = DATETIME_FORMATS.datePicker,
  timezone = getLocalTimezone(),
}) => {
  const fromMoment = convertStartDateIntervalToMoment(interval.from, format, TIME_UNITS.day);
  const toMoment = convertStopDateIntervalToMoment(interval.to, format, TIME_UNITS.day);
  const fromStartedOfDay = convertDateToStartOfDayMoment(fromMoment);
  const toStartedOfDay = convertDateToStartOfDayMoment(toMoment);

  return {
    from: convertDateToTimestampByTimezone(fromStartedOfDay, timezone),
    to: convertDateToTimestampByTimezone(toStartedOfDay, timezone),
  };
};

/**
 * Convert interval form to timestamp interval
 *
 * @param {IntervalForm} [interval = {}]
 * @param {string} [format = DATETIME_FORMATS.datePicker]
 * @param {string} [timezone = getLocalTimezone()]
 * @returns {Interval}
 */
export const convertQueryIntervalToTimestamp = ({
  interval = {},
  format = DATETIME_FORMATS.datePicker,
  timezone = getLocalTimezone(),
}) => {
  const from = convertStartDateIntervalToTimestamp(interval.from, format, TIME_UNITS.hour);
  const to = convertStopDateIntervalToTimestamp(interval.to, format, TIME_UNITS.hour);

  return {
    from: convertDateToTimestampByTimezone(from, timezone),
    to: convertDateToTimestampByTimezone(to, timezone),
  };
};
