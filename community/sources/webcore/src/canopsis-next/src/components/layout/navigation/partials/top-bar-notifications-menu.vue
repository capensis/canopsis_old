<template>
  <top-bar-menu
    :title="$tc('common.notification', 2)"
    :links="links"
  />
</template>

<script>
import { ROUTES_NAMES, USER_PERMISSIONS } from '@/constants';

import TopBarMenu from './top-bar-menu.vue';

export default {
  components: { TopBarMenu },
  computed: {
    links() {
      return [
        {
          route: { name: ROUTES_NAMES.notificationInstructionStats },
          icon: 'star_half',
          permission: USER_PERMISSIONS.technical.notification.instructionStats,
        },
      ];
    },
  },
};
</script>
