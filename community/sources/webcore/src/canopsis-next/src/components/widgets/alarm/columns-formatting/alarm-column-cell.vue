<template>
  <v-menu
    v-if="column.popupTemplate"
    v-model="opened"
    :close-on-content-click="false"
    :open-on-click="false"
    class="alarms-column-cell"
    offset-x
    disable-resize
  >
    <template #activator="{ on }">
      <v-layout
        class="alarms-column-cell__layout"
        d-inline-flex
        align-center
        v-on="on"
      >
        <div
          v-if="column.isHtml"
          v-html="sanitizedValue"
        />
        <div
          v-else
          v-bind="component.bind"
          :is="component.is"
          v-on="component.on"
        />
        <v-btn
          :class="{ 'alarms-column-cell__show-info-btn--small': small }"
          class="ma-0 alarms-column-cell__show-info-btn"
          icon
          small
          @click.stop="showInfoPopup"
        >
          <v-icon small>
            info
          </v-icon>
        </v-btn>
      </v-layout>
    </template>
    <alarm-column-cell-popup-body
      :alarm="alarm"
      :selected-tags="selectedTags"
      :template="column.popupTemplate"
      :template-id="popupTemplateId"
      @close="hideInfoPopup"
      @select:tag="$emit('select:tag', $event)"
      @remove:tag="$emit('remove:tag', $event)"
    />
  </v-menu>
  <div
    v-else-if="column.isHtml"
    v-html="sanitizedValue"
  />
  <div
    v-else
    v-bind="component.bind"
    :is="component.is"
    v-on="component.on"
  />
</template>

<script>
import { get } from 'lodash';

import { sanitizeHtml, linkifyHtml } from '@/helpers/html';
import { getAlarmWidgetColumnPopupTemplateId } from '@/helpers/entities/alarm/list';

import ColorIndicatorWrapper from '@/components/common/table/color-indicator-wrapper.vue';
import AlarmsExpandPanel from '@/components/widgets/alarm/expand-panel/alarms-expand-panel.vue';

import AlarmColumnCellPopupBody from './alarm-column-cell-popup-body.vue';
import AlarmColumnValueStatus from './alarm-column-value-status.vue';
import AlarmColumnValueExtraDetails from './alarm-column-value-extra-details.vue';

/**
 * Component to format alarms list columns
 *
 * @module alarm
 *
 * @prop {Object} alarm - Object representing the alarm
 * @prop {Object} widget - Object representing the widget
 * @prop {Object} column - Property concerned on the column
 */
export default {
  components: {
    AlarmsExpandPanel,
    AlarmColumnCellPopupBody,
    AlarmColumnValueStatus,
    AlarmColumnValueExtraDetails,
    ColorIndicatorWrapper,
  },
  props: {
    alarm: {
      type: Object,
      required: true,
    },
    widget: {
      type: Object,
      required: true,
    },
    column: {
      type: Object,
      required: true,
    },
    selectedTags: {
      type: Array,
      default: () => [],
    },
    small: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      opened: false,
    };
  },
  computed: {
    value() {
      const value = get(this.alarm, this.column.value, '');

      return this.column.filter ? this.column.filter(value) : value;
    },

    sanitizedValue() {
      return sanitizeHtml(linkifyHtml(String(this.value ?? '')));
    },

    component() {
      return this.column.getComponent(this);
    },

    popupTemplateId() {
      return getAlarmWidgetColumnPopupTemplateId(this.widget._id, this.column.value);
    },
  },
  methods: {
    showInfoPopup() {
      this.opened = true;
    },

    hideInfoPopup() {
      this.opened = false;
    },
  },
};
</script>

<style lang="scss">
.alarms-column-cell {
  &__show-info-btn {
    flex-shrink: 0 !important;

    &--small {
      width: 22px;
      height: 22px;
      max-width: 22px;
      max-height: 22px;
    }
  }

  &__layout {
    max-width: 100%;
  }
}
</style>
