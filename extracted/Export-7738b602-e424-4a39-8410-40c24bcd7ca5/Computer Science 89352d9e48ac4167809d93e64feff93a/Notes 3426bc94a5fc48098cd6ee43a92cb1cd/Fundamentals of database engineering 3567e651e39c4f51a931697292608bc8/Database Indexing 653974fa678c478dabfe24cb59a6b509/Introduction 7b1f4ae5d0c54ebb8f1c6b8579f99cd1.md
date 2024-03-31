# Introduction

Owner: Asish Shaji Thomas
Verification: Verified
Last edited time: August 10, 2023 4:23 AM

The index is a data structure that you build and assign on top of your data that looks through your data and creates kind of shortcuts.

---

You have an employee table with id and name. id is the primary key and it’s the index for this table.

## Code

```sql
create table employees(id serial not null, name varchar(30) not null );
create index on employees(id);
insert into employees(name) select substr(md5(random()::text), 0, 15) 
				from generate_series(0,1000000);
select * from employees LIMIT 10;
```

---

![Untitled](Introduction%207b1f4ae5d0c54ebb8f1c6b8579f99cd1/Untitled.png)

you are trying to find id 2000. Here the `explain` shows it’s an index-only scan. So it didn't read the data heap sequentially. It searched the index btree to find the location of the page.

Heap Fetches: 0, indicates that the query didn't need to go to the data heap to get the data. It’s because, you are selecting only the index, which is already present in the index heap. If the data is in the index heap, that’s your most efficient query. 

---

![Untitled](Introduction%207b1f4ae5d0c54ebb8f1c6b8579f99cd1/Untitled%201.png)

Here you are also selecting the name, which is not present in the index heap. So reading the data heap is required. Therefore it’s a slower query compared to prior one. 

If we run the same query again the execution time will be much smaller because of the cache. Cache at different levels, OS, DB, etc. Everyone caches

---

## explain analyze select id from employees where name = ‘Zs’;

![Untitled](Introduction%207b1f4ae5d0c54ebb8f1c6b8579f99cd1/Untitled%202.png)

It takes a lot of time, almost 3 seconds. Why?? Because you don't have an index for the name field. So the query scans the pages one by one (theoretically, most databases use concurrency to make the scan faster) till they find the actual page.