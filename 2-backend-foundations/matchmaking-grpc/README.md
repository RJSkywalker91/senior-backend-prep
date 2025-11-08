# Matchmaking GRPC

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

### Key Endpoints
| RPC           | Method      | Purpose                                   |
|:---           |:---         |:---                                       |
| PlayerService | Create      | Create a new player in the DB             |
| PlayerService | Login       | Sign and return JWT for authorized player |
| QueueService  | Queue       | Queue player using JWT from Login         |
| QueueService  | CancelQueue | Cancel queued player                      |

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
go run cmd/client/create/main.go -username=TEST -email=test@test.com -password=FAKE
```

**Run the Login client**
```bash
go run cmd/client/login/main.go -username=NAME -password=FAKE
```

**Run the Queue client**
```bash
go run cmd/client/queue/main.go
```

## 📖 Notes & Reflections
What went well?
- Everything worked in the end
- Surprised how much more I learned in this lesson (e.g. grpc/proto, .env, migrations, etc.)
- It is nice to see this beginning to look more like an actual project
What was challenging?
- Had to utilize ChatGPT for some of this
- grpc is going to take some getting used to
What you’d improve if you revisited this?
- Definitely would want to try and take a stab at an Auth Service
- Want to build my own queue manager at some point

# Useful Things
**Player Proto generator command**
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cmd/proto/player.proto
```

**Queue Proto generator command**
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cmd/proto/queue.proto
```