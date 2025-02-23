import { omit } from 'lodash';

import { ALARM_STATES, ACTION_TYPES, PATTERNS_FIELDS } from '@/constants';

import { uid } from '@/helpers/uid';
import { durationToForm } from '@/helpers/date/duration';
import { getLocalTimezone } from '@/helpers/date/date';
import { objectToTextPairs, textPairsToObject } from '@/helpers/text-pairs';

import { formToPbehavior, pbehaviorToForm, pbehaviorToRequest } from '../pbehavior/form';
import { requestToForm, formToRequest } from '../shared/request/form';
import {
  declareTicketRuleWebhookDeclareTicketToForm,
  formToDeclareTicketRuleWebhookDeclareTicket,
} from '../declare-ticket/rule/form';
import { filterPatternsToForm, formFilterToPatterns } from '../filter/form';

/**
 * @typedef {
 *   "snooze" |
 *   "unsnooze" |
 *   "pbehavior" |
 *   "pbehaviorremove" |
 *   "changestate" |
 *   "webhook" |
 *   "ack" |
 *   "ackremove" |
 *   "assocticket" |
 *   "cancel"
 * } ActionType
 */

/**
 * @typedef {Object} ActionForwardAuthorParameters
 * @property {boolean} [forward_author]
 * @property {string} [author]
 */

/**
 * @typedef {ActionForwardAuthorParameters} ActionDefaultParameters
 * @property {string} output
 */

/**
 * @typedef {ActionDefaultParameters} ActionSnoozeParameters
 * @property {Duration} duration
 */

/**
 * @typedef {ActionDefaultParameters} ActionChangeStateParameters
 * @property {number} state
 */

/**
 * @typedef {ActionDefaultParameters} AssociateTicketParameters
 * @property {string} ticket
 * @property {string} ticket_url
 * @property {string} ticket_url_title
 * @property {string} ticket_system_name
 * @property {Object} ticket_data
 * @property {boolean} ticket_resources
 */

/**
 * @typedef {ActionDefaultParameters & AssociateTicketParameters} ActionAssocTicketParameters
 */

/**
 * @typedef {AssociateTicketParameters} ActionAssocTicketParametersForm
 * @property {TextPairObject[]} mapping
 */

/**
 * @typedef {Object} ActionWebhookParameters
 * @property {Request} request
 * @property {?DeclareTicketRuleWebhookDeclareTicket} [declare_ticket]
 * @property {boolean} [forward_author]
 * @property {boolean} [empty_response]
 * @property {boolean} skip_for_child
 * @property {boolean} skip_for_instruction
 * @property {string} [author]
 * @property {string} [ticket_system_name]
 */

/**
 * @typedef {ActionWebhookParameters} ActionWebhookFormParameters
 * @property {RequestForm} request
 * @property {DeclareTicketRuleWebhookDeclareTicketForm} declare_ticket
 */

/**
 * @typedef {
 *   Pbehavior |
 *   ActionDefaultParameters |
 *   ActionSnoozeParameters |
 *   ActionChangeStateParameters |
 *   ActionWebhookParameters |
 *   ActionAssocTicketParameters
 * } ActionParameters
 */

/**
 * @typedef {FilterPatterns} Action
 * @property {ActionType} type
 * @property {boolean} drop_scenario_if_not_matched
 * @property {boolean} emit_trigger
 * @property {string} comment
 * @property {ActionParameters} parameters
 */

/**
 * @typedef {
 *   PbehaviorForm |
 *   ActionDefaultParameters |
 *   ActionSnoozeParameters |
 *   ActionChangeStateParameters |
 *   ActionWebhookFormParameters |
 *   ActionAssocTicketParametersForm
 * } ActionFormParameters
 */

/**
 * @typedef {Action} ActionForm
 * @property {FilterPatternsForm} patterns
 * @property {Object.<ActionType, ActionFormParameters>} parameters
 */

/**
 * Check action type is pbehavior
 *
 * @param {ActionType} type
 * @return {boolean}
 */
export const isPbehaviorActionType = type => type === ACTION_TYPES.pbehavior;

/**
 * Check action type is pbehaviorremove
 *
 * @param {ActionType} type
 * @return {boolean}
 */
export const isPbehaviorRemoveActionType = type => type === ACTION_TYPES.pbehaviorRemove;

/**
 * Check action type is webhook
 *
 * @param {ActionType} type
 * @return {boolean}
 */
export const isWebhookActionType = type => type === ACTION_TYPES.webhook;

/**
 * Check if action type is associate ticket
 *
 * @param {ActionType} type
 * @returns {boolean}
 */
