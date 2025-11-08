# Backend Foundations Notes

I want to create a very simple Matchmaking Ticket REST API to explore:
- REST architecture
- JWT Authentication
- Database connection
- CRUD operations
- Simple, idempotent workflow
- All of the above using GO

---

# Vertical Slice Player API
### Starting Thoughts
The goal will be to build out a complete api with database connection (PostgreSQL). This will be E2E from router, handler, service, and repository levels.
### Final Thoughts
I feel a lot more confident in file structure for GO projects after making this API. The concepts for DI are familiar from my work in C#, so that was easy to understand. It was interesting to see how everything gets put together in main.go, and overall I feel like I learned a lot from this.

---

# Simple Matchmaking gRPC
### Starting Thoughts
I want to learn how to use protobufs to create a gRPC server, along with expanding on the initial logic set up from the Player API to add a method to mimic a queue for players to find a match. 
### Final Thoughts
Protobufs are pretty easy, although the commands to build them feel longer than I would want to type often. I also leaned into using JWT authentication via an interceptor, but this was mostly done via ChatGPT. I will be learning from this boiler plate code. It was fun implementing a STREAM rpc method for queuing a player up to find a match. (also ChatGPT)

---

# Docker Learning
### Starting Thoughts
This project was just to learn the basics of docker and the various files/commands needed to run a go application using docker. 
### Final Thoughts
Docker feels easy enough to use, but there are a lot of nuanced options within, especially as soon as docker-compose and makefile get involed. Learning those will be a lot more learning on top of GO.