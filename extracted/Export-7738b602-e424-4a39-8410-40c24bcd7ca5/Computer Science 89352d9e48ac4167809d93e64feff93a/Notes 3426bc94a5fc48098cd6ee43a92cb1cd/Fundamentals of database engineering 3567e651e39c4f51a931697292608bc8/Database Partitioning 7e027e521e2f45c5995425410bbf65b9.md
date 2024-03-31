# Database Partitioning

- the table is split into multiple tables, and the DBMS decides based on the where clause which partition to hit.

![Untitled](Database%20Partitioning%207e027e521e2f45c5995425410bbf65b9/Untitled.png)

Here if the database has an index on the id, DBMS will first find the index and then go to the data on disk. 

If there is no index we will do a sequential scan on the heap data file. The worst thing that you can do.

Regardless of using an index scan or sequential scan, scanning a table with 1M entries is bad.

Break the table into smaller parts so that you know you are working with only small parts of the data.

![Untitled](Database%20Partitioning%207e027e521e2f45c5995425410bbf65b9/Untitled%201.png)

These partitions are attached to CUSTOMERS_TABLE which acts as a parent for these partitions. The parent table does not have any data associated with it. Actual data exists on the partitions.

700,001: which “partition” is customer 700,001?? it’s in 800K. Now you have only 200,000 records. That’s a smaller set of data. Now you can do a sequential or index scan.

## Vertical and Horizontal Partitioning

- Horizontal partitioning splits rows into partitions
    - as we discussed earlier
- Vertical partitioning splits columns partitions
    - Large column (blob) that you can store in a slow access drive in its own tablespace.
    - eg: If your table in which a column takes a lot of space, the column is used rarely, then you use vertical partitioning to store the column in a separate partition.

## Partitioning Types

- By Range
    - Dates, ids (eg: by log date or customer id from-to)
- By List
    - Discrete values (eg States Kerala, TN..etc) or zip codes
- By Hash
    - Hash functions

## Difference between Partitioning and Sharding

- Partitioning splits a big table into multiple tables in the same database and the database takes care of the management of the partition
- You as a client of agnostic of which partition it will hit.
- Sharding splits a big table into multiple tables across multiple database servers.
    - eg Separate shard for south India and north India
    - If the request is from South India, the database server in South India is hit.
- HP table name changes
- Sharding everything is the same but the server changes.

[Creating Partitions in Postgresql](Database%20Partitioning%207e027e521e2f45c5995425410bbf65b9/Creating%20Partitions%20in%20Postgresql%204dec18b2656b461fa1ab34905b9cd0eb.md)

## Advantages

- Increases query performance when accessing a single partition
    - Indices are smaller, thus much faster.