export const isAssociateTicketActionType = type => type === ACTION_TYPES.assocticket;

/**
 * Convert action parameters to form
 *
 * @param {ActionForwardAuthorParameters | {}} [parameters = {}]
 * @returns {ActionForwardAuthorParameters}
 */
const defaultActionForwardAuthorToForm = (parameters = {}) => ({
  forward_author: parameters.forward_author ?? true,
  author: parameters.author ?? '',
});

/**
 * Convert action parameters to form
 *
 * @param {ActionDefaultParameters | {}} [parameters = {}]
 * @returns {ActionDefaultParameters}
 */
const defaultActionParametersToForm = (parameters = {}) => ({
  ...defaultActionForwardAuthorToForm(parameters),
  output: parameters.output ?? '',
});

/**
 * Convert action webhook parameters to form
 *
 * @param {ActionWebhookParameters} [parameters = {}]
 * @returns {ActionWebhookFormParameters}
 */
const webhookActionParametersToForm = (parameters = {}) => ({
  ...defaultActionForwardAuthorToForm(parameters),
  declare_ticket: declareTicketRuleWebhookDeclareTicketToForm(parameters.declare_ticket),
  request: requestToForm(parameters.request),
  skip_for_child: parameters.skip_for_child ?? false,
  skip_for_instruction: parameters.skip_for_instruction ?? false,
  ticket_system_name: parameters.ticket_system_name ?? '',
});

/**
 * Convert action snooze parameters to form
 *
 * @param {ActionSnoozeParameters | {}} [parameters = {}]
 * @returns {ActionSnoozeParameters}
 */
const snoozeActionParametersToForm = (parameters = {}) => ({
  ...defaultActionParametersToForm(parameters),
  duration: durationToForm(parameters.duration),
});

/**
 * Convert action snooze parameters to form
 *
 * @param {ActionChangeStateParameters | {}} [parameters = {}]
 * @returns {ActionChangeStateParameters}
 */
const changeStateActionParametersToForm = (parameters = {}) => ({
  ...defaultActionParametersToForm(parameters),
  state: parameters.state ?? ALARM_STATES.minor,
});

/**
 * Convert action assoc ticket parameters to form
 *
 * @param {ActionAssocTicketParameters | {}} [parameters = {}]
 * @returns {ActionAssocTicketParametersForm}
 */
const assocTicketActionParametersToForm = (parameters = {}) => ({
  ...defaultActionParametersToForm(parameters),
  ticket: parameters.ticket ?? '',
  ticket_url: parameters.ticket_url ?? '',
  ticket_url_title: parameters.ticket_url_title ?? '',
  ticket_resources: parameters.ticket_resources ?? false,
  ticket_system_name: parameters.ticket_system_name ?? '',
  mapping: objectToTextPairs(parameters.ticket_data),
});

/**
 * Convert action pbehavior parameters to form
 *
 * @param {Pbehavior} [parameters = {}]
 * @param {string} [timezone = getLocalTimezone()]
 * @returns {PbehaviorForm}
 */
const pbehaviorActionParametersToForm = (parameters = {}, timezone = getLocalTimezone()) => {
  const pbehaviorForm = {
    ...defaultActionForwardAuthorToForm(parameters),
    ...pbehaviorToForm(parameters, null, timezone),
  };

  pbehaviorForm.start_on_trigger = !!parameters.start_on_trigger;
  pbehaviorForm.duration = durationToForm(parameters.duration);

  return pbehaviorForm;
};

/**
 * Prepare parameters for all action types
 *
 * @returns {Object.<ActionType, ActionFormParameters>}
 */
const prepareDefaultActionParameters = () => ({
  [ACTION_TYPES.snooze]: snoozeActionParametersToForm(),
  [ACTION_TYPES.unsnooze]: defaultActionParametersToForm(),
  [ACTION_TYPES.pbehavior]: pbehaviorActionParametersToForm(),
  [ACTION_TYPES.pbehaviorRemove]: defaultActionForwardAuthorToForm(),
  [ACTION_TYPES.changeState]: changeStateActionParametersToForm(),
  [ACTION_TYPES.ack]: defaultActionParametersToForm(),
  [ACTION_TYPES.ackremove]: defaultActionParametersToForm(),
  [ACTION_TYPES.assocticket]: assocTicketActionParametersToForm(),
  [ACTION_TYPES.cancel]: defaultActionParametersToForm(),
  [ACTION_TYPES.webhook]: webhookActionParametersToForm(),
});

/**
 * Convert action parameters to form
 *
 * @param {Action} action
 * @param {string} timezone
 * @returns {Object.<ActionType, ActionParameters>}
 */
