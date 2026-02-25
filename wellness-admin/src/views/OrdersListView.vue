<script setup lang="ts">
  import { onMounted } from 'vue';
  import { useOrdersStore } from '../stores/ordersStore';
  import SearchInput from "@/components/SearchInput.vue";
  import OrdersTable from "@/components/OrdersTable.vue";

  const ordersStore = useOrdersStore();

  // call data loading when opening the page
  onMounted(() => {
    ordersStore.fetchOrders();
  });
</script>

<template>
  <main>
    <section class="orders_list">
      <div class="section_wrap">
        <div class="controls-row">
          <SearchInput class="search-component" />
        </div>
        <OrdersTable />
      </div>
    </section>
  </main>
</template>

<style scoped>
  .orders_list{
    padding: 76px 0;
  }
  .controls-row {
    margin-bottom: 46px;
  }
</style>