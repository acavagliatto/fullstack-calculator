# Full-Stack Calculator

A production-quality full-stack calculator application with a React + TypeScript frontend and a Go REST API backend.

## Features

- **Operations**: Addition, Subtraction, Multiplication, Division, Exponentiation, Square Root, Percentage
- **Input Validation**: Comprehensive client-side and server-side validation
- **Error Handling**: User-friendly error messages for edge cases (division by zero, negative square roots, etc.)
- **Responsive UI**: Clean, modern interface that works on desktop and mobile
- **Comprehensive Testing**: Unit tests for both frontend and backend
- **Docker Support**: Run the entire stack with docker-compose

## Architecture

### Backend (Go)
- **Framework**: Native Go HTTP server with CORS support
- **Structure**: 
  - `calculator/`: Core calculation logic with comprehensive error handling
  - `handlers/`: HTTP request handlers with input validation
  - `main.go`: Server setup and routing
- **Port**: 8080 (configurable via `PORT` environment variable)

### Frontend (React + TypeScript)
- **Framework**: React 19 with TypeScript
- **Build Tool**: Vite 5
- **Testing**: Vitest + React Testing Library
- **Styling**: CSS with responsive design
- **Port**: 5173 (dev), 3000 (docker)

## API Contract

### Endpoints

#### `POST /api/calculate`
Performs a calculation based on the provided operation and operands.

**Request Body:**
```json
{
  "operation": "add|subtract|multiply|divide|exponentiation|sqrt|percentage",
  "operands": [number, ...]
}
```

**Response (Success):**
```json
{
  "result": 42.0
}
```

**Response (Error):**
```json
{
  "error": "error message"
}
```

**Status Codes:**
- `200 OK`: Calculation successful
- `400 Bad Request`: Invalid input or calculation error
- `405 Method Not Allowed`: Non-POST request

#### `GET /api/health`
Health check endpoint.

**Response:**
```json
{
  "status": "ok"
}
```

### Operation Requirements

| Operation | Operands | Example Request | Example Result |
|-----------|----------|----------------|----------------|
| `add` | 2+ numbers | `[5, 3]` | `8` |
| `subtract` | 2+ numbers | `[10, 3]` | `7` |
| `multiply` | 2+ numbers | `[5, 3]` | `15` |
| `divide` | 2 numbers | `[10, 2]` | `5` |
| `exponentiation` | 2 numbers (base, exponent) | `[2, 3]` | `8` |
| `sqrt` | 1 number | `[9]` | `3` |
| `percentage` | 2 numbers (value, percentage) | `[100, 10]` | `10` |

### Error Cases

- **Division by zero**: `"division by zero"`
- **Negative square root**: `"cannot compute square root of negative number"`
- **Invalid exponent**: `"invalid exponent operation"` (e.g., 0^-1 or (-4)^0.5)
- **Invalid operation**: `"invalid operation"`
- **Missing/invalid operands**: `"operands are required"`, `"add requires at least 2 operands"`, etc.

## Setup and Run

### Prerequisites
- **Go**: 1.21 or higher
- **Node.js**: 20 or higher
- **npm**: 9 or higher
- **Docker** (optional): For containerized deployment

### Backend Setup

```bash
cd backend

# Install dependencies
go mod download

# Run tests
go test ./... -v

# Run with coverage
go test ./... -cover

# Start the server
go run main.go
# Server will start on http://localhost:8080
```

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Run tests
npm test

# Run tests in watch mode
npm test -- --watch

# Run tests with UI
npm run test:ui

# Start development server
npm run dev
# Frontend will start on http://localhost:5173

# Build for production
npm run build

# Preview production build
npm run preview
```

### Docker Compose (Recommended)

Run both frontend and backend together:

```bash
# Build and start services
docker-compose up --build

# Or run in detached mode
docker-compose up -d

# Stop services
docker-compose down
```

Access the application:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080/api

## API Usage Examples (curl)

### Addition
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "add", "operands": [5, 3]}'
# Response: {"result":8}
```

### Subtraction
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "subtract", "operands": [10, 3]}'
# Response: {"result":7}
```

### Multiplication
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "multiply", "operands": [5, 3]}'
# Response: {"result":15}
```

### Division
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "divide", "operands": [10, 2]}'
# Response: {"result":5}
```

### Division by Zero (Error Example)
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "divide", "operands": [10, 0]}'
# Response: {"error":"division by zero"}
```

### Exponentiation
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "exponentiation", "operands": [2, 3]}'
# Response: {"result":8}
```

### Square Root
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operands": [9]}'
# Response: {"result":3}
```

### Negative Square Root (Error Example)
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "sqrt", "operands": [-4]}'
# Response: {"error":"cannot compute square root of negative number"}
```

### Percentage
```bash
curl -X POST http://localhost:8080/api/calculate \
  -H "Content-Type: application/json" \
  -d '{"operation": "percentage", "operands": [100, 10]}'
# Response: {"result":10}
```

### Health Check
```bash
curl http://localhost:8080/api/health
# Response: {"status":"ok"}
```

## Testing

### Backend Tests
```bash
cd backend

