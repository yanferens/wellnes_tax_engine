import { createRouter, createWebHistory } from 'vue-router'
import OrdersListView from '../views/OrdersListView.vue'
import ImportCsvView from '../views/ImportCSVView.vue'
import CreateOrderView from '../views/CreateOrderView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'orders',
      component: OrdersListView
    },
    {
      path: '/import',
      name: 'import',
      component: ImportCsvView
    },
    {
      path: '/create',
      name: 'create',
      component: CreateOrderView
    }
  ]
})

export default router