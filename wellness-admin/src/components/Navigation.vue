<script setup lang="ts">
import { useRoute } from 'vue-router';
const route = useRoute();
</script>

<template>
  <nav>
    <ul>
      <li><RouterLink :to="{ name: 'orders' }" :class="route.name === 'orders' ? 'active_btn' : 'inactive_link'">
        Список замовлень</RouterLink></li>
      <li><RouterLink :to="{ name: 'import' }" :class="route.name === 'import' ? 'active_btn' : 'inactive_link'">
        Імпорт CSV</RouterLink></li>
    </ul>
  </nav>
</template>

<style scoped>
  nav ul{
    list-style-type: none;
    display: flex;
    flex-direction: row;
    align-items: center;
    column-gap: 78px;
  }
  .active_btn {
    text-decoration: none;
    padding: 8px 16px;
    border-radius: 8px;
    background: #E8F3EC;
    color: #5E9C76;
    font-size: 18px;
    font-weight: 400;
    cursor: default;
    box-shadow: 0 3px 5px rgba(0, 0, 0, 0.25);
    transition: all 0.2s ease-in-out;
    display: inline-block;
  }
  .inactive_link {
    text-decoration: none;
    text-shadow: 0 5px 3px rgba(0, 0, 0, 0.15);
    color: #000;
    font-size: 18px;
    font-weight: 400;
    cursor: pointer;
    transition: all 0.3s ease;
    padding: 8px 16px;
    display: inline-block;
  }
  .inactive_link:hover {
    transform: translateY(-3px);
    text-shadow: 0 15px 15px rgba(0, 0, 0, 0.1);
  }
</style>