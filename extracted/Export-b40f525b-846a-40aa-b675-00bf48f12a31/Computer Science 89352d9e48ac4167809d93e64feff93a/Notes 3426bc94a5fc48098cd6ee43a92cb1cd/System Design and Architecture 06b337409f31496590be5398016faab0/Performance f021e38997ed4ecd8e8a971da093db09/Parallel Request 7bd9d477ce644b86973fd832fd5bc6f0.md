# Parallel Request

Owner: Asish Shaji Thomas
Last edited time: August 29, 2023 6:21 AM

[SPE Fundamentals · SPE BoK](https://tangowhisky37.github.io/PracticalPerformanceAnalyst/spe_fundamentals/)

## [Amdahl’s Law](https://learn.saylor.org/mod/page/view.php?id=27029)

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled.png)

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%201.png)

In the second graph, speed/throughput vs no of processors, Each line represents a parallel portion, the line close to the y axis has 95% parallel execution and rest 5% serially. Even if the serial port is very low, the graph will at some point become a line.

> Speedup measures the ratio of performance between two objects
> 

In general terms, **Amdahl’s Law** states that in parallelization, if **P** is the proportion of a system or program that can be made parallel, and **1-P** is the proportion that remains serial, then the maximum speedup **S(N)** that can be achieved using **N** processors is:

$$
  S(N)=1/((1-P)+(P/N))
$$

## Gunther’s universal scalability law

[Applying the Universal Scalability Law to organisations | the morning paper](https://blog.acolyer.org/2015/04/29/applying-the-universal-scalability-law-to-organisations/)

Whenever there are sequential requests, it will end up in a queue, for example: if we have a lock on a shared resource, the first request acquires the lock.

The subsequent requests are usually queued when they try to acquire the lock. Queuing is one of the causes of the throughput of a system.

There are two reasons which limit the concurrency of a system. 

- queuing
- coherence

Amdhal’s law only deals with queueing. Universal law deals with queueing and coherence.

### Coherence

Let’s assume that we have a 3 core system. Threads are executed on these cores. And you have a volatile variable declared in your program. When you declare a variable as volatile, the changes that you make one thread must be reflected on other threads. This comes at a performance cost.

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%202.png)

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%203.png)

So in our application if we are using lot of such variables, then it comes at a cost.

In queueing it never brings down your overall speedup. But in case of coherence, it has this annoying ability to bring down your speedup. 

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%204.png)

So for a system to be highly concurrent, you need to bring down both queueing and coherence. 

## Shared Resource Contention

- When requests are made to the server, if the server is busy, then the request is **queued to the listen or accept queue.**

[How backend accepts connections](../../Fundamentals%20of%20backend%20engineering%20626148e90d3840e88d5c8b434f9c983d/How%20backend%20accepts%20connections%2001b23bcbb64e43348bed5233ce85561a.md)

- After the request is accepted and it needs to be processed. So we need to create a thread for the request. Threads are limited for the server process. If the system is overloaded then there will be a **contention for a thread from fixed-size thread pool.**

Now the request is being handled by the thread. 

- If the request is forwarded to the database or service then there will be **contention at the connection pool**.
- If we have a lot of database writes there can be **contention at the disk level.**
- The biggest cause of contention in any application is **locks**. Locks make the code synchronous.

### Minimizing shared resource contention

- For disks, you can use RAID disks.

