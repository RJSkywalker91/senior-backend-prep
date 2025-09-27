# Networking Basics Notes

## echo-tcp
### Starting Thoughts
- I have worked with plenty of server concepts in C#, but not a low level (sockets), connection-based server where I am the one writing the logic for connections.
- I will create a simple "echo" tcp server in 4 languages: Python, C#, JS, and GO
- I do have quite a bit of experience in Python, C#, and JS. I am confident these three will be very easy to implement in a small amount of time.
- Starting out with the easiest possible implementation. There will not be a lot of error checking, try/catch, etc. 

### Python Thoughts
- Interesting how many of the socket functions use tuples
- Overall the process is incredibly simple with the usage of the 'socket' library
- It looks like the 'socket' library has low-level and high-level implementations that could be utilized (e.g. 'send' vs 'sendall'). I am going to use the high-level ones for now, but might come back and implement the server again with lower level concepts after I finish the other three languages

### C# Thoughts
- C# socket library was much easier to read EXCEPT for Bind(). It was a little difficult to understand the concept of IPEndpoint without looking it up.
- Encoding was a bit trickier

### JS Thoughts

### GO Thoughts

### Final Thoughts
- 