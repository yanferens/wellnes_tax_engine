import axios from 'axios';
import type { Order, OrderInput } from '../interfaces/order';

const apiClient = axios.create({
    // Обов'язково додаємо протокол http://
    baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8000',
    headers: {
        'Content-Type': 'application/json',
    },
});

// Interface for GET request parameters (pagination and filters)
export interface FetchOrdersParams {
    page?: number;
    limit?: number;
}

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