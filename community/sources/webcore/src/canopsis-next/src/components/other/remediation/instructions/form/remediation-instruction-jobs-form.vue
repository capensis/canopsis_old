<template>
  <v-layout
    class="mt-2"
    column
  >
    <c-alert
      v-if="!jobs.length"
      type="info"
    >
      {{ $t('remediation.instruction.emptyJobs') }}
    </c-alert>
    <c-information-block :title="$t('remediation.instruction.listJobs')">
      <c-draggable-list-field
        v-field="jobs"
        :disabled="disabled"
        :group="draggableGroup"
        handle=".job-drag-handler"
        ghost-class="white"
      >
        <v-card
          v-for="(job, index) in jobs"
          :key="job.key"
          class="my-2"
        >
          <v-card-text>
            <remediation-instruction-job-field
              v-field="jobs[index]"
              :jobs="jobsItems"
              :name="job.key"
              :job-number="index + 1"
              :disabled="disabled"
              class="py-1"
              @remove="removeItemFromArray(index)"
            />
          </v-card-text>
        </v-card>
      </c-draggable-list-field>

      <c-btn-with-error
        :error="hasJobsErrors ? $t('remediation.instruction.errors.jobRequired') : ''"
        :disabled="disabled"
        outlined
        @click="addJob"
      >
        {{ $t('remediation.instruction.addJob') }}
      </c-btn-with-error>
    </c-information-block>
  </v-layout>
</template>

<script>
import { MAX_LIMIT } from '@/constants';

import { remediationInstructionJobToForm } from '@/helpers/entities/remediation/instruction/form';

import { formArrayMixin } from '@/mixins/form';
import { entitiesRemediationJobMixin } from '@/mixins/entities/remediation/job';

import RemediationInstructionJobField from './fields/remediation-instruction-job-field.vue';

export default {
  inject: ['$validator'],
  components: {
    RemediationInstructionJobField,
  },
  mixins: [formArrayMixin, entitiesRemediationJobMixin],
  model: {
    prop: 'jobs',
    event: 'input',
  },
  props: {
    jobs: {
      type: Array,
      default: () => ([]),
    },
    name: {
      type: String,
      default: 'jobs',
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {
      isDragging: false,
      jobsItems: [],
    };
  },
  computed: {
    hasJobsErrors() {
      return this.errors.has(this.name);
    },

    draggableGroup() {
      return {
        name: 'remediation-instruction-jobs',
        pull: false,
        put: false,
      };
    },
  },
  watch: {
    jobs() {
      this.$validator.validate(this.name);
    },
  },
  mounted() {
    this.fetchList();
  },
  created() {
    this.$validator.attach({
      name: this.name,
      rules: 'min_value:1',
      getter: () => this.jobs.length,
      vm: this,
    });
  },
  beforeDestroy() {
    this.$validator.detach(this.name);
  },
  methods: {
    addJob() {
      this.addItemIntoArray(remediationInstructionJobToForm());
    },

    async fetchList() {
      const { data: jobs } = await this.fetchRemediationJobsListWithoutStore({
        params: {
          limit: MAX_LIMIT,
        },
      });

      this.jobsItems = jobs;
    },
  },
};
</script>
