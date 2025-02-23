<template>
  <widget-settings
    :submitting="submitting"
    :dirty="hasChanges"
    divider
    @submit="submit"
  >
    <field-title v-model="form.title" />
    <field-periodic-refresh
      v-model="form.parameters"
      with-live-watching
    />
    <widget-settings-group :title="$t('settings.liveReporting.title')">
      <field-live-reporting v-model="form.parameters.liveReporting" />
    </widget-settings-group>
    <widget-settings-group :title="$tc('common.column', 2)">
      <field-default-sort-column
        v-model="form.parameters.sort"
        :columns="sortablePreparedWidgetColumns"
        :columns-label="$t('settings.columnName')"
      />
      <field-columns
        v-model="form.parameters.widgetColumns"
        :template="form.parameters.widgetColumnsTemplate"
        :templates="alarmColumnsWidgetTemplates"
        :templates-pending="widgetTemplatesPending"
        :label="$t('settings.columnNames')"
        :type="$constants.ENTITIES_TYPES.alarm"
        with-template
        with-html
        with-color-indicator
        @update:template="updateWidgetColumnsTemplate"
      />
      <field-columns
        v-model="form.parameters.widgetGroupColumns"
        :template="form.parameters.widgetGroupColumnsTemplate"
        :templates="alarmColumnsWidgetTemplates"
        :templates-pending="widgetTemplatesPending"
        :label="$t('settings.groupColumnNames')"
        :type="$constants.ENTITIES_TYPES.alarm"
        with-template
        with-html
        with-color-indicator
        @update:template="updateWidgetGroupColumnsTemplate"
      />
      <field-columns
        v-model="form.parameters.serviceDependenciesColumns"
        :template="form.parameters.serviceDependenciesColumnsTemplate"
        :templates="entityColumnsWidgetTemplates"
        :templates-pending="widgetTemplatesPending"
        :label="$t('settings.trackColumnNames')"
        :type="$constants.ENTITIES_TYPES.entity"
        with-color-indicator
        @update:template="updateServiceDependenciesColumnsTemplate"
      />
      <field-resize-column-behavior v-model="form.parameters.columns" />
      <field-root-cause-settings v-model="form.parameters" />
      <field-info-popup
        v-model="form.parameters.infoPopups"
        :columns="preparedWidgetColumns"
      />
    </widget-settings-group>

    <widget-settings-group :title="$tc('common.filter', 2)">
      <field-filters
        v-if="hasAccessToFilter"
        v-model="form.parameters.mainFilter"
        :filters.sync="form.filters"
        :widget-id="widget._id"
        addable
        editable
        with-alarm
        with-entity
        with-pbehavior
      />
      <field-opened-resolved-filter v-model="form.parameters.opened" />
      <field-remediation-instructions-filters
        v-if="hasAccessToRemediationInstructionsFilter"
        v-model="form.parameters.remediationInstructionsFilters"
        addable
        editable
      />
      <field-switcher
        v-model="form.parameters.isCorrelationEnabled"
        :title="$t('common.correlation')"
      />
      <field-switcher
        v-model="form.parameters.clearFilterDisabled"
        :title="$t('settings.clearFilterDisabled')"
      />
    </widget-settings-group>

    <widget-settings-group :title="$tc('common.view', 1)">
      <field-default-elements-per-page v-model="form.parameters.itemsPerPage" />
      <field-density v-model="form.parameters.dense" />
      <widget-settings-group :title="$t('settings.kioskMode')">
        <field-switcher
          v-model="form.parameters.kiosk.hideActions"
          :title="$t('settings.kiosk.hideActions')"
        />
        <field-switcher
          v-model="form.parameters.kiosk.hideMassSelection"
          :title="$t('settings.kiosk.hideMassSelection')"
        />
        <field-switcher
          v-model="form.parameters.kiosk.hideToolbar"
          :title="$t('settings.kiosk.hideToolbar')"
        />
      </widget-settings-group>
      <field-switcher
        v-model="form.parameters.sticky_header"
        :title="$t('settings.stickyHeader')"
      />
      <field-switcher
        v-model="form.parameters.sticky_horizontal_scroll"
        :title="$t('settings.stickyHorizontalScroll')"
      />
      <field-switcher
        v-model="form.parameters.isVirtualScrollEnabled"
        :title="$t('settings.isVirtualScrollEnabled')"
      />
    </widget-settings-group>

    <widget-settings-group :title="$t('common.actionsLabel')">
      <widget-settings-group :title="$t('common.ack')">
        <field-switcher
          v-model="form.parameters.isAckNoteRequired"
          :title="$t('settings.isAckNoteRequired')"
        />
        <field-switcher
          v-model="form.parameters.isMultiAckEnabled"
          :title="$t('settings.isMultiAckEnabled')"
        />
        <field-fast-action-output
          v-model="form.parameters.fastAckOutput"
          :label="$t('settings.fastAckOutput')"
        />
      </widget-settings-group>

      <widget-settings-group :title="$t('common.cancel')">
        <field-fast-action-output
          v-model="form.parameters.fastCancelOutput"
          :label="$t('settings.fastCancelOutput')"
        />
        <field-switcher
          v-model="form.parameters.isUncancelAlarmsCommentRequired"
          :title="$t('settings.isUncancelAlarmsCommentRequired')"
        />
      </widget-settings-group>

      <fast-pbehavior-form v-model="form.parameters" />

      <field-switcher
        v-model="form.parameters.isSnoozeNoteRequired"
        :title="$t('settings.isSnoozeNoteRequired')"
      />
      <field-switcher
        v-model="form.parameters.isRemoveAlarmsFromMetaAlarmCommentRequired"
        :title="$t('settings.isRemoveAlarmsFromMetaAlarmCommentRequired')"
      />
      <field-switcher
        v-model="form.parameters.isMultiDeclareTicketEnabled"
        :title="$t('settings.isMultiDeclareTicketEnabled')"
      />
      <field-switcher
        v-model="form.parameters.isActionsAllowWithOkState"
        :title="$t('settings.isActionsAllowWithOkState')"
      />
      <field-text-editor-with-template
        :value="form.parameters.exportPdfTemplate"
        :template="form.parameters.exportPdfTemplateTemplate"
        :title="$t('settings.exportPdfTemplate')"
        :variables="exportPdfAlarmVariables"
        :default-value="defaultExportPdfTemplateValue"
        :dialog-props="{ maxWidth: 1070 }"
        :templates="alarmExportToPdfWidgetTemplates"
        addable
        removable
        @input="updateExportPdf"
      />
    </widget-settings-group>

    <widget-settings-group :title="$t('settings.expandPanel.title')">
      <widget-settings-group :title="$t('alarm.tabs.moreInfos')">
        <field-text-editor-with-template
          :value="form.parameters.moreInfoTemplate"
          :template="form.parameters.moreInfoTemplateTemplate"
          :title="$t('common.infos')"
          :variables="alarmVariables"
          :templates="alarmMoreInfosWidgetTemplates"
          addable
          removable
          @input="updateMoreInfo"
        />
        <field-grid-range-size
          v-model="form.parameters.expandGridRangeSize"
          :title="$t('common.width')"
          inline
        />
      </widget-settings-group>
      <field-availability-graph-settings v-model="form.parameters.availability" />
      <field-switcher
        v-model="form.parameters.isHtmlEnabledOnTimeLine"
        :title="$t('settings.isHtmlEnabledOnTimeLine')"
      />
    </widget-settings-group>
    <export-csv-form
      v-model="form.parameters"
      :type="$constants.ENTITIES_TYPES.alarm"
      :templates="alarmColumnsWidgetTemplates"
      :templates-pending="widgetTemplatesPending"
      :variables="columnsVariables"
      datetime-format
      with-instructions
      with-simple-template
      optional-infos-attributes
    />
    <charts-form v-model="form.parameters.charts" />
  </widget-settings>
