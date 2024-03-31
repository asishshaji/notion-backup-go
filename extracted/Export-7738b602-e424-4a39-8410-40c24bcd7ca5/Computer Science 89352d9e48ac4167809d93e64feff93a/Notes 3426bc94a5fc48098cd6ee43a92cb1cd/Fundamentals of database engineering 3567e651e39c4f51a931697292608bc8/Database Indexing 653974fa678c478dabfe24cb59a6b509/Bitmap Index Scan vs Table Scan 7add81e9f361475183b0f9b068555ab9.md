# Bitmap Index Scan vs Table Scan

Owner: Asish Shaji Thomas
Last edited time: May 27, 2023 2:22 PM

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled.png)

---

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled%201.png)

---

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled%202.png)

finds the indexes less than 100 from the index tree and goes to the heap to see the pages. This can get slower if the rows are enormous.

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled%203.png)

PostgreSQL knows that rows are going to be huge. So it doesn't go to the index, instead, it does a seq scan on the heap. In an index scan, you go to the index and then to the heap. If the rows are larger why go to two places?

> If Postgres sees that you have a few rows, it's going to use an index scan. If it’s larger then use a sequential scan.
> 

---

What if you don't have too many rows, but just enough that it's efficient to use a sequential scan?

That’s where ***bitmap scan*** comes into the picture.

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled%204.png)

[Bitmap indexes in Go: unbelievable search speed](https://medium.com/bumble-tech/bitmap-indexes-in-go-unbelievable-search-speed-bb4a6b00851)

Basically, the query goes to the index and maintains a bitmap, the bitmap(to the length of all the pages) is updated when a true row is found. 

Then bitmap can be used to fetch all pages from the heap. You only jump once to heap.

***You know exactly what pages to fetch, so you only jump once to the heap.***

When read these pages, you also get rows that don't satisfy g > 95. That’s why you have a RECHECK cond in the above query. 

---

The beauty of a bitmap index scan is that you can scan multiple indexes and maintain multiple bitmaps. You can **and** them together and will get a single bitmap. Then jump to the heap.

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled%205.png)

Two sets of index, bitmap scan. Then does a BitmapAnd to get the new bitmap. 

![Untitled](Bitmap%20Index%20Scan%20vs%20Table%20Scan%207add81e9f361475183b0f9b068555ab9/Untitled%206.png)