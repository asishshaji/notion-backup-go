# Create Index Concurrently - Avoid Blocking Production Database Writes

Owner: Asish Shaji Thomas
Last edited time: August 30, 2023 8:07 AM

When an index is being created, your table is locked.

You can read but you cannot write.

PostgreSQL has a feature to create indexes concurrently. 

```sql
create index concurrently g on grades(g);
```

Now you can read and write. The only downside is that the index creation can take time. The operation can also fail.🤏🏽