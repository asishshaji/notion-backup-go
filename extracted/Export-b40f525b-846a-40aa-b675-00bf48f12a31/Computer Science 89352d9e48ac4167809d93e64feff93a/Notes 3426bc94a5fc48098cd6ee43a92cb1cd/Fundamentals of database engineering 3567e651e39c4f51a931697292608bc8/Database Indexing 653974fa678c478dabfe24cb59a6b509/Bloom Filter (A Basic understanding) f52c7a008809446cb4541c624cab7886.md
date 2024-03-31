# Bloom Filter (A Basic understanding)

Owner: Asish Shaji Thomas
Last edited time: February 14, 2024 2:04 AM

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled.png)

Very slowww…

Use cache????

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled%201.png)

Inefficient, memory

## Bloom filter

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled%202.png)

Here 3 is not set. So it’s sure that there is no user with name “paul”. So the request need not touch the database.

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled%203.png)

Here 63 is set, which means there is a chance that Jack is in the system. If the bit is set the request can hit the database and check if Jack exists in the database.

By using bloom filter you basically limit the number of database accesses. 

## Create user bloom filter

First of all you have a blank bloom filter.

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled%204.png)

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled%205.png)

![Untitled](Bloom%20Filter%20(A%20Basic%20understanding)%20f52c7a008809446cb4541c624cab7886/Untitled%206.png)

Here 63 is already set, so there is no need to set 63 again.

[Demystifying Bloom Filters with Real-Life Examples](https://medium.com/@abhishekranjandev/demystifying-bloom-filters-with-real-life-examples-b66db7e37b37)