import { durationWithEnabledToForm, formToDurationWithEnabled } from '@/helpers/date/duration';
import { TIME_UNITS } from '@/constants';

/**
 * @typedef {Object} DataStorageJunitConfig
 * @property {DurationWithEnabled} delete_after
 */

/**
 * @typedef {Object} DataStorageRemediationConfig
 * @property {DurationWithEnabled} accumulate_after
 * @property {DurationWithEnabled} delete_after
 */

/**
 * @typedef {Object} DataStorageAlarmConfig
 * @property {DurationWithEnabled} archive_after
 * @property {DurationWithEnabled} delete_after
 */

/**
 * @typedef {Object} DataStorageEntityConfig
 * @property {boolean} archive
 * @property {boolean} archive_dependencies
 */

/**
 * @typedef {Object} DataStoragePbehaviorConfig
 * @property {DurationWithEnabled} delete_after
 */

/**
 * @typedef {Object} DataStorageConfig
 * @property {DataStorageJunitConfig} junit
 * @property {DataStorageRemediationConfig} remediation
 * @property {DataStorageAlarmConfig} alarm
 * @property {DataStorageEntityConfig} [entity]
 * @property {DataStoragePbehaviorConfig} pbehavior
 */

/**
 * @typedef {Object} HistoryWithCount
 * @property {number} archived
 * @property {number} deleted
 * @property {number} time
 */

/**
 * @typedef {Object} DataStorageHistory
 * @property {number} junit
 * @property {number} remediation
 * @property {HistoryWithCount} alarm
 * @property {HistoryWithCount} entity
 */

/**
 * @typedef {Object} DataStorage
 * @property {DataStorageConfig} config
 * @property {DataStorageHistory} history
 */

/**
 * @typedef {Object} DataStorageJunitConfigForm
 * @property {DurationWithEnabledForm} delete_after
 */

/**
 * @typedef {Object} DataStorageRemediationConfigForm
 * @property {DurationWithEnabledForm} accumulate_after
 * @property {DurationWithEnabledForm} delete_after
 */

/**
 * @typedef {Object} DataStorageAlarmConfigForm
 * @property {DurationWithEnabledForm} archive_after
 * @property {DurationWithEnabledForm} delete_after
 */

/**
 * @typedef {DataStorageEntityConfig} DataStorageEntityConfigForm
 */

/**
 * @typedef {Object} DataStoragePbehaviorConfigForm
 * @property {DurationWithEnabledForm} delete_after
 */

/**
 * @typedef {Object} DataStorageConfigForm
 * @property {DataStorageJunitConfigForm} junit
 * @property {DataStorageRemediationConfigForm} remediation
 * @property {DataStorageAlarmConfigForm} alarm
 * @property {DataStorageAlarmConfigForm} entity
 * @property {DataStoragePbehaviorConfigForm} pbehavior
 */

/**
 * @typedef {Object} DataStorageRequest
 * @property {DataStorageJunitConfigForm} junit
 */

/**
 * Convert data storage junit config to junit form object
 *
 * @param {DataStorageJunitConfig} junitConfig
 * @return {DataStorageJunitConfigForm}
 */
export const dataStorageJunitSettingsToForm = (junitConfig = {}) => ({
  delete_after: junitConfig.delete_after
    ? durationWithEnabledToForm(junitConfig.delete_after)
    : { value: 1, unit: TIME_UNITS.day, enabled: false },
});

/**
 * Convert data storage remediation config to remediation form object
 *
 * @param {DataStorageRemediationConfig} remediationConfig
 * @return {DataStorageRemediationConfigForm}
 */
export const dataStorageRemediationSettingsToForm = (remediationConfig = {}) => ({
  accumulate_after: remediationConfig.accumulate_after
    ? durationWithEnabledToForm(remediationConfig.accumulate_after)
    : { value: 1, unit: TIME_UNITS.day, enabled: false },
  delete_after: remediationConfig.delete_after
    ? durationWithEnabledToForm(remediationConfig.delete_after)
    : { value: 2, unit: TIME_UNITS.day, enabled: false },
});

