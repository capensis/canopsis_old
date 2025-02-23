<template>
  <div>
    <v-layout
      v-if="serviceEntities.length"
      class="d-inline-flex my-1"
      align-center
    >
      <v-simple-checkbox
        v-model="isAllSelected"
        :disabled="!entitiesWithActions.length"
        class="ml-4 my-2"
      />
      <v-fade-transition mode="out-in">
        <span
          v-if="!selectedEntities.length"
          class="font-italic"
        >
          {{ $t('modals.service.massActionsDescription') }}
        </span>
        <service-entity-actions
          v-else
          :actions="actions"
          :entity="service"
          @apply="applyActionForSelected"
        />
      </v-fade-transition>
    </v-layout>
    <div
      v-for="serviceEntity in serviceEntities"
      :key="serviceEntity.key"
      class="mt-2"
    >
      <service-entity
        :service-id="service._id"
        :entity="serviceEntity"
        :last-action-unavailable="hasActionError(serviceEntity)"
        :entity-name-field="entityNameField"
        :widget-parameters="widgetParameters"
        :selected="isEntitySelected(serviceEntity)"
        :actions-requests="actionsRequests"
        @add:action="addAction"
        @update:selected="updateSelected(serviceEntity, $event)"
        @remove:unavailable="removeEntityFromUnavailable(serviceEntity)"
        @refresh="$listeners.refresh"
      />
    </div>
    <c-table-pagination
      :total-items="totalItems"
      :items-per-page="options.itemsPerPage"
      :page="options.page"
      class="mt-1"
      @update:page="updatePage"
      @update:items-per-page="updateItemsPerPage"
      @input="updatePaginationOptions"
    />
  </div>
</template>

<script>
import {
  getAvailableActionsByEntities,
  getAvailableEntityActionsTypes,
  isActionTypeAvailableForEntity,
} from '@/helpers/entities/entity/actions';
import { filterById, mapIds } from '@/helpers/array';

import { widgetActionPanelServiceEntityMixin } from '@/mixins/widget/actions-panel/service-entity';

import ServiceEntityActions from './service-entity-actions.vue';
import ServiceEntity from './service-entity.vue';

export default {
  components: {
    ServiceEntityActions,
    ServiceEntity,
  },
  mixins: [widgetActionPanelServiceEntityMixin],
  props: {
    service: {
      type: Object,
      required: true,
    },
    options: {
      type: Object,
      required: true,
    },
    serviceEntities: {
      type: Array,
      default: () => [],
    },
    entityNameField: {
      type: String,
      default: 'name',
    },
    widgetParameters: {
      type: Object,
      default: () => ({}),
    },
    totalItems: {
      type: Number,
      required: false,
    },
    actionsRequests: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      selectedEntities: [],
    };
  },
  computed: {
    entitiesWithActions() {
      return this.serviceEntities.filter(entity => getAvailableEntityActionsTypes(entity).length);
    },

    isAllSelected: {
      get() {
        return this.entitiesWithActions.length > 0
          && this.entitiesWithActions.every(({ _id: id }) => this.selectedEntitiesIds.includes(id));
      },
      set(checked) {
        if (checked) {
          this.selectedEntities = [...this.entitiesWithActions];
        } else {
          this.selectedEntities = [];
        }
      },
    },

    actions() {
      return getAvailableActionsByEntities(this.selectedEntities)
        .filter(this.actionsAccessFilterHandler)
        .map(action => ({
          ...action,
          disabled: !this.hasEntityWithoutAction(action.type),
        }));
    },

    selectedEntitiesIds() {
      return mapIds(this.selectedEntities);
    },
  },
  watch: {
    serviceEntities() {
      this.selectedEntities = [];
    },
  },
  methods: {
    hasActionError(entity) {
      return this.unavailableEntitiesAction[entity._id];
    },

    hasEntityWithoutAction(type) {
      return this.selectedEntities.filter(entity => isActionTypeAvailableForEntity(type, entity));
    },

    applyActionForSelected({ type }) {
      this.applyEntityAction(type, this.selectedEntities);
    },

    updateSelected(entity, checked) {
      if (checked) {
        this.selectedEntities.push(entity);
      } else {
        this.selectedEntities = filterById(this.selectedEntities, entity);
      }
    },

    isEntitySelected(entity) {
      return this.selectedEntitiesIds.includes(entity._id);
    },

    addAction(action) {
      this.$emit('add:action', action);
    },

    updatePage(page) {
      this.$emit('update:options', { ...this.options, page });
    },

    updateItemsPerPage(itemsPerPage) {
      this.$emit('update:options', { ...this.options, itemsPerPage });
    },

    updatePaginationOptions({ page, itemsPerPage }) {
      this.$emit('update:options', { ...this.options, page, itemsPerPage });
    },
  },
};
</script>
