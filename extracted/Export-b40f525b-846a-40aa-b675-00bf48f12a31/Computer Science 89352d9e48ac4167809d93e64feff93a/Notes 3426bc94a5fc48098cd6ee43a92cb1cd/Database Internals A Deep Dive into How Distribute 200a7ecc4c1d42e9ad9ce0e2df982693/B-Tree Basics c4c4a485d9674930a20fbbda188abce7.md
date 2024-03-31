# B-Tree Basics

Storage engines have an underlying storage structure data structure. Storage structures have three common variables

**Buffering**

- defines whether or not the storage structure chooses to collect a certain amount of data in memory before putting it on disk. Since the smallest unit of data transfer to and from the disk is a block, it is preferable to write full blocks.

**Mutability (or immutability)**

- defines whether or not the storage structure reads parts of the file, updates them and writes the updated result at the same location in the file.
- Immutable structures are append-only.
- Modifications are appended to the end of the file.

**Ordering**

- defines whether or not the data records are stored in key order in the heap.
- Ordering often defines whether or not we can efficiently scan the range of records, not only locate the individual records.
- Storing out-of-order opens up some write-time optimizations.

Storage engines often allow multiple versions of data records to be present in the database.

One of the popular storage data structures is B-Tree.

## Binary Search Trees

![Untitled](B-Tree%20Basics%20c4c4a485d9674930a20fbbda188abce7/Untitled.png)

![Untitled](B-Tree%20Basics%20c4c4a485d9674930a20fbbda188abce7/Untitled%201.png)

- BST is a sorted in-memory data structure, used for efficient key-value lookup.
- Has a fanout of 2.
- Following the left node in the tree gets you the smallest element in the tree
- Following the right node gets you the largest element in the tree.

### Tree Balancing

- Insert operations don't follow any specific pattern.
- Element insertion might lead to an unbalanced tree. (One of the branches end up being longer than other one).
- An unbalanced tree looks like a linked list, so instead of getting a logarithmic complexity, we get a linear complexity.

![Balanced and an unbalanced tree](B-Tree%20Basics%20c4c4a485d9674930a20fbbda188abce7/Untitled%202.png)

Balanced and an unbalanced tree

- Without tree balancing, we lose the performance benefits of the binary tree.
- In the balanced tree, following the left or right node pointer reduces the search space in half on average. Look up time is O(log N).
- Instead of adding new elements to one of the tree branches and making it longer, while the other one remains empty, the tree is balanced after each operation.
- Due to low fanout (fanout is the maximum allowed number of children per node), we have to perform balancing, relocate nodes, and update pointers rather frequently.
- Increased maintenance makes BST impractical as on-disk data structures.

## Problems with having BST as a disk-based storage

- **Locality**
    - Since elements are added in random order, there’s no guarantee that a newly created node is written close to its parent
    - That means child pointers may span across several disk pages.
- **Tree Height**
    - Since binary trees have a fanout of just two, you need to perform O(log2 N) seeks to locate the searched element and subsequently perform the same number of disk transfers.

BST are effective in memory storage structures, small node size makes them impractical for external storage.

***In short on-disk BST are bad.***

[Disk-Based Structures](B-Tree%20Basics%20c4c4a485d9674930a20fbbda188abce7/Disk-Based%20Structures%20183fc72f9e064848b490d0c57ac45f47.md)