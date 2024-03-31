# How backend accepts connections

![Three way handshake](How%20backend%20accepts%20connections%2001b23bcbb64e43348bed5233ce85561a/Untitled.png)

Three way handshake

## Connection Establishment

- Server listens on an address:port
- Client connects to that address and port
- Kernel does the handshake creating a connection
    - When your backend application starts listening it asks the kernel to setup some things
    - It creates a socket
    - When client sends a SYN request, your kernel is responsible for sending SYN/ACK back.
    - After that the handshake completes
    - Then the kernel creates the connection
    
    [What happens in the kernel](How%20backend%20accepts%20connections%2001b23bcbb64e43348bed5233ce85561a/What%20happens%20in%20the%20kernel%20d20697cabeb84377af2a0e314cd81344.md)
    
- Only then your backend application start accepting connection. Moving the connection from kernel to the backend process.

---

[TCP Socket Listen: A Tale of Two Queues (2022)](https://arthurchiao.art/blog/tcp-listen-a-tale-of-two-queues/)

[The Daily DDoS: Ten Days of Massive Attacks](https://blog.cloudflare.com/the-daily-ddos-ten-days-of-massive-attacks/)

![https://yqintl.alicdn.com/0f72fe628f4e77b59f92d149c8584476f8d6fc5d.png](https://yqintl.alicdn.com/0f72fe628f4e77b59f92d149c8584476f8d6fc5d.png)

![Untitled](How%20backend%20accepts%20connections%2001b23bcbb64e43348bed5233ce85561a/Untitled%201.png)

- **A SYN queue**: used for storing connection requests (one connection request per SYN packet);
- **An accept queue**: used to store fully established (but haven’t been `accept()`ed by the server application) connections