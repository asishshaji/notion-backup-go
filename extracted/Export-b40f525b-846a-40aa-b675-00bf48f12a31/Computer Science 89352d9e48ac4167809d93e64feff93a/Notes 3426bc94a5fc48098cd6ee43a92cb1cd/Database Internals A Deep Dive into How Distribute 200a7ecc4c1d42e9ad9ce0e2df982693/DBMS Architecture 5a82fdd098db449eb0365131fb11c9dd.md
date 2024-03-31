# DBMS Architecture

- follows client/server architecture.

![Untitled](DBMS%20Architecture%205a82fdd098db449eb0365131fb11c9dd/Untitled.png)

- Client requests arrive through the transport subsystem. The request comes in the form of queries, often expressed in some query language.
- The transport subsystem also has a cluster communication module, which uses a remote execution module in the execution engine to read and write data to and from other nodes.
- The query request is parsed, validated and interpreted by the query parser.
- The query optimizer then optimizes the parsed query but removes redundant queries and then attempts to find the optimized query plan. A query is usually represented in the form of an execution plan (query plan). A query can have multiple query plans, the query optimizer’s job is to select the most optimum query plan for the query.
- The execution plan is then forwarded to the execution engine, which collects the the results of execution of local and remote operations.
    - Remote execution involves writing and reading data to and from other nodes in the cluster, and replication.
    - Local queries are executed in the storage engine.
- Transaction Manager
    - schedules transactions and ensures they cannot leave the database in a logically inconsistent state.
- Lock Manager
    - manages locks on the database table, it ensures that concurrent operations do not violate physical data integrity.
- Access Methods
    - These manage access and organise data on disk.
    - Access methods include heap files and storage structures such as B-Trees.
- Buffer Manager
    - caches data pages in memory
- Recovery manager
    - maintains the operation log and restores the system state in case of failure
    - eg: WAL in Postgresql