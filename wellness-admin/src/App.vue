<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <nav style="display: flex; gap: 20px; padding: 20px; background: #f0f0f0; margin-bottom: 20px;">
      <RouterLink to="/">Список замовлень</RouterLink>
      <RouterLink to="/import">Імпорт CSV</RouterLink>
      <RouterLink to="/create">Створити вручну</RouterLink>
    </nav>
  </header>

  <main style="padding: 20px;">
    <RouterView />
  </main>
</template>