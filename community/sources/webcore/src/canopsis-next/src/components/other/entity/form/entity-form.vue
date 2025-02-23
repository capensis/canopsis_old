<template>
  <v-tabs
    slider-color="primary"
    centered
  >
    <v-tab>{{ $t('entity.form') }}</v-tab>
    <v-tab-item eager>
      <v-layout
        class="mt-3"
        column
      >
        <c-name-field
          v-field="form.name"
          disabled
        />
        <c-description-field v-field="form.description" autofocus />
        <v-layout justify-space-between>
          <v-flex xs3>
            <c-enabled-field v-field="form.enabled" />
          </v-flex>
          <v-flex xs9>
            <v-layout>
              <v-flex
                class="pr-3"
                xs3
              >
                <c-impact-level-field
                  v-field="form.impact_level"
                  required
                />
              </v-flex>
              <v-flex
                class="pr-3"
                xs3
              >
                <c-alarm-state-field
                  v-field="form.sli_avail_state"
                  :label="$t('entity.availabilityState')"
                  required
                />
              </v-flex>
              <v-flex xs6>
                <c-entity-type-field
                  v-field="form.type"
                  required
                  disabled
                />
              </v-flex>
            </v-layout>
          </v-flex>
        </v-layout>
        <c-coordinates-field
          v-field="form.coordinates"
          row
        />
        <entity-state-setting
          v-if="hasStateSetting"
          :form="form"
          :preparer="prepareStateSettingForm"
        />
      </v-layout>
    </v-tab-item>
    <v-tab>{{ $t('entity.manageInfos') }}</v-tab>
    <v-tab-item eager>
      <manage-infos v-field="form.infos" />
    </v-tab-item>
  </v-tabs>
</template>

<script>
import { computed } from 'vue';

import { isEntityComponentType } from '@/helpers/entities/entity/form';

import ManageInfos from '@/components/widgets/context/manage-infos.vue';
import EntityStateSetting from '@/components/other/state-setting/entity-state-setting.vue';

export default {
  inject: ['$validator'],
  components: {
    ManageInfos,
    EntityStateSetting,
  },
  model: {
    prop: 'form',
    event: 'input',
  },
  props: {
    form: {
      type: Object,
      required: true,
    },
    prepareStateSettingForm: {
      type: Function,
      default: data => data,
    },
  },
  setup(props) {
    const hasStateSetting = computed(() => isEntityComponentType(props.form.type));

    return {
      hasStateSetting,
    };
  },
};
</script>
