# Internet Protocol

Regardless of the protocol that you use the data will eventually sit into something called an IP packet. 

## IP Building Blocks

Packet - Layer 3 {dest_ip, source_ip}  assume for now

### IP Address

- Layer 3 property
- Can be set automatically or statically
- Network and Host portion
- 4 bytes in IPv4 - 32 bits

### Network vs Host

- a.b.c.d/x (a,b,c,d are integers) x is the network bit and the remaining are the hosts
- eg 192.168.254.0/24
- The first 3 bytes are network and the rest 1 byte is for the host.
- This means we can have 2^24 networks and each network has 2^8 (255) hosts
- Also called a subnet.

### Subnet Mask

- 192.168.254.0/24 is also called a subnet
- The subnet has a mask 255.255.255.0
- A subnet mask is used to determine whether an IP is in the same subnet.

### Default Gateway

- Most networks consist of hosts and a default gateway.
- When host A wants to talk to B directly if both are in the same subnet(They can use the MAC address to communicate).
- Otherwise, A sends it to some who might know, the gateway.
- The gateway has an IP address and each host should know its gateway.
- Gateway happens to be a device which has multiple network interfaces.

![Untitled](Internet%20Protocol%20552abfeaa5664fb6859440abdecab957/Untitled.png)

![Untitled](Internet%20Protocol%20552abfeaa5664fb6859440abdecab957/Untitled%201.png)

## IP Packet

- Has two sections, header and data
- Header 20 bytes to 60 bytes
- Data section can go up to 65536

![Untitled](Internet%20Protocol%20552abfeaa5664fb6859440abdecab957/Untitled%202.png)

![Untitled](Internet%20Protocol%20552abfeaa5664fb6859440abdecab957/Untitled%203.png)