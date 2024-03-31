# Understanding Database Internals

## How tables and indices are stored on disk

![Untitled](Understanding%20Database%20Internals%206ae0efa163f943508474716cdf51454f/Untitled.png)

## ROW_ID

- Internal and system maintained
- In certain databases, it is the same as the primary key but other databases like  Postgres have a system column row_id (tuple_id)

![Untitled](Understanding%20Database%20Internals%206ae0efa163f943508474716cdf51454f/Untitled%201.png)

## Page

[Database Pages — A deep dive](https://medium.com/@hnasr/database-pages-a-deep-dive-38cdb2c79eb5)

- Depending on the storage model (row vs column store), the rows are stored and read in logical pages
- The database doesn't read a single row, it reads a page or more in a single IO and we get a lot of rows in that IO.
- Each page has a size (8KB in Postgres, 16KB in MySQL)
- Assume each page holds 3 rows in this example. with 1001 rows, you will have 1001/3 = 333 ~ pages

 

![Untitled](Understanding%20Database%20Internals%206ae0efa163f943508474716cdf51454f/Untitled%202.png)

![Untitled](Understanding%20Database%20Internals%206ae0efa163f943508474716cdf51454f/Untitled%203.png)

## IO

- IO operation (input/output) is a read request to the disk
- We try to minimize this as much as possible (IO is like the currency of the database, use it wisely)
- An IO can fetch 1 page or more depending on the disk partitions and other factors
- An Io cannot read a single row, it's a page with many rows in them, you get them for free.
- You want to minimize the number of IOs as they are expensive.
- Some IOs in OS goes to the operating system cache and not the disk

## Heap

- A heap is a data structure where the table is stored with all its pages one after another.
- This is where the actual data is stored including everything.
- Traversing the heap is expensive as we need to read so much data to find what we want.
- That is why we need indexes that help to tell us exactly what part of the heap we need to read. What page(s) of the heap do we need to pull?

![Untitled](Understanding%20Database%20Internals%206ae0efa163f943508474716cdf51454f/Untitled%204.png)

## Index

- An Index is another data structure separate from the heap that has “pointers” to the heap
- It has part of the data and is used to quickly search for something
- You can index one column or more
- Once you find a value of the index, you go to the heap to fetch more information where everything is there.
- Index tells you EXACTLY which page to fetch in the heap instead of taking the hit to scan every page in the heap.
- The index is also stored as a page and costs IO to pull the index entries.
- The smaller the index, the more it can fit in memory the faster the search
- The popular data structure for the index is b-tress

![Untitled](Understanding%20Database%20Internals%206ae0efa163f943508474716cdf51454f/Untitled%205.png)

### Example: Index

**No index**

Select * from EMP where EMP_ID = 10000;

Queries pages from 0 till you get the match. Very expensive. You can use multithreading to scan it. But in general, this idea itself is expensive

**With index**

Select * from EMP where EMP_ID = 10000;

Index details are searched in the b-tree on index heaps. This finds the page number and row of the data page heap.

# Primary Key vs Secondary Key

[What is the difference between Clustered and Non-Clustered Indexes in SQL Server?](https://www.sqlshack.com/what-is-the-difference-between-clustered-and-non-clustered-indexes-in-sql-server/)

[Clustered Vs Non Clustered Index](https://medium.com/fintechexplained/clustered-vs-non-clustered-index-8efed55ed7b9)

Data exists in the heap. Normal querying with a key is very expensive.

Clustering: the idea of organising a table(heap) around an order. 

A c**lustered index** is an index which defines the physical order in which table records are stored in a database. Since there can be only one way in which records are physically stored in a database table, there can be only one clustered index per table. By default, a clustered index is created on a primary key column.

With clustered indexes, the database manager attempts to keep the data in the data pages in the same order as the corresponding keys in the index pages. Thus the database manager attempts to insert rows with similar keys onto the same pages. If the table is reorganized, data is inserted into the data pages in the order of the index keys. The database manager does not maintain any order of the data when compared to the order of the corresponding keys of a non-clustered index.

**Secondary Key**: you maintain a separate ordered structure somewhere else. Your data heap can be random in order. We want to query, you go to the maintained heap and get the pages tuple and then get the actual rows from the data heap. Nonclustered indexes have a structure separate from the data rows. A nonclustered index contains the nonclustered index key values and each key value entry has a pointer to the data row that contains the key value.