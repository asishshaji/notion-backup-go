# Understanding the SQL query planner and Optimizer with Explain

Owner: Asish Shaji Thomas
Last edited time: August 10, 2023 4:55 AM

Take into consideration a table ***grades***

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled.png)

id, grade and name

Both grade and id have indexes against them.

### Code

```sql
create index g_idx on grades(grade);
create index id_idx on grades(id);
```

---

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%201.png)

This query is selecting everything (*) and there is no filter, so the query goes inside the heap and fetches all the pages.

[Performance Tips](https://www.postgresql.org/docs/8.1/performance-tips.html#USING-EXPLAIN)

The **cost** usually has two numbers separated by .. . First number represents the time in milliseconds to fetch the first page. If the cost is larger, that means you are doing a lot of work before fetching the data.

The second number represents(estimated) the total amount of time in milliseconds to finish the query.

**Rows** represent the approximate (estimated) number of rows that the query will have. 

**Width** is the average width (in bytes) of rows output by this query.

---

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%202.png)

Here the **cost** went up a little bit. Postgres did some work before fetching the rows. That work is ordering. An index scan is used here.

---

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%203.png)

Here the name has no index.

> Always read the query plan from the bottom.
> 

So parallel seq scan first. Reads the entire table. 

Now we sort on the name. 

Since we have two workers we need to merge the results at the end, that is the gather merge.

---

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%204.png)

width is 4, id is an integer which is 4 bytes.

---

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%205.png)

The width is 19. 

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%206.png)

First, the index is scanned to get the location of the page containing the grade with id 10 and then reads the page. 

![Untitled](Understanding%20the%20SQL%20query%20planner%20and%20Optimizer%20%2036da221766b4401688dbc08642784a90/Untitled%207.png)

Here the scan is index only, so we don't even need to jump to the heap to fetch the data.