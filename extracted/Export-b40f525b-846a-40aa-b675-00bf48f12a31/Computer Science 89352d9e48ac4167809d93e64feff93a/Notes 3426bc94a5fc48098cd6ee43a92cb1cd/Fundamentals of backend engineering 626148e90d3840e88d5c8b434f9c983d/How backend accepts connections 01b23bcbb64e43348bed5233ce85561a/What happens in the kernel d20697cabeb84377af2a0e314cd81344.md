# What happens in the kernel

- Kernel creates a socket and two queues SYN and accept
- These queues live in kernel
- Client sends a SYN
- Kernel adds SYN(With the source ip/port , dest ip/port) to the SYN queue and replies with SYN/ACK
- Client replies with ACK
- Kernel gets the ACK, and find the SYN matching that ACK (the source port is unique so that way you can match and find) and then connection is moved to the accept queue and is removed from the SYN queue
- Backend calls the Accept function and that removes the connection from the accept queue to the application process memory.
- A file descriptor for the connection is created.

![Untitled](What%20happens%20in%20the%20kernel%20d20697cabeb84377af2a0e314cd81344/Untitled.png)