<template>
  <v-form @submit.prevent="submit">
    <modal-wrapper close>
      <template #title="">
        <span>{{ $t('modals.createDeclareTicketEvent.title') }}</span>
      </template>
      <template #text="">
        <c-enabled-field
          v-if="config.items.length > 1"
          v-model="singleMode"
          :label="$t('declareTicket.oneTicketForAlarms')"
        />
        <declare-ticket-events-form
          v-model="form"
          :alarms="config.items"
          :tickets-by-alarms="config.ticketsByAlarms"
          :alarms-by-tickets="config.alarmsByTickets"
          :hide-ticket-resource="!isAllComponentAlarms"
          :single-mode="singleMode"
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
          {{ $t('common.submit') }}
        </v-btn>
      </template>
    </modal-wrapper>
  </v-form>
</template>

<script>
import { keyBy } from 'lodash';

import { MODALS, VALIDATION_DELAY } from '@/constants';

import {
  alarmsToDeclareTicketEventForm,
  formToDeclareTicketEvents,
} from '@/helpers/entities/declare-ticket/event/form';
import { isEntityComponentType } from '@/helpers/entities/entity/form';
import { isSuccessStepTicketType } from '@/helpers/entities/alarm/step/entity';

import { modalInnerMixin } from '@/mixins/modal/inner';
import { submittableMixinCreator } from '@/mixins/submittable';
import { confirmableModalMixinCreator } from '@/mixins/confirmable-modal';

import DeclareTicketEventsForm from '@/components/other/declare-ticket/form/declare-ticket-events-form.vue';

import ModalWrapper from '../modal-wrapper.vue';

/**
 * Modal to declare a ticket
 */
export default {
  name: MODALS.createDeclareTicketEvent,
  $_veeValidate: {
    validator: 'new',
    delay: VALIDATION_DELAY,
  },
  components: { DeclareTicketEventsForm, ModalWrapper },
  mixins: [
    modalInnerMixin,
    submittableMixinCreator(),
    confirmableModalMixinCreator(),
  ],
  data() {
    const { alarmsByTickets } = this.modal.config;

    return {
      singleMode: false,
      form: alarmsToDeclareTicketEventForm(alarmsByTickets),
    };
  },
  computed: {
    isAllComponentAlarms() {
      return this.config.items.every(({ entity }) => isEntityComponentType(entity.type));
    },
  },
  methods: {
    isSomeOneSuccessTicketExist() {
      const alarmsById = keyBy(this.config.items, '_id');

      return Object.entries(this.form.alarms_by_tickets).some(([ticketRuleId, alarmIds]) => alarmIds.some((alarmId) => {
        const alarm = alarmsById[alarmId];

        return alarm.v.tickets?.some(
          ticket => ticketRuleId === ticket.ticket_rule_id
            && isSuccessStepTicketType(ticket._t),
        );
      }));
    },

    async callActionAndHideModal() {
      const preparedForm = formToDeclareTicketEvents(this.form, this.singleMode);

      await this.config.action?.(preparedForm, this.singleMode);

      this.$modals.hide();
    },

    async submit() {
      const isFormValid = await this.$validator.validateAll();

      if (isFormValid) {
        if (this.isSomeOneSuccessTicketExist()) {
          const isSingleAlarm = this.config.items.length === 1;

          this.$modals.show({
            name: MODALS.confirmation,
            config: {
              title: isSingleAlarm
                ? this.$t('modals.confirmationCreateNewTicketForAlarm.title')
                : this.$t('modals.confirmationCreateNewTicketForAlarms.title'),
              text: isSingleAlarm
                ? this.$t('modals.confirmationCreateNewTicketForAlarm.text')
                : this.$t('modals.confirmationCreateNewTicketForAlarms.text'),
              action: this.callActionAndHideModal,
            },
          });
        } else {
          await this.callActionAndHideModal();
        }
      }
    },
  },
};
</script>
