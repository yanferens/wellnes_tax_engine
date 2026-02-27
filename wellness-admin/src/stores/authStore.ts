import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { authService } from '../api/orders';

export const useAuthStore = defineStore('auth', () => {
    // state
    const token = ref<string | null>(localStorage.getItem('auth_token'));
    const userEmail = ref<string | null>(localStorage.getItem('user_email'));
    const isLoading = ref(false);
    const errorMessage = ref<string | null>(null);

    // getters
    const isAuthenticated = computed(() => !!token.value);

    // actions
    // const login = async (email: string, password: string) => {
    //     isLoading.value = true;
    //     errorMessage.value = null;
    //
    //     try {
    //         // ІМІТАЦІЯ ЗАПИТУ (затримка 1 сек):
    //         await new Promise(resolve => setTimeout(resolve, 1000));
    //
    //         if (password === '12345678') { // Тимчасова перевірка для тесту
    //             const fakeToken = 'bearer-token-xyz';
    //
    //             // Зберігаємо дані
    //             token.value = fakeToken;
    //             userEmail.value = email;
    //             localStorage.setItem('auth_token', fakeToken);
    //             localStorage.setItem('user_email', email);
    //
    //             return true; // Успіх
    //         } else {
    //             throw new Error('Невірний логін або пароль');
    //         }
    //
    //     } catch (error: any) {
    //         errorMessage.value = error.message || 'Помилка входу';
    //         return false; // Невдача
    //     } finally {
    //         isLoading.value = false;
    //     }
    // };

    const login = async (email: string, password: string) => {
        isLoading.value = true;
        errorMessage.value = null;

        try {
            const data = await authService.login(email, password);
            const accessToken = data.access_token;
            token.value = accessToken;
            userEmail.value = email;
            localStorage.setItem('auth_token', accessToken);
            localStorage.setItem('user_email', email);

            return true;
        } catch (error: any) {
            if (error.response) {
                // Отримуємо повідомлення "Incorrect username or password" з FastAPI
                const serverMessage = error.response.data?.detail;

                if (error.response.status === 400 || error.response.status === 401) {
                    errorMessage.value = serverMessage || 'Невірний логін або пароль';
                } else if (error.response.status === 422) {
                    errorMessage.value = 'Помилка формату даних (422)';
                } else {
                    // ВИПРАВЛЕНО: додано зворотні лапки ``
                    errorMessage.value = `Помилка сервера: ${error.response.status}`;
                }
            } else {
                // Це спрацьовує при ERR_CONNECTION_REFUSED
                errorMessage.value = 'Сервер недоступний. Перевірте IP адресу в конфігурації.';
            }
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