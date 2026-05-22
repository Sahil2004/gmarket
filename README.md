# 📈 GMarket

GMarket is a real-time stock watchlist, charting, and trading platform built with **Angular** on the frontend and **Go (Fiber)** on the backend. It combines live market data, technical indicators, portfolio and order management, and account tooling in a single service-driven app.

---

## 🚀 Features

### 📊 Market Data

- Real-time LTP updates
- Market depth visualization with bids, asks, and spread
- Candle data across multiple ranges
- Symbol status and live polling
- Viewport-aware market fetching to avoid unnecessary requests

### 📈 Charting

- Candlestick chart for selected symbols
- Range switching across 1D, 1W, 1M, and 1Y views
- Moving average overlays on the price chart
- Fullscreen chart mode for focused analysis
- Responsive layout that moves the strategy panel below the chart on smaller screens

### 🧠 Algo Strategies

- RSI strategy configuration
- Moving average crossover configuration
- Combined RSI + MA signal evaluation
- Live indicator preview with current RSI and MA values
- Save and reload strategy settings per symbol and exchange

### ⭐ Watchlists

- Create and switch between watchlists
- Add and remove symbols from a watchlist
- Persist watchlists per user
- Watchlist-aware polling so active charts do not fight background refreshes

### 💼 Trading and Account Management

- Portfolio and positions view
- Orders screen for order history and management
- Order preview, place order, and cancel order actions
- Funds management with deposit and withdraw flows
- Bank account management for linked funding sources

### 👤 Authentication and Profile

- Register, login, and logout flows
- Session-based authentication
- Protected routes for authenticated users
- Profile updates and password changes
- User deletion and auth status checks

### 🖼 Media and Profile Assets

- Image upload support
- Profile avatar storage and replacement

### ⚡ Frontend Performance

- Angular Signals for reactive state
- Angular CDK Virtual Scroll
- Optimized polling strategy
- Resolver-driven route prefetching
- Material + Tailwind UI composition

---

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
- Lightweight Charts

### Backend

- Go
- Fiber
- PostgreSQL
- database/sql
- JWT / session tokens

### Infrastructure

- Docker
- REST APIs
- Polling-based real-time updates

---

## ⚙️ Setup Instructions

### 🖥 Backend and Database

1. Clone the repository.

```bash
git clone https://github.com/Sahil2004/gmarket.git
cd gmarket
```

2. Configure environment variables.
   Create a `.env` file from your example config before starting the backend.

3. Start the full stack with Docker.

```bash
docker compose up
```

For development with file watching:

```bash
docker compose watch
```

To inspect logs in another terminal:

```bash
docker compose logs -f
```

4. Run only the backend locally.

```bash
cd server
go mod tidy
go run main.go
```

Backend runs on:

```bash
http://localhost:3000
```

### 🖥 Frontend

```bash
cd client
npm install
npm start
```

Frontend runs on:

```bash
http://localhost:4200
```

---

## License

This project is under GPL-v2. You can view it at [LICENSE](./LICENSE).

## Author

For any information or just to say hi, you can contact me at: me.sahil.gg@gmail.com.
