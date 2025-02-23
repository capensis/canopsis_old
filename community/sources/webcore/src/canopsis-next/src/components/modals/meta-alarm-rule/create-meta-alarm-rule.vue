<template>
  <v-form @submit.prevent="submit">
    <modal-wrapper close>
      <template #title="">
        {{ title }}
      </template>
      <template #text="">
        <meta-alarm-rule-form
          v-model="form"
          ref="formElement"
          :active-step.sync="activeStep"
          :disabled-id-field="config.isDisabledIdField"
          :alarm-infos="alarmInfos"
          :entity-infos="entityInfos"
        />
      </template>
      <template #actions="">
        <v-btn
          depressed
          text
          @click="$modals.hide"
        >
          {{ $t('common.cancel') }}
        </v-btn>
        <v-btn
          v-if="isLastStep"
          key="submit"
          :disabled="isDisabled"
          :loading="submitting"
          class="primary"
          type="submit"
        >
          {{ $t('common.submit') }}
        </v-btn>
        <v-btn
          v-else
          key="next"
          :disabled="!isStepValid"
          type="button"
          class="primary"
          @click="next"
        >
          {{ $t('common.next') }}
        </v-btn>
      </template>
    </modal-wrapper>
  </v-form>
</template>

<script>
import { MODALS, VALIDATION_DELAY, META_ALARMS_FORM_STEPS } from '@/constants';

import { formToMetaAlarmRule, metaAlarmRuleToForm } from '@/helpers/entities/meta-alarm/rule/form';

import { modalInnerMixin } from '@/mixins/modal/inner';
import { submittableMixinCreator } from '@/mixins/submittable';
import { confirmableModalMixinCreator } from '@/mixins/confirmable-modal';
import { entitiesInfosMixin } from '@/mixins/entities/infos';

import MetaAlarmRuleForm from '@/components/other/meta-alarm-rule/form/meta-alarm-rule-form.vue';

import ModalWrapper from '../modal-wrapper.vue';

export default {
  name: MODALS.createMetaAlarmRule,
  $_veeValidate: {
    validator: 'new',
    delay: VALIDATION_DELAY,
  },
  components: {
    MetaAlarmRuleForm,
    ModalWrapper,
  },
  mixins: [
    modalInnerMixin,
    entitiesInfosMixin,
    submittableMixinCreator(),
    confirmableModalMixinCreator(),
  ],
  data() {
    return {
      activeStep: META_ALARMS_FORM_STEPS.general,
      isStepValid: false,
      form: metaAlarmRuleToForm(this.modal.config.rule),
    };
  },
  computed: {
    title() {
      return this.config.title ?? this.$t('modals.metaAlarmRule.create.title');
    },

    isLastStep() {
      return this.activeStep === META_ALARMS_FORM_STEPS.parameters;
    },
  },
  mounted() {
    this.fetchInfos();

    const handleValidationChanged = () => {
      this.isStepValid = this.isCurrentStepValid();
    };

    handleValidationChanged();

    this.$watch(() => this.$refs.formElement.hasGeneralError, handleValidationChanged);
    this.$watch(() => this.$refs.formElement.hasTypeError, handleValidationChanged);
    this.$watch(() => this.$refs.formElement.hasParametersError, handleValidationChanged);
  },
  methods: {
    isCurrentStepValid() {
      const { hasGeneralError, hasTypeError, hasParametersError } = this.$refs.formElement ?? {};

      return {
        [META_ALARMS_FORM_STEPS.general]: !hasGeneralError,
        [META_ALARMS_FORM_STEPS.type]: !hasTypeError,
        [META_ALARMS_FORM_STEPS.parameters]: !hasParametersError,
      }[this.activeStep];
    },

    validateCurrentStepValid() {
      const {
        validateGeneralChildren,
        validateTypeChildren,
        validateParametersChildren,
      } = this.$refs.formElement ?? {};

      const func = {
        [META_ALARMS_FORM_STEPS.general]: validateGeneralChildren,
        [META_ALARMS_FORM_STEPS.type]: validateTypeChildren,
        [META_ALARMS_FORM_STEPS.parameters]: validateParametersChildren,
      }[this.activeStep];

      return func?.();
    },

    async next() {
      const isValid = await this.validateCurrentStepValid();

      if (isValid) {
        this.activeStep += 1;
      }
    },

    async submit() {
      const isFormValid = await this.$validator.validateAll();

      if (isFormValid) {
        await this.config.action(formToMetaAlarmRule(this.form));

        this.$modals.hide();
      }
    },
  },
};
</script>
