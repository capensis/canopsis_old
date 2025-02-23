<template>
  <v-layout column>
    <v-layout>
      <text-editor-field
        v-field="form.app_title"
        :disabled="disabled"
        :label="$t('userInterface.appTitle')"
        class="fill-width"
        with-default-variables
      />
    </v-layout>
    <c-duration-field
      v-field="form.popup_timeout.info"
      :label="$t('userInterface.infoPopupTimeout')"
      name="popup_timeout.info"
    />
    <c-duration-field
      v-field="form.popup_timeout.error"
      :label="$t('userInterface.errorPopupTimeout')"
      name="popup_timeout.error"
    />
    <v-layout>
      <c-language-field
        v-field="form.language"
        :label="$t('userInterface.language')"
      />
    </v-layout>
    <v-layout>
      <c-number-field
        v-field="form.max_matched_items"
        :label="$t('userInterface.maxMatchedItems')"
        :min="1"
        name="max_matched_items"
      >
        <template #append="">
          <c-help-icon
            :text="$t('userInterface.tooltips.maxMatchedItems')"
            color="grey darken-1"
            icon="help"
            left
          />
        </template>
      </c-number-field>
    </v-layout>
    <v-layout>
      <c-number-field
        v-field="form.check_count_request_timeout"
        :label="$t('userInterface.checkCountRequestTimeout')"
        :min="1"
        name="check_count_request_timeout"
      >
        <template #append="">
          <c-help-icon
            :text="$t('userInterface.tooltips.checkCountRequestTimeout')"
            color="grey darken-1"
            icon="help"
            left
          />
        </template>
      </c-number-field>
    </v-layout>
    <v-layout>
      <c-timezone-field
        v-field="form.timezone"
        disabled
      />
    </v-layout>
    <v-layout>
      <v-flex xs6>
        <c-enabled-field
          v-field="form.allow_change_severity_to_info"
          :label="$t('userInterface.allowChangeSeverityToInfo')"
        />
      </v-flex>
      <v-flex xs6>
        <c-enabled-field
          v-field="form.show_header_on_kiosk_mode"
          :label="$t('userInterface.showHeaderOnKioskMode')"
        />
      </v-flex>
    </v-layout>
    <v-layout>
      <v-flex xs6>
        <c-enabled-field
          v-field="form.required_instruction_approve"
          :label="$t('userInterface.requiredInstructionApprove')"
        />
      </v-flex>
      <v-flex xs6>
        <v-layout>
          <c-enabled-field
            v-field="form.disabled_transitions"
            :label="$t('userInterface.disabledTransitions')"
          >
            <template #append>
              <c-help-icon
                :text="$t('userInterface.disabledTransitionsTooltip')"
                color="grey darken-1"
                icon="help"
                top
              />
            </template>
          </c-enabled-field>
        </v-layout>
      </v-flex>
    </v-layout>
    <v-layout>
      <v-flex xs6>
        <c-enabled-field
          v-field="form.auto_suggest_pbehavior_name"
          :label="$t('userInterface.autoSuggestPbehaviorName')"
        />
      </v-flex>
      <v-flex xs6>
        <c-theme-field
          v-field="form.default_color_theme"
          :label="$t('userInterface.defaultTheme')"
        />
      </v-flex>
    </v-layout>
    <v-layout>
      <v-flex>
        <text-editor-field
          v-field="form.version_description"
          :label="$t('userInterface.versionDescriptionTooltip')"
          :config="textEditorConfig"
          :variables="versionDescriptionVariables"
          with-default-variables
          public
        />
      </v-flex>
    </v-layout>
    <v-layout>
      <v-flex>
        <text-editor-field
          v-field="form.footer"
          :label="$t('userInterface.footer')"
          :config="textEditorConfig"
          with-default-variables
          public
        />
      </v-flex>
    </v-layout>
    <v-layout class="mt-3">
      <v-flex>
        <text-editor-field
          v-field="form.login_page_description"
          :label="$t('userInterface.description')"
          :config="textEditorConfig"
          with-default-variables
          public
        />
      </v-flex>
    </v-layout>
    <v-layout class="mt-3">
      <v-flex>
        <span class="v-label file-selector__label">{{ $t('userInterface.logo') }}</span>
        <v-layout>
          <file-selector
            ref="fileSelectorElement"
            :max-file-size="maxFileSize"
            :disabled="disabled"
            class="mt-1"
            accept="image/*"
            name="logo"
            with-files-list
            @change="changeLogoFile"
          />
        </v-layout>
      </v-flex>
    </v-layout>
  </v-layout>
</template>

<script>
import { computed, ref } from 'vue';

import { MAX_ICON_SIZE_IN_KB } from '@/constants';

import { useModelField } from '@/hooks/form/model-field';

import FileSelector from '@/components/forms/fields/file-selector.vue';
import TextEditorField from '@/components/forms/fields/text-editor-field.vue';

export default {
  components: {
    FileSelector,
    TextEditorField,
  },
  model: {
    prop: 'form',
    event: 'input',
  },
  props: {
    form: {
      type: Object,
      default: () => ({}),
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  },
  setup(props, { emit }) {
    const { updateField } = useModelField(props, emit);

    const maxFileSize = MAX_ICON_SIZE_IN_KB;

    const fileSelectorElement = ref(null);

    const versionDescriptionVariables = [
      'edition',
      'versionUpdated',
      'serialName',
    ].map(value => ({ value, text: value }));

    const textEditorConfig = computed(() => ({ disabled: props.disabled }));

    /**
     * Updates the 'logo' field with the provided file.
     *
     * @param {Array} [file=[]] - An array containing the file to be set as the logo.
     */
    const changeLogoFile = ([file] = []) => updateField('logo', file);

    /**
     * Clears the file selector element.
     *
     * This function accesses the `fileSelectorElement` and calls its `clear` method,
     * if the element is defined, to reset the file selection.
     */
    const reset = () => fileSelectorElement.value?.clear();

    return {
      maxFileSize,
      fileSelectorElement,
      textEditorConfig,
      versionDescriptionVariables,

      changeLogoFile,
      reset,
    };
  },
};
</script>

<style lang="scss" scoped>
.file-selector {
  &__label {
    font-size: .85em;
    display: block;
  }
}
</style>
