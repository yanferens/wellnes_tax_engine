// declare module 'vue-virtual-scroller' {
//     import { DefineComponent } from 'vue';
//
//     interface RecycleScrollerProps {
//         items: readonly any[];
//         direction?: 'vertical' | 'horizontal';
//         itemSize?: number | null;
//         gridItems?: number;
//         itemSecondarySize?: number;
//         minItemSize?: number;
//         sizeField?: string;
//         typeField?: string;
//         keyField?: string;
//         pageMode?: boolean;
//         prerender?: number;
//         buffer?: number;
//         emitUpdate?: boolean;
//         [key: string]: any;
//     }
//
//     interface DynamicScrollerProps extends RecycleScrollerProps {
//         minItemSize: number;
//         [key: string]: any;
//     }
//
//     interface DynamicScrollerItemProps {
//         item: any;
//         active: boolean;
//         sizeDependencies?: any[];
//         watchData?: boolean;
//         tag?: string;
//         emitResize?: boolean;
//         [key: string]: any;
//     }
//
//
//     export const RecycleScroller: DefineComponent<
//         RecycleScrollerProps,
//         {},
//         any
//     >;
//
//     export const DynamicScroller: DefineComponent<
//         DynamicScrollerProps,
//         {},
//         any
//     >;
//
//     export const DynamicScrollerItem: DefineComponent<
//         DynamicScrollerItemProps,
//         {},
//         any
//     >;
//
//     export const IdState: DefineComponent<{}, {}, any>;
// }