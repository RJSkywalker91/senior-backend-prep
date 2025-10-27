# UDP Chat Service

## 📌 Overview
What is it?
- A simple, bare-bones chat service over UDP using custom packets
Why did I build it?
- Diving deeper into protocols built over UDP + custom server logic 

---

## 🎯 Learning Goals
- [x] Learn how to Marshal/Unmarhsal custom packets on top of UDP
- [x] Learn more about goroutines and how/when to use them
---

## 🛠️ Tech Stack
- Language: Go
- Libraries/Frameworks: "net"
- Custom Packets
---

## 🚀 How to Run
**Clone the repo:**
```bash
git clone https://github.com/RJSkywalker91/senior-backend-prep.git
cd 1-networking-basics/udp-custom-chat/src
```
**Run the Server**
```bash
go run server/server.go
```
**Test with Client**
```bash
go run client/client.go
```

## 📖 Notes & Reflections
What went well?
- Implementing the custom MessagePacket in server/client
What was challenging?
- Creating the actual Marshal and UnMarshal code
What you’d improve if you revisited this?
- Pretty much everything

## Workflows
- TBD
