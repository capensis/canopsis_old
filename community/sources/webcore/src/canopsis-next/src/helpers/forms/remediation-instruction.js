import { isUndefined, omit, pick, cloneDeep } from 'lodash';

import {
  REMEDIATION_INSTRUCTION_APPROVAL_TYPES,
  REMEDIATION_INSTRUCTION_TYPES,
  TIME_UNITS,
  WORKFLOW_TYPES,
} from '@/constants';

import uuid from '@/helpers/uuid';
import { durationToForm, formToDuration } from '@/helpers/date/duration';
import { flattenErrorMap } from '@/helpers/forms/flatten-error-map';

import { enabledToForm } from './shared/common';

/**
 * @typedef {Object} RemediationInstructionStepOperation
 * @property {string} name
 * @property {string} description
 * @property {Array} jobs
 * @property {Duration} time_to_complete
 */

/**
 * @typedef {RemediationInstructionStepOperation} RemediationInstructionStepOperationForm
 * @property {DurationForm} time_to_complete
 * @property {string} [key]
 */

/**
 * @typedef {Object} RemediationInstructionStep
 * @property {string} endpoint
 * @property {string} name
 * @property {boolean} stop_on_fail
 * @property {RemediationInstructionStepOperation[]} operations
 */

/**
 * @typedef {RemediationInstructionStep} RemediationInstructionStepForm
 * @property {RemediationInstructionStepOperationForm[]} operations
 * @property {string} [key]
 */

/**
 * @typedef {Object} RemediationInstructionApproval
 * @property {User} [user]
 * @property {Role} [role]
 * @property {string} comment
 * @property {string} requested_by
 */

/**
 * @typedef {RemediationInstructionApproval} RemediationInstructionApprovalForm
 * @property {boolean} need_approve
 * @property {number} type
 */

/**
 * @typedef {RemediationInstructionApproval} RemediationInstructionJob
 * @property {RemediationJob} job
 * @property {boolean} stop_on_fail
 */

/**
 * @typedef {RemediationInstructionJob} RemediationInstructionJobForm
 * @property {string} [key]
 */

/**
 * @typedef {Object} RemediationInstruction
 * @property {number} type
 * @property {string} name
 * @property {number} priority
 * @property {boolean} enabled
 * @property {string} description
 * @property {Duration} timeout_after_execution
 * @property {Array} alarm_patterns
 * @property {Array} entity_patterns
 * @property {string[]} active_on_pbh
 * @property {string[]} disabled_on_pbh
 * @property {RemediationInstructionStep[]} steps
 * @property {RemediationInstructionJob[]} jobs
 * @property {RemediationInstructionApproval} approval
 */

/**
 * @typedef {RemediationInstruction} RemediationInstructionForm
 * @property {DurationForm} timeout_after_execution
 * @property {RemediationInstructionStepForm[]} steps
 * @property {RemediationInstructionJobForm[]} jobs
 * @property {RemediationInstructionApprovalForm} approval
 */

/**
 * @typedef {Object | undefined} RemediationInstructionApprovalRequest
 * @property {string} [user]
 * @property {string} [role]
 * @property {string} comment
 */

/**
 * Convert a remediation instruction step operation to form
 *
 * @param {RemediationInstructionStepOperation} [operation]
 * @returns {RemediationInstructionStepOperationForm}
 */
export const remediationInstructionStepOperationToForm = (operation = {}) => ({
  name: operation.name || '',
  description: operation.description || '',
  time_to_complete: operation.time_to_complete
    ? durationToForm(operation.time_to_complete)
    : { value: 0, unit: TIME_UNITS.minute },
  jobs: operation.jobs ? cloneDeep(operation.jobs) : [],
  key: uuid(),
});

/**
 * Convert a remediation instruction step operation array to form array
 *
 * @param {RemediationInstructionStepOperation[]} [operations = [undefined]]
 * @returns {Array}
 */
const remediationInstructionStepOperationsToForm = (operations = [undefined]) =>
  operations.map(remediationInstructionStepOperationToForm);

/**
 * Convert a remediation instruction step to form
 *
 * @param {RemediationInstructionStep} [step]
 * @returns {RemediationInstructionStepForm}
 */
export const remediationInstructionStepToForm = (step = {}) => ({
  endpoint: step.endpoint || '',
  name: step.name || '',
  stop_on_fail: !isUndefined(step.stop_on_fail) ? step.stop_on_fail : WORKFLOW_TYPES.stop,
  operations: remediationInstructionStepOperationsToForm(step.operations),
  key: uuid(),
});

/**
 * Convert a remediation instruction steps array to form array
 *
 * @param {RemediationInstructionStep[]} [steps = [undefined]]
 * @returns {RemediationInstructionStepForm[]}
 */
const remediationInstructionStepsToForm = (steps = [undefined]) => steps.map(remediationInstructionStepToForm);

/**
 * Convert a remediation instruction approval to form
 *
 * @param {RemediationInstructionApproval} approval
 * @return {RemediationInstructionApprovalForm}
 */
const remediationInstructionApprovalToForm = (approval = {}) => ({
  need_approve: !!approval.comment,
  type: approval.user && approval.user._id
    ? REMEDIATION_INSTRUCTION_APPROVAL_TYPES.user
    : REMEDIATION_INSTRUCTION_APPROVAL_TYPES.role,
  user: approval.user,
  role: approval.role,
  comment: approval.comment || '',
});

