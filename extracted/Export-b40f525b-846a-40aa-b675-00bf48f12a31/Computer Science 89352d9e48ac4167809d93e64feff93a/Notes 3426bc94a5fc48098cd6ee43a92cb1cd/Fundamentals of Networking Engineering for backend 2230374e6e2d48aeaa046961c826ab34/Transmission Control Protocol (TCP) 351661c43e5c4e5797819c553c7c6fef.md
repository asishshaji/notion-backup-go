# Transmission Control Protocol (TCP)

[How Does TCP Work?](https://sookocheff.com/post/networking/how-does-tcp-work/)

- Layer 4 protocol
- Ability to address processes in a host using ports
- “Controls” the transmission, unlike UDP which is a firehose.
- Connection : (Some state is saved in both client and server.)
- Requires handshake
- 20 bytes headers segment
- Stateful

## Use cases

- Reliable communication
- Remote shell
- Database connections
- Web communications
- Any bidirectional communication
    
    ![Untitled](Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef/Untitled.png)
    

## TCP Connection

- Connection is a Layer 5 (Session),
- Connection is an agreement between the client and the server
- Must create a connection to send data
- Connection is identified by
    - sourceip - sourceport
    - destinatationip - destinationport
- Can’t send outside of the connection
- sometimes called socket or file descriptor
- Requires 3-way handshake
- Segments are sequenced and ordered
- Segments are acknowledged
- Lost segments are retransmitted

![Untitled](Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef/Untitled%201.png)

![Untitled](Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef/Untitled%202.png)

![Untitled](Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef/Untitled%203.png)

If you acknowledge with 3, that means you have received 1,2 and 3. If acknowledge with 2, that means you only received 1 and 2.

![Untitled](Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef/Untitled%204.png)

What if a segment is lost?

The server in above figure receives two segments, third one doesn’t reach the server somehow. The server acknowledge with 2, so the client knows that the server got only 1 and 2. So then, the server retransmits segment 3.

### Closing a connection

![Untitled](Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef/Untitled%205.png)