export const actionParametersToForm = (action, timezone) => {
  const parameters = prepareDefaultActionParameters();

  if (!action.type || !action.parameters) {
    return parameters;
  }

  const parametersPreparers = {
    [ACTION_TYPES.snooze]: snoozeActionParametersToForm,
    [ACTION_TYPES.webhook]: webhookActionParametersToForm,
    [ACTION_TYPES.pbehavior]: pbehaviorActionParametersToForm,
    [ACTION_TYPES.assocticket]: assocTicketActionParametersToForm,
    [ACTION_TYPES.changeState]: changeStateActionParametersToForm,
  };

  const prepareParametersToFormFunction = parametersPreparers[action.type];

  parameters[action.type] = prepareParametersToFormFunction
    ? prepareParametersToFormFunction(action.parameters, timezone)
    : defaultActionParametersToForm({ ...action.parameters });

  return parameters;
};

/**
 * Convert action to form
 *
 * @param {Action} [action = {}]
 * @param {string} [timezone = getLocalTimezone()]
 * @returns {ActionForm}
 */
export const actionToForm = (action = {}, timezone = getLocalTimezone()) => ({
  type: action.type ?? ACTION_TYPES.snooze,
  key: uid(),
  parameters: actionParametersToForm(action, timezone),
  drop_scenario_if_not_matched: !!action.drop_scenario_if_not_matched,
  emit_trigger: !!action.emit_trigger,
  comment: action.comment ?? '',
  patterns: filterPatternsToForm(action, [PATTERNS_FIELDS.alarm, PATTERNS_FIELDS.entity]),
});

/**
 * Convert pbehavior parameters to action
 *
 * @param {ActionWebhookFormParameters | {}} [parameters = {}]
 * @return {ActionWebhookParameters}
 */
export const formToWebhookActionParameters = (parameters = {}) => ({
  declare_ticket: formToDeclareTicketRuleWebhookDeclareTicket(parameters.declare_ticket),
  request: formToRequest(parameters.request),
  empty_response: parameters.empty_response,
  skip_for_child: parameters.skip_for_child,
  skip_for_instruction: parameters.skip_for_instruction,
  ticket_system_name: parameters.ticket_system_name,
});

/**
 * Convert pbehavior parameters to action
 *
 * @param {PbehaviorForm | {}} [parameters = {}]
 * @param {string} [timezone = getLocalTimezone()]
 * @return {PbehaviorRequest}
 */
export const formToPbehaviorActionParameters = (parameters = {}, timezone = getLocalTimezone()) => {
  const pbehavior = formToPbehavior(omit(parameters, ['start_on_trigger', 'duration']), timezone);

  if (parameters.start_on_trigger) {
    pbehavior.start_on_trigger = parameters.start_on_trigger;
    pbehavior.duration = parameters.duration;
  }

  return pbehaviorToRequest(pbehavior);
};

/**
 * Convert form object to associate ticket action parameters object
 *
 * @param {ActionAssocTicketParametersForm} form
 * @return {ActionAssocTicketParameters}
 */
export const formToAssociateTicketParameters = form => ({
  ...omit(form, ['mapping']),
  ticket_data: textPairsToObject(form.mapping),
});

/**
 * Convert form to action parameters
 *
 * @param {ActionForm} form
 * @param {string} [timezone]
 * @returns {ActionParameters}
 */
const formToActionParameters = (form, timezone) => {
  const {
    forward_author: forwardAuthor,
    author,
    ...parametersByCurrentType
  } = form.parameters[form.type];

  const parametersPreparers = {
    [ACTION_TYPES.webhook]: formToWebhookActionParameters,
    [ACTION_TYPES.pbehavior]: formToPbehaviorActionParameters,
    [ACTION_TYPES.assocticket]: formToAssociateTicketParameters,
  };

  const prepareParametersToAction = parametersPreparers[form.type];
  const parameters = prepareParametersToAction
    ? prepareParametersToAction(parametersByCurrentType, timezone)
    : parametersByCurrentType;

  if (!isPbehaviorActionType(form.type)) {
    parameters.forward_author = forwardAuthor;

    if (!forwardAuthor) {
      parameters.author = author;
    }
  }

  return parameters;
};

/**
 * Convert form to action
 *
 * @param {ActionForm} form
 * @param {string} [timezone]
 * @returns {Action}
 */
export const formToAction = (form, timezone) => ({
  ...omit(form, ['key', 'patterns']),
  ...formFilterToPatterns(form.patterns, [PATTERNS_FIELDS.alarm, PATTERNS_FIELDS.entity]),
  parameters: formToActionParameters(form, timezone),
});
