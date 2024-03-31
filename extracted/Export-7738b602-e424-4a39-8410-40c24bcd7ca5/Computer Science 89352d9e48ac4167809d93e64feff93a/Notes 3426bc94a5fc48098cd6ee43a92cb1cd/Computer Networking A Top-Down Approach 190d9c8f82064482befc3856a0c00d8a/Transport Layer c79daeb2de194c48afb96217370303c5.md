# Transport Layer

[TCP IP(Network series blog 3)](https://medium.com/srendevops/tcp-ip-network-series-blog-3-f0421578177)

The **Network Layer** handles the logical communication between hosts (using IP addresses)

The **Transport Layer** handles logical communication between processes (using ports)

- Internet Protocol (IP) provides a packet delivery service across an internet
- however, IP cannot distinguish between multiple *processes* (applications) running on the same computer
- fields in the IP datagram header identify only *computers*
- a protocol that allows an application to serve as an *end-point* of communication is known as a *transport protocol* or an *end-to-end protocol*

---

- a *socket* is the interface through which a process (application) communicates with the transport layer
- each process can potentially use many sockets
- the transport layer in a receiving machine receives a sequence of segments from its network layer
- delivering segments to the correct socket is called *demultiplexing*
- assembling segments with the necessary information and passing them to the network layer is called *multiplexing*
- multiplexing and demultiplexing are need whenever a communications channel is *shared*

# Transport layer multiplexing and demultiplexing

[Transport Layer](https://www.dcs.bbk.ac.uk/~ptw/teaching/IWT/transport-layer/notes.html)

Host receives IP datagrams. Each datagram has source and destination IP addresses. Each datagram carries one transport-layer segment with source and destination port numbers.

A host uses IP addresses and port numbers to direct the segments to appropriate sockets.

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled.png)

UDP segments are demultiplexed at the receiver side by their port numbers. 

TCP segments are demultiplexed at the receiver end by four tuples representing each socket. (src/dest IP, src/dest port).

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled%201.png)

![https://miro.medium.com/v2/resize:fit:786/format:webp/1*tUOFoRHeIut2-uU_p7QFSg.png](https://miro.medium.com/v2/resize:fit:786/format:webp/1*tUOFoRHeIut2-uU_p7QFSg.png)

![https://miro.medium.com/v2/resize:fit:828/format:webp/1*BIt3bTjMbNOwLjN5dM_DDw.png](https://miro.medium.com/v2/resize:fit:828/format:webp/1*BIt3bTjMbNOwLjN5dM_DDw.png)

---

# Principles of Reliable Data Transfer

How do two distributed entities reliably communicate through a network??? Hoowww???

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled%202.png)

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled%203.png)

Complexity of reliable data transfer protocol will depend on characteristics of unreliable channel (lose, corrupt, reorder data).

Sender and receiver do not know the “state” of each other eg : was a message received. So for solving this the receiver should somehow inform the sender that some packets are lost. 

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled%204.png)

# Designing reliable data transfer model.

We will: 

- incrementally develop sender, receiver sides of reliable data transfer protocol (rdt)
- consider only unidirectional data transfer
    - but control info will flow on both directions.
- Use finite state machines (FSM) to specify sender, receiver

[What is a Finite State Machine?](https://medium.com/@mlbors/what-is-a-finite-state-machine-6d8dec727e2c)

## RDT1.0: reliable transfer over a reliable channel.

- underlying channel perfectly reliable
    - no bit errors
    - no loss of packets
- Separate FSM for sender, receiver
    - sender sends data into underlying channel
    - receiver reads data from underlying channel

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled%205.png)

![Untitled](Transport%20Layer%20c79daeb2de194c48afb96217370303c5/Untitled%206.png)

Sender has only one state, it waits for call from above. (Application layer calls transport layer).

Transport layer after getting this call from the application layer, creates a packet out of this message.

And then send the packet to the underlying unreliable channel using udt_send(packet).

Receiver also has only one state. It wait for call from below. rd_rcv is the event, on event extract and deliver_data is called and data is passed to the application layer.

rdt1.0 is very easy. Because the underlying channel is perfectly reliable. So it didn't have to worry about packet corruption, lost, wrong order. But that’s not the case in the real world. Channels are not always reliable. 

## RDT2.0: Channel with bit errors

- underlying channel may flip bits in packet
    - checksum to detect bit erros
- the question: **how to recover from errors?**

How do humans recover from “errors” during conversations? Well they ask again. Simple.