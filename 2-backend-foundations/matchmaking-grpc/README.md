# Matchmaking REST API

## 📌 Overview
What is it?
- A GRPC built with GO
- Built to simulate matchmaking services
---

## 🎯 Learning Goals
- Authentication + JWT
- Using .env in GO
- Using migrations with Postgres
- Implementing GRPC
---

## 🛠️ Tech Stack
- Language: Go
- Framework: grpc
---

## Design

### Key Entities
- Player
- MatchTicket
- Match

### Key Endpoints
| RPC Client             | Method | Purpose                                   |
|:---                    |:---    |:---                                       |
| NewPlayerServiceClient | Create | Create a new player in the DB             |
| NewPlayerServiceClient | Login  | Sign and return JWT for authorized player |

## 🚀 How to Run
**Clone the repo:**
```bash
git clone https://github.com/RJSkywalker91/senior-backend-prep.git
cd 2-backend-foundations/matchmaking-grpc
```
**Run the Server**
```bash
go run cmd/server/main.go
```

**Run the Create client**
```bash
go run cmd/client/create/main.go
```

**Run the Login client**
```bash
go run cmd/client/login/main.go
```

## 📖 Notes & Reflections
What went well?
What was challenging?
What you’d improve if you revisited this?