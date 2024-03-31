# Routing Example

![Untitled](Routing%20Example%20b1191291d59c4c68a00c41e2cfda0a75/Untitled.png)

## A→B

A need to know the mac address of B. So it first looks up in its arp table. If it doesn’t find an entry it send an arp broadcast to the network, asking for the MAC address of B. Some device ACK reply back. B,C, D, R gets this arp broadcast.

> Other thing is, switch is smart. Whenever any request from any device goes through the switch. It sort identifies which port corresponds to which device, since the switch is a layer 2 device. So it will send request from A→B directly.
> 

## D→X

Completely different network. D first checks if the X is in the same subnet by subnet masking. It’s clearly not. So it send the request to gateway. So for that it’ll need the MAC address of gateway. D will do a ARP broadcast and gets the MAC address of gateway. 

IP packet has destination IP of X and destination mac of R. The R gets the message. R checks the dest mac address and verifies it for him. It sees that the packet wants to go to X’s IP. The router sends an ARP, and gets X’s mac address. Then R routes the packet to X. 

X then replies back to R. R then routes back to D.

## B→G