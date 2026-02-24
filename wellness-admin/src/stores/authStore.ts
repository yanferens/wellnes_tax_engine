import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';

export const useAuthStore = defineStore('auth', () => {
    const router = useRouter();

    // state
    const token = ref<string | null>(localStorage.getItem('auth_token'));
    const userEmail = ref<string | null>(localStorage.getItem('user_email'));
    const isLoading = ref(false);
    const errorMessage = ref<string | null>(null);

    // getters
    const isAuthenticated = computed(() => !!token.value);

    // actions
    const login = async (email: string, password: string) => {
        isLoading.value = true;
        errorMessage.value = null;

        try {
            // THERE WILL BE A REQUEST FOR THE BACKEND
            // const response = await axios.post('/api/login', { email, password });

            // REQUEST SIMULATION:
            await new Promise(resolve => setTimeout(resolve, 1000));

            if (password === '12345678') {
                const fakeToken = 'bearer-token-xyz';
                token.value = fakeToken;
                userEmail.value = email;
                localStorage.setItem('auth_token', fakeToken);
                localStorage.setItem('user_email', email);
                return true;
            } else {
                throw new Error('Невірний логін або пароль');
            }

        } catch (error: any) {
            errorMessage.value = error.message || 'Помилка входу';
            return false;
        } finally {
            isLoading.value = false;
        }
    };

    const logout = () => {
        token.value = null;
        userEmail.value = null;
        localStorage.removeItem('auth_token');
        localStorage.removeItem('user_email');
        window.location.href = '/login';
    };

    return { token, userEmail, isAuthenticated, isLoading, errorMessage, login, logout };
});