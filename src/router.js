import Vue from 'vue';
import Router from 'vue-router';

import Login from '@/views/login.vue';
import Alarm from '@/views/alarm.vue';
import Home from '@/views/home.vue';
import About from '@/views/about.vue';
import Filter from '@/views/filter.vue';
import Rrule from '@/components/other/rrule/rrule-form.vue';

// EXAMPLES

Vue.use(Router);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/login',
      name: 'login',
      component: Login,
    },
    {
      path: '/about',
      name: 'about',
      component: About,
    },
    {
      path: '/alarms',
      name: 'alarms',
      component: Alarm,
    },
    {
      path: '/filter',
      name: 'filter',
      component: Filter,
    },
    {
      path: '/rrule',
      name: 'rrule',
      component: Rrule,
    },
  ],
});
