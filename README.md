# 🧠 MoodAPI 

Uma API REST desenvolvida em Go para monitoramento de humor. Pacientes registram como estão se sentindo e psicólogos acessam relatórios semanais.

## 📦 Tecnologias Utilizadas

- Go
- Chi (router)
- PostgreSQL
- JWT (autenticação)
- Arquitetura em camadas: `handler`, `service`, `repository`

---

## 📁 Estrutura do Projeto

```bash
.
├── cmd/
│   └── api/
│       └── main.go          # Ponto de entrada da aplicação
├── database/
│   └── database.sql         # Script para criação do banco de dados
├── internal/
│   ├── handler/             # Controladores HTTP
│   ├── service/             # Regras de negócio
│   ├── repository/          # Comunicação com o banco
│   ├── models/              # Structs e entidades
│   └── middleware/          # Middlewares
├── go.mod
└── .env                     # Variáveis de ambiente
```
---
## ⚙️ Pré-requisitos

- Go 1.21+
- PostgreSQL 13+
- psql instalado e configurado

---

## 📄 Variáveis de Ambiente
Crie um arquivo .env na raiz do projeto com o conteúdo:

``` bash
DATABASE_URL=postgres://usuario:senha@localhost:5432/mooddb?sslmode=disable
JWT_SECRET=supertokenseguro
```
---

## 🛠️ Como rodar o projeto
1. Clone o repositório:

``` bash
git clone https://github.com/seu-usuario/moodapi.git
cd moodapi
```

2. Crie o banco de dados e tabelas:
``` bash
psql -U seu_usuario -d postgres -f database/database.sql
```
O script cria o banco mooddb (caso não exista) e as tabelas users e mood_entries.

3. Instale as dependências do Go:

``` bash
go mod tidy
```

4. Execute o servidor:

``` bash
go run cmd/api/main.go
```
O servidor estará disponível em: http://localhost:8080

---

## 📲 Endpoints
### 🔐 Autenticação
| Método | Rota   | Descrição                |
|--------|--------|--------------------------|
| POST   | `/signup` | Cadastrar novo usuário   |
| POST   | `/login`  | Login e retorno do token |

### 😌 Humor (Paciente)
| Método | Rota    | Descrição                        |
|--------|---------|----------------------------------|
| POST   | `/mood`   | Registrar nível de humor (1 a 5) |
| GET    | `/report` | Ver relatório semanal            |