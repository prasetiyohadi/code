# frozen_string_literal: true

require "socket"

server = TCPServer.new("0.0.0.0", 8080)

loop do
  # When client connects, return a TCPSocket
  socket = server.accept

  # Read the first line of the request (the Request-Line)
  request = socket.gets

  # Log the request to the console for debugging
  STDERR.puts request

  response = "Hello, World!\n"

  # Response the HTTP request
  # We need to include the Content-Type and Content-Length headers
  socket.print "HTTP/1.1 200 OK\r\n" +
    "Content-Type: text/plain\r\n" +
    "Content-Length: #{response.bytesize}\r\n" +
    "Connection: close\r\n"

  # Print a blank line to separate the header from the response body
  # as required by the protocol.
  socket.print "\r\n"

  # Print the actual response body
  socket.print response

  # Close the socket, terminating the connection
  socket.close
end
