# Flow Control

## How much data the receiver can handle?

Flow control tries to answer that.

![It is very sloooowwww….](Flow%20Control%20997adfc2f4634d2d8ab789d69ea41c0a/Untitled.png)

It is very sloooowwww….

We need to be able to send multiple segments and receive a single ack.

![Untitled](Flow%20Control%20997adfc2f4634d2d8ab789d69ea41c0a/Untitled%201.png)

The question is, how much A can send? how much B can handle?

Flow control is the window that A needs to maintain so that B is not overloaded.

![Untitled](Flow%20Control%20997adfc2f4634d2d8ab789d69ea41c0a/Untitled%202.png)

If the receiver (B) is overwhelmed with segments (ie, the receive buffer is about to be filled), the packets will be lost. So the solution here is to let the sender know the server is overwhelmed. 

![Untitled](Flow%20Control%20997adfc2f4634d2d8ab789d69ea41c0a/Untitled%203.png)

Window size is included in ack. 

![Untitled](Flow%20Control%20997adfc2f4634d2d8ab789d69ea41c0a/Untitled%204.png)

![Untitled](Flow%20Control%20997adfc2f4634d2d8ab789d69ea41c0a/Untitled%205.png)