You are my senior full-stack engineer. Build a production-quality full-stack calculator app.

Hard requirements:
- Frontend: React + TypeScript. Clean UI, responsive, input validation, error handling.
- Backend: Go REST microservice. Validate input, handle edge cases (division by zero, invalid data), JSON responses.
- Operations: add, subtract, multiply, divide. Also implement exponentiation, square root, percentage.
- Unit tests for backend and frontend covering key functionality.
- Documentation: README with setup instructions, run steps, API examples (curl), and design rationale/assumptions.
- Provide a repo-ready file tree and ALL code files with exact paths and contents.
- Optional: Dockerfile + docker-compose to run frontend + backend together (do it if reasonable).

Accuracy / quality protocol (do not skip):
1) Start by proposing the API contract (endpoints + request/response + error schema).
2) Then implement backend with table-driven tests first (or alongside), and show `go test ./...` commands and expected output format.
3) Then implement frontend consuming the API. Include input validation rules and error UI behavior.
4) Add frontend tests (React Testing Library + Jest/Vitest). Show `npm test` commands.
5) Provide coverage instructions for both layers.
6) Before final output, do a self-check list: edge cases, tests, lint/typecheck, consistent error formats.

Constraints:
- Keep code idiomatic and maintainable.
- Use simple dependencies.
- Use a consistent JSON error format like { "error": { "code": "...", "message": "...", "details": ... } }.

Output format:
- First: API contract section
- Second: repo file tree
- Third: each file in a separate fenced code block labeled with its path
- Fourth: run instructions + test/coverage commands
- Fifth: README content (complete)
