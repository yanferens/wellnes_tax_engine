from fastapi import FastAPI, UploadFile, File, Depends, HTTPException
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
from fastapi.middleware.cors import CORSMiddleware
from sqlalchemy.orm import Session
import pandas as pd
from sqlalchemy import text, desc, asc
import io
import uuid
from . import models, schemas, services, database, auth


models.Base.metadata.create_all(bind=database.engine)

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")

app = FastAPI(title="Wellness Tax Engine Gateway")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

def get_current_user(token: str = Depends(oauth2_scheme)):
    try:
        payload = auth.jwt.decode(token, auth.SECRET_KEY, algorithms=[auth.ALGORITHM])
        username: str = payload.get("sub")
        if username is None:
            raise HTTPException(status_code=401, detail="Could not validate credentials")
        return username
    except auth.JWTError:
        raise HTTPException(status_code=401, detail="Invalid token")

@app.post("/token")
async def login(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(database.get_db)):
    user = db.query(models.User).filter(models.User.username == form_data.username).first()
    if not user or not auth.verify_password(form_data.password, user.hashed_password):
        raise HTTPException(status_code=400, detail="Incorrect username or password")
    
    access_token = auth.create_access_token(data={"sub": user.username})
    return {"access_token": access_token, "token_type": "bearer"}

@app.post("/orders")
async def create_order(order: schemas.OrderCreate, current_user: str = Depends(get_current_user)):

    order_id = str(uuid.uuid4())


    services.send_order_to_queue(order_id, order.latitude, order.longitude, order.subtotal)

    return {
        "message": "Замовлення відправлено на розрахунок податків",
        "order_id": order_id
    }

@app.post("/orders/import")
async def import_orders(file: UploadFile = File(...), current_user: str = Depends(get_current_user)):
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
def list_orders(
    page: int = 1, 
    limit: int = 50, 
    sort_by: str = "timestamp",
    sort_order: str = "desc",
    db: Session = Depends(database.get_db), 
    current_user: str = Depends(get_current_user)
):
    skip = (page - 1) * limit

    sort_columns = {
        "id": models.Order.id,
        "subtotal": models.Order.subtotal,
        "tax_amount": models.Order.total_amount,
        "jurisdictions": models.Order.jurisdictions,
        "timestamp": models.Order.timestamp
    }

    sort_column = sort_columns.get(sort_by, models.Order.timestamp)

    query = db.query(models.Order)

    if sort_order == "asc":
        query = query.order_by(asc(sort_column))
    else:
        query = query.order_by(desc(sort_column))

    orders = query.offset(skip).limit(limit).all()
    total = db.query(models.Order).count()

    return {
        "total": total,
        "page": page,
        "limit": limit,
        "data": orders
    }
