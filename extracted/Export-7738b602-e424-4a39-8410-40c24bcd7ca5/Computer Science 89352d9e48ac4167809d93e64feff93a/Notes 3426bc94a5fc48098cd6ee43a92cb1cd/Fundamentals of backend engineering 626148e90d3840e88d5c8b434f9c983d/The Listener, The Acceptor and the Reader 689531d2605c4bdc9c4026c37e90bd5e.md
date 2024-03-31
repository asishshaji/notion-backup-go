# The Listener, The Acceptor and the Reader

## Listener

- thread or process calls the listen and passes the port and gets back the socket id

> A socket is not a connection but a place to connect to. Many connections can be connected to the socket. The listener is the process where socket lives.
> 

## Acceptor

- whoever have the access to socket id and calls the accept function to get the socket file descriptor pointing to connection.
- With a socket in hand, the backend can call the OS function *accept* passing this socket to accept any available connections on this socket. The *accept* function returns a [file descriptor](https://en.wikipedia.org/wiki/File_descriptor) representing the connection. Connections have to be actively accepted by the application in order to serve clients. Otherwise the connections will remain in the OS *accept* queue unused. The acceptor is the thread or process that calls the *accept* function.

## Reader

- You can also call it worker too, this is where the work happens.
- When we have a connection, the backend application is also responsible to read data sent to this connection otherwise the buffer allocated by the OS for this connection will fill up and client won’t be able to send any more data.
- I call the process that reads the TCP stream on the connection the Reader.
- The reader passes in the file descriptor (connection) to read the stream and act on it.