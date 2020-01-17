import alarmModule from './alarm';
import entityModule from './entity';
import watcherModule from './watcher';
import pbehaviorModule from './pbehavior';
import userPreferenceModule from './user-preference';
import viewModule from './view';
import statsModule from './stats';
import roleModule from './role';
import userModule from './user';
import rightModule from './right';
import eventFilterRuleModule from './event-filter-rule';
import webhookModule from './webhook';
import infoModule from './info';
import filterHintModule from './filter-hint';
import snmpRuleModule from './snmp/rule';
import snmpMibModule from './snmp/mib';
import actionModule from './action';
import heartbeatModule from './heartbeat';
import keepaliveModule from './keepalive';
import dynamicInfoModule from './dynamic-info';
import alarmColumnFiltersModule from './alarm-column-filters';
import pbehaviorReasonsModule from './pbehavior-reasons';

export default {
  alarm: alarmModule,
  entity: entityModule,
  watcher: watcherModule,
  pbehavior: pbehaviorModule,
  userPreference: userPreferenceModule,
  view: viewModule,
  stats: statsModule,
  role: roleModule,
  user: userModule,
  right: rightModule,
  eventFilterRule: eventFilterRuleModule,
  webhook: webhookModule,
  info: infoModule,
  snmpRule: snmpRuleModule,
  snmpMib: snmpMibModule,
  action: actionModule,
  heartbeat: heartbeatModule,
  keepalive: keepaliveModule,
  dynamicInfo: dynamicInfoModule,
  filterHint: filterHintModule,
  pbehaviorReasons: pbehaviorReasonsModule,
  alarmColumnFilters: alarmColumnFiltersModule,
};
