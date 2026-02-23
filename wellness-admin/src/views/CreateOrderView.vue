<script setup lang="ts">
  import { ref } from 'vue';
  import { useOrdersStore } from '../stores/ordersStore';

  const ordersStore = useOrdersStore();

  const form = ref({
    latitude: 0,
    longitude: 0,
    subtotal: 0,
  });

  const submitForm = async () => {
    const orderData = {
      ...form.value,
      timestamp: new Date().toISOString(), // Додаємо поточний час
    };
    await ordersStore.createOrder(orderData);
    alert('Замовлення створено та розраховано!');
  };
</script>

<template>
  <div>
    <h1>Створити замовлення вручну</h1>
    <form @submit.prevent="submitForm">
      <div>
        <label>Latitude:</label>
        <input type="number" step="any" v-model="form.latitude" required />
      </div>
      <div>
        <label>Longitude:</label>
        <input type="number" step="any" v-model="form.longitude" required />
      </div>
      <div>
        <label>Subtotal ($):</label>
        <input type="number" step="0.01" v-model="form.subtotal" required />
      </div>
      <button type="submit" :disabled="ordersStore.isLoading">Створити</button>
    </form>
  </div>
</template>

<style scoped>

</style>