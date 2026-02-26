<script setup lang="ts">
  import { onMounted, computed } from 'vue';
  import { useOrdersStore } from '../stores/ordersStore';
  import SearchInput from "@/components/SearchInput.vue";
  import OrdersTable from "@/components/OrdersTable.vue";
  import ErrorMessageBlock from '@/components/ErrorMessageBlock.vue';

  const ordersStore = useOrdersStore();
  const hasError = computed(() => !!ordersStore.errorMessage);

  const closeErrorModal = () => {
    ordersStore.errorMessage = null;
  };

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
      <ErrorMessageBlock :show="hasError" :message="ordersStore.errorMessage ||
      'Сталася помилка під час завантаження даних'" @close="closeErrorModal"/>
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
  @media screen and (max-width: 600px){
    .orders_list{
      padding: 56px 0;
    }
    .orders_list .section_wrap{
      padding: 0 30px;
    }
  }
</style>