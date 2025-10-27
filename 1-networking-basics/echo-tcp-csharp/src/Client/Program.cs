using System.Net;
using System.Net.Sockets;
using System.Text;

internal class Program
{
  private const string HOST = "127.0.0.1";
  private const int PORT = 80;
  private const int BUFFER_SIZE = 1024;

  private static void Main(string[] args)
  {
    // Create the socket
    var client = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);

    // Connect the socket to the server
    var ipAddress = IPAddress.Parse(HOST);
    var ipEndpoint = new IPEndPoint(ipAddress, PORT);
    client.Connect(ipEndpoint);

    // Send data to the server
    var message = "Hello world!";
    var encoding = Encoding.UTF8;
    client.Send(encoding.GetBytes(message));
    Console.WriteLine("Message sent to server");

    // Get response from the server
    var response = new byte[BUFFER_SIZE];
    client.Receive(response);
    Console.WriteLine($"Response form server: {encoding.GetString(response)}");

    // Close sockets
    client.Close();
    Console.WriteLine("Client socket closed.");
  }
}