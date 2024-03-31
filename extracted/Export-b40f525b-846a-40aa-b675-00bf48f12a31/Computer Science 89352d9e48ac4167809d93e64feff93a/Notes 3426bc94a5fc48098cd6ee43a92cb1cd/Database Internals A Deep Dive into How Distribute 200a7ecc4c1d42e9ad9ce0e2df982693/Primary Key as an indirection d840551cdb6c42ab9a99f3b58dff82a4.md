# Primary Key as an indirection

Indexes can be primary and secondary. You can specify which keep you want to be primary or most dbms create a primary key. 

If the data in the heap is organised in sequence as per the key, then it’s called a clustered key. 

Data records in the clustered case are usually stored in the same file or in a clustered file, where the key order is preserved.

If the data is stored in a separate file, and its order does not follow the key order, the index is called nonclustered

There are different opinions on whether the data records should be referenced by data offsets or by directly using the primary key.

By referencing data directly by using data offset, we can reduce the number of disk seeks, but have to pay the cost of updating the pointers whenever the record is updated or relocated during a maintenance process. 

Using indirection in the form of a primary index allows us to reduce the cost of updating pointers, but has a higher cost on the read path. (You read the secondary index to find the primary key, and then use the primary to locate the data tuple  in the heap)

> For example, MySQL InnoDB uses a primary index and performs two lookups: one in the secondary index, and one in a
primary index when performing a query. This adds an overhead of a primary index lookup instead of following the offset directly from the secondary index.
> 

If the application is primarily read-heavy, you can stick to the first approach, data offset. When the application is write-heavy, updating the pointer becomes a bottleneck.

![Referencing data tuples directly versus using a primary index as indirection](Primary%20Key%20as%20an%20indirection%20d840551cdb6c42ab9a99f3b58dff82a4/Untitled.png)

Referencing data tuples directly versus using a primary index as indirection

It is also possible to use a hybrid approach and store both data file offsets and primary keys. First, you check if the data offset is still valid and pay the extra cost of going through the primary key index if it has changed, updating the index file after finding a new offset.