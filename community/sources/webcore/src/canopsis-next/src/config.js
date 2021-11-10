import { removeTrailingSlashes } from '@/helpers/url';

export const {
  BASE_URL,
  VUE_APP_API_HOST,
  VUE_APP_LOCAL_STORAGE_ACCESS_TOKEN_KEY,
  VUE_APP_PAGINATION_LIMIT,
} = {
  ...process.env,
  ...window.injectedEnv,
};

export const APP_HOST = removeTrailingSlashes(`${window.location.origin}${BASE_URL}`);

export const API_HOST = VUE_APP_API_HOST || removeTrailingSlashes(`${window.location.origin}/backend`);

export const SOCKET_HOST = API_HOST.replace(/^http(s?)/, 'wss');

export const SOCKET_ROUTE = '/api/v4/ws';

export const SOCKET_URL = removeTrailingSlashes(`${SOCKET_HOST}${SOCKET_ROUTE}`);

export const ROUTER_MODE = 'history';

export const LOCAL_STORAGE_ACCESS_TOKEN_KEY = VUE_APP_LOCAL_STORAGE_ACCESS_TOKEN_KEY || 'accessToken';

export const PAGINATION_LIMIT = parseInt(VUE_APP_PAGINATION_LIMIT, 10);

export const PAGINATION_PER_PAGE_VALUES = [5, 10, 20, 50, 100];

export const PAGINATION_TOTAL_VISIBLE = 7;

export const DEFAULT_MAX_MULTI_SORT_COLUMNS_COUNT = 3;

export const DEFAULT_WEATHER_LIMIT = 120;

export const DEFAULT_LOCALE = 'en';

export const DEFAULT_VIEW_STATS_INTERVAL = 120000;

export const LOCALE_PRIORITIES = {
  default: 1,
  global: 2,
  personal: 3,
};

export const POPUP_AUTO_CLOSE_DELAY = 3000;

export const ACTIVE_LOGGED_USERS_COUNT_FETCHING_INTERVAL = 60000;

export const ACTIVE_BROADCAST_MESSAGE_FETCHING_INTERVAL = 60000;

export const VUETIFY_ANIMATION_DELAY = 300;

export const POPUP_TICK_DELAY = 100;

export const SIDE_BAR_WIDTH = 250;

export const TOP_BAR_HEIGHT = 48;

export const EXPAND_DEFAULT_MAX_LETTERS = 50;

export const SCHEMA_EMBEDDED_KEY = '_embedded';

export const MEDIA_QUERIES_BREAKPOINTS = {
  m: 450,
  t: 900,
  l: 1400,
  xl: 'Infinity',
};

export const MAX_LOGO_SIZE_IN_KB = 16384;

export const DEFAULT_APP_TITLE = 'Canopsis';

export const EXPORT_VIEWS_AND_GROUPS_FILENAME_PREFIX = 'canopsis_groups_views-';

