// https://nightwatchjs.org/guide/#working-with-page-objects

const el = require('../../helpers/el');

const commands = {
  clickSubmitAlarms() {
    return this.customClick('@submitAlarms');
  },

  setEnableHtml(checked = false) {
    return this.getAttribute('@enableHtmlInput', 'aria-checked', ({ value }) => {
      if (value !== String(checked)) {
        this.customClick('@enableHtml');
      }
    });
  },

  clickAckGroup() {
    return this.customClick('@ackGroup');
  },

  setIsAckNoteRequired(checked = false) {
    return this.getAttribute('@isAckNoteRequiredInput', 'aria-checked', ({ value }) => {
      if (value !== String(checked)) {
        this.customClick('@isAckNoteRequired');
      }
    });
  },

  setIsMultiAckEnabled(checked = false) {
    return this.getAttribute('@isMultiAckEnabledInput', 'aria-checked', ({ value }) => {
      if (value !== String(checked)) {
        this.customClick('@isMultiAckEnabled');
      }
    });
  },

  clickFastAckOutput() {
    return this.customClick('@fastAckOutput');
  },

  setFastAckOutputSwitch(checked = false) {
    return this.getAttribute('@fastAckOutputSwitchInput', 'aria-checked', ({ value }) => {
      if (value !== String(checked)) {
        this.customClick('@fastAckOutputSwitch');
      }
    });
  },

  clickFastAckOutputText() {
    return this.customClick('@fastAckOutputField');
  },

  clearFastAckOutputText() {
    return this.customClearValue('@fastAckOutputField');
  },

  setFastAckOutputText(value) {
    return this.customSetValue('@fastAckOutputField', value);
  },

  clickCreateLiveReporting() {
    return this.customClick('@liveReportingCreateButton');
  },

  clickEditLiveReporting() {
    return this.customClick('@liveReportingEditButton');
  },

  clickDeleteLiveReporting() {
    return this.customClick('@liveReportingDeleteButton');
  },

  el,
};

module.exports = {
  elements: {
    submitAlarms: sel('submitAlarms'),

    optionSelect: '.menuable__content__active .v-select-list [role="listitem"]:nth-of-type(%s)',

    enableHtml: `${sel('isHtmlEnabledOnTimeLine')} div${sel('switcherField')}`,
    enableHtmlInput: `${sel('isHtmlEnabledOnTimeLine')} input${sel('switcherField')}`,

    isAckNoteRequired: `${sel('isAckNoteRequired')} div${sel('switcherField')}`,
    isAckNoteRequiredInput: `${sel('isAckNoteRequired')} input${sel('switcherField')}`,
    isMultiAckEnabled: `${sel('isMultiAckEnabled')} div${sel('switcherField')}`,
    isMultiAckEnabledInput: `${sel('isMultiAckEnabled')} input${sel('switcherField')}`,
    ackGroup: sel('ackGroup'),
    fastAckOutput: sel('fastAckOutput'),
    fastAckOutputSwitch: `div${sel('fastAckOutputSwitch')} .v-input--selection-controls__ripple`,
    fastAckOutputSwitchInput: `input${sel('fastAckOutputSwitch')}`,
    fastAckOutputField: sel('fastAckOutputField'),

    liveReportingCreateButton: `${sel('liveReporting')} + div > ${sel('createButton')}`,
    liveReportingEditButton: `${sel('liveReporting')} + div > ${sel('editButton')}`,
    liveReportingDeleteButton: `${sel('liveReporting')} + div > ${sel('deleteButton')}`,
  },
  commands: [commands],
};
