# TCP Segment

- the header is 20 bytes and can go up to 60bytes
- TCP segment slides into the IP packet as data
- Port are 16 bit

![Untitled](TCP%20Segment%2039a4f72e328e445da1f40c1cf5e5f09d/Untitled.png)

- Source port and destination port: 16 bits each
- Sequence number: Every time you send a segment you include the segment number
- The TCP receive **window size** is the amount of received data (in bytes) that can be buffered during a connection.
- The **window** or the **receive window** is the number of bytes that the sender of this segment is willing to receive. A key feature of TCP flow control.

### Maximum Segment Size

- depends on the MTU of the network.
- Usually, 512 bytes can go up to 1460
- The default MTU on the internet is 1500 (results in MSS 1460, 1500 minus the 20 bytes header, 20 bytes in IP)