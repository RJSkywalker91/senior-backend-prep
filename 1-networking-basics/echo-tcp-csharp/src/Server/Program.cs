using System.Net;
using System.Net.Sockets;

internal class Program
{
  private const string HOST = "127.0.0.1";
  private const int PORT = 80;
  private const int MAX_CONNECTIONS = 2;
  private const int BUFFER_SIZE = 1024;

  private static void Main(string[] args)
  {
    // Create the socket
    var server = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);

    // Bind the socket
    var ipAddress = IPAddress.Parse(HOST);
    var ipEndpoint = new IPEndPoint(ipAddress, PORT);
    server.Bind(ipEndpoint);

    // Listen for incoming connection (up to max)
    server.Listen(MAX_CONNECTIONS);
    Console.WriteLine($"Server listening on Port: {PORT}");

    // Accept a connection
    var client = server.Accept();
    Console.WriteLine($"Client connected");

    // Echo received data back to connection
    var data = new byte[BUFFER_SIZE];
    client.Receive(data);
    Console.WriteLine("Message received from client");
    client.Send(data);
    Console.WriteLine("Message echoed back to client");

    // Close sockets
    client.Close();
    server.Close();
    Console.WriteLine("Client and Server sockets closed.");
  }
}