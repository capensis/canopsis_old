<template>
  <v-sheet class="remediation-instruction-assigned-jobs-expand-panel px-3 py-2">
    <div v-if="isFailedJob">
      {{ $t('remediation.instructionExecute.jobs.failedReason') }}:&nbsp;
      <span
        v-html="job.fail_reason"
        class="pre-wrap"
      />
    </div>
    <div>
      {{ $t('remediation.instructionExecute.jobs.output') }}:&nbsp;
      <span
        v-html="output"
        class="pre-wrap"
      />
    </div>
  </v-sheet>
</template>

<script>
import { createNamespacedHelpers } from 'vuex';

const { mapGetters } = createNamespacedHelpers('remediationJobExecution');

export default {
  props: {
    job: {
      type: Object,
      default: () => ({}),
    },
  },
  computed: {
    ...mapGetters(['getOutputById']),

    output() {
      return this.getOutputById(this.job._id);
    },

    isFailedJob() {
      return !!this.job.fail_reason;
    },
  },
};
</script>

<style lang="scss">
.v-sheet.remediation-instruction-assigned-jobs-expand-panel {
  background: var(--v-background-darken1);
}
</style>
