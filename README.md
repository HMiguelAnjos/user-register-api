# SOLID + Hexagonal Architecture Go API

A small, production-ready example of a REST API in Go using **SOLID** principles and **Hexagonal Architecture** (Ports & Adapters).

## Highlights
- **Domain-centric** design (entities have no framework deps).
- **Ports & Adapters**: `ports` define interfaces; `adapters` implement them (HTTP, in-memory repo, id generator, logger).
- **Use Cases** in `internal/app` expose application logic.
- **Controllers** in `internal/adapters/http` translate HTTP <-> DTOs.
- **Dependency Inversion**: main wires interfaces to concrete adapters.
- **Patterns used**: Repository, Factory Method, DTO, Strategy (ID generator), Adapter, Controller, Service, Mapper.

## Run
```bash
cd user-register-api
go run ./cmd/api
```
Server starts on `:8080`.

## Endpoints
- `POST /v1/tasks` — create task `{ "title": "...", "description": "..." }`
- `GET /v1/tasks` — list tasks
- `GET /v1/tasks/{id}` — get by id
- `PUT /v1/tasks/{id}` — update `{ "title": "...", "description": "...", "done": true }`
- `DELETE /v1/tasks/{id}` — delete

## Tests
```bash
go test ./...
```
