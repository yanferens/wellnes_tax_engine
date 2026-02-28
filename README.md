# Instant Wellness Kits - Tax Engine (BetterMe Test Task)

## 📌 Business Problem & Solution

**The Challenge:** Sales tax in NYS is highly dependent on precise locations (State + County + City + Special Jurisdictions like MCTD). A simple bounding box approach is inaccurate.

**The Solution:** This project solves the business problem by utilizing **PostGIS** for precise spatial queries (`ST_Contains`). By importing official NYS shapefiles, the system maps any given latitude/longitude directly to its corresponding tax jurisdictions and calculates the exact `composite_tax_rate` and `tax_amount`.

---

## 🌐 Live Deployment & Demo
The application has been deployed and is ready for testing.
* **URL:** [http://134.122.82.46:5173](http://134.122.82.46:5173)
* **Admin Credentials**:
    * **Login:** `admin@gmail.com`
    * **Password:** `admin`

---

## 🏗️ Architecture & Tech Stack
We chose a microservices approach optimized for heavy spatial data processing and high-throughput CSV ingestion:
* **Backend API (Python/FastAPI):** Handles REST endpoints for the frontend.
* **Data Processor (Go Worker):** Parses the large CSV file concurrently and performs heavy spatial intersections via PostGIS.
* **Database (PostgreSQL + PostGIS):** Stores orders and handles geometric spatial queries.
* **Frontend (Vue.js):** A responsive SPA fulfilling the admin panel requirements (Import, Manual Create, Orders List with pagination/filters).

---

## 🚀 Local Setup & Installation

Follow these steps to run the application locally.

### 1. Clone the repository
```bash
git clone https://github.com/yanferens/wellnes_tax_engine.git
cd wellnes_tax_engine
```
### 2. Configure Environment Variables
The project uses a .env file to manage database connections, secrets, and service URLs. A template is provided as .env.example.
```bash
cp .env.example .env
```
Ensure your .env file contains the following default values (optimized for the Docker Compose network):
```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=tax_db
DATABASE_URL=postgresql://postgres:postgres@db:5432/tax_db?sslmode=disable
REDIS_URL=redis:6379
JWT_SECRET=your_super_secret_key
```
### 3. Launch with Docker Compose
The entire infrastructure is fully containerized. You can launch all services (API, Go Worker, Database, Redis, and Frontend) with a single command:
```bash
docker compose up -d --build
```