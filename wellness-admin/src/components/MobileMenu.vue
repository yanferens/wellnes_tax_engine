<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';
import { useRoute } from 'vue-router';
import CreateOrderExit from "@/components/CreateOrderExit.vue";

const route = useRoute();

const isOpen = ref(false);
const menuRef = ref<HTMLElement | null>(null);
const buttonRef = ref<HTMLElement | null>(null);

const toggleMenu = () => {
  isOpen.value = !isOpen.value;
  updateHeaderClass();
};

const closeMenu = () => {
  if (isOpen.value) {
    isOpen.value = false;
    updateHeaderClass();
  }
};

const updateHeaderClass = () => {
  const header = document.querySelector('header');
  if (header) {
    if (isOpen.value) {
      header.classList.add('menu_active');
    } else {
      header.classList.remove('menu_active');
    }
  }
};

const handleClickOutside = (event: MouseEvent) => {
  const target = event.target as Node;
  if (
      menuRef.value &&
      !menuRef.value.contains(target) &&
      buttonRef.value &&
      !buttonRef.value.contains(target)
  ) {
    closeMenu();
  }
};

const handleLinkClick = () => {
  closeMenu();
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
  const header = document.querySelector('header');
  header?.classList.remove('menu_active');
});
</script>

<template>
  <nav class="menu" :class="{ active: isOpen }" ref="menuRef">
    <ul>
      <li @click="handleLinkClick" :class="route.name === 'orders' ? 'active_btn' : 'inactive_link'">
        <RouterLink :to="{ name: 'orders' }">
        Список замовлень</RouterLink>
      </li>
      <li @click="handleLinkClick" :class="route.name === 'import' ? 'active_btn' : 'inactive_link'">
        <RouterLink :to="{ name: 'import' }">
        Імпорт CSV</RouterLink>
      </li>
      <li>
        <CreateOrderExit/>
      </li>
    </ul>
  </nav>

  <button class="toggle_menu mobile" :class="{ active: isOpen }" @click="toggleMenu" ref="buttonRef">
    <span class="sandwich">
      <span class="sw-topper"></span>
      <span class="sw-bottom"></span>
      <span class="sw-footer"></span>
    </span>
  </button>
</template>

<style scoped>
.toggle_menu {
  width: 32px;
  height: 32px;
  z-index: 110;
  border: none;
  right: 0;
  margin-left: 10px;
  top: 0;
  padding: 0;
  cursor: pointer;
  border-radius: 50%;
  position: relative;
  -webkit-animation-delay: 1s;
  animation-delay: 1s;
  background-color: transparent
}
.toggle_menu:focus {
  outline: 0
}
.toggle_menu span {
  display: block
}
.sandwich {
  width: 28px;
  height: 24px;
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  margin: auto;
  z-index: 200
}
.sw-topper {
  position: relative;
  top: 0;
  width: 28px;
  height: 4px;
  background: #000000;
  border: none;
  -webkit-transition: -webkit-transform .5s, top .2s;
  -webkit-transition: top .2s, -webkit-transform .5s;
  -o-transition: transform .5s, top .2s;
  transition: transform .5s, top .2s, -webkit-transform .5s;
  display: block;
  border-radius: 3px
}
.sw-bottom {
  position: relative;
  width: 28px;
  height: 4px;
  left: 0;
  top: 6px;
  background: #000000;
  border: none;
  -webkit-transition: -webkit-transform .5s, top .2s;
  -webkit-transition: top .2s, -webkit-transform .5s;
  -o-transition: transform .5s, top .2s;
  transition: transform .5s, top .2s, -webkit-transform .5s;
  border-radius: 3px
}
.sw-footer {
  position: relative;
  width: 28px;
  height: 4px;
  left: 0;
  top: 12px;
  background: #000000;
  border: none;
  -webkit-transition: all .5s;
  -o-transition: all .5s;
  transition: all .5s;
  -webkit-transition-delay: .1s;
  -o-transition-delay: .1s;
  transition-delay: .1s;
  border-radius: 3px
}
.toggle_menu.active {
  border: 2px solid #000000;
}
.toggle_menu.active .sandwich .sw-topper {
  top: 11px;
  left: 5px;
  width: 18px;
  height: 2px;
  -webkit-transform: rotate(-45deg);
  -ms-transform: rotate(-45deg);
  transform: rotate(-45deg);
  -webkit-clip-path: polygon(100% 0, 0 0, 0 100%, 100% 100%);
  clip-path: polygon(100% 0, 0 0, 0 100%, 100% 100%)
}
.toggle_menu.active .sandwich .sw-footer {
  top: 7px;
  left: 5px;
  height: 2px;
  width: 18px;
  -webkit-transform: rotate(45deg);
  -ms-transform: rotate(45deg);
  transform: rotate(45deg);
  -webkit-clip-path: polygon(100% 0, 0 0, 0 100%, 100% 100%);
  clip-path: polygon(100% 0, 0 0, 0 100%, 100% 100%)
}
.toggle_menu.active .sandwich .sw-bottom {
  opacity: 0;
  top: 0;
  height: 2px;
  -webkit-transform: rotate(180deg);
  -ms-transform: rotate(180deg);
  transform: rotate(180deg);
  -webkit-clip-path: polygon(100% 0, 0 0, 0 100%, 100% 100%);
  clip-path: polygon(100% 0, 0 0, 0 100%, 100% 100%)
}
.menu.active {
  top: 96px;
  background-color: white;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: flex-start;
}
nav.menu{
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(4px);
  padding: 18px 20px;
  position: absolute;
  top: -120vh;
  right: 0;
  left: 0;
  max-width: 100%;
  z-index: 20;
  overflow-y: auto;
  text-align: center;
  -webkit-transition: all .5s;
  -o-transition: all .5s;
  transition: all .5s;
  justify-content: flex-start;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.05);
  border-top: 1px solid #E5E7EB;
}
nav{
  display: none;
}
nav.menu ul{
  width: 100%;
  list-style-type: none;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}
.active_btn {
  display: none;
}
.inactive_link a{
  display: inline-block;
  text-decoration: none;
  text-shadow: 0 5px 3px rgba(0, 0, 0, 0.15);
  color: #000;
  font-size: 16px;
  font-weight: 400;
  cursor: pointer;
  transition: all 0.3s ease;
}
.inactive_link:hover a,  .inactive_link:active a{
  color: #5E9C76;
}
</style>