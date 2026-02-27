<script setup lang="ts">
import { ref, computed } from 'vue';
import { useOrdersStore } from '../stores/ordersStore';
import FileDropZone from '@/components/FileDrop.vue';
import ErrorMessageBlock from '@/components/ErrorMessageBlock.vue';

const ordersStore = useOrdersStore();
const selectedFile = ref<File | null>(null);
const errorMessage = ref<string>('');
const hasError = computed(() => errorMessage.value.length > 0);

const setFile = (file: File) => {
  selectedFile.value = file;
  ordersStore.errorMessage = null;
};

const handleError = (message: string) => {
  errorMessage.value = message;
};

const clearError = () => {
  errorMessage.value = '';
};

const submitUpload = async () => {
  if (!selectedFile.value) return;
  await ordersStore.uploadCsv(selectedFile.value);
  if (!ordersStore.errorMessage) {
    alert('Succesfully uploaded!');
    selectedFile.value = null;
  } else {
    handleError(ordersStore.errorMessage);
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
          <ErrorMessageBlock :show="hasError" :message="errorMessage" @close="clearError"/>
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
  flex-direction: column;
  align-items: center;
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
  border: none;
  padding: 8px 16px;
  border-radius: 5px;
  background: #32D571;
  color: #FFFEFE;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.25);
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  z-index: 1;
}
.upload-btn:after{
  content: "";
  position: absolute;
  top: 0;
  left: -150%;
  width: 80%;
  height: 100%;
  background: linear-gradient(
      to right,
      rgba(255, 255, 255, 0) 0%,
      rgba(255, 255, 255, 0.2) 50%,
      rgba(255, 255, 255, 0) 100%
  );
  transform: skewX(-25deg);
  animation: shine 4s ease-in-out infinite;
  pointer-events: none;
}
.upload-btn:hover, .upload-btn:active{
  transform: translateY(-3px);
  box-shadow: 0 8px 15px rgba(50, 213, 113, 0.4);
  background: #28c063;
}
@keyframes shine {
  0% { left: -150%; opacity: 0; }
  10% { opacity: 1; }
  40% { left: 150%; opacity: 1; }
  50% { opacity: 0; }
  100% { left: 150%; opacity: 0; }
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
  .error-message {
    padding: 30px 20px;
  }
}
</style>