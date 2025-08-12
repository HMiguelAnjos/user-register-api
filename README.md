# SOLID + Hexagonal Architecture Go API — User Registration

Uma API REST de exemplo, pronta para produção, para **registro e gerenciamento de usuários** em Go, aplicando **princípios SOLID** e **Arquitetura Hexagonal** (Ports & Adapters).

## 🏗 Destaques
- **Design centrado no domínio** — entidades puras, sem dependência de frameworks.
- **Ports & Adapters** — `ports` definem interfaces; `adapters` implementam (HTTP, repositório em memória ou banco de dados, gerador de IDs, logger).
- **Casos de Uso** em `internal/app` contendo toda a lógica de aplicação.
- **Controllers** em `internal/adapters/http` fazem a tradução HTTP <-> DTOs.
- **Inversão de Dependência** — o `main` conecta interfaces a implementações concretas.
- **Padrões utilizados**: Repository, Factory Method, DTO, Strategy (gerador de ID), Adapter, Controller, Service, Mapper.

## 🚀 Executando o projeto
```bash
cd user-register-api
go run ./cmd/api
```
Servidor disponível em `:8080`.

## 📌 Endpoints
- **POST /v1/users** — cria usuário  
  **Body:**  
  ```json
  { "name": "John Doe", "email": "john@example.com", "password": "123456" }
  ```
- **GET /v1/users** — lista todos os usuários
- **GET /v1/users/{id}** — busca usuário pelo ID
- **PUT /v1/users/{id}** — atualiza dados do usuário  
  **Body:**  
  ```json
  { "name": "John Updated", "email": "john.updated@example.com" }
  ```
- **DELETE /v1/users/{id}** — remove usuário

## 🧪 Rodando os testes
```bash
go test ./...
```