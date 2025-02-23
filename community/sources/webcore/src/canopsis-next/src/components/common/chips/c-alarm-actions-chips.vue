<template>
  <v-layout
    :key="wrapperKey"
    :class="{ 'c-alarm-actions-chips--small': small }"
    class="c-alarm-actions-chips"
    wrap
    align-center
  >
    <c-alarm-action-chip
      v-for="item in inlineItems"
      :key="item[itemValue]"
      :class="itemClass"
      :color="item.color"
      :small="small"
      :closable="isClosableItem(item)"
      :outlined="outlined"
      :text-color="textColor"
      @click="selectItem(item)"
      @close="closeItem(item)"
    >
      <slot
        :item="item"
        name="item"
      >
        <span>{{ item[itemText] }}</span>
      </slot>
    </c-alarm-action-chip>
    <v-menu
      v-if="dropDownItems.length"
      key="more"
      max-height="400px"
      bottom
      left
      @input="$emit('activate')"
    >
      <template #activator="{ on }">
        <v-btn
          class="c-alarm-actions-chips__more-btn ma-0"
          color="grey"
          icon
          v-on="on"
        >
          <v-icon
            color="white"
            size="20"
          >
            more_horiz
          </v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-card-text>
          <v-layout
            :class="{ 'c-alarm-actions-chips--small': small }"
            class="c-alarm-actions-chips__more"
            wrap
          >
            <c-alarm-action-chip
              v-for="item in dropDownItems"
              :key="item[itemValue]"
              :class="itemClass"
              :color="item.color"
              :closable="isClosableItem(item)"
              :text-color="textColor"
              :outlined="outlined"
              class="mx-0"
              @click="selectItem(item)"
              @close="closeItem(item)"
            >
              <slot
                :item="item"
                name="item"
              >
                <span>{{ item[itemText] }}</span>
              </slot>
            </c-alarm-action-chip>
          </v-layout>
        </v-card-text>
      </v-card>
    </v-menu>
  </v-layout>
</template>

<script>
import { uid } from '@/helpers/uid';

export default {
  props: {
    items: {
      type: Array,
      default: () => [],
    },
    activeItem: {
      type: String,
      required: false,
    },
    activeItems: {
      type: Array,
      default: () => [],
    },
    inlineCount: {
      type: [Number, String],
      default: 2,
    },
    closable: {
      type: Boolean,
      default: false,
    },
    closableActive: {
      type: Boolean,
      default: false,
    },
    small: {
      type: Boolean,
      default: false,
    },
    outlined: {
      type: Boolean,
      default: false,
    },
    itemValue: {
      type: String,
      default: 'text',
    },
    itemText: {
      type: String,
      default: 'text',
    },
    itemClass: {
      type: [String, Object],
      required: false,
    },
    returnObject: {
      type: Boolean,
      default: false,
    },
    textColor: {
      type: String,
      default: 'white',
    },
  },
  data() {
    return {
      wrapperKey: uid(),
    };
  },
  computed: {
    sortedItems() {
      return [...this.items].sort((first, second) => {
        if (first[this.itemValue] === this.activeItem || this.activeItems.includes(first[this.itemValue])) {
          return -1;
        }

        if (second[this.itemValue] === this.activeItem || this.activeItems.includes(second[this.itemValue])) {
          return 0;
        }

        if (first[this.itemText] < second[this.itemText]) {
          return -1;
        }

        if (first[this.itemText] > second[this.itemText]) {
          return 1;
        }

        return 0;
      });
    },

    inlineItems() {
      return this.sortedItems.slice(0, this.inlineCount);
    },

    dropDownItems() {
      return this.sortedItems.slice(this.inlineCount);
    },
  },
  watch: {
    inlineItems() {
      this.wrapperKey = uid();
    },
  },
  methods: {
    selectItem(item) {
      this.$emit('select', this.returnObject ? item : item[this.itemValue]);
    },

    closeItem(item) {
      this.$emit('close', this.returnObject ? item : item[this.itemValue]);
    },

    isClosableItem(item) {
      const isActive = this.activeItem === item[this.itemValue] || this.activeItems.includes(item[this.itemValue]);

      return this.closable || (this.closableActive && isActive);
    },
  },
};
</script>

<style lang="scss">
.c-alarm-actions-chips {
  &, &__more {
    column-gap: 8px;
    row-gap: 4px;
  }

  &--small {
    column-gap: 4px;
  }

  & &__more-btn {
    width: 24px;
    height: 24px;

    .theme--light & {
      background-color: var(--v-application-background-darken2);
    }

    .theme--dark & {
      background-color: var(--v-application-background-lighten4);
    }
  }
}
</style>
