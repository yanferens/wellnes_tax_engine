<script setup lang="ts">
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css';
import { useOrdersStore } from '@/stores/ordersStore';
import {
  DynamicScroller as _DynamicScroller,
  DynamicScrollerItem as _DynamicScrollerItem
} from 'vue-virtual-scroller';
import 'vue-virtual-scroller/dist/vue-virtual-scroller.css';

const store = useOrdersStore();

const DynamicScroller = _DynamicScroller as any;
const DynamicScrollerItem = _DynamicScrollerItem as any;

const setSort = async (key: string) => {
  const currentSort = store.filters.sort_by;
  const currentOrder = store.filters.sort_order || 'asc';

  let newOrder = 'asc';
  if (currentSort === key) {
    newOrder = currentOrder === 'asc' ? 'desc' : 'asc';
  }

  store.setFilters({
    ...store.filters,
    sort_by: key,
    sort_order: newOrder
  });
};

const toggle = (item: any) => {
  item.isOpen = !item.isOpen;
}

const isHighlighted = (item: any) => {
  const search = store.filters.search;
  if (!search) return false;

  const searchLower = search.toLowerCase();

  const parts = searchLower.split(',').map((p: string) => p.trim());

  if (parts.length === 3) {
    const [sId, sLat, sLon] = parts;

    const matchesId = String(item.id).toLowerCase().includes(sId);
    const matchesLat = String(item.latitude).startsWith(sLat); // startsWith точніше для координат
    const matchesLon = String(item.longitude).startsWith(sLon);

    return matchesId && matchesLat && matchesLon;
  }

  if (parts.length === 2) {
    const [sLat, sLon] = parts;

    const matchesLat = String(item.latitude).startsWith(sLat);
    const matchesLon = String(item.longitude).startsWith(sLon);

    return matchesLat && matchesLon;
  }

  return String(item.id).toLowerCase().includes(searchLower) ||
      String(item.latitude).includes(searchLower) ||
      String(item.longitude).includes(searchLower);
}

const columns = [
  { key: 'id', label: 'ID' },
  { key: 'subtotal', label: 'Subtotal' },
  { key: 'tax_amount', label: 'Total Amount' },
  { key: 'jurisdictions', label: 'Jurisdictions' },
];
</script>

<template>
  <div class="table-wrapper">
    <div v-if="store.isLoading" class="loading">Завантаження даних... ⏳</div>
    <div v-else-if="store.errorMessage" class="error">{{ store.errorMessage }}</div>

    <div v-else class="custom-table">
      <div class="table-header">
        <div v-for="col in columns"
             :key="col.key"
             @click="setSort(col.key)"
             :class="{ active: store.filters.sort_by === col.key }"
             class="header-cell"
        >
          {{ col.label }}
          <span v-if="store.filters.sort_by === col.key">
            {{ store.filters.sort_order === 'asc' ? '↑' : '↓' }}
          </span>
        </div>
      </div>

      <DynamicScroller
          class="scroller"
          :items="store.orders"
          :min-item-size="60"
          key-field="id"
      >
        <template v-slot="{ item, index, active }">
          <DynamicScrollerItem
              :item="item"
              :active="active"
              :size-dependencies="[item.isOpen]"
              :data-index="index"
          >
            <div class="table-row-container">

              <div class="table-row main-info" @click="toggle(item)" :class="{ 'is-open': item.isOpen,
              'is-highlighted': isHighlighted(item) }">
                <div class="cell id-cell">
                  <span class="arrow" :class="{ rotated: item.isOpen }">▼</span>
                  {{ item.id }}
                </div>
                <div class="cell">${{ item.subtotal }}</div>
                <div class="cell">${{ item.total_amount }}</div>
                <div class="cell">{{ item.jurisdictions }}</div>
              </div>

              <div class="accordion-wrapper" :class="{ 'is-open': item.isOpen }">
                <div class="accordion-inner">
                  <div class="row-details">
                    <div class="cell">
                      <span>Tax Rate (Composite)</span>
                      {{item.composite_tax_rate}}
                    </div>
                    <div class="cell">
                      <span>Tax Amount</span>
                      {{item.tax_amount}}
                    </div>
                    <div class="cell">
                      <span>State_r</span>
                      {{item.state_rate}}
                    </div>
                    <div class="cell">
                      <span>Country_r</span>
                      {{item.county_rate}}
                    </div>
                    <div class="cell">
                      <span>City_r</span>
                      {{item.city_rate}}
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </DynamicScrollerItem>
        </template>
      </DynamicScroller>

      <div v-if="store.orders.length === 0" class="empty-state">
        Даних не знайдено
      </div>

      <div class="pagination">
        <button :disabled="store.currentPage === 1" @click="store.setPage(store.currentPage - 1)">
          <svg xmlns="http://www.w3.org/2000/svg" width="41" height="15" viewBox="0 0 41 15" fill="none">
            <path d="M0.292892 8.07112C-0.0976295 7.6806 -0.0976295 7.04743 0.292892 6.65691L6.65685 0.292946C7.04738 -0.0975785 7.68054 -0.0975785 8.07107 0.292946C8.46159 0.68347 8.46159 1.31664 8.07107 1.70716L2.41422 7.36401L8.07107 13.0209C8.46159 13.4114 8.46159 14.0446 8.07107 14.4351C7.68054 14.8256 7.04738 14.8256 6.65685 14.4351L0.292892 8.07112ZM41 7.36401V8.36401H1V7.36401V6.36401H41V7.36401Z" fill="black"/>
          </svg>
        </button>
        <span>Сторінка {{ store.currentPage }} з {{ Math.ceil(store.totalOrders / store.limit) || 1 }}</span>
        <button :disabled="store.currentPage >= Math.ceil(store.totalOrders / store.limit)" @click="store.setPage(store.currentPage + 1)">
          <svg xmlns="http://www.w3.org/2000/svg" width="41" height="15" viewBox="0 0 41 15" fill="none">
            <path d="M40.7071 8.07112C41.0976 7.6806 41.0976 7.04743 40.7071 6.65691L34.3431 0.292946C33.9526 -0.0975785 33.3195 -0.0975785 32.9289 0.292946C32.5384 0.68347 32.5384 1.31664 32.9289 1.70716L38.5858 7.36401L32.9289 13.0209C32.5384 13.4114 32.5384 14.0446 32.9289 14.4351C33.3195 14.8256 33.9526 14.8256 34.3431 14.4351L40.7071 8.07112ZM0 7.36401V8.36401H40V7.36401V6.36401H0V7.36401Z" fill="black"/>
          </svg>
        </button>
      </div>

    </div>
  </div>
