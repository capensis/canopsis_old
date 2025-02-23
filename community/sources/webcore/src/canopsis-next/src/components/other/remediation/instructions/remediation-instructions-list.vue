<template>
  <c-advanced-data-table
    :headers="headers"
    :items="remediationInstructions"
    :loading="pending"
    :total-items="totalItems"
    :options="options"
    :select-all="removable"
    search
    advanced-pagination
    @update:options="$emit('update:options', $event)"
  >
    <template #mass-actions="{ selected, selectedKeys }">
      <c-action-btn
        v-if="removable"
        type="delete"
        @click="$emit('remove-selected', selected)"
      />
      <c-db-export-btn :ids="selectedKeys" instruction />
    </template>
    <template #headerCell="{ header }">
      <span class="c-table-header__text--multiline">{{ header.text }}</span>
    </template>
    <template #enabled="{ item }">
      <c-enabled :value="item.enabled" />
    </template>
    <template #status="{ item }">
      <remediation-instructions-approval-icon :instruction="item" />
    </template>
    <template #type="{ item }">
      {{ $t(`remediation.instruction.types.${item.type}`) }}
    </template>
    <template #last_modified="{ item }">
      {{ item.last_modified | date }}
    </template>
    <template #last_executed_on="{ item }">
      {{ item.last_executed_on | date }}
    </template>
    <template #actions="{ item }">
      <v-layout justify-end>
        <c-action-btn
          v-if="isInstructionDismissed(item)"
          :tooltip="$t('remediation.instruction.approvalDismissed')"
          icon="warning"
          color="error"
          @click="$emit('edit', item)"
        />
        <c-action-btn
          v-if="item.approval && isNeedApproveByCurrentUser(item)"
          :tooltip="$t('remediation.instruction.needApprove')"
          icon="notification_important"
          color="error"
          @click="$emit('approve', item)"
        />
        <c-action-btn
          v-if="updatable"
          type="edit"
          @click="$emit('edit', item)"
        />
        <c-action-btn
          v-if="duplicable"
          type="duplicate"
          @click="$emit('duplicate', item)"
        />
        <c-action-btn
          v-if="removable"
          type="delete"
          @click="$emit('remove', item)"
        />
        <c-db-export-btn :id="item._id" instruction />
      </v-layout>
    </template>
  </c-advanced-data-table>
</template>

<script>
import { isUserHasRole } from '@/helpers/entities/user/entity';
import { isApproveRequested, isInstructionDismissed } from '@/helpers/entities/remediation/instruction/form';

import { authMixin } from '@/mixins/auth';

import RemediationInstructionsApprovalIcon
  from '@/components/other/remediation/instructions/partials/remediation-instructions-approval-icon.vue';

export default {
  components: { RemediationInstructionsApprovalIcon },
  mixins: [authMixin],
  props: {
    remediationInstructions: {
      type: Array,
      required: true,
    },
    pending: {
      type: Boolean,
      default: false,
    },
    totalItems: {
      type: Number,
      required: false,
    },
    options: {
      type: Object,
      required: true,
    },
    removable: {
      type: Boolean,
      default: false,
    },
    updatable: {
      type: Boolean,
      default: false,
    },
    duplicable: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    headers() {
      return [
        {
          text: this.$t('common.name'),
          value: 'name',
        },
        {
          text: this.$t('common.author'),
          value: 'author.display_name',
        },
        {
          text: this.$t('common.enabled'),
          value: 'enabled',
        },
        {
          text: this.$t('common.type'),
          value: 'type',
        },
        {
          text: this.$t('common.lastModifiedOn'),
          value: 'last_modified',
        },
        {
          text: this.$t('common.priority'),
          value: 'priority',
        },
        {
          text: this.$t('common.status'),
          value: 'status',
        },
        {
          text: this.$t('remediation.instruction.table.monthExecutions'),
          value: 'month_executions',
          sortable: false,
        },
        {
          text: this.$t('remediation.instruction.table.lastExecutedOn'),
          value: 'last_executed_on',
        },
        {
          text: this.$t('common.actionsLabel'),
          value: 'actions',
          sortable: false,
        },
      ];
    },
  },
  methods: {
    isInstructionDismissed(remediationInstruction) {
      return isInstructionDismissed(remediationInstruction);
    },

    isNeedApproveByCurrentUser(remediationInstruction) {
      return isApproveRequested(remediationInstruction)
        && (remediationInstruction.approval.user?._id === this.currentUser._id
        || isUserHasRole(this.currentUser, remediationInstruction.approval.role));
    },
  },
};
</script>
