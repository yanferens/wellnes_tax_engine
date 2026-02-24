export interface OrderInput {
    latitude: number;
    longitude: number;
    subtotal: number;
    timestamp: string;
}

export interface TaxBreakdown {
    state_rate: number;
    county_rate: number;
    city_rate: number;
    special_rates: number;
    jurisdictions?: string[];
}

export interface Order extends OrderInput {
    id: string;
    composite_tax_rate: number;
    tax_amount: number;
    total_amount: number;
    breakdown: TaxBreakdown;
}