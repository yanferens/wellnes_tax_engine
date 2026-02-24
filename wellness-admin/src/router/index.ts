import { createRouter, createWebHistory } from 'vue-router'
import OrdersListView from '../views/OrdersListView.vue'
import ImportCSVView from "../views/ImportCSVView.vue"
import LoginVue from "@/views/LoginVue.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginVue,
      meta: { guestOnly: true }
    },
    {
      path: '/',
      name: 'orders',
      component: OrdersListView,
      meta: { requiresAuth: true }
    },
    {
      path: '/import',
      name: 'import',
      component: ImportCSVView,
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem('auth_token');
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'login' });
  }
  else if (to.meta.guestOnly && isAuthenticated) {
    next({ name: 'orders' });
  }
  else {
    next();
  }
});

export default router