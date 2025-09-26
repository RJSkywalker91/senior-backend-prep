import socket

# Constants added so I can get used to all functions without
# figuring out what all the numbers / other constants mean.
IP4_SOCKET = socket.AddressFamily.AF_INET
TCP_SOCKET = socket.SocketKind.SOCK_STREAM
SERVER_HOST = '127.0.0.1'
SERVER_PORT = 80
BUFFER_SIZE = 1024

def client():
  
  # Create the socket
  client_socket = socket.socket(IP4_SOCKET, TCP_SOCKET)
  
  # Connect the socket to the server
  client_socket.connect((SERVER_HOST, SERVER_PORT))

  # Send data to the server
  message = "Hello world!"
  client_socket.sendall(message.encode('utf-8'))
  print("Message sent to server")

  # Get response from the server
  response = client_socket.recv(BUFFER_SIZE)
  print(f"Response from server: {response.decode('utf-8')}")
  
  # Close socket
  client_socket.close()
  print("Client socket closed")

if __name__ == "__main__":
  client()