<script setup lang="ts">
import { useOrdersStore } from '@/stores/ordersStore';
const store = useOrdersStore();
</script>

<template>
  <div class="table-wrapper">
    <div v-if="store.isLoading" class="loading">Завантаження даних... ⏳</div>
    <div v-else-if="store.errorMessage" class="error">{{ store.errorMessage }}</div>

    <div v-else>
      <table>

        <thead>
        <tr>
          <th>ID</th>
          <th>Subtotal</th>
          <th>Total Amount</th>
          <th>Jurisdictions</th>
        </tr>
        </thead>

        <tbody>
        <tr v-for="order in store.orders" :key="order.id">
          <td>{{ order.id }}</td>
          <td>${{ order.subtotal }}</td>
          <td>${{ order.total_amount }}</td>
          <td>{{order.jurisdictions}}</td>
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
table{
  width: 100%;
  border-radius: 12px;
  border: 1px solid #E5E7EB;
  background: #FFF;
  box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.05);
  padding: 14px 0;
}
table tbody tr td, table thead tr th {
  padding: 0 20px;
}
table thead tr th{
  color: #6B7280;
  font-size: 13px;
  font-style: normal;
  font-weight: 700;
  line-height: normal;
  padding-bottom: 15px;
  text-align: left;
}
table thead tr th:first-child{
  width: 351.55px;
}
table thead tr th:nth-child(2){
  width: 241.22px;
}
table thead tr th:nth-child(3){
  width: 320.73px;
}
table thead tr th:last-child{
  width: 286.5px;
}
table tbody tr td{
  color: #1F2937;
  font-style: normal;
  line-height: normal;
  text-align: left;
  border-top: 1px solid #F3F4F6;
}
table tbody tr td:first-child, table tbody tr td:nth-child(3){
  font-size: 13.5px;
  font-weight: 700;
  padding: 16.5px 20px;
}
table tbody tr td:nth-child(2){
  font-size: 13.9px;
  font-weight: 400;
}
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