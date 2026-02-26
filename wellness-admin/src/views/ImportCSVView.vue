<script setup lang="ts">
  import { ref } from 'vue';
  import { useOrdersStore } from '../stores/ordersStore';
  import FileDropZone from '@/components/FileDrop.vue';

  const ordersStore = useOrdersStore();
  const selectedFile = ref<File | null>(null);

  const setFile = (file: File) => {
    selectedFile.value = file;
    ordersStore.errorMessage = null;
  };

  const handleError = (msg: string) => {
    alert(msg);
  };

  const submitUpload = async () => {
    if (!selectedFile.value) return;
    await ordersStore.uploadCsv(selectedFile.value);
    if (!ordersStore.errorMessage) {
      alert('Succesfully uploaded!');
      selectedFile.value = null;
    }
  };
</script>

<template>
  <main>
    <section class="import_csv_file_section">
      <div class="section_wrap">
        <div class="import_icon">
          <img src="../assets/import.png" alt="import"/>
        </div>
        <h1>Iмпорт замовлень CSV</h1>
        <p>Завантажте ваші файли у віконце, щоб порахувати податок автоматично</p>
        <div class="dropzone-container">
          <FileDropZone v-if="!selectedFile" @file-selected="setFile" @error="handleError"/>
          <div v-else class="file-preview">
            <div class="file-info">
              📄 {{ selectedFile.name }} ({{ (selectedFile.size / 1024 / 1024).toFixed(2) }} MB)
            </div>
            <button @click="selectedFile = null" class="remove-btn">Видалити</button>
          </div>
        </div>
        <div class="actions" v-if="selectedFile">
          <button @click="submitUpload" :disabled="ordersStore.isLoading" class="upload-btn">
            {{ ordersStore.isLoading ? 'Обробка...' : 'Завантажити на сервер' }}
          </button>
        </div>
        <div class="example-cont">
          <div class="example-text"><span>Обов’язковий формат CSV:</span> довгота, широта,  вартість без податку</div>
          <div class="example-text"><span>Зразок:</span></div>
          <ol>
            <li>40.712776, -74.005974, 125.50</li>
            <li>40.758896, -73.985130, 89.99</li>
          </ol>
        </div>
      </div>
    </section>
  </main>
</template>

<style scoped>
  .import_csv_file_section{
    margin: 170px 0 74px;
  }
  .import_csv_file_section .section_wrap{
    max-width: 897px;
    border-radius: 12px;
    border: 1px solid #E5E7EB;
    background: #FFF;
    box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.25);
    display: flex;
    flex-direction: column;
    align-items: center;
    padding: 100px 32px;
    row-gap: 24px;
  }
  .import_icon{
    padding: 8px 16px;
    border-radius: 12px;
    border: 1px solid #32D571;
    background: #E8F3EC;
    width: fit-content;
  }
  .import_icon img{
    width: 157px;
    height: 157px;
  }
  h1{
    color: #000;
    font-size: 32px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
  p{
    color: #A7A3A3;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
  .dropzone-container {
    width: 100%;
    display: flex;
    justify-content: center;
  }
  .file-preview {
    width: 100%;
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 8px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #f9fafb;
  }
  .upload-btn {
    background-color: #000;
    color: white;
    padding: 12px 24px;
    border-radius: 6px;
    border: none;
    cursor: pointer;
    font-size: 16px;
  }
  .remove-btn {
    color: red;
    background: none;
    border: none;
    cursor: pointer;
    text-decoration: underline;
  }
  .example-cont{
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    align-self: flex-start;
    width: fit-content;
    margin-left: 97px;
  }
  .example-cont .example-text{
    color: #000;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
  .example-cont .example-text span{
    color: #000;
    font-size: 20px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
  .example-cont ol{
    list-style-position: inside;
  }

  @media screen and (max-width: 600px){
    .import_csv_file_section{
      margin: 145px 0 49px;
    }
    .import_csv_file_section .section_wrap{
      padding: 60px 30px;
      row-gap: 18px;
    }
    h1{
      font-size: 28px;
      text-align: center;
    }
    p{
      text-align: center;
    }
    .example-cont{
      margin-left: 0;
    }
  }
</style>