<script setup lang="ts">
import { ref } from 'vue';
import OrderForm from '@/components/OrderForm.vue';

const isModalOpen = ref(false);

const openForm = () => {
  isModalOpen.value = true;
};
</script>

<template>
  <div class="create_order_wrap">
    <button class="create_order_btn" type="button" @click="openForm">+ Створити вручну</button>
    <a href="" class="exit_btn"><img src="../assets/exit.png" alt="exit"></a>
    <OrderForm
        :is-open="isModalOpen"
        @close="isModalOpen = false"
    />
  </div>
</template>

<style scoped>
  .create_order_wrap{
    display: flex;
    flex-direction: row;
    align-items: center;
    column-gap: 50px;
  }
  .create_order_btn{
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
  .create_order_btn::after {
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
  .create_order_btn:hover, .create_order_btn:active{
    transform: translateY(-3px);
    box-shadow: 0 8px 15px rgba(50, 213, 113, 0.4);
    background: #28c063;
  }
  @keyframes shine {
    0% {
      left: -150%;
      opacity: 0;
    }
    10% {
      opacity: 1;
    }
    40% {
      left: 150%;
      opacity: 1;
    }
    50% {
      opacity: 0;
    }
    100% {
      left: 150%;
      opacity: 0;
    }
  }
  .exit_btn img{
    object-fit: contain;
    width: 38px;
    height: 38px;
    filter: drop-shadow(0px 3px 2px rgba(0, 0, 0, 0.25));
  }
</style>