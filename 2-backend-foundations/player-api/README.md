# Player REST API

## 📌 Overview
What is it?
- A REST API built with GO
- A simple API to register and retrieve player objects from a database
---

## 🎯 Learning Goals
- [x] Idiomatic Project Structure
- [x] HTTP Routing
- [x] Extrapolation of logic by domain
- [x] Static Config of project
- [x] Integration with DB (postgres) 
---

## 🛠️ Tech Stack
- Language: Go
- DB: Postgres
---

## Design

### Key Entities
- Player

### Key Endpoints
| Method | Endpoint      | Purpose              |
|:---    |:---           |:---                  |
| POST   | /register     | Register a Player    |
| GET    | /players/{id} | Get Player data      |

## 🚀 How to Run
**Clone the repo:**
```bash
git clone https://github.com/RJSkywalker91/senior-backend-prep.git
cd 2-backend-foundations/player-api-go
```
**Run the Server**
```bash
go run cmd/player-api/main.go
```

## 📖 Notes & Reflections
What went well?
- Learned A LOT about how to create APIs in GO
What was challenging?
- Project structure in go, specifically with domain driven design in mind