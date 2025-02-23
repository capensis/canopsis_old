<template>
  <tr
    :class="classes"
    :data-id="alarm._id"
    class="alarm-list-row"
    v-on="listeners"
  >
    <td v-if="!preparedVisible" :style="emptyTdStyle" :colspan="availableHeaders.length + Number(hasRowActions)" />
    <template v-if="localBooted">
      <td
        v-if="hasRowActions"
        v-show="preparedVisible"
        class="alarm-list-row__icons pr-0"
      >
        <v-layout
          align-center
          justify-space-between
        >
          <v-layout class="alarm-list-row__checkbox">
            <template v-if="selectable">
              <v-simple-checkbox
                v-if="isAlarmSelectable"
                v-field="selected"
                class="ma-0"
                color="primary"
                hide-details
              />
              <v-simple-checkbox
                v-else
                disabled
                hide-details
              />
            </template>
          </v-layout>
          <v-layout
            v-if="hasAlarmInstruction"
            align-center
          >
            <alarms-list-row-instructions-icon :alarm="alarm" />
          </v-layout>
          <v-layout
            v-if="hasBookmark"
            align-center
          >
            <alarms-list-row-bookmark-icon />
          </v-layout>
          <alarms-expand-panel-btn
            v-if="expandable"
            :expanded="expanded"
            :alarm="alarm"
            :widget="widget"
            :small="small"
            :search="search"
            @input="$emit('expand', $event)"
          />
        </v-layout>
      </td>
      <td
        v-show="preparedVisible"
        v-for="header in availableHeaders"
        :key="header.value"
        class="alarm-list-row__cell"
      >
        <c-booted-placeholder-loader
          v-if="header.value === 'actions'"
          :eager="eager"
          class="layout align-center"
          async-booting-provider="$asyncBootingActionsPanel"
        >
          <actions-panel
            :item="alarm"
            :widget="widget"
            :parent-alarm="parentAlarm"
            :refresh-alarms-list="refreshAlarmsList"
            :small="small"
            :ignore-media-query="actionsIgnoreMediaQuery"
            :inline-count="actionsInlineCount"
          />
        </c-booted-placeholder-loader>
        <alarm-column-value
          v-else
          :alarm="alarm"
          :widget="widget"
          :column="header"
          :selected-tags="selectedTags"
          :small="small"
          :parent-alarm="parentAlarm"
          @activate="activateRow"
          @select:tag="$emit('select:tag', $event)"
          @remove:tag="$emit('remove:tag', $event)"
          @click:state="$emit('click:state', $event)"
        />
        <span
          v-if="resizing"
          class="alarms-list-table__resize-handler"
          @mousedown.prevent="$emit('start:resize', header.value)"
          @click.stop=""
        />
      </td>
    </template>
  </tr>
</template>

<script>
import { flow, isNumber } from 'lodash';

import featuresService from '@/services/features';

import { isActionAvailableForAlarm } from '@/helpers/entities/alarm/form';

import { formBaseMixin } from '@/mixins/form';

import ActionsPanel from '../actions/actions-panel.vue';
import AlarmColumnValue from '../columns-formatting/alarm-column-value.vue';
import AlarmsExpandPanelBtn from '../expand-panel/alarms-expand-panel-btn.vue';

import AlarmsListRowInstructionsIcon from './alarms-list-row-instructions-icon.vue';
import AlarmsListRowBookmarkIcon from './alarms-list-row-bookmark-icon.vue';

