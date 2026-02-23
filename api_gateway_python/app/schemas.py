from pydantic import BaseModel
from datetime import datetime
from typing import Optional

class OrderBase(BaseModel):
    latitude: float
    longitude: float
    subtotal: float

class OrderCreate(OrderBase):
    pass

class OrderResponse(OrderBase):
    id: int
    timestamp: datetime
    composite_tax_rate: float
    tax_amount: float
    total_amount: float
    state_rate: float
    county_rate: float
    city_rate: float
    special_rates: float
    jurisdictions: Optional[str]

    class Config:
        from_attributes = True