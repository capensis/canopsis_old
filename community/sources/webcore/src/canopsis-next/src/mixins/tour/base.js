import { createNamespacedHelpers } from 'vuex';
import { formToUser, userToForm } from '@/helpers/forms/user';

const { mapActions: authMapActions, mapGetters: authMapGetters } = createNamespacedHelpers('auth');
const { mapActions: userMapActions } = createNamespacedHelpers('user');

export const tourBaseMixin = {
  computed: {
    ...authMapGetters(['currentUser']),
  },
  methods: {
    ...authMapActions(['fetchCurrentUser']),
    ...userMapActions({
      updateUser: 'update',
    }),

    async finishTourByName(tourName) {
      const userForm = userToForm(this.currentUser);
      const user = formToUser({
        ...userForm,
        ui_tours: {
          ...this.currentUser.ui_tours,

          [tourName]: true,
        },
      });

      await this.updateUser({
        data: user,
        id: this.currentUser._id,
      });

      this.fetchCurrentUser();
    },
  },
};
