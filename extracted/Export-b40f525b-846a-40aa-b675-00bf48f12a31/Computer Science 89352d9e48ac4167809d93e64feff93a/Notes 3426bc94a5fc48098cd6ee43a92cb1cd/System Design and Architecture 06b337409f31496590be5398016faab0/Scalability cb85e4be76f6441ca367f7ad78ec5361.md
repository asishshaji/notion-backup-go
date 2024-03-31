# Scalability

**Performance**: a measure of the latency and throughput of a system. Low latency and high throughput. FIXED LOAD

In **scalability**, we discuss how to increase the system's throughput. VARIABLE LOAD

The ability of a system to increase its throughput by adding more hardware capacity.

Vertical scaling: increase the h/w resources of a system. It is easier to achieve, but limited scalability.

![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled.png)

Horizontal scaling: increase the instances of the application. It’s hard to orchestrate. The number of cases limits your scalability.

![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%201.png)

It’s easier to scale down as well.

## Reverse Proxy

![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%202.png)

In the first figure, the client needs to remember all the IP addresses of the servers. But it’s not practical. So we add a reverse proxy. The client needs to know only the address of the reverse proxy. Reverse Proxy can also act as a load balancer. 

---

![System for scalability discussion](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%203.png)

System for scalability discussion

## Scalability Principles

- Decentralization - Monolith is an antipattern for scalability.
    - you increase your workforce essentially, instead of upskilling your only employee.
    - More workers - Instances, thread
    - Specialised workers - Services
- Independence - Multiple workers are as good as a single worker if they can’t work independently.
    - They must work concurrently to the maximum extent
    - Independence is impeded by
        - shared resource
        - shared mutable data

## Replication

- When the load on a system increases, we can increase the scalability of the system vertically or horizontally.
- More instances of applications are created.
- Now user load can be shared.

## Stateful replication in web applications

- Used when low latency is required
- When a user tries to access the profile page on a social media app.
- The request comes to the load balancer, which delegates the request to a server. The application will communicate with the service layer or worst case the database. Here there are multiple hops.
- So in a stateful system, we can reduce these hops, by storing the user information in the memory of the server.
- When the first time request comes the server instance won’t have the data in its memory so it will ask the database. Before sending the response back, the server will store the data in its memory. So that any subsequent requests can be satisfied by the server and hops to the database are avoided.
- Now if the same user’s request is delegated to some other server instance, it has to again repeat the steps to get the user info on its memory.
- This can be avoided by using sticky sessions. When the data is saved in the instances memory, a session is created and a cookie which has details about the session id and instance id are sent to the client.
- A client can use this information to make subsequent requests.

![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%204.png)

- Each session data occupies memory
- Session clustering for reliability
    
    ![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%205.png)
    
    - What if Node 1 goes down? All session data stored in Node 1 are lost.
    - There are ways in which you can mitigate this issue, by clustering the session data across the nodes.
    
    > Is it really a good idea?
    > 
    
    ## Stateless replication in web applications
    
    In stateful applications, the limitation was we were being limited by the session memory of the instance. So we will now switch to a stateless approach.
    
    How do we introduce some sort of statefulness in this? We can use caching to store the session data.
    
    ![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%206.png)
    
    We moved the in-memory session storage to a separate entity. It’s obisouly not faster than in-memory storage.
    
    ## Database Replication
    
    We can use separate databases to handle read and write load. 
    
    ![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%207.png)
    
    - Master Slave (Primary-Secondary)
        - One master and a bunch of slaves
        - Master/primary can serve both read and write
        - Reads are exclusively handled by the slave/secondary
        
        ![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%208.png)
        
        - Data replication is uni-directional
        - Any changes made in the master are eventually propagated to the slave
            - This can be synchronous or asynchronous
        - High read scalability
        - High read availability
        - No write conflicts
    - Master-Master (No-Master/ Peer to Peer)
        - A client can read/write on any database.
        - Any changes made to one database are propagated to the other database.
        
        ![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%209.png)
        
        - There are chances of writing conflicts.
            - Happens when both instances try to write to the same row.
        - High read scalability
        - High read write availability
        - Write conflicts
    
    ## Database Partitioning
    
    ![Untitled](Scalability%20cb85e4be76f6441ca367f7ad78ec5361/Untitled%2010.png)