<script setup lang="ts">
import { ref, computed } from 'vue';
import { useOrdersStore } from '../stores/ordersStore';
import ErrorMessageBlock from '@/components/ErrorMessageBlock.vue';

defineProps<{
  isOpen: boolean
}>();

const emit = defineEmits<{
  (e: 'close'): void
}>();

const ordersStore = useOrdersStore();

interface FormState {
  latitude: number | null;
  longitude: number | null;
  subtotal: number | null;
}

const form = ref<FormState>({
  latitude: null,
  longitude: null,
  subtotal: null,
});

const errorMessage = ref<string>('');
const hasError = computed(() => errorMessage.value.length > 0);

const isProcessing = ref(false);

const handleError = (message: string) => {
  errorMessage.value = message;
};

const clearError = () => {
  errorMessage.value = '';
};

const submitForm = async () => {
  clearError();
  isProcessing.value = true;

  const orderData = {
    latitude: form.value.latitude ?? 0,
    longitude: form.value.longitude ?? 0,
    subtotal: form.value.subtotal ?? 0,
    timestamp: new Date().toISOString(),
  };

  try {
    const response = (await ordersStore.createOrder(orderData)) as any;

    if (ordersStore.errorMessage) {
      handleError(ordersStore.errorMessage);
      isProcessing.value = false;
      return;
    }

    const orderId = response?.order_id || response?.data?.order_id;

    if (!orderId) {
      throw new Error('Не вдалося отримати ID замовлення від сервера');
    }

    let isProcessed = false;
    let attempts = 0;
    const maxAttempts = 30;

    while (!isProcessed && attempts < maxAttempts) {
      attempts++;

      await ordersStore.fetchOrders();
      const order = ordersStore.orders.find((o: any) => o.id === orderId);

      if (order && order.total_amount !== null && order.total_amount !== undefined) {
        isProcessed = true;
      }

      if (!isProcessed) {
        await new Promise(resolve => setTimeout(resolve, 2000));
      }
    }

    if (!isProcessed) {
      throw new Error('Час очікування розрахунку вийшов. Спробуйте пізніше.');
    }

    closeModal();
    window.location.reload();

  } catch (e: any) {
    handleError(e?.message || 'Виникла помилка під час обробки замовлення');
    isProcessing.value = false;
  }
};

const closeModal = () => {
  emit('close');
  clearError();
  isProcessing.value = false;
};
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div v-if="isOpen" class="modal-overlay" @click.self="closeModal">
        <div class="modal-content">
          <form @submit.prevent="submitForm" class="create_order_form">
            <button class="close_btn" type="button" @click="closeModal">
              <img src="@/assets/close.png" alt="X">
            </button>
            <h2>Створити нове замовлення</h2>
            <div class="form_inputs">
              <div>
                <label>Широта:</label>
                <input type="number" placeholder="40.712776" step="any" v-model="form.latitude" required :disabled="isProcessing" />
              </div>
              <div>
                <label>Довгота:</label>
                <input type="number" placeholder="-74.005974" step="any" v-model="form.longitude" required :disabled="isProcessing" />
              </div>
              <div>
                <label>Сума без податку ($):</label>
                <input type="number" placeholder="100.07" step="0.01" v-model="form.subtotal" required :disabled="isProcessing" />
              </div>
            </div>
            <button class="submit_btn" type="submit" :disabled="isProcessing">
              {{ isProcessing ? 'Обробка замовлення...' : 'Створити замовлення' }}
            </button>
          </form>
        </div>
      </div>
    </Transition>
  </Teleport>
  <ErrorMessageBlock :show="hasError" :message="errorMessage" @close="clearError"/>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}
.modal-content {
  width: 534px;
  padding: 22px 23px 64px;
  border-radius: 16px;
  border: 1px solid rgba(94, 156, 118, 0.40);
  background: #FFF;
  box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.25);
  position: relative;
}
.create_order_form {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  height: 100%;
}
.close_btn {
  border: none;
  background: none;
  align-self: flex-end;
  margin-bottom: 11px;
  cursor: pointer;
  transition: 0.3s ease;
}
.close_btn:hover {
  transform: scale(1.1);
}
.close_btn img {
  width: 28px;
  height: 28px;
}
h2 {
  color: #000;
  font-size: 32px;
  font-weight: 400;
  margin-bottom: 30px;
  text-align: center;
}
.form_inputs {
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 16px;
  text-align: center;
}
.form_inputs div {
  display: flex;
  flex-direction: column;
  align-items: center;
  row-gap: 7px;
}
.form_inputs div label {
  color: #000;
  font-size: 20px;
}
.form_inputs div input {
  border-radius: 12px;
  border: 1px solid #9CA3AF;
  background: #F9FAFB;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  width: 319px;
  height: 46px;
  padding: 10px 22px;
  font-size: 18px;
  text-align: center;
}
.submit_btn {
  border-radius: 12px;
  background: #32D571;
  padding: 8px 16px;
  color: #FFFEFE;
  font-size: 24px;
  border: none;
  margin-top: 30px;
  cursor: pointer;
  transition: 0.3s ease;
}
.submit_btn:hover, .submit_btn:active {
  background: #4ae68a;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.15);
}
</style>

<style>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}
.modal-enter-active .modal-content {
  transition: all 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}

.modal-leave-active .modal-content {
  transition: all 0.2s cubic-bezier(0.25, 1, 0.5, 1);
}
.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
.modal-enter-from .modal-content,
.modal-leave-to .modal-content {
  opacity: 0;
  transform: scale(0.9) translateY(20px);
}
@media screen and (max-width: 600px){
  .modal-content{
    width: 100% !important;
    max-width: 340px;
    padding: 22px 23px 57px !important;
  }
  .modal-content form .close_btn img{
    width: 24px;
    height: 24px;
  }
  .modal-content form h2{
    font-size: 28px;
  }
  .modal-content form .form_inputs{
    width: 100%;
  }
  .modal-content form .form_inputs div{
    width: 100%;
  }
  .modal-content form .form_inputs div input{
    width: 100%;
  }
}
</style>