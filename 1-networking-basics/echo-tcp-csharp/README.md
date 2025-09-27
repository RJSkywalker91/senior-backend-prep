# TCP Echo Server in C#

## 📌 Overview
What is it?
- A simple TCP echo server that accepts client connections and sends back messages.
Why did I build it?
- Exploring the basics again after almost 10 years

---

## 🎯 Learning Goals
- [x] Create a server in c#
- [x] Utilize Socket libraries

---

## 🛠️ Tech Stack
- Language: C#
- Libraries/Frameworks: .NET Sockets

---

## 🚀 How to Run
**Clone the repo:**
```bash
git clone https://github.com/RJSkywalker91/senior-backend-prep.git
cd 1-networking-basics/echo-tcp-csharp/src
```
**Run the Server**
```bash
dotnet run --project Server/Server.csproj
```
**Test with Client**
```bash
dotnet run --project Client/Client.csproj
```

## 📖 Notes & Reflections
What went well?
- Went quicker after doing the python
- Didn't have much trouble getting back into writing C# code
What was challenging?
- Haven't created a dotnet solution outside of Visual Studio in a while. Took about 30 minutes of troublshooting to get everything working in VS Code / terminal
- Encoding/Decoding is more complicated in C# compared to python
What you’d improve if you revisited this?
- Add error handling