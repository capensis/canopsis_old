import { EVENT_FILTER_TYPES, EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES, EVENT_FILTER_FAILURE_TYPES } from '@/constants';

export default {
  externalData: 'External data',
  actionsRequired: 'Please add at least one action',
  configRequired: 'No configuration defined. Please add at least one config parameter',
  idHelp: 'If no id is specified, a unique id will be generated automatically on rule creation',
  editPattern: 'Edit pattern',
  advanced: 'Advanced',
  addAField: 'Add a field',
  simpleEditor: 'Simple editor',
  field: 'Field',
  value: 'Value',
  advancedEditor: 'Advanced editor',
  comparisonRules: 'Comparison rules',
  editActions: 'Edit actions',
  addAction: 'Add an action',
  editAction: 'Edit an action',
  actions: 'Actions',
  onSuccess: 'On success',
  onFailure: 'On failure',
  configuration: 'Configuration',
  resource: 'Resource ID or template',
  component: 'Component ID or template',
  connector: 'Connector ID or template',
  connectorName: 'Connector name or template',
  duringPeriod: 'Applied during this period only',
  enrichmentOptions: 'Enrichment options',
  changeEntityOptions: 'Change entity options',
  eventsFilteredSinceLastUpdate: 'Events filtered since last update',
  errorsSinceLastUpdate: 'Errors since last update',
  markAsRead: 'Mark as read',
  filterByType: 'Filter by type',
  copyEventToClipboard: 'Copy event to clipboard',
  event: 'Event',
  eventCopied: 'Event copied to clipboard',
  syntaxIsValid: 'Syntax is valid',
  types: {
    [EVENT_FILTER_TYPES.drop]: 'Drop',
    [EVENT_FILTER_TYPES.break]: 'Break',
    [EVENT_FILTER_TYPES.enrichment]: 'Enrichment',
    [EVENT_FILTER_TYPES.changeEntity]: 'Change entity',
  },
  failureTypes: {
    [EVENT_FILTER_FAILURE_TYPES.invalidPattern]: 'Invalid pattern',
    [EVENT_FILTER_FAILURE_TYPES.invalidTemplate]: 'Invalid template',
    [EVENT_FILTER_FAILURE_TYPES.externalDataMongo]: 'Mongo',
    [EVENT_FILTER_FAILURE_TYPES.externalDataApi]: 'External API',
    [EVENT_FILTER_FAILURE_TYPES.other]: 'Other',
  },
  tooltips: {
    addValueRuleField: 'Add value rule field',
    editValueRuleField: 'Edit value rule field',
    addObjectRuleField: 'Add object rule field',
    editObjectRuleField: 'Edit object rule field',
    removeRuleField: 'Remove rule field',
  },
  actionsTypes: {
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.copy]: {
      text: 'Copy a value from a field of event to another',
      message: 'This action is used used to copy the value of a control in an event.',
      description: 'The parameters of the action are:\n- value: the name of the control whose value must be copied. It can be an event field, a subgroup of a regular expression, or an external data.\n- description (optional): the description.\n- name: the name of the event field into which the value must be copied.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.copyToEntityInfo]: {
      text: 'Copy a value from a field of an event to an info of an entity',
      message: 'This action is used to copy the field value of an event to the field of an entity.',
      description: 'The parameters of the action are:\n- description (optional): the description.\n- name: the name of the field of an entity.\n- value: the name of the control whose value must be copied. It can be an event field, a subgroup of a regular expression, or an external data.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setEntityInfo]: {
      text: 'Set an info of an entity to a constant',
      message: 'This action is used to set the dynamic information from an entity corresponding to the event.',
      description: 'The parameters of the action are:\n- description (optional): the description.\n- name: the name of the field.\n- value: the value of a field.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setEntityInfoFromTemplate]: {
      text: 'Set a string info of an entity using a template',
      message: 'This action is used to modify the dynamic information from an entity corresponding to the event.',
      description: 'The parameters of the action are:\n- description (optional): the description\n- name: the name of the field.\n- value: the template used to determine the value of the data item.\nTemplates {{.Event.NomDuChamp}}, regular expressions or external data can be used.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setField]: {
      text: 'Set a field of an event to a constant',
      message: 'This action can be used to modify a field of the event.',
      description: 'The parameters of the action are:\n- description (optional): the description.\n- name: the name of the field.\n- value: the new value of the field.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setFieldFromTemplate]: {
      text: 'Set a string field of an event using a template',
      message: 'This action allows you to modify an event field from a template.',
      description: 'The parameters of the action are:\n- description (optional): the description.\n- name: the name of the field.\n- value: the template used to determine the value of the field.\n Templates {{.Event.NomDuChamp}}, regular expressions or external data can be used.',
    },
    [EVENT_FILTER_ENRICHMENT_ACTIONS_TYPES.setEntityInfoFromDictionary]: {
      text: 'Set entity info from a dictionary',
      message: 'This action can be used for setting entity infos from event fields with a dictionary type node.',
      description: 'The parameters of the action are:\n- Value: the event field from which the infos are retrieved. The value must contain an array of name: value pairs.\n- Description (optional): the description which is used for the entity infos description. If not defined, the entity infos description is left empty.',
    },
  },
};
