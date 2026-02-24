<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import Logo from "@/components/Logo.vue";
import Navigation from "@/components/Navigation.vue";
import CreateOrderExit from "@/components/CreateOrderExit.vue";

const route = useRoute();
const isLoginPage = computed(() => route.name === 'login');
</script>

<template>
  <header>
    <div class="header_wrap" :class="{ 'login_page_content': isLoginPage }">
      <Logo />
      <template v-if="!isLoginPage">
        <Navigation/>
        <CreateOrderExit/>
      </template>
      <div class="admin_panel">Панель адміна</div>
    </div>
  </header>
</template>

<style scoped>
  header{
    background-color: white;
    position: fixed;
    top: 0;
    width: 100%;
    box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.05);
    z-index: 9;
  }
  .header_wrap{
    width: 100%;
    max-width: 1200px;
    margin: 0 auto;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 18px 0;
  }
  .admin_panel{
    display: none;
  }
  .login_page_content.header_wrap{
    max-width: 897px;
  }
  .login_page_content .admin_panel{
    display: block;
    padding: 8px 16px;
    border-radius: 5px;
    background: #32D571;
    color: #FFFEFE;
    font-size: 16px;
    font-style: normal;
    font-weight: 400;
    line-height: normal;
  }
</style>