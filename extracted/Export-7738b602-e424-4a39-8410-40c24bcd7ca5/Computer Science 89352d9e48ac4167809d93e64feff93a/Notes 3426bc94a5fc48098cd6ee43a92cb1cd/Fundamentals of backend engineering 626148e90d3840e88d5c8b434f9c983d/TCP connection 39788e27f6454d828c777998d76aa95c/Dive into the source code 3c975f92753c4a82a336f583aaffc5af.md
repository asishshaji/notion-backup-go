# Dive into the source code

## Server receives a syn packet

The handler for tcp_protocol is tcp_v4_rcv

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%201.png)

sk_buff is the socket buffer, sk_buff is a metadata structure and does not hold any packet data. All the data is held in associated buffers.

Lookup skbuff.h

When a request is received, the listening socket is looked up. 

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%202.png)

The socket must be in a listening state, ie waiting for new connections

[https://documentation.help/Microchip-TCP.IP-Stack/TCP_STATE.html](https://documentation.help/Microchip-TCP.IP-Stack/TCP_STATE.html)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%203.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%204.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%205.png)

You reach tcp_rcv_state_process.

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%206.png)

This procedure is effectively a state machine applied to a socket that processes a current packet, it first checks a socket state, then runs some logic based on it, and advances the state accordingly.

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%207.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%208.png)

It looks at the state and then runs some logic

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%209.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2010.png)

Here it checks if the packet is syn, if so it calls conn_request, which is tcp_v4_conn_request

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2011.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2012.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2013.png)

tcp_v4_conn_request calls tcp_conn_request. 

We want a scenario where syncookies are not used. 

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2014.png)

If syncookies is 2, then syncookies is enabled. It also checks the receive queue is full

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2015.png)

A request socket is created, the request socket is not a full-blown connection socket used for data exchange, it’s a connection request or a mini socket.

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2016.png)

The ireq_state is TCP_NEW_SYN_RECV, and a listener socket is attached to it and syn packet is saved. tcp_reqsk_record_syn

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2017.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2018.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2019.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2020.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2021.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2022.png)

The request sock is cast to a sock struct and is added to a queue. The length of the queue is also increased.

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2023.png)

At the end, a synack segment is sent.

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2024.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2025.png)

As of now no socket is created, only a mini socket is created, which is request_sock disguised as a normal socket.

## Server receives an ack packet

We are looking at tcp_v4_rcv function because it’s the handler for an TCP requests.

During the syn packet processing step, a socket is looked up. It looks up a socket in an ehash.

[](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af.md) 

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2026.png)

We already have a socket, a temporary request_socket cast to a sock struct which was added during the syn packet processing stage.

Its state is still in TCP_NEW_SYN_RECV [The ireq_state is TCP_NEW_SYN_RECV, and a listener socket is attached to it and syn packet is saved. tcp_reqsk_record_syn](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af.md) 

So inside the tcp_v4_rcv it goes inside the if condition 

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2027.png)

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2028.png)

tcp_check_req is called.

![Untitled](Dive%20into%20the%20source%20code%203c975f92753c4a82a336f583aaffc5af/Untitled%2029.png)

Inside that function, this sort of creates a full-blown socket by cloning a listening socket.

Right after the normal socket is added, the mini socket(request socket) is removed. Right after that in the return, some function called hashdance is called.