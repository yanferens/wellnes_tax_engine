<script setup lang="ts">
import { computed } from 'vue';
import { ref, onMounted, onUnmounted } from 'vue';
import { useRoute } from 'vue-router';
import Logo from "@/components/Logo.vue";
import Navigation from "@/components/Navigation.vue";
import CreateOrderExit from "@/components/CreateOrderExit.vue";
import MobileMenu from "@/components/MobileMenu.vue";

const route = useRoute();
const isLoginPage = computed(() => route.name === 'login');

const isMobile = ref(false);
const checkScreenSize = () => {
  isMobile.value = window.innerWidth <= 768;
};
onMounted(() => {
  checkScreenSize();
  window.addEventListener('resize', checkScreenSize);
});
onUnmounted(() => {
  window.removeEventListener('resize', checkScreenSize);
});
</script>

<template>
  <header>
    <div class="header_wrap" :class="{ 'login_page_content': isLoginPage }">
      <Logo />
      <template v-if="!isLoginPage">
        <MobileMenu v-if="isMobile"/>
        <div v-else class="desktop-nav">
          <Navigation/>
          <CreateOrderExit/>
        </div>
      </template>
      <div v-else class="admin_panel">Панель адміна</div>
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
    padding: 0 20px;
  }
  .header_wrap{
    width: 100%;
    max-width: 1200px;
    margin: 0 auto;
    display: flex;
    flex-direction: row;
    column-gap: 176px;
    align-items: center;
    padding: 18px 0;
  }
  .header_wrap div.desktop-nav{
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    column-gap: 176px;
  }
  .admin_panel{
    display: none;
  }
  .login_page_content.header_wrap{
    max-width: 897px;
    column-gap: 0;
    justify-content: space-between;
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

  @media screen and (max-width: 600px){
    .header_wrap{
      width: 100%;
      max-width: 1200px;
      margin: 0 auto;
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      align-items: center;
      padding: 18px 0;
      column-gap: 0;
    }
  }
</style>