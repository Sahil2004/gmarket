# 📈 GMarket

GMarket is a real-time stock watchlist and market monitoring platform built with **Angular** (frontend) and **Go (Fiber)** (backend). It provides live stock price updates, market depth visualization, watchlist management, and scalable service-driven architecture.

---

## 🚀 Features

### 📊 Market Monitoring
- Real-time LTP (Last Traded Price) updates
- Market depth visualization (bids & asks)
- Spread calculation
- Efficient viewport-based data fetching

### ⭐ Watchlist Management
- Create & switch between watchlists
- Add / remove stocks
- Persist watchlists per user

### 🔄 High Performance Rendering
- Angular CDK Virtual Scroll
- Fetches only visible stock data
- Optimized polling strategy
- Signal-based reactive UI

### 🔐 Authentication & Session Handling
- Secure login/logout
- Session token management
- Cookie-based auth support

### 🖼 Media Services
- Image upload support
- Profile avatar storage


## 🧠 Design Principles

- Single DB connection pool
- Dependency injection via constructors
- Viewport-driven data fetching
- Separation of concerns
- Stateless controllers
- Signal-based UI reactivity
- Service-driven backend architecture

---

## 🧰 Tech Stack

### Frontend
- Angular
- Angular Signals
- Angular CDK Virtual Scroll
- Angular Material
- TailwindCSS

### Backend
- Go
- Fiber
- PostgreSQL
- database/sql
- JWT / Session tokens

### Infrastructure
- Docker
- REST APIs
- Polling-based real-time updates

---

## ⚙️ Setup Instructions

---

### 🖥 Backend Setup

#### 1. Clone repository
```bash
git clone https://github.com/Sahil2004/gmarket.git
cd gmarket/server
```
#### 2. Configure environment variables
Create .env from example.env
#### 3. To run the complete backend and frontend with database:
The database is setup in the docker compose itself.
```bash
docker compose up
```
To run for development:
```bash
docker compose watch
```
and in another terminal you can view the logs:
```bash
docker compose logs -f
```
#### 3. To run only backend
```bash
go mod tidy
go run main.go
```
#### Server runs on:
```bash
http://localhost:3000
```
### To run only frontend
```bash
cd client
npm install
npm start
```
#### Frontend runs on:
```bash
http://localhost:4200
```

---

## License
This project is under GPL-v2. You can view it at [LICENSE](./LICENSE).

## Author

For any information or just to say hi, you can contact me at: `me.sahil.gg@gmail.com`.
