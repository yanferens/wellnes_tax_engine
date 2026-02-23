<script setup lang="ts">
  import { onMounted } from 'vue';
  import { useOrdersStore } from '../stores/ordersStore';

  const ordersStore = useOrdersStore();

  // call data loading when opening the page
  onMounted(() => {
    ordersStore.fetchOrders();
  });
</script>

<template>
  <div>
    <h1>Список замовлень</h1>

    <div v-if="ordersStore.isLoading">Завантаження даних... ⏳</div>
    <div v-else-if="ordersStore.errorMessage" style="color: red;">{{ ordersStore.errorMessage }}</div>

    <div v-else>
      <table border="1" cellpadding="10" style="border-collapse: collapse; width: 100%; text-align: left;">
        <thead>
        <tr>
          <th>ID</th>
          <th>Subtotal</th>
          <th>Tax Amount</th>
          <th>Total</th>
          <th>Tax Rate</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="order in ordersStore.orders" :key="order.id">
          <td>{{ order.id }}</td>
          <td>${{ order.subtotal }}</td>
          <td>${{ order.tax_amount }}</td>
          <td><b>${{ order.total_amount }}</b></td>
          <td>{{ (order.composite_tax_rate * 100).toFixed(3) }}%</td>
        </tr>
        </tbody>
      </table>

      <div style="margin-top: 20px; display: flex; gap: 10px; align-items: center;">
        <button
            :disabled="ordersStore.currentPage === 1"
            @click="ordersStore.setPage(ordersStore.currentPage - 1)"
        >
          Попередня
        </button>
        <span>Сторінка {{ ordersStore.currentPage }} з {{ Math.ceil(ordersStore.totalOrders / ordersStore.limit) }}</span>
        <button
            :disabled="ordersStore.currentPage >= Math.ceil(ordersStore.totalOrders / ordersStore.limit)"
            @click="ordersStore.setPage(ordersStore.currentPage + 1)"
        >
          Наступна
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>