/**
 * Convert a remediation instruction job to form
 *
 * @param {RemediationInstructionJob} [job = {}]
 * @returns {RemediationInstructionJobForm}
 */
export const remediationInstructionJobToForm = (job = {}) => ({
  job: job.job,
  stop_on_fail: !isUndefined(job.stop_on_fail) ? job.stop_on_fail : WORKFLOW_TYPES.stop,
  key: uuid(),
});

/**
 * Convert a remediation instruction jobs array to form array
 *
 * @param {RemediationInstructionJob[]} [jobs = [undefined]]
 * @returns {RemediationInstructionJobForm[]}
 */
const remediationInstructionJobsToForm = (jobs = [undefined]) => jobs.map(remediationInstructionJobToForm);

/**
 * Convert a remediation instruction object to form object
 *
 * @param {RemediationInstruction} remediationInstruction
 * @returns {RemediationInstructionForm}
 */
export const remediationInstructionToForm = (remediationInstruction = {}) => ({
  name: remediationInstruction.name || '',
  priority: remediationInstruction.priority || 0,
  type: !isUndefined(remediationInstruction.type) ? remediationInstruction.type : REMEDIATION_INSTRUCTION_TYPES.manual,
  enabled: enabledToForm(remediationInstruction.enabled),
  timeout_after_execution: durationToForm(remediationInstruction.timeout_after_execution),
  active_on_pbh: remediationInstruction.active_on_pbh
    ? cloneDeep(remediationInstruction.active_on_pbh)
    : [],
  disabled_on_pbh: remediationInstruction.disabled_on_pbh
    ? cloneDeep(remediationInstruction.disabled_on_pbh)
    : [],
  alarm_patterns: remediationInstruction.alarm_patterns
    ? cloneDeep(remediationInstruction.alarm_patterns)
    : [],
  entity_patterns: remediationInstruction.entity_patterns
    ? cloneDeep(remediationInstruction.entity_patterns)
    : [],
  description: remediationInstruction.description || '',
  steps: remediationInstructionStepsToForm(remediationInstruction.steps),
  approval: remediationInstructionApprovalToForm(remediationInstruction.approval),
  jobs: remediationInstructionJobsToForm(remediationInstruction.jobs),
});


/**
 * Convert a remediation instruction step operations form array to a API compatible operation array
 *
 * @param {RemediationInstructionJobForm[]} jobs
 * @returns {RemediationInstructionJob[]}
 */
const formJobsToRemediationInstructionJobs = (jobs = []) => jobs.map(job => ({
  ...omit(job, ['key']),

  job: job.job._id,
}));


/**
 * Convert a remediation instruction step operations form array to a API compatible operation array
 *
 * @param {RemediationInstructionStepOperationForm[]} operations
 * @returns {RemediationInstructionStepOperation[]}
 */
const formOperationsToRemediationInstructionOperations = operations => operations.map(operation => ({
  ...omit(operation, ['key']),

  time_to_complete: formToDuration(operation.time_to_complete),
  jobs: operation.jobs.map(({ _id }) => _id),
}));

/**
 * Convert a remediation instruction steps form array to a API compatible array
 *
 * @param {RemediationInstructionStepForm[]} steps
 * @returns {RemediationInstructionStep[]}
 */
const formStepsToRemediationInstructionSteps = steps => steps.map(step => ({
  ...omit(step, ['key']),

  operations: formOperationsToRemediationInstructionOperations(step.operations),
}));

/**
 * Convert a remediation instruction approval form
 *
 * @param {RemediationInstructionApprovalForm} approval
 * @returns {RemediationInstructionApprovalRequest}
 */
const formApprovalToRemediationInstructionApproval = (approval) => {
  if (!approval.need_approve) {
    return undefined;
  }

  const data = pick(approval, ['comment']);

  if (approval.type === REMEDIATION_INSTRUCTION_APPROVAL_TYPES.role) {
    data.role = approval.role._id;
  } else {
    data.user = approval.user._id;
  }

  return data;
};

/**
 * Convert a remediation instruction form object to a API compatible object
 *
 * @param {RemediationInstructionForm} form
 * @returns {RemediationInstruction}
 */
export const formToRemediationInstruction = (form) => {
  const {
    steps, jobs, priority, ...instruction
  } = form;

  if (form.type === REMEDIATION_INSTRUCTION_TYPES.manual) {
    instruction.steps = formStepsToRemediationInstructionSteps(steps);
  } else {
    instruction.jobs = formJobsToRemediationInstructionJobs(jobs);
    instruction.priority = priority;
  }

  return {
    ...instruction,
    timeout_after_execution: formToDuration(form.timeout_after_execution),
    alarm_patterns: form.alarm_patterns.length ? form.alarm_patterns : undefined,
    entity_patterns: form.entity_patterns.length ? form.entity_patterns : undefined,
    approval: formApprovalToRemediationInstructionApproval(form.approval),
  };
};

/**
 * Convert error structure to form structure
 *
 * @param {FlattenErrors} errors
 * @param {RemediationInstructionForm} form
 * @return {FlattenErrors}
 */
export const remediationInstructionErrorsToForm = (errors, form) => flattenErrorMap(errors, (errorsObject) => {
  const { jobs, ...errorMessages } = errorsObject;

  if (jobs) {
    errorMessages.jobs = jobs.reduce((acc, messages, index) => {
      const job = form.jobs[index];
      acc[job.key] = messages;

      return acc;
    }, {});
  }

  return errorMessages;
});
