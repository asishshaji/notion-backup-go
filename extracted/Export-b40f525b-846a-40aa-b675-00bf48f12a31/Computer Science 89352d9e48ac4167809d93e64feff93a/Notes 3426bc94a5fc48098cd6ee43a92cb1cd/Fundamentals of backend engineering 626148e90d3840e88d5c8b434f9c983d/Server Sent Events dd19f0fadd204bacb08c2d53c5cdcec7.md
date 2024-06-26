# Server Sent Events

[Golang and Server-Sent Events (SSE)](https://dev.to/rafaelgfirmino/golang-and-sse-3l56)

![https://res.cloudinary.com/practicaldev/image/fetch/s--RQ--xQUi--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ot1nf5wz5zaacyjo2a2z.png](https://res.cloudinary.com/practicaldev/image/fetch/s--RQ--xQUi--/c_imagga_scale,f_auto,fl_progressive,h_420,q_auto,w_1000/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/ot1nf5wz5zaacyjo2a2z.png)

> *SSE is a lightweight protocol on top of HTTP Streaming that allows super light-weight subscribe-only capabilities to clients. Unlike WebSockets, SSE does not provide the capability of two way communication, but can be used by the server to push data to the clients in real-time.*
> 

Polling strategies might not be the best approach to doing things, especially when we have a situation where the server just has to send new information to its clients at specific intervals or when a change occurs in its data state.

SSE is unidirectional. The data direction of data flow is only from server to client. The other way is not possible.

In a nutshell, we could define the idea behind SSE as follows:

> A web app “subscribes” to a stream of updates generated by a server and, whenever a new event occurs, a notification is sent to the client.
> 

For a server to send a stream of events, the data and response have to be in a very specific format:

1. The content type of the response header should be `text/event-stream` .
2. The response should contain a `data` field followed by the message you’d like to send and should always be terminated by a `double carriage return` such as follows:

```
data: example message\n\n
```

3. One can send data from multiple data points through a single event, but each message should be a value with a `data` key and except for the last message each data should be terminated by a single carriage return. The last message should be terminated by a double carriage return. For example:

```
data: message item 1\n
data: message item 2\n
data: message item 3\n\n
```

These multiple data from a single event could then be parsed by the client to split them into separate values and can be utilized by them for their needs.

```jsx
// this is just to demo the concept of SSE, not intended for production usage.

const http = require("http");
var os = require("os");

const host = "127.0.0.1";
const port = 8080;

// A simple dataSource that changes over time
let dataSource = 0;
const updateDataSource = () => {
  const delta = Math.random();
  dataSource += delta;
}

const requestListener = function (req, res) {
  if (req.url === '/ticker') {
    res.statusCode = 200;
    res.setHeader("Access-Control-Allow-Origin", "*");
    res.setHeader("Cache-Control", "no-cache");
    res.setHeader("connection", "keep-alive");
    res.setHeader("Content-Type", "text/event-stream");

    setInterval(() => {
      const data = JSON.stringify({ ticker: dataSource });
      res.write(`id: ${(new Date()).toLocaleTimeString()}\ndata: ${data}\n\n`);
    }, 1000);
  } else {
    res.statusCode = 404;
    res.end("resource does not exist");
  }
};

const server = http.createServer(requestListener);
server.listen(port, host, () => {
  setInterval(()=> updateDataSource(), 500);
  console.log(`server running at http://${host}:${port}`);
});
```

# **Limitations of SSE:**

The following are some of the limitations of SSE:

1. By design, it is only meant to be used for unidirectional data flow, and that too only from the server to the client. If you require a bidirectional data flow, then WebSockets would be the ideal choice.
2. If you are using the `HTTP/1` protocol, which is the default that many apps use at the time of writing this article, there is a limitation on the number of subscribers from a single active browser window. ( Only 6 tabs can connect to a specific stream per domain). A seventh simultaneous tab would fail to subscribe to the stream. This could be overcome by using the `HTTP/2`protocol, where the default limit for simultaneous subscription per domain on a single active browser window is 100 tabs.
3. 

![https://miro.medium.com/v2/resize:fit:1100/format:webp/1*ZOvd7h41rtYPVvxUcyP5Kw.png](https://miro.medium.com/v2/resize:fit:1100/format:webp/1*ZOvd7h41rtYPVvxUcyP5Kw.png)