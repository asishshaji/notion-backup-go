# B-Tree vs B+Tree in Production Database Systems

Why do we need B-Tree or B+Tree in the first place???

Because if don't use a tree you need to go and do a full table scan, which tends to be expensive.

## Original B-Tree

[](http://carlosproal.com/ir/papers/p121-comer.pdf)

Actual B-Tree proposal paper

[Intro to Algorithms: CHAPTER 19: B-TREES](http://staff.ustc.edu.cn/~csli/graduate/algorithms/book6/chap19.htm)

[](https://cs.kangwon.ac.kr/~leeck/file_processing/B-tree.pdf)

- Balanced Data structure for fast traversal
- B-Tree has Nodes
- In the B-Tree of “m” degrees some nodes can have (m) child nodes
- Node has up to (m-1) elements
- Each element has a key and value
- The value is usually a data pointer to the row
- Data pointers can point to the primary key or tuple
- The root node, internal node and leaf nodes
- A node = disk page

## How the Original B-Tree Helps Performance

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled.png)

Indexes less than 2 are kept on the left side. Indexes between 2 and 4 are kept in the middle. Indexes greater than 4 are kept on the right side of 4.

## Limitation of B-Tree

- Key and values are stored in nodes.

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled%201.png)

- Internal nodes take more space thus requiring more IO and can slow down traversal.
    - If you create an index on a string, the string values are stored in the nodes. They take up more space. More space means fewer elements in a node. Fewer elements mean you will need to search more parts to get to your target element. **MORE IO.. SLOOOOWWW..**
- Range queries are slow because of random access (give me all values 1-5).
    - first, find one, then jump over and find 2 and so on. This is very slow.
- B+Tree solves both of these problems.

## Why is B-Tree is bad for range queries?

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled%202.png)

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled%203.png)

To solve these challenges we have B+Tree.

## B+Tree

- Exactly like B-Tree but only stores the keys in the nodes
- Values are stored only in the leaf nodes
- Internal nodes are smaller since they only store keys and they can fit more elements (keys)
- One I/O to fetch one node, one single page gets you way more keys.
- Leaf nodes are linked, so when you find a key you can find all the values before and after that key
- Great for range queries

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled%204.png)

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled%205.png)

Here the keys are duplicated.

## B+Tree & DBMS Considerations

- Cost of leaf pointer (cheap)
- 1 Node fits a DBMS page (most DBMS)
- Can fit internal nodes easily in memory for fast traversal
    - A single IO gives you far more elements
- Leaf nodes can live in data files in the heap
- Most DBMS systems use B+Tree
- B+Tree should fit in memory.

![Untitled](B-Tree%20vs%20B+Tree%20in%20Production%20Database%20Systems%2084f3335c5cea40e49fa5dc70c97dabbc/Untitled%206.png)

The internal nodes can be loaded into the memory since those nodes are much smaller since data is not stored in the nodes.

## B+Tree Storage Cost in MYSQL vs Postgres

- B+Trees secondary index values can either point directly to the tuple(Postgres) or to the primary key (MySQL)
- If the Primary key data type is expensive this can cause bloat in all secondary indexes for databases such as MySQL(InnoDB)