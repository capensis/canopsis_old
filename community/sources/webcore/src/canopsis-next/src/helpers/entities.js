import { omit, pick, isObject, groupBy, map } from 'lodash';

import i18n from '@/i18n';
import {
  ENTITIES_STATES,
  ENTITIES_STATUSES,
  DATETIME_FORMATS,
  WIDGET_TYPES,
} from '@/constants';

import uid from './uid';
import uuid from './uuid';
import { convertDateToString } from './date/date';
import { formToWidget, widgetToForm } from './forms/widgets/common';

/**
 * Convert default columns from constants to columns with prepared by i18n label
 *
 * @param {{ labelKey: string, value: string }[]} [columns = []]
 * @returns {{ label: string, value: string }[]}
 */
export function defaultColumnsToColumns(columns = []) {
  return columns.map(({ labelKey, value }) => ({
    label: i18n.tc(labelKey),
    value,
  }));
}

/**
 * Checks if alarm is resolved
 * @param alarm - alarm entity
 * @returns {boolean}
 */
export const isResolvedAlarm = alarm => [ENTITIES_STATUSES.closed, ENTITIES_STATUSES.cancelled]
  .includes(alarm.v.status.val);

/**
 * Checks if alarm have critical state
 *
 * @param alarm - alarm entity
 * @returns {boolean}
 */
export const isWarningAlarmState = alarm => ENTITIES_STATES.ok !== alarm.v.state.val;

/**
 * Add uniq key field in each entity.
 *
 * @param {Array} entities
 * @return {Array}
 */
export const addKeyInEntities = (entities = []) => entities.map(entity => ({
  ...entity,
  key: uid(),
}));

/**
 * Remove key field from each entity.
 *
 * @param {Array} entities
 * @return {Array}
 */
export const removeKeyFromEntities = (entities = []) => entities.map(entity => omit(entity, ['key']));

/**
 * Get id from entity
 *
 * @param {Object} entity
 * @param {string} idField
 * @return {string}
 */
export const getIdFromEntity = (entity, idField = '_id') => (isObject(entity) ? entity[idField] : entity);

/**
 * Get grouped steps by date
 *
 * @param {AlarmEvent[]} steps
 * @return {Object.<string, AlarmEvent[]>}
 */
export const groupAlarmSteps = (steps) => {
  const orderedSteps = [...steps].reverse();

  return groupBy(orderedSteps, step => convertDateToString(step.t, DATETIME_FORMATS.short));
};

/**
 * Return entities ids
 *
 * @param {Array} entities
 * @param {string} [idKey = '_id']
 */
export const mapIds = (entities, idKey = '_id') => map(entities, idKey);

/**
 * Pick id field from entities
 *
 * @param {U[]} entities
 * @param {string} idKey
 * @returns {PartialObject<U>[]}
 */
export const pickIds = (entities = [], idKey = '_id') => entities.map(entity => pick(entity, [idKey]));

/**
 * Filter entities by ids
 *
 * @param {Object[]} items
 * @param {Object} item
 * @param {string} [idKey = '_id']
 */
export const filterById = (items, item, idKey = '_id') => items
  .filter(({ [idKey]: itemId }) => item[idKey] !== itemId);

/**
 * Generate alarm list widget form with default parameters.
 *
 * @return {WidgetForm}
 */
export const generateDefaultAlarmListWidgetForm = () => widgetToForm({ type: WIDGET_TYPES.alarmList });

/**
 * Generate alarm list widget with default parameters.
 *
 * @return {Widget}
 */
export const generateDefaultAlarmListWidget = () => ({
  ...formToWidget(generateDefaultAlarmListWidgetForm()),

  _id: uuid(),
});

/**
 * Generate context widget with default parameters.
 *
 * @return {Widget}
 */
export const generateDefaultContextWidget = () => ({
  ...formToWidget(widgetToForm({ type: WIDGET_TYPES.context })),

  _id: uuid(),
});

/**
 * Generate alarm details id by widgetId
 *
 * @param {string} alarmId
 * @param {string} widgetId
 * @returns {string}
 */
export const generateAlarmDetailsId = (alarmId, widgetId) => `${alarmId}_${widgetId}`;

/**
 * Get dataPreparer for alarmDetails entity
 *
 * @param {string} widgetId
 * @returns {Function}
 */
export const getAlarmDetailsDataPreparer = widgetId => data => (
  data.map(item => ({
    ...item,

    /**
     * We are generating new id based on alarmId and widgetId to avoiding collision with two widgets
     * on the same view with opened expand panel on the same alarm
     */
    _id: generateAlarmDetailsId(item._id, widgetId),
  }))
);
