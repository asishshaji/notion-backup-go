# Creating Partitions in Postgresql

```sql
create table grades_org(id serial not null, g int not null );
insert into grades_org(g) select floor(random() * 100) from generate_series(0,10000);
```

Create an index

```sql
create index grades_org_index on grades_org(g);
```

## How to create partitions

```sql
create table grades_parts(id serial not null, g int not null) partition by range(g);

// creating partitions
// creates table with same structure as grades_parts
create table g0035 (like grades_parts including indexes);
create table g3560 (like grades_parts including indexes);
create table g6080 (like grades_parts including indexes);
create table g80100 (like grades_parts including indexes);

// attach partition to the main table 
alter table grades_parts attach partition g0035 for values from (0) to (35);
alter table grades_parts attach partition g3560 for values from (35) to (60);
alter table grades_parts attach partition g6080 for values from (60) to (80);
alter table grades_parts attach partition g80100 for values from (80) to (100);
```

![Untitled](Creating%20Partitions%20in%20Postgresql%204dec18b2656b461fa1ab34905b9cd0eb/Untitled.png)

```sql
insert into grades_parts select * from grades_org;
```

The data from `grades_org` is written to grades_parts. The value is inserted to the appropriate partition depending upon where the values lye. 

The `grades_org` is a virtual table, it doesnot really hold the data.

```sql
create index grades_parts_idx on grades_parts(g);
```

![Untitled](Creating%20Partitions%20in%20Postgresql%204dec18b2656b461fa1ab34905b9cd0eb/Untitled%201.png)

When an index is created on `grades_parts` table, indices are created for all the partitions, 

![Untitled](Creating%20Partitions%20in%20Postgresql%204dec18b2656b461fa1ab34905b9cd0eb/Untitled%202.png)

![Untitled](Creating%20Partitions%20in%20Postgresql%204dec18b2656b461fa1ab34905b9cd0eb/Untitled%203.png)

Above query is done on the g0035 partition on its index.

The below image shows the size occupied.

![Untitled](Creating%20Partitions%20in%20Postgresql%204dec18b2656b461fa1ab34905b9cd0eb/Untitled%204.png)

gxxxx_g_idx are  way smaller than the original index, which is way faster.