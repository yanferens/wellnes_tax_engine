import { defineStore } from 'pinia';
import { ordersService, type FetchOrdersParams } from '../api/orders';
import type { Order, OrderInput } from '../interfaces/order';

export const useOrdersStore = defineStore('orders', {
    // all data is stored here
    state: () => ({
        orders: [] as Order[],
        totalOrders: 0,
        currentPage: 1,
        limit: 10,
        filters: {} as Record<string, any>,
        isLoading: false,
        errorMessage: null as string | null,
    }),

    // functions for updating status and calling API
    actions: {
        // getting a list of orders
        async fetchOrders() {
            this.isLoading = true;
            this.errorMessage = null;
            try {
                const params: FetchOrdersParams = {
                    page: this.currentPage,
                    limit: this.limit,
                    ...this.filters,
                };
                const response = await ordersService.getOrders(params);
                this.orders = response.data;
                this.totalOrders = response.total;
            } catch (error) {
                this.errorMessage = 'Помилка під час завантаження замовлень';
                console.error(error);
            } finally {
                this.isLoading = false;
            }
        },

        // upload CSV file
        async uploadCsv(file: File) {
            this.isLoading = true;
            this.errorMessage = null;
            try {
                await ordersService.importCsv(file);
                this.currentPage = 1;
                await this.fetchOrders();
            } catch (error) {
                this.errorMessage = 'Помилка під час обробки CSV-файлу';
                console.error(error);
                throw error;
            } finally {
                this.isLoading = false;
            }
        },

        // manual order creation
        async createOrder(orderData: OrderInput) {
            this.isLoading = true;
            this.errorMessage = null;
            try {
                await ordersService.createOrder(orderData);

                this.currentPage = 1;
                await this.fetchOrders();

            } catch (error) {
                this.errorMessage = 'Помилка під час створення замовлення';
                console.error(error);
                throw error;
            } finally {
                this.isLoading = false;
            }
        },

        // update pagination settings
        setPage(page: number) {
            this.currentPage = page;
            this.fetchOrders();
        },

        // update filters
        setFilters(newFilters: Record<string, any>) {
            this.filters = newFilters;
            this.currentPage = 1;
            this.fetchOrders();
        }
    },
});