/**
 * Convert data storage alarm config to alarm form object
 *
 * @param {DataStorageAlarmConfig} alarmConfig
 * @return {DataStorageAlarmConfigForm}
 */
export const dataStorageAlarmSettingsToForm = (alarmConfig = {}) => ({
  archive_after: alarmConfig.archive_after
    ? durationWithEnabledToForm(alarmConfig.archive_after)
    : { value: 1, unit: TIME_UNITS.year, enabled: false },
  delete_after: alarmConfig.delete_after
    ? durationWithEnabledToForm(alarmConfig.delete_after)
    : { value: 2, unit: TIME_UNITS.year, enabled: false },
});

/**
 * Convert data storage pbehavior config to pbehavior form object
 *
 * @param {DataStoragePbehaviorConfig} pbehaviorConfig
 * @return {DataStoragePbehaviorConfigForm}
 */
export const dataStoragePbehaviorSettingsToForm = (pbehaviorConfig = {}) => ({
  delete_after: pbehaviorConfig.delete_after
    ? durationWithEnabledToForm(pbehaviorConfig.delete_after)
    : { value: 1, unit: TIME_UNITS.year, enabled: false },
});

/**
 * Convert data storage entity config to entity form object
 *
 * @param {DataStorageEntityConfig} entityConfig
 * @return {DataStorageEntityConfigForm}
 */
export const dataStorageEntitySettingsToForm = (entityConfig = {}) => ({
  archive: entityConfig.archive || false,
  archive_dependencies: entityConfig.archive_dependencies || false,
});

/**
 * Convert data storage object to data storage form
 *
 * @param {DataStorageConfig} dataStorage
 * @return {DataStorageConfigForm}
 */
export const dataStorageSettingsToForm = (dataStorage = {}) => ({
  junit: dataStorageJunitSettingsToForm(dataStorage.junit),
  remediation: dataStorageRemediationSettingsToForm(dataStorage.remediation),
  alarm: dataStorageAlarmSettingsToForm(dataStorage.alarm),
  entity: dataStorageEntitySettingsToForm(dataStorage.entity),
  pbehavior: dataStoragePbehaviorSettingsToForm(dataStorage.pbehavior),
});

/**
 * Convert junit data storage form to junit data storage object
 *
 * @param {DataStorageJunitConfigForm} form
 * @return {DataStorageJunitConfig}
 */
export const formJunitToDataStorageSettings = (form = {}) => ({
  delete_after: formToDurationWithEnabled(form.delete_after),
});

/**
 * Convert remediation data storage form to remediation data storage object
 *
 * @param {DataStorageRemediationConfigForm} form
 * @return {DataStorageRemediationConfig}
 */
export const formToRemediationDataStorageSettings = (form = {}) => ({
  delete_after: formToDurationWithEnabled(form.delete_after),
  accumulate_after: formToDurationWithEnabled(form.accumulate_after),
});

/**
 * Convert alarm data storage form to alarm data storage object
 *
 * @param {DataStorageAlarmConfigForm} form
 * @return {DataStorageAlarmConfig}
 */
export const formToAlarmDataStorageSettings = (form = {}) => ({
  delete_after: formToDurationWithEnabled(form.delete_after),
  archive_after: formToDurationWithEnabled(form.archive_after),
});

/**
 * Convert pbehavior data storage form to pbehavior data storage object
 *
 * @param {DataStoragePbehaviorConfigForm} form
 * @return {DataStoragePbehaviorConfig}
 */
export const formToPbehaviorDataStorageSettings = (form = {}) => ({
  delete_after: formToDurationWithEnabled(form.delete_after),
});

/**
 * Convert data storage form to data storage object
 *
 * @param {DataStorageConfigForm} form
 * @return {DataStorageConfig}
 */
export const formToDataStorageSettings = (form = {}) => ({
  junit: formJunitToDataStorageSettings(form.junit),
  remediation: formToRemediationDataStorageSettings(form.remediation),
  alarm: formToAlarmDataStorageSettings(form.alarm),
  pbehavior: formToPbehaviorDataStorageSettings(form.pbehavior),
});
