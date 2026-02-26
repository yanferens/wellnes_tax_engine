/// <reference types="vite/client" />
declare module 'vue-virtual-scroller' {
    import type { DefineComponent } from 'vue';

    // <any, any, any> критично важливо, щоб дозволити слоти
    export const DynamicScroller: DefineComponent<any, any, any>;
    export const DynamicScrollerItem: DefineComponent<any, any, any>;
    export const RecycleScroller: DefineComponent<any, any, any>;
}