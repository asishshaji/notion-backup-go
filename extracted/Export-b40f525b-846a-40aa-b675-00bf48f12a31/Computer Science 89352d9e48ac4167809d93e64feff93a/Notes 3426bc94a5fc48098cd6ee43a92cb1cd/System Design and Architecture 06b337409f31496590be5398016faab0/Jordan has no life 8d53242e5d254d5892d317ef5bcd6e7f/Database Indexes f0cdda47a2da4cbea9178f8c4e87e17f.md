# Database Indexes

## Hash indexes

Keep an in-memory hash table of key mapped to the corresponding data, occasionally write to disk. 

### Pros

Easy to implement and very fast(ram being faster than disk)

### Cons

All the keys must fit in the memory

Bad for range query because of the random distribution of the data due to the hash function.

![Untitled](Database%20Indexes%20f0cdda47a2da4cbea9178f8c4e87e17f/Untitled.png)

Why not keep a hash table in disk??? Random reads are expensive, the disk has to rotate.

> Use this when large number of writes per key, but it’s feasible to keep all the keys in memory.
When there are large writes, but not many distinct keys
> 

## SSTables and LSM Trees

Hash Indexes were unordered. To sort data by key, we can use disk-based B-Tree or memory-based structures like AVL or red-black tree. Memory based are easier to implement. These allow writing in any order and reading back in sorted order.

![Untitled](Database%20Indexes%20f0cdda47a2da4cbea9178f8c4e87e17f/Untitled%201.png)

![Untitled](Database%20Indexes%20f0cdda47a2da4cbea9178f8c4e87e17f/Untitled%202.png)

Memtable is a in-memory balanced tree. When the write comes it’s first written to the mem table. When memtable gets bigger than some threshold, it’s written out to the disk as an SSTable file. This is very efficient because the tree maintains a sorted version of the data. The new sstable file becomes the recent segment of the database.

While the new SSTable is written to the disk, new writes can be done to the new in-memory tree instance.

### How to read

![Untitled](Database%20Indexes%20f0cdda47a2da4cbea9178f8c4e87e17f/Untitled%203.png)

- First look in the in-memory tree
- if not found look in the recent sstable in disk, if not look at next-older segment etc.

We run a merging and compaction process in the background to combine the segment files, and discard overwritten or deleted values.

Even if the database crashes, you still have the sstable in disk, and the commit logs. You can use the commit log to generate the in-memory memtable, which was lost during the crash from memory.

> Storage engines that are based on the principle of merging and compacting sorted files are called LSM (Log-structured Merge-tree) storage engines.
> 

LSM trees perform very poorly when you look for data which is not the storage. You look at the memtable, then all the segments to finally find out that the data doesn’t exist. You can use a bloom filter to check if the data is not in storage. Thus saving lot of unnecessary disk io.

[Bloom Filter (A Basic understanding)](../../Fundamentals%20of%20database%20engineering%203567e651e39c4f51a931697292608bc8/Database%20Indexing%20653974fa678c478dabfe24cb59a6b509/Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886.md)

Maintaining a sparse index in memory is an effective optimization technique.

This method significantly improves the efficiency of range queries. Essentially, you group records and maintain a common index.

For example, if you are searching for 'Andy', you would refer to your sparse index. Here, you would find that 'Andy' falls between 'Alice' and 'Bob' in the disk offset. This allows you to only load the block for 'A'. And then you run a binary search O(nlogn) on the block.

Each entry in the sparse index corresponds to the start of a compressed block.

![Untitled](Database%20Indexes%20f0cdda47a2da4cbea9178f8c4e87e17f/Untitled%204.png)

### Pros

- writes are faster due to writes going in to memory buffer
- good for range queries due to internal sorting of data

### Cons

- Slow reads, especially when the data does not exists in the database
- merging and compaction is done as a background process, which can consume resource

## B-Tree

Most commonly used type of index. A disk data structure. 

Model your data such that it’s a tree on disk.

# Verdict

![Untitled](Database%20Indexes%20f0cdda47a2da4cbea9178f8c4e87e17f/Untitled%205.png)