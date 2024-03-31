# Data Management: Choosing the right database

## Polyglot Persistence Principle in Microservices

[Building Microservices with Polyglot Persistence Using Spring Cloud and Docker](https://www.kennybastani.com/2015/08/polyglot-persistence-spring-cloud-docker.html)

- Every microservice has its data, and the data integrity and data consistency be considered very carefully
- Microservices should not share databases

[Microservices Pattern: Database per service](https://microservices.io/patterns/data/database-per-service.html)

![Untitled](Data%20Management%20Choosing%20the%20right%20database%20398ab043f0514e32a924150a3e2a808f/Untitled.png)

- Microservice share data over the application microservices using REST  API.
- We can pick different databases for different services as per the usecase.
- When you are sharing the data between multiple microservices you loose your acid properties.
- Challenging to implement queries and transactions that visits serveral microservices.

---

## Database-per-Service pattern

- Microservices should be loosely coupled, scalable and independent.
- Use distributed data model with many smaller databases for particular microservice.

![Untitled](Data%20Management%20Choosing%20the%20right%20database%20398ab043f0514e32a924150a3e2a808f/Untitled%201.png)

- Changes to an individual database schema won’t affect the other services running.
- Individual databases are easier to scale.
- We can use polyglot persistence. (use different technologies per service)
- Services will need to communicate with other services to access their data, this will need inter-service communication, which can cause a bottleneck.
- Distributed transactions across microservices can negatively impact consistency and atomicity.
- No simple way to execute join queries on multiple data stores.

## The Shared Database: An Anti-Pattern

![Untitled](Data%20Management%20Choosing%20the%20right%20database%20398ab043f0514e32a924150a3e2a808f/Untitled%202.png)

- Follows ACID property
- hard to scale
- When we use a shared database, the microservice lose their core properties: scalability, resilience and independence.
- It can block microservices due to single-point-of-failure.

---

![Untitled](Data%20Management%20Choosing%20the%20right%20database%20398ab043f0514e32a924150a3e2a808f/Untitled%203.png)

## When to use relational databases?

- ACID compliance, Data consistency
- If the application data is predictable, using the structure of data, the size and the frequency of access
- When data is highly structured.
- Relational databases are scaled vertically.

## When NoSQL databases

- flexible schema and dynamic data
- Use case of an IOT platform that stores data from different sensors with frequently changing attributes.
- When you have a high-volume workload and need to scale horizontally with low latency.
- NoSQL databases are designed for the cloud
- NoSQL can handle data coming in high volume.
- NoSQL database is eventually consistent and doesn’t support transactions
- NoSQL focuses on high-volume data and horizontal scaling.
- Writes are much faster compared to SQL because they are eventually consistent. Eventually, consistency is leveraged to attain fast writes.
- Not good for complex join queries

## Best practices when choosing databases

- Don’t use relational databases everywhere.
- Don’t use single database technology
- Use polyglot persistence.
- Focus on the data type that you need to store
    - Store JSON documents in NoSQL DB
    - transactional data into relational DB
    - time series data into telemetry databases
    - blob data into blob storage
    - application logs into elastic search dbs

## How to choose the best database for your micro-service

- Do we need strict-string consistency or eventual consistency?
    - strong consistency: in banks, you can use SQL DBS here.
- Do we need ACID compliance? It’s preferable to follow eventual consistency in micro-services to gain high scalability and availability.
- Fixed or Flexible Schema ???
- High or low data volume, Predictable or unpredictable data?
    - High volume, unpredictable: NoSQL, because it’s designed for the cloud
- Relational or non-relational data?? Are there complex join queries?
- Deployments, Centralized or de-centralized structures?
- Do we need to achieve fast read-write performance?
- Do we need high scalability??
    - You will need to give up strong consistency here.

## CAP theorem

![Untitled](Data%20Management%20Choosing%20the%20right%20database%20398ab043f0514e32a924150a3e2a808f/Untitled%204.png)

In a distributed system, consistency, availability and partitioning cannot  be achieved at the same time.

Don’t confuse database partitioning with partition in CAP theorem.

Partitioning: In distributed system, when you have multiple services which cannot talk to each other, then there is a partition between them.

So partition tolerance means it can support partition tolerence.

No partition tolerence means there is only one machine.

Distributed systems should sacrifice between consistency, availability and partition tolerance.

Any database can only guarantee two of these. 

In distribued architecture, Partition Tolerence is a must-have feature, it’s necessary to choose between consistency or availability.

If a system needs to be fully consistent, it must sacrifise high availability.

Microservices choose partition tolerance with high availability and eventual consistency.

[My thoughts on the CAP theorem](https://youtu.be/KmGy3sU6Xw8?t=418)

[An Illustrated Proof of the CAP Theorem](https://mwhittaker.github.io/blog/an_illustrated_proof_of_the_cap_theorem/)

The system uses networks, so it’s partitioned application. So the architecture is partition tolerant. 

![cap_theorem.drawio.png](Data%20Management%20Choosing%20the%20right%20database%20398ab043f0514e32a924150a3e2a808f/cap_theorem.drawio.png)

When a write comes to the master node, we get two choices.

1. Immediately acidify and completely commit to the master and return to the client with an acknowledgement(Asynchronous). Hence my system as a whole is available.
    
    Is it consistent though??? The new updates haven’t been propogated to the read replicas yet. So any read requests to the read replicas will return stale data. So it’s not consistent
    
    This **AP** part of CAP.
    
2. You block the client till, the data is written in the replica set. You wait till all the writes are written in all the replica set. (Synchronous).
Is it available though??? Since the system supports partition tolerence, the request to one the read replica can fail. So the master nodes keep trying, the read replica is already down. Since the write is not commited to the replica, the client is still waiting for the acknowledgement. 
Therefore the system is not available. 

If there is partition tolerence, you get consistency and availability.

[Database Partitioning](../Fundamentals%20of%20database%20engineering%203567e651e39c4f51a931697292608bc8/Database%20Partitioning%207e027e521e2f45c5995425410bbf65b9.md)

[Database Sharding](../Fundamentals%20of%20database%20engineering%203567e651e39c4f51a931697292608bc8/Database%20Sharding%20edf58b69ee4e429a81e1c5f7d7098ede.md)

---