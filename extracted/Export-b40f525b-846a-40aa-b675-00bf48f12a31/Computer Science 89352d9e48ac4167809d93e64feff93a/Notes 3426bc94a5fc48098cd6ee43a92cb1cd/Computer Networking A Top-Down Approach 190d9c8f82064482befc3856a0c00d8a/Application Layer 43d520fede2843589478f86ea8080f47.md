# Application Layer

The process is a program running within a host

Within the same host, two processes communicate using interprocess communication.

Processes in two different systems communicate by exchanging messages.

## [Socket](https://medium.com/swlh/understanding-socket-connections-in-computer-networking-bac304812b5c)

- process sends/receives messages to/from its sockets

![Untitled](Application%20Layer%2043d520fede2843589478f86ea8080f47/Untitled.png)

- socket is analogous to a door
    - sending process shoves messages out of the door
    - sending process relies on transport infrastructure on the other side of the door to deliver messages to the socket at receiving process.
    - two sockets are involved: one on each side

## Addressing Processes

- to receive messages, a process must have an identifier
- host device has a unique 32-bit IP address.
- Does the IP address of the host on which the process runs suffice for identifying the process
    - No, many processes can be running on the same host. Identifier includes both the IP address and port numbers associated with the process on the host.

![Untitled](Application%20Layer%2043d520fede2843589478f86ea8080f47/Untitled%201.png)

## Sockets

It’s the only API existing between the application layer and the network layer.

Two socket types for two transport services: TCP and UDP

### TCP

- the server process must be already running
- the server must have created a socket for welcoming the client’s contact.
- The client will create a socket at its end, specifying the server IP, port
- Now the client establishes a TCP connection to the server TCP.
- When contacted by the client, server TCP creates a new socket for the server process to communicate with that particular client
    - This allows the server to communicate with multiple clients
    - source port is used to distinguish between clients