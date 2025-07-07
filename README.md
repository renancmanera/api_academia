# API Academia – Backend em Go

Este projeto é uma API RESTful para gerenciamento de academia, desenvolvida em Go (Gin + GORM + PostgreSQL).  
Inclui autenticação JWT, CRUD completo para usuários, treinos e exercícios, e associação entre treinos e exercícios.

## Funcionalidades

- Cadastro e login de usuário com JWT
- CRUD de usuários, treinos e exercícios
- Associação de exercícios a treinos (muitos-para-muitos)
- Proteção de rotas com middleware JWT

## Principais Endpoints

| Método | Rota                        | Descrição                       |
|--------|-----------------------------|---------------------------------|
| POST   | /cadastro                   | Cadastro de usuário             |
| POST   | /login                      | Login e geração de token JWT    |
| GET    | /usuarios                   | Listar usuários                 |
| PUT    | /usuarios/:id               | Atualizar usuário               |
| DELETE | /usuarios/:id               | Deletar usuário                 |
| POST   | /treinos                    | Cadastrar treino                |
| GET    | /treinos                    | Listar treinos do usuário       |
| PUT    | /treinos/:id                | Atualizar treino                |
| DELETE | /treinos/:id                | Deletar treino                  |
| POST   | /exercicios                 | Cadastrar exercício             |
| GET    | /exercicios                 | Listar exercícios               |
| PUT    | /exercicios/:id             | Atualizar exercício             |
| DELETE | /exercicios/:id             | Deletar exercício               |
| POST   | /treinos/:id/exercicios     | Associar exercícios ao treino   |
| GET    | /treinos/:id/exercicios     | Listar exercícios do treino     |

## Como rodar

1. Clone o repositório e entre na pasta do projeto
2. Crie um arquivo `.env` com as variáveis do seu banco e JWT
3. Instale as dependências:  
   `go mod tidy`
4. Rode o servidor:  
   `go run cmd/main.go`