</template>

<script>
import { SIDE_BARS, ALARM_UNSORTABLE_FIELDS, ALARM_FIELDS_TO_LABELS_KEYS, ALARM_PAYLOADS_VARIABLES } from '@/constants';

import { formToWidgetColumns } from '@/helpers/entities/widget/column/form';
import { getWidgetColumnLabel, getWidgetColumnSortable } from '@/helpers/entities/widget/list';

import { widgetSettingsMixin } from '@/mixins/widget/settings';
import { entitiesInfosMixin } from '@/mixins/entities/infos';
import { alarmVariablesMixin } from '@/mixins/widget/variables';
import { widgetTemplatesMixin } from '@/mixins/widget/templates';
import { permissionsWidgetsAlarmsListFilters } from '@/mixins/permissions/widgets/alarms-list/filters';
import {
  permissionsWidgetsAlarmsListRemediationInstructionsFilters,
} from '@/mixins/permissions/widgets/alarms-list/remediation-instructions-filters';
import { payloadVariablesMixin } from '@/mixins/payload/variables';

import ALARM_EXPORT_PDF_TEMPLATE from '@/assets/templates/alarm-export-pdf.html';

import FieldAvailabilityGraphSettings from '../availability/form/fields/availability-graph-settings.vue';
import WidgetSettingsGroup from '../partials/widget-settings-group.vue';
import WidgetSettings from '../partials/widget-settings.vue';
import FieldTitle from '../form/fields/title.vue';
import FieldDefaultSortColumn from '../form/fields/default-sort-column.vue';
import FieldColumns from '../form/fields/columns.vue';
import FieldPeriodicRefresh from '../form/fields/periodic-refresh.vue';
import FieldDefaultElementsPerPage from '../form/fields/default-elements-per-page.vue';
import FieldFilters from '../form/fields/filters.vue';
import FieldTextEditorWithTemplate from '../form/fields/text-editor-with-template.vue';
import FieldSwitcher from '../form/fields/switcher.vue';
import FieldGridRangeSize from '../form/fields/grid-range-size.vue';
import ExportCsvForm from '../form/export-csv.vue';
import ChartsForm from '../chart/form/charts-form.vue';
import FieldRootCauseSettings from '../form/fields/root-cause-settings.vue';

