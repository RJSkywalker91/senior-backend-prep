# UDP Time Server

## 📌 Overview
What is it?
- A server that broadcasts the current time across UDP sockets, and a client that receives the broadcasts and prints the time.
Why did I build it?
- Looking at how UDP sockets work

---

## 🎯 Learning Goals
- [x] Learn how to keep a socket open until I shut down the program
- [x] Learn how UDP implementation is different than TCP
---

## 🛠️ Tech Stack
- Language: Go
- Libraries/Frameworks: "net"

---

## 🚀 How to Run
**Clone the repo:**
```bash
git clone https://github.com/RJSkywalker91/senior-backend-prep.git
cd 1-networking-basics/udp-time-server/src
```
**Run the Server**
```bash
go run server/time-server.go
```
**Test with Client**
```bash
go run client/client.go
```

## 📖 Notes & Reflections
What went well?
- Felt easier than TCP (but maybe that's cuase I already did TCP)
What was challenging?
- DialUDP used new params that I wasn't quite sure how to implement
What you’d improve if you revisited this?
- N/A

## Workflows
- TBD
