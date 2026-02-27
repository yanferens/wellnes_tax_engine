import axios, { type InternalAxiosRequestConfig } from 'axios';
import type { Order, OrderInput } from '../interfaces/order';

const apiClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL || 'http://134.122.82.46:8000',
    headers: {
        'Content-Type': 'application/json',
    },
});

// INTERCEPTOR
apiClient.interceptors.request.use((config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('auth_token');
    if (token) {
        if (!config.headers) {
            // @ts-ignore
            config.headers = {};
        }
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
}, (error) => {
    return Promise.reject(error);
});

// INTERCEPTOR
apiClient.interceptors.request.use((config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('auth_token');
    if (token) {
        if (!config.headers) {
            config.headers = {} as any;
        }
        config.headers.Authorization = 'Bearer ${token}';
    }
    return config;
}, (error) => {
    return Promise.reject(error);
});

// Interface for GET request parameters (pagination and filters)
export interface FetchOrdersParams {
    page?: number;
    limit?: number;
    search?: string;
    sort_by?: string;
    sort_order?: 'asc' | 'desc';
}

export const authService = {
    async login(email: string, password: string): Promise<{ access_token: string, token_type: string }> {
        const params = new URLSearchParams();
        params.append('username', email);
        params.append('password', password);
        const response = await apiClient.post('/token', params, {
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            }
        });
        return response.data;
    }
};

export const ordersService = {
    /**
     * POST /orders/import
     * Sends the CSV file to the server for processing
     */
    async importCsv(file: File): Promise<{ message: string }> {
        const formData = new FormData();
        formData.append('file', file);

        const response = await apiClient.post('/orders/import', formData, {
            headers: { 'Content-Type': 'multipart/form-data' },
        });
        return response.data;
    },

    /**
     * POST /orders
     * Sends data to create one order manually
     */
    async createOrder(orderData: OrderInput): Promise<Order> {
        const response = await apiClient.post('/orders', orderData);
        return response.data;
    },

    /**
     * GET /orders
     * Gets a list of orders with pagination and filters
     */
    async getOrders(params?: FetchOrdersParams): Promise<{ data: Order[]; total: number }> {
        const response = await apiClient.get('/orders', { params });
        return response.data;
    },
};