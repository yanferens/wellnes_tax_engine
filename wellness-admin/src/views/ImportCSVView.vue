<script setup lang="ts">
  import { ref } from 'vue';
  import { useOrdersStore } from '../stores/ordersStore';

  const ordersStore = useOrdersStore();
  const selectedFile = ref<File | null>(null);

  const handleFileChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files && target.files.length > 0) {
      selectedFile.value = target.files.item(0);
    }
  };

  const submitUpload = async () => {
    if (!selectedFile.value) return;
    await ordersStore.uploadCsv(selectedFile.value);
    alert('Файл успішно завантажено!');
    selectedFile.value = null;
  };
</script>

<template>
  <div>
    <h1>Завантажити CSV</h1>
    <input type="file" accept=".csv" @change="handleFileChange" />
    <button :disabled="!selectedFile || ordersStore.isLoading" @click="submitUpload">
      Відправити на сервер
    </button>
  </div>
</template>

<style scoped>

</style>