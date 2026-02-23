import json
import redis
import os
from dotenv import load_dotenv
from datetime import datetime

load_dotenv()


redis_client = redis.Redis.from_url(os.getenv("REDIS_URL", "redis://localhost:6379/0"))

def send_order_to_queue(order_id: str, lat: float, lon: float, subtotal: float):

    task_data = {
        "order_id": order_id,
        "latitude": lat,
        "longitude": lon,
        "subtotal": subtotal,
        "timestamp": datetime.utcnow().isoformat()
    }


    redis_client.lpush("orders_queue", json.dumps(task_data))