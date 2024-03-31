# Durability

If while doing a transaction the database fails, I should see the data after the startup.

A lot of dbs write to memory and occasionally take a snapshot and write to the disk,

this makes the writes faster. 

- Changes made by committed transactions must be persisted in durable non-volatile storage.
- Durability techniques
    - WAL- Write ahead log
    - Asynchronous snapshot: as we write data is inserted into the memory, but in the background, we asynchronously take snapshots and flush them to disk.
    - Append Only File (AOF)
    
    # WAL
    
    [Write-Ahead Log](https://martinfowler.com/articles/patterns-of-distributed-systems/wal.html)
    
    - Writing a lot of data to disk is expensive (indices, data, files, columns, rows ..etc)
    - That is why DBMS persist in a compressed version of the changes known as WAL.
    
    # OS Cache
    
    - A write request in OS usually goes to the OS cache.
    - write requests are first written to the OS cache.
    - Because OS wants to batch these writes and flush to the disk at once for performance reasons.
    - OS can tell the database that it has written the data even if it’s to the cache.
    - When the writes go to the OS cache, an OS crashes, machine restart could lead to loss of data.
    - Fsync OS command forces write to always go to disk.
    - Fsync can be expensive and slow down commits.