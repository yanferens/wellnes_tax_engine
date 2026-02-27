<script setup lang="ts">
defineProps<{
  show: boolean;
  message: string;
}>();

const emit = defineEmits<{
  (e: 'close'): void;
}>();

const closeModal = () => {
  emit('close');
};
</script>

<template>
  <Transition name="modal">
    <div v-if="show" class="modal-overlay" @click.self="closeModal">
      <div class="error-message">
        <h3>Помилка запиту</h3>
        <img src="@/assets/error.svg" alt="error"/>
        <p class="error-text">{{ message }}</p>
        <button @click="closeModal" class="clear-error-btn">Спробувати знову</button>
      </div>
    </div>
  </Transition>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 9999;
}
.error-message {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: #FFF;
  padding: 40px;
  border-radius: 16px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  width: 90%;
  max-width: 450px;
  row-gap: 20px;
}
.error-message h3{
  color: #1A1A1A;
  text-align: center;
  font-size: 27px;
  font-style: normal;
  font-weight: 500;
  line-height: normal;
  margin-bottom: 10px;
}
.error-message img{
  width: 66px;
  height: 66px;
}
.error-text {
  color: #757575;
  text-align: center;
  font-size: 15px;
  font-style: normal;
  font-weight: 400;
  line-height: 24px;
  margin: 0;
}
.error-message button{
  width: 100%;
  padding: 16px 0;
  text-align: center;
  border-radius: 12px;
  background: #FF4D4D;
  color: #FFF;
  font-size: 20px;
  font-style: normal;
  font-weight: 500;
  line-height: normal;
  border: none;
  margin-top: 10px;
  cursor: pointer;
  transition: all 0.3s ease;
}
.error-message button:hover, .error-message button:active{
  background: #df4242;
}
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
.modal-enter-active .error-message,
.modal-leave-active .error-message {
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
}
.modal-enter-from .error-message,
.modal-leave-to .error-message {
  transform: translateY(30px) scale(0.9);
  opacity: 0;
}
</style>