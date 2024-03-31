# NGINX Architecture

URL: https://medium.com/@hnasr/the-architecture-of-nginx-2b32fc0b7877
Category: Proxy

![https://miro.medium.com/v2/resize:fit:1200/1*wDs1IPqtub2KJbMXViLeTg.png](https://miro.medium.com/v2/resize:fit:1200/1*wDs1IPqtub2KJbMXViLeTg.png)

---

![NGINX%20Architecture%2053aae4b774a346388cdeec557aa6629e/1J0NC2PUN71Spe_88oLQjrw.png](NGINX%20Architecture%2053aae4b774a346388cdeec557aa6629e/1J0NC2PUN71Spe_88oLQjrw.png)

NGINX is an open source reverse proxy and web server designed for scale. It exploded in popularity as the first line of defense to backend infrastructures. Whether as a caching layer, load balancer, an API gateway or Web server, NGINX does it all.

In this article I explore the internal architecture of NGINX focusing on the following topics:

- NGINX worker processes
- Connection Management
- Request Processing

# Worker Processes

The major part of NGINX architecture is the master process. When you spin up NGINX, it starts a process that manages the rest of processes in NGINX.

The master process creates two processes for cache management. One for reading the cache from disk and another for refreshing the caches. What interests me the most are the worker processes. These are the power horses that do most of the work in NGINX.

![NGINX%20Architecture%2053aae4b774a346388cdeec557aa6629e/1EwpLk8kb4NjaVNizvQmdbA.png](NGINX%20Architecture%2053aae4b774a346388cdeec557aa6629e/1EwpLk8kb4NjaVNizvQmdbA.png)

When you set NGINX to auto mode (the default), the application spins up a number of worker processes based on how many CPU cores you have. This changes based on whether you have Hyper-Threading enabled or not.

> Hyper-Threading technology allows a single physical core to appear as two logical cores to the operating system. This can boost performance because the operating system can efficiently schedule processes on multiple hardware CPU threads.
> 
> 
> If you have two-core CPU and hyper-threading is enabled, the OS will see four logical cores or hardware threads. This is a choice you make based on performance.
> 

Having a one-to-one cardinality between worker process and CPU is a backend choice. This assumes the worker process and requests are CPU-bound which are not necessarily the case. If requests are I/O bound e.g. reading from another backend or database the process spends most of its time doing I/O and not using the CPU. This of course also depends on the protocol used in communication.

# Connection Management

The worker processes listen on their configured port, say it’s 80 for HTTP. The act of…