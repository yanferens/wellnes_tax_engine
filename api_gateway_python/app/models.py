from sqlalchemy import Column, Float, String, DateTime, Integer
from .database import Base
import datetime

class Order(Base):
    __tablename__ = "orders"


    id = Column(String, primary_key=True, index=True)
    latitude = Column(Float, nullable=False)
    longitude = Column(Float, nullable=False)
    subtotal = Column(Float, nullable=False)
    timestamp = Column(DateTime, default=datetime.datetime.utcnow)

    composite_tax_rate = Column(Float)
    tax_amount = Column(Float)
    total_amount = Column(Float)

    state_rate = Column(Float)
    county_rate = Column(Float)
    city_rate = Column(Float)
    special_rates = Column(Float)

    jurisdictions = Column(String, nullable=True)

class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True, nullable=False)
    hashed_password = Column(String, nullable=False)