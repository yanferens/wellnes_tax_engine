<script setup lang="ts">
import { ref } from 'vue';
import { useOrdersStore } from '@/stores/ordersStore';

const store = useOrdersStore();
const searchQuery = ref('');
let timeout: ReturnType<typeof setTimeout>;

const onInput = () => {
  clearTimeout(timeout);
  timeout = setTimeout(() => {
    store.setFilters({ ...store.filters, search: searchQuery.value });
  }, 500);
};
</script>

<template>
  <div class="search_input_wrap">
    <img src="../assets/search.svg" alt="search" class="search_img" v-show="searchQuery === ''"/>
    <input v-model="searchQuery" @input="onInput" type="text" placeholder="Знайти за ID, широтою, довготою">
  </div>
</template>

<style scoped>
  .search_input_wrap{
    position: relative;
    width: 100%;
    display: flex;
  }
  input{
    width: 100%;
    height: 64px;
    padding: 10px 22px;
    border-radius: 12px;
    border: 1px solid #9CA3AF;
    background: #F9FAFB;
    box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.25);
    color: #9CA3AF;
    text-align: center;
    font-size: 18px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
  input::placeholder{
    color: #9CA3AF;
    text-align: center;
    font-size: 18px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
  .search_img{
    position: absolute;
    top: 31%;
    left: 36%;
    width: 27px;
    height: 27px;
    flex-shrink: 0;
    color: #9CA3AF;
  }
</style>