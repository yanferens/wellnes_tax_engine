<script setup lang="ts">
import { useOrdersStore } from '@/stores/ordersStore';
const store = useOrdersStore();

const columns = [
  { key: 'id', label: 'ID' },
  { key: 'subtotal', label: 'Subtotal' },
  { key: 'tax_amount', label: 'Tax Amount' },
  { key: 'total_amount', label: 'Total' },
  { key: 'composite_tax_rate', label: 'Tax Rate' },
];

const setSort = (key: string) => {
  const currentSort = store.filters.sortBy;
  const currentOrder = store.filters.sortOrder || 'asc';

  let newOrder = 'asc';
  if (currentSort === key) {
    newOrder = currentOrder === 'asc' ? 'desc' : 'asc';
  }

  store.setFilters({
    ...store.filters,
    sortBy: key,
    sortOrder: newOrder
  });
};
</script>

<template>
  <div class="filter-bar">
    <button
        v-for="col in columns"
        :key="col.key"
        @click="setSort(col.key)"
        :class="{ active: store.filters.sortBy === col.key }"
        class="filter-btn"
    >
      {{ col.label }}
      <span v-if="store.filters.sortBy === col.key">
        {{ store.filters.sortOrder === 'asc' ? '↑' : '↓' }}
      </span>
    </button>
  </div>
</template>

<style scoped>

</style>