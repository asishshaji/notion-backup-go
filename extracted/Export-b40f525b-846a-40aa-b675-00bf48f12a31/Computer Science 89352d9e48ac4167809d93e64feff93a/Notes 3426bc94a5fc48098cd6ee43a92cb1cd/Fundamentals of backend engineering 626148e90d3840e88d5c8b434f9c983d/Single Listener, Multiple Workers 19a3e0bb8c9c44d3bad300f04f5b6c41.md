# Single Listener, Multiple Workers

eg : Memcached

![https://miro.medium.com/v2/resize:fit:1100/format:webp/1*RyAQMEP7iR5YCi7JqMpBrA.png](https://miro.medium.com/v2/resize:fit:1100/format:webp/1*RyAQMEP7iR5YCi7JqMpBrA.png)

Multi threading can be used to build performant backend applications that takes advantage of all CPU cores. The price you pay as an engineer is complexity (nothing is free) but sometimes its worth it.

In this architecture, you still have a single listener thread and that same thread usually also accepts connections. The difference here is each accepted connection is handed over to another thread.

Once a thread has a connection it will call read on the connection file descriptor