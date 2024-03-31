# A TCP Flow and HTTP Flow Guide: Changing HTTP Messages | HackerNoon

URL: https://hackernoon.com/a-tcp-flow-and-http-flow-guide-changing-http-messages
Category: General, Network

![https://hackernoon.imgix.net/images/SnVhdDNm3fMGfftKZO7mnXldvQm2-2f03bnp.jpeg?mark-pad=0&mark=http://hackernoon.imgix.net/HackerNoon%20Rounded%20Horizontal.png?w=400](https://hackernoon.imgix.net/images/SnVhdDNm3fMGfftKZO7mnXldvQm2-2f03bnp.jpeg?mark-pad=0&mark=http://hackernoon.imgix.net/HackerNoon%20Rounded%20Horizontal.png?w=400)

---

## Too Long; Didn't Read

The net library has a special instance Server. Let’s use that to create a server instance and tell it to listen to the port 3000 on the host. To do this, the client and server need to establish a TCP connection first. The thing is, HTTP messaging sessions work over TCP sockets. Once a connection is established, the server and client can communicate with the help of the sockets. In the old version of the protocol, there was a separate request for each request for a new connection to be established.

### Companies Mentioned

![A%20TCP%20Flow%20and%20HTTP%20Flow%20Guide%20Changing%20HTTP%20Messa%202bc4381bcdcb4ff2a7540bfc54e44bc2/SnVhdDNm3fMGfftKZO7mnXldvQm2-2f03bnp.jpeg](A%20TCP%20Flow%20and%20HTTP%20Flow%20Guide%20Changing%20HTTP%20Messa%202bc4381bcdcb4ff2a7540bfc54e44bc2/SnVhdDNm3fMGfftKZO7mnXldvQm2-2f03bnp.jpeg)

**Georgii Kliukovkin**

Imagine this situation: we have an HTTP server and a client working with it. Let’s take a closer look under the hood at what the process of changing HTTP messages is.

## TCP Flow

When we speak about HTTP servers and clients, we mean communications between server and client by sending HTTP messages between them. But, to do this, the client and server need to establish a TCP connection first. Let’s create an instance of a server and establish a TCP connection. We will use only utils from libuv, to understand the process under the loupe.

The net library has a special instance Server. Let’s use that to create a server instance and tell it to listen to the port 3000 on the host 0.0.0.0 which is literally called localhost.

```
const net = require('net');
const HOST = '0.0.0.0';
const PORT = 3000;

const server = new net.Server();

server.listen(PORT, HOST);

```

To check that everything works we can add handlers on “connection” and “listening” events or life cycles of the server. Because the Server inherits Emmiter class, we can do this:

```
server.on("listening", () => {
  console.log("server is listening");
})
server.on("connection", () => {
  console.log("connection established");
});

```

Now you can establish a TCP connection over port 3000. To do this let’s use telnet.

❯ telnet localhost 3000

Trying ::1...

Connected to localhost.

Escape character is '^]'.

Great. Connection is established. Now we need a way to read the data. The thing is, HTTP messaging sessions work over TCP sockets. From Node.JS documentation:

[<net.Socket>](https://nodejs.org/api/net.html?ref=hackernoon.com#class-netsocket) The connection object

We can get access to it while creating a server.

```
const net = require('net');

const HOST = '0.0.0.0';
const PORT = 3000;

const server = new net.Server((socket) => {

});

```

So, once again. We create a server and ask it to listen to the port 3000 of the localhost. Also, we asked the socket, which handles this port to log received data. Now we can add some event handlers to the socket:

```
socket.on("data", (data) => {
  console.log(`receive data: ${data}`);
});

```

Now we can, once again, connect to the server and send a message.

❯ telnet localhost 3000

Trying ::1...

Connected to localhost.

Escape character is '^]'.

**hello**

In the console, we will see this message `receive data: hello`.

One important thing to understand, at this juncture, is that once a TCP connection is established, the client and server can communicate with the help of the sockets. That’s right, sockets appear on both sides.

To understand this better, instead of using telnet let’s write a client using the net library.

Here is the code of our server and client.

```
const net = require("net");
const HOST = "0.0.0.0";
const PORT = 3000;

const server = new net.Server((socket) => {
  socket.once("data", (data) => {
    console.log(`receive from client: "${data}"`);
    socket.write("hello");
  });
});

server.listen(PORT, HOST);

const client = new net.Socket();
client.connect(PORT, HOST);
client.write("hello");
client.on("data", (data) => {
  console.log(`receive from server: "${data}"`);
});

```

Web browsers, such as Google Chrome are basically the same client, which makes for all this magic under the hood. If you run `lsof -n -i TCP | grep Google` you will see how many TCP connections are established by your Chrome browser.

Another important thing here is there can be multiple clients connected to this server's TCP socket.

If clients make multiple TCP connections to the server then the client OS must generate different random source ports, so that server can see them as unique connections.

## Difference between HTTP versions

The difference worth noting regards the `Connection` header. In the old version of the protocol, for each message exchange, there was a separate connection established. Client request smth from the server, TCP connection established client got response and connection got closed. Next time, the client would need to send another request for a new TCP connection to be established. Not very optimal, right? So that’s why `Connection` the header was created. If we pass `keep-alive` using `Connection` , the header will tell the server not to quit this connection.

## HTTP flow

HTTP is an application layer protocol. The flow is quite simple:

1. You type a URL in your browser and press Enter
2. The browser looks up the IP address for the domain
3. The browser initiates a TCP connection with the server
4. The browser sends the HTTP request to the server
5. Server processes request and sends back a response
6. The browser renders the content

## What is an HTTP message?

> 
> 
> 
> HTTP messages are simple, line-oriented sequences of characters. Because they are plain text, not binary, they are easy for humans to read and write.
> 

**HTTP messages consist of three parts:**

***Start line***

The first line of the message is the start line, indicating what a request requires to be done or what happened to warrant a response.

Request messages ask servers to do something to a resource. The start line for a request message, or *request line*, contains a method describing what operation the server should perform and a request URL describing the resource on which to perform the operation. The request line also includes an HTTP *version* which tells the server what dialect of HTTP the client is speaking.

***Header fields***

Zero or more header fields follow the start line. Each header field consists of a name and a value, separated by a colon (:) for easy parsing. The headers end with a blank line. Adding a header field is as easy as adding another line.

***Body***

After the blank line, an optional message body follows, containing any kind of data. Request bodies carry data to the web server; response bodies carry data back to the client. Unlike the start lines and headers, which are textual and structured, the body can contain arbitrary binary data (e.g., images, videos, audio tracks, software applications). Of course, the body can also contain text.

## Status codes

There is no need to know all status codes, but you need to, at least, know what category each code belongs to.

| Category | Examples | Details |
| --- | --- | --- |
| 1xx (informational) | N/A | HTTP/1.0 doesn’t define any 1xx status codes but does define the category. |
| 2xx (successful) | 200 OK | This code is the standard response code for a successful request. |
|  
 | 201 Created | This code should be returned for a POST request. |
|  
 | 204 No content | The request has been accepted and processed, butthere’’s no BODY response to send back. |
| 3xx(redirection) | 300 Multiple choices | This code isn’t used directly. It explains that the 3xx category implies that the resource is available at one (or more) locations, and the exact response provides more details on where it is. |
|  
 | 301 Moved permanently | The Location HTTP response header should provide the new URL of the resource. |
|  
 | 302 Moved temporarily | The Location HTTP response header should provide the new URL of the resource. |
| 4xx(client error) | 400 Bad request | The request couldn’t be understood and should be changed before resending. |
|  
 | 401 Unauthorized | This code usually means that you’re not authenticated. |
|  
 | 403 Forbidden | This code usually means that you’re authenticated, but your credentials don’t have access. |
| 5xx(server error) | 500 Internal server error | The request couldn’t be completed due to a server-side error. |
|  
 | 503 Service unavailable | The server is unable to fulfill the request, perhaps because the server is overloaded or down for maintenance. |

## Important headers

Connection: keep-alive

Host: www.example.com -→ ***this is for virtual hosting***

User-Agent: MyAwesomeWebBrowser

## Dictionary

### Cookie, or session cookie

A **session cookie** is a temporary cookie that keeps track of settings and preferences as a user navigates a site.

A session cookie is deleted when the user exits the browser.

The cookie contains an arbitrary list of *name=value* information, and it is attached to the user using the **Set-Cookie** or **Set-Cookie2** HTTP response headers. Cookies can be restricted by domain or path

### URL

URLs describe the specific location of a resource on a particular server. They tell you exactly how to fetch a resource from a precise, fixed location. URL consists of *protocol, server, and local resource*

protocols: HTTPS, HTTP, FTP, and so on

server domain → DNS

local resource - e.g. /index.html

### What does HTTP port mean?

Well, actually it is not an HTTP port. It is a socket. For details take a look wiki

[https://en.wikipedia.org/wiki/Port_(computer_networking)](https://en.wikipedia.org/wiki/Port_(computer_networking)?ref=hackernoon.com)

Here are the default ones:

22 - ssh

80 - http

443 - https

### Session

It is quite obvious why it is important to know what the user server is talking to (*e*.*g.* *targeted recommendations, session tracking, etc).*

Ways to store session:

- inside [cookie](https://app.hackernoon.com/mobile/2knVUMbc3a0gi05HHZ4P?ref=hackernoon.com#h-cookie-or-session-cookie)
- dedicated datastore(e.g. Redis)
- sticky sessions with mapping on load balancer

### Virtual hosting

An IP address consists of 4 integers from 0 to 255. We can calculate how many sites can be placed separately with a dedicated IP address. But the reality is we have much more sites than unique IP addresses. This is why virtual hosting was invented. The key is to provide the ability to place several sites on one IP address. But how will the server based on this IP know which HTML page to provide? With host header that will contain required domain like google.com

### Idempotent HTTP methods

An HTTP method is **idempotent** if an identical request can be made once or several times in a row with the same effect while leaving the server in the same state. In other words, an idempotent method should not have any side effects (except for keeping statistics). Implemented correctly, the `GET`, `HEAD`, `PUT`, and `DELETE` methods are **idempotent**, but not the `POST` method.

## Used materials:

[https://www.oreilly.com/library/view/http-the-definitive/1565925092/ -](https://www.oreilly.com/library/view/http-the-definitive/1565925092/ - bible :)) HTTP bible :)

[https://developer.mozilla.org/en-US/docs/Glossary/Idempotent](https://developer.mozilla.org/en-US/docs/Glossary/Idempotent?ref=hackernoon.com)

[https://en.wikipedia.org/wiki/Port_(computer_networking)](https://en.wikipedia.org/wiki/Port_(computer_networking)?ref=hackernoon.com)

[https://aws.amazon.com/blogs/mobile/what-happens-when-you-type-a-url-into-your-browser/](https://aws.amazon.com/blogs/mobile/what-happens-when-you-type-a-url-into-your-browser/?ref=hackernoon.com)

[The Web Development Writing Contest](https://www.contests.hackernoon.com/web-development-writing-contest?ref=hackernoon.com) is brought to you by IONOS in partnership with HackerNoon. Publish your stories on #web-development to win from $6000!!!

Experience coding the way it's meant to be with IONOS Cloud Cubes - powerful KVM-based virtual servers designed for maximum flexibility and on-demand scalability. Start for free: .