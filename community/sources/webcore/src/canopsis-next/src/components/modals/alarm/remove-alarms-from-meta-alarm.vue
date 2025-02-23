<template>
  <v-form @submit.prevent="submit">
    <modal-wrapper close>
      <template #title="">
        <span>{{ config.title }}</span>
      </template>
      <template #text="">
        <alarm-general-table :items="config.items" class="mb-4" />
        <c-name-field
          v-model="form.comment"
          :label="$tc('common.comment')"
          :required="isCommentRequired"
          name="comment"
          autofocus
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
          :loading="submitting"
          :disabled="isDisabled"
          class="primary"
          type="submit"
        >
          {{ $t('common.saveChanges') }}
        </v-btn>
      </template>
    </modal-wrapper>
  </v-form>
</template>

<script>
import { MODALS, VALIDATION_DELAY } from '@/constants';

import { mapIds } from '@/helpers/array';

import { modalInnerMixin } from '@/mixins/modal/inner';
import { submittableMixinCreator } from '@/mixins/submittable';
import { confirmableModalMixinCreator } from '@/mixins/confirmable-modal';

import AlarmGeneralTable from '@/components/widgets/alarm/alarm-general-list.vue';

import ModalWrapper from '../modal-wrapper.vue';

/**
 * Modal to remove alarms from meta alarm
 */
export default {
  name: MODALS.removeAlarmsFromMetaAlarm,
  $_veeValidate: {
    validator: 'new',
    delay: VALIDATION_DELAY,
  },
  components: { AlarmGeneralTable, ModalWrapper },
  mixins: [
    modalInnerMixin,
    submittableMixinCreator(),
    confirmableModalMixinCreator(),
  ],
  data() {
    return {
      form: {
        comment: '',
      },
    };
  },
  computed: {
    isCommentRequired() {
      return this.config.isCommentRequired ?? true;
    },
  },
  methods: {
    async submit() {
      const isFormValid = await this.$validator.validateAll();

      if (isFormValid) {
        await this.config?.action?.({
          ...this.form,

          alarms: mapIds(this.config.items),
        });

        this.$modals.hide();
      }
    },
  },
};
</script>
