<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/authStore';

const router = useRouter();
const authStore = useAuthStore();

const email = ref('');
const password = ref('');

const handleLogin = async () => {
  const success = await authStore.login(email.value, password.value);

  if (success) {
    router.push({ name: 'orders' });
  }
};
</script>

<template>
  <section class="login-section">
    <div class="section_wrap">
      <img class="big_logo" src="../assets/logo.png" alt="logo">
      <h1>“Назва”<br>вхід адміністратора</h1>
      <h2>“Назва” податкові операції</h2>
      <form @submit.prevent="handleLogin">
        <div class="input_wrap">
          <label for="email">Email address</label>
          <input v-model="email" type="email" placeholder="youremail@email.com" name="email" required>
        </div>
        <div class="input_wrap">
          <label for="password">Password</label>
          <input v-model="password" type="password" placeholder="*********" name="password" required>
        </div>
        <p v-if="authStore.errorMessage" style="color: red; width: 100%;">
          {{ authStore.errorMessage }}
        </p>
        <button class="submit_btn" type="submit" :disabled="authStore.isLoading">
          {{ authStore.isLoading ? 'Вхід...' : 'Увійти на сайт' }}</button>
      </form>
    </div>
  </section>
</template>

<style scoped>
.login-section{
  margin: 170px 0 74px;
}
.login-section .section_wrap{
  max-width: 897px;
  border-radius: 12px;
  border: 1px solid #E5E7EB;
  background: #FFF;
  box-shadow: 0 4px 20px 0 rgba(0, 0, 0, 0.25);
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 100px 118px;
  row-gap: 24px;
}
.big_logo{
  width: 213px;
  height: 152px;
}
h1{
  color: #000;
  text-align: center;
  font-size: 48px;
  font-style: normal;
  font-weight: 600;
  line-height: normal;
}
h2{
  color: #989898;
  text-align: center;
  font-size: 30px;
  font-style: normal;
  font-weight: 600;
  line-height: normal;
  margin-bottom: 18px;
}
form{
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  row-gap: 24px;
}
.input_wrap{
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  row-gap: 24px;
}
.input_wrap label{
  color: #000;
  font-size: 24px;
  font-style: normal;
  font-weight: 600;
  line-height: normal;
}
.input_wrap input{
  width: 661px;
  height: 46px;
  padding: 10px 22px;
  border-radius: 12px;
  border: 1px solid #9CA3AF;
  background: #F9FAFB;
  box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.25);
  color: #9CA3AF;
  font-size: 18px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  text-align: center;
}
.input_wrap input::placeholder{
  color: #9CA3AF;
  font-size: 18px;
  font-style: normal;
  font-weight: 400;
  line-height: normal;
  text-align: center;
}
.submit_btn {
  border-radius: 12px;
  background: #32D571;
  padding: 8px 16px;
  color: #FFFEFE;
  font-size: 24px;
  border: none;
  cursor: pointer;
  transition: 0.3s ease;
  margin: 27px auto 0 auto;
}
.submit_btn:hover, .submit_btn:active {
  background: #4ae68a;
  box-shadow: 0 3px 5px rgba(0, 0, 0, 0.15);
}
.submit_btn:disabled {
  background: #ccc;
  color: #000000;
  cursor: not-allowed;
}
</style>