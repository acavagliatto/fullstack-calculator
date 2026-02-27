# Full-Stack Calculator

A production-quality full-stack calculator application built with a React + TypeScript frontend and a Go REST API backend.

This project demonstrates clean architecture, strong input validation, comprehensive testing, and Dockerized deployment.

---

## 🚀 Features

- ✅ Basic & advanced arithmetic operations  
- ✅ Client-side and server-side validation  
- ✅ Robust error handling for edge cases  
- ✅ Responsive modern UI (desktop + mobile)  
- ✅ Unit & integration tests (frontend and backend)  
- ✅ Coverage reporting  
- ✅ Dockerized full-stack deployment  

---

## 🏗 Architecture

### Backend (Go)

- Native `net/http` server (minimal dependencies)
- Clear separation of concerns:
  - `calculator/` – Core business logic
  - `handlers/` – HTTP layer + validation
  - `main.go` – Server bootstrap
- JSON-based REST API
- Configurable port (default: `8080`)

### Frontend (React + TypeScript)

- React 19 + TypeScript
- Vite 5 build tool
- Vitest + React Testing Library
- Dynamic inputs based on selected operation
- Strong client-side validation
- Dev port: `5173`
- Docker port: `3000`

---

## 📡 API Contract

### Base URL

http://localhost:8080

---

### POST `/api/calculate`

Performs a calculation based on the specified operation and operands.

#### Request Body

```json
{
  "operation": "add | subtract | multiply | divide | exponentiation | sqrt | percentage",
  "operands": [number, ...]
}
```

#### Success Response

```json
{
  "result": 42.0
}
```

#### Error Response

```json
{
  "error": "error message"
}
```

#### Status Codes

- `200 OK` – Calculation successful
- `400 Bad Request` – Invalid input or calculation error
- `405 Method Not Allowed` – Non-POST request

---

### GET `/api/health`

Health check endpoint.

#### Response

```json
{
  "status": "ok"
}
```

---

## 🧮 Supported Operations

| Operation         | Operands Required | Example Input | Example Result |
|------------------|------------------|--------------|---------------|
| `add`            | 2+ numbers        | `[5, 3]`     | `8`           |
| `subtract`       | 2+ numbers        | `[10, 3]`    | `7`           |
| `multiply`       | 2+ numbers        | `[5, 3]`     | `15`          |
| `divide`         | Exactly 2         | `[10, 2]`    | `5`           |
| `exponentiation` | Exactly 2         | `[2, 3]`     | `8`           |
| `sqrt`           | Exactly 1         | `[9]`        | `3`           |
| `percentage`     | Exactly 2         | `[100, 10]`  | `10`          |

---

## ⚠️ Error Handling

Handled edge cases include:

- Division by zero
- Negative square roots
- Invalid exponentiation (e.g., `0^-1`)
- Missing operands
- Invalid operation names
- Non-numeric inputs

All validation errors return:

HTTP 400 Bad Request

---

# 🛠 Setup & Run

## Prerequisites

- Go 1.22+
- Node.js 20+
- npm 9+
- Docker (optional)

---

## ▶ Backend

```bash
cd backend

go mod download
go test ./...
go test ./... -cover
go run main.go
```

Backend runs at:

http://localhost:8080

---

## ▶ Frontend

```bash
cd frontend

npm install
npm test
npm run dev
```

Frontend runs at:

http://localhost:5173

---

# 🐳 Docker Deployment (Recommended)

Run full stack with:

```bash
docker compose up --build
```

Detached mode:

```bash
docker compose up -d
```

Stop services:

```bash
docker compose down
```

### Access

- Frontend: http://localhost:3000  
- Backend API: http://localhost:8080/api  

---

# 📬 API Usage Examples

### Addition

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"add","operands":[5,3]}'
```

### Division by Zero (Error)

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"divide","operands":[10,0]}'
```

### Square Root

```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation":"sqrt","operands":[9]}'
```

### Health Check

```bash
curl http://localhost:8080/api/health
```

---

# 🧪 Testing

## Backend

```bash
cd backend
go test ./...
go test ./... -v
```

### Coverage

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Generates:
- coverage.out
- coverage.html

---

## Frontend

```bash
cd frontend
npm test
```

### Coverage

```bash
npm run coverage
```

Generates:
- /coverage folder with HTML report

---

# 🧠 Design Rationale

## Architectural Decisions

- Backend and frontend are fully decoupled via REST API.
- Go chosen for simplicity, performance, and strong typing.
- React + TypeScript chosen for maintainability and type safety.
- Validation occurs both client-side and server-side.
- Comprehensive unit and integration tests ensure correctness.

---

# 📌 Assumptions

- All numbers handled as float64 / number.
- Floating point precision limitations apply.
- Percentage calculates X% of Y.
- CORS allows all origins (development mode).
- No authentication (assumed trusted environment).

---

# 🔐 Security Notes

- JSON-only API
- Input validation on both layers
- CORS should be restricted in production
- Rate limiting recommended for production
- HTTPS required in real deployments

---

# 📂 Project Structure

```
fullstack-calculator/
├── backend/
│   ├── calculator/
│   ├── handlers/
│   ├── main.go
│   ├── go.mod
│   └── Dockerfile
├── frontend/
│   ├── src/
│   ├── index.html
│   ├── package.json
│   ├── vite.config.ts
│   ├── nginx.conf
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

---


# 📈 Production Readiness Checklist

- [x] Comprehensive input validation
- [x] Error handling for edge cases
- [x] Unit tests (frontend & backend)
- [x] Coverage reporting
- [x] Dockerized deployment
- [x] Health check endpoint
- [ ] Rate limiting (recommended)
- [ ] Monitoring/logging (recommended)
- [ ] HTTPS (required for production)

---

# 📜 License

MIT License