[RAID (Redundant Arrays of Independent Disks) - GeeksforGeeks](https://www.geeksforgeeks.org/raid-redundant-arrays-of-independent-disks/)

- We generally need not do anything about the listen queue. You will need to go the operating system kernel to modify the queue size. It’s the last thing you do.
- If the rate at which the requests are consumed from the accept queue is very slow, then we can increase the size of the thread pool.
- Smaller thread pool will increase contention
- Thread Pool Size can depend on
    - Wait Time: If the IO to the services is taking significant time, then we can **increase the size of the thread pool**. Why??? Because these IO operations can cause context switches and threads will get time
    - CPU Time; If the application is doing a lot of calculations on the CPU, it doesn't make sense to increase the thread pool size, **it should be more or less close to the number of cores**.
    - Number Of CPU
- Connection Pool size
    - If the Thread Pool size is 100, then the connection pool should have at least 100 connections in it.
    
    [Improve database performance with connection pooling](https://stackoverflow.blog/2020/10/14/improve-database-performance-with-connection-pooling/)
    

### Minimizing locking-related contention

- Reduce the duration for which a lock is held
    - Move out the code, out of the critical section, that doesn't require lock (especially locks)
    - It increases the chances of context switching
        
        # Does adding IO operation in critical section increases the change of context switching?
        
        Yes, adding Input/Output (IO) operations in a critical section can increase the likelihood of context switching. 
        
        Now, when an IO operation is added to a critical section, the following scenario can occur:
        
        1. Thread A enters the critical section and starts performing some IO operation, such as reading from a file or waiting for network data.
        2. While Thread A is waiting for the IO operation to complete, the CPU is essentially idle during that time. The operating system might decide to perform a context switch and allow another ready thread (Thread B) to run on the CPU, utilizing the CPU time more efficiently.
        3. Once the IO operation of Thread A completes, the operating system will schedule Thread A again to continue its execution. This involves another context switch from Thread B back to Thread A.
        
        The additional context switches between threads due to IO operations within the critical section can lead to an overhead in terms of performance. The more time threads spend waiting for IO operations to complete, the more context switches are likely to happen, and this can potentially impact the overall efficiency of the system.
        
        To minimize the impact of IO operations in critical sections, developers often use techniques like asynchronous IO or non-blocking IO to allow threads to continue executing other tasks while waiting for the IO operation to finish. This approach reduces the idle time and context switching caused by blocking IO operations, leading to better overall performance.
        
    - Lock Splitting is used to split locks into lower granularity locks that are experiencing moderate contention.
    - Locking Striping is used to split locks for each partition of data like in Concurrent Hashmap.
    
    [Introduction to Lock Striping | Baeldung](https://www.baeldung.com/java-lock-stripping)
    

- Replace exclusive locks with coordination mechanisms
    - use ReadWriteLock/Stamped locks
        - Multiple read locks can be placed on a variable, but write locks are blocked till all the read locks are unlocked.
    
    [How to understand & implement read-write locks and bounded buffers](https://www.youtube.com/watch?v=7zI_4CKk-3Y)
    
    - use atomic variables (protected by Compare and Swap)

When it’s raining, we assume that it’s gonna rain and you take an umbrella. The pressimistic approach is to take the umbrella. Optimistic approach is to assume that it’s not going to rain. I don’t care if it rains. I’ll get an umbrella somehow. I’ll deal with it when it happens. 

When two transaction trying to access the same row (writing). [Isolation](../../Fundamentals%20of%20database%20engineering%203567e651e39c4f51a931697292608bc8/ACID%20aabd352e2e354da1839b540dc978ae49/Isolation%20e46137f8713d46078c0a47425ee05da1.md) 

## Pessimistic Locking

Much larger critical section. Locks are held longer. Threads must wait to acquire locks.

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%205.png)

FOR UPDATE is used to acquire the lock in the above figure.

This locking is used when contention is very high.

[Pessimistic concurrency control vs Optimistic concurrency control in Database Systems Explained](https://www.youtube.com/watch?v=I8IlO0hCSgY)

## Optimistic Locking

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%206.png)

If the value changes, then you again retry to update.

Used when contention is low to moderate.

[Topic 06, Part 07 - Optimistic vs. Pessimistic Locking](https://www.youtube.com/watch?v=x1wZPXKz40k&ab_channel=Dr.DanielSoper)

[Optimistic vs. Pessimistic Locking - Vlad Mihalcea](https://vladmihalcea.com/optimistic-vs-pessimistic-locking/)

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%207.png)

# Compare And Swap (CAS)

It’s an optimistic locking mechanism supported by OS. All modern hardware (CPU) supports this. Languages build CAS on top of it.

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%208.png)

# Deadlocks

Deadlock is defined as a situation where a set of processes are blocked because each process holding a resource and waiting to acquire a resource held by another process.

Two reasons

- Lock Ordering Related
    - Result of threads trying to acquire multiple locks
        - Simultaneous money transfer from X and Y accounts by thread T1 and T2.
        - X is locked by T1 and Y is locked by T2
        
        ![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%209.png)
        
        - T1 tries to lock Y, but is blocked since Y is locked by T2.
        - T2 tries to lock X, but is blocked since X is locked by T1.
    - The above-mentioned issue can be solved by acquiring locks in a fixed global order
        - Both threads sort the accounts
        - They lock the variables in the same order
        - T1 and T2 first try to lock X, and then Y
- Request Load Related
    - Let’s assume we have a gateway with 10 threads in a thread pool. The request flow is 1-2-3-4 and back.
    
    ![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%2010.png)
    
    - 10 clients make a request to the gateway. The 10 threads in the thread pool are now used.
    - 1-2 connection is established.
    - 2-3 connections won’t be established, since there are no free threads in the thread pool.
    - This results in a deadlock.
    - This issue happens only in high-load situations.

## Coherence related delays

![Untitled](Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0/Untitled%2011.png)

Let’s assume we have an object shared by multiple threads. Each thread running on the CPU will have its own copy of the object in its cache. Change in one CPU won’t be reflected in the other.

We can add Visibility(volatile) or Lock (synchronized), these guarantee that the object is always read from the main memory and written back to the main memory when updated in the CPU. 

When we modify a variable guarded by synchronized code, any change in one CPU will make the cache of the other CPU dirty. So the CPU will be forced to read the value from the main memory. This latency is called coherence.