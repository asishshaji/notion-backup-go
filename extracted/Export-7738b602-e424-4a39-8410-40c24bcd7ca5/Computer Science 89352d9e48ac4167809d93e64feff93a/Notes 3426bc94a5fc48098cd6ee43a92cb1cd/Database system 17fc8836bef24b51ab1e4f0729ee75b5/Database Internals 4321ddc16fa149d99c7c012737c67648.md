# Database Internals

## Data Files and Index Files

Data Files stores the data records. Index files store the record metadata and use it to locate records in the data file.

Data files can be implemented as heap-organized tables(heap files), hash-organized tables(hashed files), or index-organized tables(IOT).

**Heap-files**: Records in heap files are not required to follow any order. Mostly they follow the write order, (append order). Therefore it doesn’t incur any additional work for reorganistation. Heap files require additional index files pointing to a record's location, to make them searchable. 

**Hash-files**: Records are stored in buckets, the hash value of the key determines which bucket the records go to. Records in a bucket can be stored in write or sorted order to increase lookup speed. 

**IOT**: Stores the record in the index file itself, thus reducing the disk reads to the disk by at least one.

If the order of the data records follows the key order, the index is called a clustered index. Data records in the clustered case are usually stored in the same or in a *clustered file,* where key order is preserved.

# Buffering, Immutability and Ordering

Storage structure is based on these data structures

## Buffering

the storage structure chooses to collect a certain amount of data in memory before persisting it to disk. It’s desirable to write a full block to disk.

## Mutability(or Immutability)

This defines whether the storage engine chooses to read part of the file, update it and write the updated results at the same location in the file. Immutable structures are write-only, the updates are append at the end of the file.

## Ordering

describes whether the data is stored on the page based on the key order or not. The keys that sort closely are stored in a contiguous memory location. Ordering often defines whether or not we can efficiently scan the range of records. Storing records out of order opens up for some write-time optimization.

[B-Tree](Database%20Internals%204321ddc16fa149d99c7c012737c67648/B-Tree%2040a007e7f7a1475e9c47ee6f9e0435cc.md)