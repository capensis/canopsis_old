import moment from 'moment';
import i18n from '@/i18n';
import { PAGINATION_LIMIT } from '@/config';
import { WIDGET_TYPES } from '@/constants';

import uuid from './uuid';

export function generateWidgetByType(type) {
  const widget = {
    type,
    _id: uuid(`widget_${type}`),
    title: '',
    minColumns: 6,
    maxColumns: 12,
    parameters: {},
  };

  let specialParameters = {};

  switch (type) {
    case WIDGET_TYPES.alarmList:
      specialParameters = {
        itemsPerPage: PAGINATION_LIMIT,
        moreInfoTemplate: '',
        alarmsStateFilter: {},
        widgetColumns: [
          {
            label: i18n.t('tables.alarmGeneral.connector'),
            value: 'alarm.connector',
          },
          {
            label: i18n.t('tables.alarmGeneral.connectorName'),
            value: 'alarm.connector_name',
          },
          {
            label: i18n.t('tables.alarmGeneral.component'),
            value: 'alarm.component',
          },
          {
            label: i18n.t('tables.alarmGeneral.resource'),
            value: 'alarm.resource',
          },
          {
            label: i18n.t('tables.alarmGeneral.output'),
            value: 'alarm.output',
          },
          {
            label: i18n.t('tables.alarmGeneral.extraDetails'),
            value: 'extra_details',
          },
          {
            label: i18n.t('tables.alarmGeneral.state'),
            value: 'alarm.state.val',
          },
          {
            label: i18n.t('tables.alarmGeneral.status'),
            value: 'alarm.status.val',
          },
        ],
        viewFilters: [],
        infoPopups: [],
        periodicRefresh: {
          enabled: false,
        },
        sort: {
          order: 'ASC',
        },
      };
      break;

    case WIDGET_TYPES.context:
      specialParameters = {
        itemsPerPage: PAGINATION_LIMIT,
        widgetColumns: [],
        selectedTypes: [],
        sort: {
          order: 'ASC',
        },
      };
      break;

    case WIDGET_TYPES.weather:
      specialParameters = {
        blockTemplate: '',
        modalTemplate: '',
        entityTemplate: '',
        columnSM: 0,
        columnMD: 0,
        columnLG: 0,
        columnHG: 0,
      };
      break;
    case WIDGET_TYPES.statsHistogram:
      specialParameters = {
        mfilter: {},
        duration: '1m',
        tstop: moment().unix(),
        groups: [],
        stats: {},
        statsColors: {},
      };
      break;
    case WIDGET_TYPES.statsCurves:
      specialParameters = {
        mfilter: {},
        duration: '1m',
        tstop: moment().unix(),
        periods: 1,
        stats: {},
      };
      break;
    case WIDGET_TYPES.statsTable:
      specialParameters = {
        duration: '1m',
        tstop: moment().unix(),
        stats: {},
        mfilter: {},
      };
      break;
  }

  widget.parameters = { ...widget.parameters, ...specialParameters };

  return widget;
}

export function generateRow() {
  return {
    _id: uuid('row'),
    title: '',
    widgets: [],
  };
}

export function generateView() {
  return {
    title: '',
    name: '',
    description: '',
    group_id: null,
    rows: [],
    tags: [],
    enabled: true,
  };
}

export default {
  generateWidgetByType,
  generateRow,
  generateView,
};