export default {
  inject: ['$system', '$intersectionObserver'],
  components: {
    ActionsPanel,
    AlarmColumnValue,
    AlarmsExpandPanelBtn,
    AlarmsListRowInstructionsIcon,
    AlarmsListRowBookmarkIcon,
  },
  mixins: [formBaseMixin],
  model: {
    prop: 'selected',
    event: 'input',
  },
  props: {
    selected: {
      type: Boolean,
      default: false,
    },
    selectable: {
      type: Boolean,
      default: false,
    },
    expandable: {
      type: Boolean,
      default: false,
    },
    alarm: {
      type: Object,
      required: true,
    },
    widget: {
      type: Object,
      required: true,
    },
    headers: {
      type: Array,
      required: true,
    },
    columnsFilters: {
      type: Array,
      default: () => [],
    },
    parentAlarm: {
      type: Object,
      default: null,
    },
    refreshAlarmsList: {
      type: Function,
      default: () => {},
    },
    selecting: {
      type: Boolean,
      default: false,
    },
    selectedTags: {
      type: Array,
      default: () => [],
    },
    medium: {
      type: Boolean,
      default: false,
    },
    small: {
      type: Boolean,
      default: false,
    },
    resizing: {
      type: Boolean,
      default: false,
    },
    expanded: {
      type: Boolean,
      default: false,
    },
    showInstructionIcon: {
      type: Boolean,
      default: false,
    },
    search: {
      type: String,
      default: '',
    },
    actionsInlineCount: {
      type: Number,
      required: false,
    },
    actionsIgnoreMediaQuery: {
      type: Boolean,
      default: false,
    },
    booted: {
      type: Boolean,
      default: false,
    },
    virtualScroll: {
      type: Boolean,
      default: false,
    },
    eager: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    const visible = this.booted || this.eager;

    return {
      active: false,
      localBooted: visible,
      visible,
      height: null,
    };
  },
  computed: {
    emptyTdStyle() {
      if (!this.height) {
        return {};
      }

      return {
        height: `${this.height}px`,
      };
    },

    preparedVisible() {
      return (this.virtualScroll && this.visible) || (!this.virtualScroll && this.localBooted);
    },

    hasBookmark() {
      return !!this.alarm.bookmark;
    },

    hasRowActions() {
      return this.selectable || this.expandable || this.showInstructionIcon || this.hasBookmark;
    },

    hasAlarmInstruction() {
      const { children_instructions: parentAlarmChildrenInstructions = false } = this.parentAlarm || {};
      const { assigned_instructions: assignedInstructions = [] } = this.alarm;

      const hasAssignedInstructions = !!assignedInstructions.length;

      if (parentAlarmChildrenInstructions && hasAssignedInstructions) {
        return true;
      }

      return hasAssignedInstructions || isNumber(this.alarm.instruction_execution_icon);
    },

    isAlarmSelectable() {
      return isActionAvailableForAlarm(this.alarm, this.widget);
    },

    isNotFiltered() {
      return this.alarm.filtered === false;
    },

    listeners() {
      let listeners = {};

      if (featuresService.has('components.alarmListRow.computed.listeners')) {
        listeners = featuresService.call('components.alarmListRow.computed.listeners', this, {});
      }

      if (this.selecting) {
        listeners.mousedown = flow([this.mouseSelecting, listeners.mouseenter].filter(Boolean));
      }

      return listeners;
    },

    classes() {
      const classes = { 'alarm-list-row--not-filtered': this.isNotFiltered, 'grey lighten-3': this.active };

      if (featuresService.has('components.alarmListRow.computed.classes')) {
        return featuresService.call('components.alarmListRow.computed.classes', this, classes);
      }

      return classes;
    },

    availableHeaders() {
      return this.headers.filter(({ value }) => value);
    },
  },
  created() {
    this.watchOnceBooted();
  },
  mounted() {
    this.$intersectionObserver?.observe(this.alarm._id, this.$el, this.intersectionHandler);
  },
  beforeDestroy() {
    this.$intersectionObserver?.unobserve(this.alarm._id);
  },
  methods: {
    intersectionHandler(entry) {
      if (entry.isIntersecting !== this.visible) {
        this.visible = entry.isIntersecting;

        if (!this.visible) {
          this.height = this.$el.offsetHeight;
        }
      }
    },

    watchOnceBooted() {
      const unwatch = this.$watch(() => this.booted, (booted) => {
        if (booted) {
          this.localBooted = booted;

          this.$nextTick(() => unwatch());
        }
      }, { immediate: true });
    },

    mouseSelecting(event) {
      if (event.ctrlKey && event.buttons) {
        event.preventDefault();

        this.updateModel(!this.selected);
      }

      return event;
    },

    activateRow(value) {
      this.active = value;
    },
  },
};
</script>

<style lang="scss">
.alarm-list-row {
  min-height: 24px;

  &__checkbox {
    width: 24px;
    max-width: 24px;
    height: 24px;

    .v-input--selection-controls__input {
      margin: 0;
    }
  }

  &__icons {
    width: 100px;
  }

  &__cell {
    position: relative;
  }

  &--not-filtered {
    opacity: .4;
    transition: opacity .3s linear;

    &:hover {
      opacity: 1;
    }
  }
}
</style>
