<script setup lang="ts">
import { ref } from 'vue';

const emit = defineEmits<{
  (e: 'file-selected', file: File): void
  (e: 'error', message: string): void
}>();

const fileInput = ref<HTMLInputElement | null>(null);
const isDragging = ref(false);

const triggerInput = () => {
  fileInput.value?.click();
};

const onFileChanged = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files.length > 0) {
    validateAndEmit(target.files[0]);
  }
};

const onDragOver = (e: DragEvent) => {
  isDragging.value = true;
};

const onDragLeave = (e: DragEvent) => {
  isDragging.value = false;
};

const onDrop = (e: DragEvent) => {
  isDragging.value = false;
  if (e.dataTransfer?.files && e.dataTransfer.files.length > 0) {
    validateAndEmit(e.dataTransfer.files[0]);
  }
};

const validateAndEmit = (file: File) => {
  const isCsv = file.type === 'text/csv' || file.type === 'application/vnd.ms-excel' || file.name.endsWith('.csv');
  if (!isCsv) {
    emit('error', 'Please upload a file in .csv format.');
    return;
  }
  if (file.size > 50 * 1024 * 1024) {
    emit('error', 'File too large (max. 50 MB)');
    return;
  }
  emit('file-selected', file);
};
</script>

<template>
  <div class="drop-zone" :class="{ 'dragging': isDragging }" @click="triggerInput" @dragover.prevent="onDragOver"
       @dragleave.prevent="onDragLeave" @drop.prevent="onDrop">
    <input ref="fileInput" type="file" accept=".csv" class="hidden-input" @change="onFileChanged"/>

    <div class="icon-wrapper">
      <img src="@/assets/import.png" alt="import"/>
    </div>

    <p class="main-text">Натисніть, або перетягніть сюди.</p>
    <p class="sub-text">Лише CSV (макс. 50 МБ)</p>
  </div>
</template>

<style scoped>
.drop-zone {
  width: 100%;
  max-width: 667px;
  padding: 33px 0;
  background-color: #E8F3EC;
  border: 1px solid #32D571;
  border-radius: 16px;
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 18px;
  cursor: pointer;
  transition: all 0.2s ease;
  text-align: center;
}
.drop-zone:hover {
  background-color: #cde3d5;
}
.drop-zone.dragging {
  background-color: #C8E6C9;
  border: 2px dashed #2E7D32;
  transform: scale(1.01);
}
.hidden-input {
  display: none;
}
.icon-wrapper img {
  width: 102px;
  height: 102px;
}
.main-text {
  color: #000;
  font-size: 16px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
}
.sub-text {
  color: #A7A3A3;
  font-size: 13px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
}
</style>