import FieldRemediationInstructionsFilters from './form/fields/remediation-instructions-filters.vue';
import FieldDensity from './form/fields/density.vue';
import FieldLiveReporting from './form/fields/live-reporting.vue';
import FieldFastActionOutput from './form/fields/fast-action-output.vue';
import FieldOpenedResolvedFilter from './form/fields/opened-resolved-filter.vue';
import FieldInfoPopup from './form/fields/info-popup.vue';
import FieldResizeColumnBehavior from './form/fields/resize-column-behavior.vue';
import FastPbehaviorForm from './form/fast-pbehavior-form.vue';

/**
 * Component to regroup the alarms list settings fields
 */
export default {
  name: SIDE_BARS.alarmSettings,
  components: {
    FieldAvailabilityGraphSettings,
    FieldRootCauseSettings,
    WidgetSettingsGroup,
    WidgetSettings,
    FieldTitle,
    FieldDefaultSortColumn,
    FieldColumns,
    FieldLiveReporting,
    FieldPeriodicRefresh,
    FieldDefaultElementsPerPage,
    FieldOpenedResolvedFilter,
    FieldFilters,
    FieldTextEditorWithTemplate,
    FieldSwitcher,
    FieldFastActionOutput,
    FieldGridRangeSize,
    FieldRemediationInstructionsFilters,
    FieldInfoPopup,
    FieldDensity,
    FastPbehaviorForm,
    ExportCsvForm,
    ChartsForm,
    FieldResizeColumnBehavior,
  },
  mixins: [
    payloadVariablesMixin,
    widgetSettingsMixin,
    entitiesInfosMixin,
    alarmVariablesMixin,
    widgetTemplatesMixin,
    permissionsWidgetsAlarmsListFilters,
    permissionsWidgetsAlarmsListRemediationInstructionsFilters,
  ],
  computed: {
    preparedWidgetColumns() {
      return formToWidgetColumns(this.form.parameters.widgetColumns).map(column => ({
        ...column,

        text: getWidgetColumnLabel(column, ALARM_FIELDS_TO_LABELS_KEYS),
      }));
    },

    sortablePreparedWidgetColumns() {
      return this.preparedWidgetColumns.filter(column => getWidgetColumnSortable(column, ALARM_UNSORTABLE_FIELDS));
    },

    defaultExportPdfTemplateValue() {
      return ALARM_EXPORT_PDF_TEMPLATE;
    },

    columnsVariables() {
      return [
        ...this.alarmPayloadVariables,
        {
          value: ALARM_PAYLOADS_VARIABLES.infosValue,
          text: this.$t('alarm.fields.alarmInfos'),
        },
      ];
    },
  },
  mounted() {
    this.fetchInfos();
  },
  methods: {
    updateWidgetColumnsTemplate(template, columns) {
      this.$set(this.form.parameters, 'widgetColumnsTemplate', template);
      this.$set(this.form.parameters, 'widgetColumns', columns);
    },

    updateWidgetGroupColumnsTemplate(template, columns) {
      this.$set(this.form.parameters, 'widgetGroupColumnsTemplate', template);
      this.$set(this.form.parameters, 'widgetGroupColumns', columns);
    },

    updateServiceDependenciesColumnsTemplate(template, columns) {
      this.$set(this.form.parameters, 'serviceDependenciesColumnsTemplate', template);
      this.$set(this.form.parameters, 'serviceDependenciesColumns', columns);
    },

    updateMoreInfo(content, template) {
      this.$set(this.form.parameters, 'moreInfoTemplate', content);

      if (template && template !== this.form.parameters.moreInfoTemplateTemplate) {
        this.$set(this.form.parameters, 'moreInfoTemplateTemplate', template);
      }
    },

    updateExportPdf(content, template) {
      this.$set(this.form.parameters, 'exportPdfTemplate', content);

      if (template && template !== this.form.parameters.exportPdfTemplateTemplate) {
        this.$set(this.form.parameters, 'exportPdfTemplateTemplate', template);
      }
    },
  },
};
</script>
