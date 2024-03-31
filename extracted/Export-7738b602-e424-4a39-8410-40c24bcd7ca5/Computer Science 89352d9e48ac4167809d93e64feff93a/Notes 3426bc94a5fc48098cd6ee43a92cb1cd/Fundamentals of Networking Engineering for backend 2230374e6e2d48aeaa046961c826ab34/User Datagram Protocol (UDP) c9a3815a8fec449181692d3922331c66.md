# User Datagram Protocol (UDP)

- Layer 4 protocol (Transport)
- Ability to address processes in a host using ports
- Prior communication is not required (no 3-way handshake)
- Stateless no knowledge is stored on the host
- 8-byte header datagram

## Use cases

- Video streaming
- VPN
- DNS
- WebRTC

## UDP Datagram Anatomy

- UDP header is 8 bytes {Source port - 2 byte, Destination port - 2 byte, Length - 2 byte, Checksum - 2 byte}
- Datagram slides into an IP packet as “data”. (From layer 3 to layer 4)
    - The Protocol IP header is changed to UDP. [Protocol : Routers are essentially a layer 3 device. They wouldn’t need to parse the entire packet to find the protocol. IP packet already has a 8 bit for setting the protocol. (ICMP, TCP, UDP)](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0.md)
- A port is 16-bit (0 - 65535)

![Untitled](User%20Datagram%20Protocol%20(UDP)%20c9a3815a8fec449181692d3922331c66/Untitled.png)

Each process has a unique port. Length represents the length of the data, Checksum represents the checksum of the whole datagram.

## Pros & Cons

### Pros

- Simple protocol
- The header size of small so datagrams are small
- Uses less bandwidth
- Stateless
- Consumes less memory
- Low latency - no handshake, order, retransmission or guaranteed delivery.

### Cons

- No acknowledgement
- No guaranteed delivery
- Connection-less anyone can send data without prior knowledge
- No flow control
- No congestion control
- No ordered packets
- Security - can be easily spoofed.