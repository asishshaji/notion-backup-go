# Reading and Sending data

## Send and Receive buffers

- The kernel actually creates 4 queues(SYN, Accept, Send and Receive Buffer)
- Client sends data on establishing connection, ie writes to socket file descriptor.
    - NIC receive the frame, frame is parsed as IP packet and is given to the OS.
    - Segment has the port. And the data is put in receive queue.
    - Kernel acks and update the window.
    - Client sends more data
    - If the buffer is full, kernel drops the data
- Therefore backend application should call read() quickly to remove the data from the buffer
- App calls read to copy the data from the receive buffer into its process memory.
- The data is encrypted. It’s the job of backend application to decrypt the data not the kernel.
- Now the app reads the request.

![Untitled](Reading%20and%20Sending%20data%204de49aacbfa847e7a9273d9aef0e7ec3/Untitled.png)

When you want to write the response back, you call send() function from your backend application. This writes the response into the send buffer. The response is encrypted raw bytes.

![Untitled](Reading%20and%20Sending%20data%204de49aacbfa847e7a9273d9aef0e7ec3/Untitled%201.png)