# Run all tests
go test ./...

# Run with verbose output
go test ./... -v

# Run with coverage
go test ./... -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Frontend Tests
```bash
cd frontend

# Run tests once
npm test -- --run

# Run in watch mode
npm test

# Run with UI
npm run test:ui

# Run with coverage
npm run coverage
```

## Design Rationale & Assumptions

### Architecture Decisions

1. **Separation of Concerns**: Backend and frontend are completely decoupled, communicating only via REST API. This allows independent scaling and deployment.

2. **Backend (Go)**:
   - **Why Go**: High performance, excellent concurrency support, simple deployment (single binary), strong standard library
   - **No Framework**: Using native `net/http` keeps dependencies minimal and reduces attack surface
   - **Table-Driven Tests**: Comprehensive test coverage with clear, maintainable test cases
   - **Error Types**: Custom error types allow precise error handling and user-friendly messages

3. **Frontend (React + TypeScript)**:
   - **Why React**: Component-based architecture, excellent ecosystem, wide adoption
   - **Why TypeScript**: Type safety reduces runtime errors, improves developer experience
   - **Why Vite**: Fast development server, optimized builds, excellent DX
   - **Dynamic Inputs**: Operation-specific input fields improve UX
   - **Client-Side Validation**: Immediate feedback reduces unnecessary API calls

4. **Testing Strategy**:
   - **Backend**: Unit tests for business logic, integration tests for HTTP handlers
   - **Frontend**: Component tests covering user interactions, success/error flows, edge cases
   - **Coverage**: Comprehensive test coverage for core functionality

### Assumptions

1. **Number Format**: All numbers are handled as `float64` (Go) / `number` (TypeScript) for flexibility with decimals

2. **Operation Behavior**:
   - `add`, `subtract`, `multiply`: Support 2+ operands
   - `divide`, `exponentiation`, `percentage`: Require exactly 2 operands
   - `sqrt`: Requires exactly 1 operand
   - Multiple operands (add/subtract/multiply) are processed left-to-right

3. **Error Handling**:
   - Division by zero returns error (not infinity)
   - Square root of negative numbers returns error (no complex numbers)
   - Invalid exponentiation (0^-1, (-4)^0.5) returns error
   - All validation errors return 400 Bad Request

4. **CORS**: Backend allows all origins (`*`) for development. In production, this should be restricted to specific domains.

5. **Precision**: Floating-point arithmetic may have small precision errors. For financial applications, consider using decimal libraries.

6. **Percentage**: Calculates X% of Y (e.g., 10% of 100 = 10), not percentage increase/decrease.

### Security Considerations

- Input validation on both client and server
- JSON-only API (no code execution)
- CORS configuration should be restricted in production
- No authentication/authorization (assumes trusted environment)
- Rate limiting should be added for production use

### Production Readiness Checklist

- [x] Comprehensive input validation
- [x] Error handling for edge cases
- [x] Unit tests (backend and frontend)
- [x] Integration tests (HTTP handlers)
- [x] Docker support
- [x] Responsive UI
- [x] Health check endpoint
- [ ] Rate limiting (recommended for production)
- [ ] Logging/monitoring (recommended for production)
- [ ] Authentication (if needed)
- [ ] HTTPS/TLS (required for production)

## Project Structure

```
fullstack-calculator/
├── backend/
│   ├── calculator/
│   │   ├── calculator.go       # Core calculation logic
│   │   └── calculator_test.go  # Unit tests
│   ├── handlers/
│   │   ├── handlers.go         # HTTP handlers
│   │   └── handlers_test.go    # Handler tests
│   ├── main.go                 # Server entry point
│   ├── go.mod                  # Go dependencies
│   ├── go.sum                  # Go dependency checksums
│   └── Dockerfile              # Backend Docker image
├── frontend/
│   ├── src/
│   │   ├── Calculator.tsx      # Main calculator component
│   │   ├── Calculator.test.tsx # Component tests
│   │   ├── Calculator.css      # Component styles
│   │   ├── api.ts              # API client
│   │   ├── types.ts            # TypeScript types
│   │   ├── App.tsx             # Root component
│   │   ├── App.css             # App styles
│   │   ├── index.css           # Global styles
│   │   ├── main.tsx            # Entry point
│   │   └── setupTests.ts       # Test configuration
│   ├── public/                 # Static assets
│   ├── index.html              # HTML template
│   ├── package.json            # npm dependencies
│   ├── vite.config.ts          # Vite configuration
│   ├── tsconfig.json           # TypeScript configuration
│   ├── nginx.conf              # Nginx config for Docker
│   └── Dockerfile              # Frontend Docker image
├── docker-compose.yml          # Docker Compose configuration
└── README.md                   # This file
```

## Contributing

1. Make changes to backend or frontend
2. Run tests: `go test ./...` (backend), `npm test` (frontend)
3. Ensure all tests pass before committing
4. Follow existing code style and patterns

## License

MIT License - See LICENSE file for details

## Coverage
1. Backend

cd backend
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html

2. Frontend 
cd frontend
npm ci
npm run coverage
# se genera /coverage (dependiendo del config de vitest)
