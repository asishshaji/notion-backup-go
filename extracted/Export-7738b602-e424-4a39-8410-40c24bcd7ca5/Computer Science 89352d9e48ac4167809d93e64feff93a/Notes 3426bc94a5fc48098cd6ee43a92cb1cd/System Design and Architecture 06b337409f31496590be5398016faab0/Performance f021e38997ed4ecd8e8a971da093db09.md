# Performance

![Untitled](Performance%20f021e38997ed4ecd8e8a971da093db09/Untitled.png)

![Untitled](Performance%20f021e38997ed4ecd8e8a971da093db09/Untitled%201.png)

Consider the database to be an RDBMS

Performance measures how fast or responsive the system is under a given workload (Backend data, request volume) or given hardware (kind, capacity). The performance problem results from some queue building somewhere, eg: Network socket queue, DB IO queue, OS run queue etc…

Reasons for the queue build-up are inefficient slow processing, serial resource access, and limited resource capacity. 

![Untitled](Performance%20f021e38997ed4ecd8e8a971da093db09/Untitled%202.png)

## Performance Principles

- Efficiency
    - Efficient Resource Utilisation
        - IO- Memory, Network, Disk
        - CPU
    - Efficient Logic
        - Algorithms
        - DB Queries
    - Efficient Data Storage
        - Data Structures
        - DB Schema
    - Caching
- Concurrency
    - Hardware
    - Software
        - Queuing
        - Coherence
- Capacity

Latency depends on wait/idle time and processing time. Throughput is the rate of requests that a system can handle.

![Untitled](Performance%20f021e38997ed4ecd8e8a971da093db09/Untitled%203.png)

[Serial Request](Performance%20f021e38997ed4ecd8e8a971da093db09/Serial%20Request%20fd1ad98cc30346ada407ca6ae4f8b719.md)

[Parallel Request ](Performance%20f021e38997ed4ecd8e8a971da093db09/Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0.md)

[Caching](Performance%20f021e38997ed4ecd8e8a971da093db09/Caching%209a3a9a858cb34792a6710c92f882b2a5.md)