# Basic GO API in Docker

## 📌 Overview
What is it?
- A bare-bones GO API built to learn how to use Docker
---

## 🎯 Learning Goals
- Docker
- Docker Compose
- Makefile
---

## 🛠️ Tech Stack
- Language: Go
- Framework: Docker
---

## 🚀 How to Run
**Clone the repo:**
```bash
git clone https://github.com/RJSkywalker91/senior-backend-prep.git
cd 2-backend-foundations/docker
```
**Run the Server locally**
```bash
go run cmd/main.go
```

**Run the Server in Docker**
```bash
docker run --rm -p 8080:8080 -e APP_MESSAGE="Howdy from Make!" hello:dev
```

**Run the Server with Docker Compose**
```bash
docker compose -f build/docker/docker-compose.yaml up --build
```

**Run the Server with Makefile**
```bash
make run
```
or
```bash
make up
```

## 📖 Notes & Reflections
### What went well?
- Easy to create and understand GO API now
### What was challenging?
- Lots of concepts with Docker