# Serial Request

Owner: Asish Shaji Thomas
Verification: Verified
Tags: Backend, Performance
Last edited time: July 24, 2023 1:27 PM

Requests to a server are sent serially.ie, send the first request, wait for its result, and then send the next. Sequential requests.

## Network latency

When we connect from browser to server, the communication is slower, because of hops between different networks which might have different speeds. Intranet communications are much faster.

![Untitled](Serial%20Request%20fd1ad98cc30346ada407ca6ae4f8b719/Untitled.png)

### Why the latency??

- Data is sent over the wire.
- TCP connection (We are considering TCP only). The three-way handshake is very expensive.

[Transmission Control Protocol (TCP)](../../Fundamentals%20of%20Networking%20Engineering%20for%20backend%202230374e6e2d48aeaa046961c826ab34/Transmission%20Control%20Protocol%20(TCP)%20351661c43e5c4e5797819c553c7c6fef.md)

- If you have SSL on top of it, the latency is significant. Creating a TCP connection and then an SSL connection.

### Minimizing network transfer latency

- Connection overheads
    - Reusing connections between server and DB. Maintain a connection pool
    - Preheat the TCP connection with the server and DB.
    - From the client make a persistent connection to the server.
    - If we are using http1.1, make use of domain sharding to get around [the browser’s 6 TCP connections per server](https://www.youtube.com/watch?v=Xkr2nm6UPN8&ab_channel=HusseinNasser).
- Data Transfer overheads
    - Reduce the size of the data
    - Avoid transfers that we don't really need. ie cache the data.
    - Use static data cache at the client.
    - Cache data wherever possible
    - Can use data format and compression.
        - eg use binary communication between services at the backend, like grpc, Avro, thrift
    - Compress the data, and of course, you will have an overhead to compress and decompress the data. This requires CPU cycles, but, those overheads are not costly.
    - SSL Session Caching
        - session parameters are cached. This can reduce the overhead.

## Memory Latency

- For processes using heap memory, the GC will start running, mostly at finite intervals of time, which can cause the running process to slow down.
- Large Heap memory, if the heap memory is large, the extra space is taken from the hard disk, which has some overhead. If the heap memory is large, the GC has to do extra work to clean up the memory.
- The selection of the GC algorithm is important

### Minimizing Memory Latency

![Untitled](Serial%20Request%20fd1ad98cc30346ada407ca6ae4f8b719/Untitled%201.png)

- Control the number of allocations on the heap as possible.
- Weak/Soft References, sort of like setting priorities for the variables on the heap. So if memory runs out, which object should GC start deleting first?

[Strong, Weak, Soft, and Phantom References in Java | Baeldung](https://www.baeldung.com/java-reference-types)

- Multiple smaller processes are better than a large process.
- Garbage collection algorithms

## Disk Latency

All systems one way or another disk access.

![Untitled](Serial%20Request%20fd1ad98cc30346ada407ca6ae4f8b719/Untitled%202.png)

### Minimizing Disk Latency

- Logging
    - Logging are essentially append-only database. It’s a sequential write operation, so that’s faster.
    - Batching log messages are better
        - If we are doing a computation followed by a log, then we proceed to do it 4 times. There are chances for context switching. So if there is a way to group or batch the logs together, then it’s a better approach.
- Web Content Caching
- Page cache, Zero copy
- DB Disk Access
    - add a cache in front of the database
    - Schema Optimization
        - Denormalisation. Why?? If your Finite Buffer Memory of the DB is large, you can denormalise the data opposed to normalised data.
        - Most database related optimisations can be improved by adding indexes.
    - Use SSD for disk IO expensive db.
    - Use higher IOPS for random accesses.

## CPU Latency

Causes are inefficient algorithms and context switching. Inefficent algorithms can be improved by writing better code.

Context switching is not so straightforward. 

![Untitled](Serial%20Request%20fd1ad98cc30346ada407ca6ae4f8b719/Untitled%203.png)

So how did the context switch happen?? 

- It can happen if the process is waiting, ie it’s doing some IO.
- It can also happen if the time slice for the process is expired.

### Minimizing CPU Latency (Improving context switching)

- Use batch/async IO where ever possible
    - Combine writes or reads together to avoid context switching as possible.
    - eg: we log in our application to have a record of various states of our application. When you log, you are performing an IO operation. Let assume this can cause a context switching. So it’s better to batch these together. You can move the logger to a separate thread and let that thread handle the logging operation.
- Single threaded model
    - they are used in v8 engine, nodejs, voltdb, nginx
    - IO operations are delegated to separate async IO threads.
- Thread Pool Size
    - We should not have a large thread pool.
    - Let’s say we have a two core machine, and the thread pool is using 100 threads, these threads most probably won’t even get time to run on these core. 100 : 2
- Multi-process in virtual env

## Latency Costs

![Untitled](Serial%20Request%20fd1ad98cc30346ada407ca6ae4f8b719/Untitled%204.png)