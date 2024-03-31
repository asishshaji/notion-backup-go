# Internet Protocol

- Layer 3 property
- Can be set automatically or statically
- Network and Host portion
- 4 byte in IPv4 - 32 bits
- first 24 bits are the network part and the rest 8 bit is host part
- A subnet or subnetwork is a smaller network inside a large network
- Subnet mask is used to determine whether an IP is in the same network.

The **Internet Protocol (IP)** is a network layer protocol or collection of rules that allow data packets to be routed and addressed to pass through networks and reach the correct destination.

IP packets are a fundamental unit of data sent from one machine to another, or we can say IP packets are the building blocks of communication between machines. It is usually made of up bytes. IP packets have two main sections, IP header, and data.

## ICMP

- Internet Control Message Protocol
- Lives in layer 3
- Designed for informational messages
    - Host unreachable, port unreachable
    - packet expired
- PING and traceroute uses it
- doesn’t require listeners or ports to be opened.

### Ping

![Untitled](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/Untitled.png)

The routers decrement in the ttl. Here there are total 4 routers between the client and router so ttl is reduced to 96.

If ttl reaches 0 before reaching the destination, it gives host unreachable message and responds back. In that response we get the router mac address at which the ttl turned zero.

![Untitled](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/Untitled%201.png)

So this can be used to find all the routers through which the request reached the server. That’s the idea behind traceroute.

![Untitled](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/Untitled%202.png)

---

## The Anatomy of IP Packet

- has data and header sections
- header is 20 bytes
- data section can go up to 65536

![Imagine like this (P.S not the actual representation)](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/Untitled%203.png)

Imagine like this (P.S not the actual representation)

![Untitled](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/Untitled%204.png)

Header : 4 bytes(row) * 5 col = 20 bytes

IHL : Internet Header length defines how long is the options. 

Total Length : 16 bit data + header

IP Packets need to be fragmented if it cannot fit in the frame.

Identification : id of the fragment 

[Fragmentation In Networking](https://nikhilvkn.medium.com/fragmentation-in-networking-a329e9485e98)

Flag : tells the router to fragment

[IP Fragmentation in Detail - Packet Pushers](https://packetpushers.net/ip-fragmentation-in-detail/)

Time To Live : Every IP packet has 8 bit which represents a counter. To avoid infinite looping. TTL is reduced by 1 when it reaches a router. 

Protocol : Routers are essentially a layer 3 device. They wouldn’t need to parse the entire packet to find the protocol. IP packet already has a 8 bit for setting the protocol. (ICMP, TCP, UDP)

ECN : Explicit Congestion Notification, set by router, lets know the client that the server is facing congestion in the accept queue.

---

[ARP (Address Resolution Protocol)](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/ARP%20(Address%20Resolution%20Protocol)%209d8bdac65fa3479f9017f0bc104b109f.md)

[Capturing IP, ARP and ICMP packets ](Internet%20Protocol%20447ab277f69a4cd294a410412d70fda0/Capturing%20IP,%20ARP%20and%20ICMP%20packets%208d62a3e597be49d285b32a454cc9dfd2.md)