export const API_ROUTES = {
  login: '/api/v4/login',
  logout: '/api/v4/logout',
  loggedUserCount: '/api/v4/logged-user-count',
  viewStats: '/api/v4/view-stats',
  currentUser: '/api/v4/account/me',
  alarmList: '/api/v4/alarms',
  entity: '/api/v4/entities',
  entityBasics: '/api/v4/entitybasics',
  service: '/api/v4/entityservices',
  serviceDependencies: '/api/v4/entityservice-dependencies',
  serviceImpacts: '/api/v4/entityservice-impacts',
  weatherService: '/api/v4/weather-services',
  alarmListExport: '/api/v4/alarm-export',
  contextExport: '/api/v4/entity-export',
  actions: '/api/v2/actions',
  event: '/api/v4/event',
  userPreferences: '/rest/userpreferences/userpreferences',
  view: '/api/v4/views',
  bulkView: '/api/v4/bulk/views',
  viewPosition: '/api/v4/view-positions',
  viewGroup: '/api/v4/view-groups',
  bulkViewGroup: '/api/v4/bulk/view-groups',
  stats: '/api/v2/stats',
  permissions: '/api/v4/permissions',
  users: '/api/v4/users',
  roles: '/api/v4/roles',
  eventFilterRules: '/api/v4/eventfilter/rules',
  file: '/api/v4/file',
  fileAccess: '/api/v4/file-access',
  snmpRule: {
    list: '/snmprule',
    create: '/snmprule/put',
  },
  snmpMib: {
    list: '/snmpmib',
    distinct: '/snmpmibdistinct',
    upload: '/uploadmib',
  },
  infos: {
    login: '/api/v4/internal/login_info',
    app: '/api/v4/internal/app_info',
    userInterface: '/api/v4/internal/user_interface',
  },
  associativeTable: '/api/v4/associativetable',
  sessionTracePath: '/api/v2/session_tracepath',
  sessionsCount: '/api/v4/sessions-count',
  broadcastMessage: {
    list: '/api/v4/broadcast-message',
    activeList: '/api/v4/active-broadcast-message',
  },
  counter: '/api/v4/alarm-counters',
  playlist: '/api/v4/playlists',
  pbehavior: {
    timespan: '/api/v4/pbehavior-timespans',
    exceptions: '/api/v4/pbehavior-exceptions',
    types: '/api/v4/pbehavior-types',
    pbehaviors: '/api/v4/pbehaviors',
    pbehaviorsCount: '/api/v4/pbehaviors/count',
    pbehaviorComments: '/api/v4/pbehavior-comments',
    pbehaviorById: '/api/v4/entities/pbehaviors',
    reasons: '/api/v4/pbehavior-reasons',
  },
  engineRunInfo: '/api/v4/engine-runinfo',
  cas: {
    login: '/api/v4/cas/login',
    loggedin: '/api/v4/cas/loggedin',
  },
  saml: {
    auth: '/api/v4/saml/auth',
  },
  scenario: {
    scenarios: '/api/v4/scenarios',
    checkPriority: '/api/v4/scenarios/check-priority',
  },
  entityCategories: '/api/v4/entity-categories',
  stateSetting: '/api/v4/state-settings/',
  dataStorage: '/api/v4/data-storage',
  notification: '/api/v4/notification/',
  idleRules: '/api/v4/idle-rules',
  idleRulesCount: '/api/v4/idle-rules/count',
  flappingRules: '/api/v4/flapping-rules',
  resolveRules: '/api/v4/resolve-rules',
  messageRateStats: '/api/v4/message-rate-stats',

  /**
   * Cat routes
   */
  dynamicInfo: '/api/v4/cat/dynamic-infos',
  metaAlarmRule: '/api/v4/cat/metaalarmrules',
  remediation: {
    instructions: '/api/v4/cat/instructions',
    instructionStats: '/api/v4/cat/instruction-stats',
    jobs: '/api/v4/cat/jobs',
    configurations: '/api/v4/cat/job-configs',
    executions: '/api/v4/cat/executions',
    jobExecutions: '/api/v4/cat/job-executions',
    pausedExecutions: '/api/v4/cat/account/paused-executions',
    instructionComments: '/api/v4/cat/instruction-comments',
  },
  junit: {
    directory: '/api/v4/cat/junit/directory',
    testSuites: '/api/v4/cat/junit/test-suites',
    entityGantt: '/api/v4/cat/junit/test-suites-entity-gantt',
    history: '/api/v4/cat/junit/test-suites-history',
    widget: '/api/v4/cat/junit/test-suites-widget',
    file: '/api/v4/cat/junit/test-cases-file',
  },
  healthcheck: {
    engines: '/api/v4/cat/healthcheck',
    status: '/api/v4/cat/healthcheck/status',
    parameters: '/api/v4/cat/healthcheck/parameters',
  },
};

export const COLORS = {
  primary: '#2fab63',
  secondary: '#2b3e4f',
  state: {
    ok: '#00a65a',
    minor: '#fcdc00',
    major: '#ff9900',
    critical: '#f56954',
    pause: '#808080',
  },
  status: {
    closed: '#2fab63',
    stealthy: '#757575',
    flapping: '#2b3E4F',
    ongoing: '#f9A825',
    cancelled: '#757575',
    noEvents: '#ff5252',
    unknown: 'black',
  },
  entitiesEvents: {
    ack: '#9c27b0',
    ackRemove: '#9c27b0',
    declareTicket: '#2196f3',
    snooze: '#e91e63',
    done: '#00a65a',
    pbhenter: '#18ffff',
    pbhleave: '#18ffff',
    comment: '#ba68c8',
    metaalarmattach: '#6d4c41',
  },
  testSuiteStatuses: {
    passed: '#2fab63',
    error: '#ffa800',
    failed: '#ff5252',
    skipped: '#5a6D80',
  },
  healthcheck: {
    error: '#FF5252',
    warning: '#FFA800',
    unknown: '#C4C4C4',
    edgeGray: '#979797',
    edgeBlack: '#000000',
  },
  statsDefault: '#dddddd',
  impactState: [
    '#2FAB63',
    '#7CB342',
    '#CDDC39',
    '#DCE775',
    '#FFF59D',
    '#FFF176',
    '#FFEB3B',
    '#FDD835',
    '#FBC02D',
    '#FFB300',
    '#F9A825',
    '#FF8F00',
    '#F57C00',
    '#FF6F00',
    '#EF6C00',
    '#FF6A3C',
    '#FF5722',
    '#F4511E',
    '#F94C16',
    '#F44336',
    '#E64A19',
    '#D84315',
    '#E53935',
    '#D32F2F',
    '#C62828',
    '#BF360C',
    '#B71C1C',
    '#C51162',
    '#C2185B',
    '#AD1457',
    '#880E4F',
  ],
};

export const FILE_BASE_URL = `${API_HOST}${API_ROUTES.file}`;

export const DOCUMENTATION_BASE_URL = 'https://doc.canopsis.net/';

export const INSTRUCTION_EXECUTE_FETCHING_INTERVAL = 10000;

export const INSTRUCTION_EXECUTE_FETCHING_INTERVAL_SECONDS = 10;

export const EXPORT_FETCHING_INTERVAL = 2000;

export const DEFAULT_CATEGORIES_LIMIT = 3;

export const HEALTHCHECK_HISTORY_FILENAME_PREFIX = 'healthcheck_history-';

export const HEALTHCHECK_LAST_HOUR_FILENAME_PREFIX = 'healthcheck_last_hour-';

export const LOGIN_INFOS_FETCHING_INTERVAL = 10000;