</template>

<style scoped>
.custom-table {
  width: 100%;
  border-radius: 12px;
  border: 1px solid #E5E7EB;
  background: #FFF;
  box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.05);
  overflow: hidden;
}
.table-header, .table-row {
  display: grid;
  grid-template-columns: 351.55px 241.22px 320.73px 286.5px;
  align-items: center;
  padding: 0 20px;
}
.table-header {
  border-bottom: 1px solid #E5E7EB;
  padding-top: 15px;
  padding-bottom: 15px;
}
.header-cell {
  color: #6B7280;
  font-size: 13px;
  font-weight: 700;
  text-align: left;
  cursor: pointer;
}
.scroller {
  height: 675px;
  overflow-y: auto;
}
.table-row-container {
  border-bottom: 1px solid #F3F4F6;
}
.table-row:hover, .table-row:hover, .table-row.is-open, .table-row.is-highlighted{
  background: #F0FDF4;
}
.table-row.main-info {
  cursor: pointer;
  min-height: 50px;
}
.cell {
  color: #1F2937;
  font-size: 13.9px;
  font-weight: 400;
  padding: 16.5px 0;
  text-align: left;
}
.cell.id-cell,
.cell:nth-child(3) {
  font-size: 13.5px;
  font-weight: 700;
}
.cell.id-cell {
  display: flex;
  align-items: center;
  gap: 13px;
}
.arrow {
  color: #9CA3AF;
  font-size: 10px;
  font-weight: 700;
  flex-shrink: 0;
  transition: transform 0.3s ease;
  transform: rotate(-90deg);
}
.arrow.rotated {
  transform: rotate(0deg);
}
.accordion-wrapper {
  display: grid;
  grid-template-rows: 0fr;
  transition: grid-template-rows 0.3s ease-out;
}
.accordion-wrapper.is-open {
  grid-template-rows: 1fr;
}
.accordion-inner {
  overflow: hidden;
  min-height: 0;
}
.row-details {
  background-color: #FAFAFA;
  padding: 20px 60px;
  display: grid;
  grid-template-columns: repeat(5, 1fr);
  align-items: center;
}
.row-details .cell{
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  row-gap: 4px;
  color: #1F2937;
  font-size: 13px;
  font-style: normal;
  font-weight: 500;
  line-height: normal;
}
.row-details .cell span{
  color: #9CA3AF;
  font-size: 11px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  letter-spacing: 0.22px;
  text-transform: uppercase;
}
.empty-state {
  text-align: center;
  padding: 40px;
  color: #6B7280;
}
.pagination {
  padding: 38.5px 0;
  display: flex;
  flex-direction: row;
  gap: 23px;
  align-items: center;
  justify-content: center;
  border-top: 1px solid #E5E7EB;
}
.pagination button {
  border: none;
  background: none;
  width: 40px;
  flex-shrink: 0;
  cursor: pointer;
}
.pagination button:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}
.pagination span {
  color: #000;
  font-size: 18px;
}
.error { color: red; text-align: center;}
.loading { text-align: center; padding: 20px; }
</style>