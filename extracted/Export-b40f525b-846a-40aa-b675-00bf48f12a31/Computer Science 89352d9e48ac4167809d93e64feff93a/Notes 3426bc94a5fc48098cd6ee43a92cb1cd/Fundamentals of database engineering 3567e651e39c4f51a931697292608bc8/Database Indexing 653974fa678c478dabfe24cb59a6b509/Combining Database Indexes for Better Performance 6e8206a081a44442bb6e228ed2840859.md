# Combining Database Indexes for Better Performance

Owner: Asish Shaji Thomas
Last edited time: May 30, 2023 3:23 AM

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled.png)

---

Creates an index on a and an index on b.

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%201.png)

---

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%202.png)

You’ve too many rows, so it’s very bad to jump back and forth between index tree and heap. 

So they use a bitmap index scan to get a bitmap containing pages where the condition is met. 

If you limit the number of results that you want, the bitmap index scan will be changed to an index scan.

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%203.png)

---

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%204.png)

The bitmap index scan on test_b_idx is done. 

---

What if I query on both a and b at the same time????

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%205.png)

Postgresql does a bitmapAnd on bitmaps of both test_idx_a and test_idx_b. Each bitmap will have page numbers which satisfy conditions on each index. And then doing an AND will give you a bitmap which satisfies both conditions. 

An OR does a bitmapOR.

---

Create a composite on a and b. Drop individual indexes.

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%206.png)

This is very efficient when you are doing a query on both a and b.

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%207.png)

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%208.png)

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%209.png)

This a and b index will only be used if the query has either only a or both a and b.

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%2010.png)

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%2011.png)

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%2012.png)

If I create a separate index on b and then do the slow queries

You’ll have a faster query.

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%2011.png)

![Untitled](Combining%20Database%20Indexes%20for%20Better%20Performance%206e8206a081a44442bb6e228ed2840859/Untitled%2013.png)

a or b. Here for a the composite index is used and b already have an index created for it.