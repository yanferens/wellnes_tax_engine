<script setup lang="ts">
import { useOrdersStore } from '@/stores/ordersStore';
const store = useOrdersStore();
</script>

<template>
  <div class="table-wrapper">
    <div v-if="store.isLoading" class="loading">Завантаження даних... ⏳</div>
    <div v-else-if="store.errorMessage" class="error">{{ store.errorMessage }}</div>

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
        <tr v-for="order in store.orders" :key="order.id">
          <td>{{ order.id }}</td>
          <td>${{ order.subtotal }}</td>
          <td>${{ order.tax_amount }}</td>
          <td><b>${{ order.total_amount }}</b></td>
          <td>{{ (order.composite_tax_rate * 100).toFixed(3) }}%</td>
        </tr>
        <tr v-if="store.orders.length === 0">
          <td colspan="5" style="text-align: center;">Даних не знайдено</td>
        </tr>
        </tbody>
      </table>

      <div class="pagination">
        <button
            :disabled="store.currentPage === 1"
            @click="store.setPage(store.currentPage - 1)"
        >
          ← Попередня
        </button>

        <span>Сторінка {{ store.currentPage }} з {{ Math.ceil(store.totalOrders / store.limit) || 1 }}</span>

        <button
            :disabled="store.currentPage >= Math.ceil(store.totalOrders / store.limit)"
            @click="store.setPage(store.currentPage + 1)"
        >
          Наступна →
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.pagination {
  margin-top: 20px;
  display: flex;
  gap: 10px;
  align-items: center;
  justify-content: center;
}
.error { color: red; }
.loading { text-align: center; padding: 20px; }
</style>