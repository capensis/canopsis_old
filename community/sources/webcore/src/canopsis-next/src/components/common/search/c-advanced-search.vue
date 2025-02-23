<template>
  <v-layout class="c-advanced-search" align-end>
    <c-advanced-search-field
      v-if="advancedSearchActive"
      v-model="localValue"
      :fields="fields"
      :conditions="conditions"
    />
    <c-search-field
      v-else
      v-model="localValue"
      :items="savedItems"
      :tooltip="tooltip"
      :combobox="combobox"
      @submit="submit"
      @remove:item="removeItem"
      @toggle-pin:item="togglePinItem"
    />
    <c-action-btn
      :tooltip="$t('common.search')"
      icon="search"
      @click="submit"
    />
    <c-action-btn
      :tooltip="$t('common.clearSearch')"
      icon="clear"
      @click="clear"
    />
    <c-action-btn
      :tooltip="advancedSearchActiveTooltip"
      :input-value="advancedSearchActive"
      icon="tune"
      @click="toggleAdvancedSearchActive"
    />
  </v-layout>
</template>

<script>
import { computed, ref, toRef } from 'vue';

import { ADVANCED_SEARCH_CONDITIONS } from '@/constants';

import { useI18n } from '@/hooks/i18n';

import { useSearchLocalValue } from './hooks/search';

export default {
  model: {
    prop: 'value',
    event: 'input',
  },
  props: {
    value: {
      type: String,
      default: '',
    },
    fields: {
      type: Array,
      default: () => [],
    },
    conditions: {
      type: Array,
      default: () => Object.values(ADVANCED_SEARCH_CONDITIONS),
    },
    savedItems: {
      type: Array,
      default: () => [],
    },
    tooltip: {
      type: String,
      default: '',
    },
    combobox: {
      type: Boolean,
      default: false,
    },
  },
  setup(props, { emit }) {
    const { t } = useI18n();

    const addItem = search => emit('add:item', search);
    const removeItem = search => emit('remove:item', search);
    const togglePinItem = search => emit('toggle-pin:item', search);

    const {
      localValue,
      submit,
      clear,
    } = useSearchLocalValue({
      value: toRef(props, 'value'),
      fields: toRef(props, 'fields'),
      onSubmit: addItem,
    }, emit);

    const advancedSearchActive = ref(false);
    const advancedSearchActiveTooltip = computed(() => (
      advancedSearchActive.value
        ? t('advancedSearch.switchAdvancedSearchActiveToFalse')
        : t('advancedSearch.switchAdvancedSearchActiveToTrue')
    ));

    const toggleAdvancedSearchActive = () => advancedSearchActive.value = !advancedSearchActive.value;

    return {
      localValue,
      advancedSearchActive,
      advancedSearchActiveTooltip,

      toggleAdvancedSearchActive,
      submit,
      clear,
      removeItem,
      togglePinItem,
    };
  },
};
</script>

<style lang="scss" scoped>
.c-advanced-search::v-deep {
  --advanced-search-input-min-width: 170px;

  .v-input, .v-input__control .v-input__slot, .v-select__selections {
    width: unset;
    max-width: unset;
  }

  .v-select__slot {
    width: 100%;
    max-width: unset;

    input {
      min-width: var(--advanced-search-input-min-width);
      width: var(--advanced-search-input-min-width);
    }

    &, .v-select__selections {
      &:first-child.input {
        width: 100%;
      }
    }
  }
}
</style>
