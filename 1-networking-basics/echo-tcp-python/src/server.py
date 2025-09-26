import socket

# Constants added so I can get used to all functions without
# figuring out what all the numbers / other constants mean.
IP4_SOCKET = socket.AddressFamily.AF_INET
TCP_SOCKET = socket.SocketKind.SOCK_STREAM
HOST = '127.0.0.1'
PORT = 80
MAX_CONNECTIONS = 5
BUFFER_SIZE = 1024

def server():
  
  # Create the socket
  server_socket = socket.socket(IP4_SOCKET, TCP_SOCKET)
  
  # Bind the socket
  server_socket.bind((HOST, PORT))
  
  # Listen for incoming connection (up to max)
  server_socket.listen(MAX_CONNECTIONS)
  print(f"Server listening on Port: {PORT}")

  # Accept a connection
  [client_conn, client_address] = server_socket.accept()
  print(f"Client connected with IP: {client_address[0]} on Port: {client_address[1]}")
  
  # Echo received data back to connection
  data = client_conn.recv(BUFFER_SIZE)
  print(f"Message received from client: {data.decode('utf-8')}")
  client_conn.send(data)
  print("Message echoed back to client")
  
  # Close sockets
  client_conn.close()
  server_socket.close()
  print("Client and Server sockets closed")

if __name__ == "__main__":
  server()