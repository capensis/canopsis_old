import { omit, isUndefined, isEmpty } from 'lodash';

import { PAGINATION_LIMIT } from '@/config';
import { WIDGET_TYPES, LIVE_REPORTING_INTERVALS } from '@/constants';

import { prepareMainFilterToQueryFilter } from './filter';

/**
 * WIDGET CONVERTERS
 */

/**
 * This function converts widget.default_sort_column to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertSortToQuery({ parameters }) {
  const { sort } = parameters;

  if (sort && sort.column && sort.order) {
    return { sortKey: sort.column, sortDir: sort.order };
  }

  return { sortKey: null, sortDir: null };
}

/**
 * This function converts widget with type 'AlarmsList' to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertAlarmWidgetToQuery(widget) {
  const {
    alarmsStateFilter = {},
    widgetColumns,
    itemsPerPage,
    mainFilter,
  } = widget.parameters;

  const query = {
    page: 1,
    opened: alarmsStateFilter.opened || false,
    resolved: alarmsStateFilter.resolved || false,
    limit: itemsPerPage || PAGINATION_LIMIT,
  };

  if (!isEmpty(mainFilter)) {
    query.filter = mainFilter.filter;
  }

  if (query.resolved) {
    query.interval = LIVE_REPORTING_INTERVALS.last30Days;
  }

  if (widgetColumns) {
    query.active_columns = widgetColumns.map(v => v.value);
  }

  return { ...query, ...convertSortToQuery(widget) };
}

/**
 * This function converts widget with type 'Context' to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertContextWidgetToQuery(widget) {
  const {
    itemsPerPage,
    selectedTypes,
  } = widget.parameters;

  const query = {
    page: 1,
    limit: itemsPerPage || PAGINATION_LIMIT,
    selectedTypes,
  };

  return { ...query, ...convertSortToQuery(widget) };
}

export function convertWeatherWidgetToQuery(widget) {
  const query = {
    filter: widget.parameters.mfilter.filter,
  };

  return query;
}

/**
 * This function converts widget with type stats field to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertWidgetStatsParameterToQuery(widget) {
  const statsList = Object.keys(widget.parameters.stats).reduce((acc, stat) => {
    acc[stat] = {
      ...widget.parameters.stats[stat],
      stat: widget.parameters.stats[stat].stat.value,
    };
    return acc;
  }, {});

  return {
    ...widget.parameters,

    stats: statsList,
  };
}

/**
 * This function converts widget with type 'StatsCalendar' to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertStatsCalendarWidgetToQuery(widget) {
  const {
    filters,
    alarmsStateFilter,
    considerPbehaviors,
  } = widget.parameters;

  const query = {
    considerPbehaviors,
    filters: filters || [],
  };

  if (alarmsStateFilter) {
    if (!isUndefined(alarmsStateFilter.opened)) {
      query.opened = alarmsStateFilter.opened;
    }

    if (!isUndefined(alarmsStateFilter.resolved)) {
      query.resolved = alarmsStateFilter.resolved;
    }
  }

  return query;
}

/**
 * This function converts widget with type 'Stats number' to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertStatsNumberWidgetToQuery(widget) {
  const { stat } = widget.parameters;
  const query = {
    ...omit(widget.parameters, ['statColors', 'criticityLevels', 'yesNoMode', 'statName']),

    trend: true,
  };

  if (stat) {
    query.stats = {
      [stat.title]: {
        parameters: stat.parameters,
        stat: stat.stat.value,
        trend: true,
      },
    };
  }

  query.trend = true;

  return query;
}

/**
 * USER_PREFERENCE CONVERTERS
 */

/**
 * This function converts userPreference with widgetXtype 'AlarmsList' to query Object
 *
 * @param {Object} userPreference
 * @returns {{}}
 */
export function convertAlarmUserPreferenceToQuery({ widget_preferences: widgetPreferences }) {
  const query = {};
  const {
    itemsPerPage,
    mainFilter,
    mainFilterCondition,
  } = widgetPreferences;

  if (itemsPerPage) {
    query.limit = itemsPerPage;
  }

  if (!isEmpty(mainFilter)) {
    query.filter = prepareMainFilterToQueryFilter(mainFilter, mainFilterCondition);
  }

  return query;
}

/**
 * This function converts userPreference with widgetXtype 'Context' to query Object
 *
 * @param {Object} userPreference
 * @returns {{}}
 */
export function convertContextUserPreferenceToQuery({ widget_preferences: widgetPreferences = {} }) {
  const query = {};
  const {
    itemsPerPage,
    mainFilter,
    mainFilterCondition,
    selectedTypes,
  } = widgetPreferences;

  if (itemsPerPage) {
    query.limit = itemsPerPage;
  }

  if (!isEmpty(mainFilter)) {
    query.mainFilter = prepareMainFilterToQueryFilter(mainFilter, mainFilterCondition);
  }

  if (!isEmpty(selectedTypes)) {
    query.typesFilter = {
      $or: selectedTypes.map(type => ({ type })),
    };
  }

  return query;
}

/**
 * MAIN CONVERTERS
 */

/**
 * This function converts userPreference to query Object
 *
 * @param {Object} userPreference
 * @returns {{}}
 */
export function convertUserPreferenceToQuery(userPreference) {
  switch (userPreference.widgetXtype) {
    case WIDGET_TYPES.alarmList:
      return convertAlarmUserPreferenceToQuery(userPreference);
    case WIDGET_TYPES.context:
      return convertContextUserPreferenceToQuery(userPreference);
    default:
      return {};
  }
}

/**
 * This function converts widget to query Object
 *
 * @param {Object} widget
 * @returns {{}}
 */
export function convertWidgetToQuery(widget) {
  switch (widget.type) {
    case WIDGET_TYPES.alarmList:
      return convertAlarmWidgetToQuery(widget);
    case WIDGET_TYPES.context:
      return convertContextWidgetToQuery(widget);
    case WIDGET_TYPES.weather:
      return convertWeatherWidgetToQuery(widget);
    case WIDGET_TYPES.statsCurves:
    case WIDGET_TYPES.statsHistogram:
    case WIDGET_TYPES.statsTable:
    case WIDGET_TYPES.text:
      return convertWidgetStatsParameterToQuery(widget);
    case WIDGET_TYPES.statsNumber:
      return convertStatsNumberWidgetToQuery(widget);
    case WIDGET_TYPES.statsCalendar:
      return convertStatsCalendarWidgetToQuery(widget);
    default:
      return {};
  }
}

export default {
  convertContextWidgetToQuery,
  convertContextUserPreferenceToQuery,
  convertUserPreferenceToQuery,
  convertWidgetToQuery,
};
