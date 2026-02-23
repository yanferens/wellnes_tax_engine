from fastapi import FastAPI, UploadFile, File, Depends, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy.orm import Session
import pandas as pd
import io
import uuid
from . import models, schemas, services, database


models.Base.metadata.create_all(bind=database.engine)

app = FastAPI(title="Wellness Tax Engine Gateway")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
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
@app.get("/orders/stats")
def get_order_stats(db: Session = Depends(database.get_db)):

    total_processed = db.query(models.Order).count()

    #
    ny_orders = db.query(models.Order).filter(
        models.Order.jurisdictions != "Out of State"
    ).count()


    out_of_state = total_processed - ny_orders


    avg_tax = 0
    if ny_orders > 0:
        from sqlalchemy import func
        avg_tax = db.query(func.avg(models.Order.composite_tax_rate)).filter(
            models.Order.jurisdictions != "Out of State"
        ).scalar()

    return {
        "summary": {
            "total_processed": total_processed,
            "ny_state_orders": ny_orders,
            "out_of_state_orders": out_of_state,
            "success_rate_percentage": round((ny_orders / total_processed * 100), 2) if total_processed > 0 else 0
        },
        "averages": {
            "avg_composite_rate": round(float(avg_tax), 5) if avg_tax else 0
        }
    }