from fastapi import FastAPI, UploadFile, File, Depends, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy.orm import Session
import pandas as pd
from sqlalchemy import text
import io
import uuid
from . import models, schemas, services, database


models.Base.metadata.create_all(bind=database.engine)

app = FastAPI(title="Wellness Tax Engine Gateway")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:5173"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.post("/orders")
async def create_order(order: schemas.OrderCreate):

    order_id = str(uuid.uuid4())


    services.send_order_to_queue(order_id, order.latitude, order.longitude, order.subtotal)

    return {
        "message": "Замовлення відправлено на розрахунок податків",
        "order_id": order_id
    }

@app.post("/orders/import")
async def import_orders(file: UploadFile = File(...)):
    if not file.filename.endswith('.csv'):
        raise HTTPException(status_code=400, detail="Тільки файли .csv дозволені")

    contents = await file.read()
    df = pd.read_csv(io.BytesIO(contents))

    count = 0
    for _, row in df.iterrows():
        order_id = str(uuid.uuid4())
        services.send_order_to_queue(
            order_id=order_id,
            lat=row['latitude'],
            lon=row['longitude'],
            subtotal=row['subtotal']
        )
        count += 1

    return {"message": f"Успішно відправлено {count} замовлень в чергу на обробку"}

@app.get("/orders")
def list_orders(page: int = 1, limit: int = 50, db: Session = Depends(database.get_db)):
    skip = (page - 1) * limit

    orders = db.query(models.Order).order_by(models.Order.timestamp.desc()).offset(skip).limit(limit).all()
    total = db.query(models.Order).count()

    return {
        "total": total,
        "page": page,
        "limit": limit,
        "data": orders
    }
