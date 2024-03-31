# How do database optimizers decide to use indexes?

Owner: Asish Shaji Thomas
Last edited time: May 30, 2023 4:48 AM

![Untitled](How%20do%20database%20optimizers%20decide%20to%20use%20indexes%203435e6ea10964751879401f03331b503/Untitled.png)

## Case 1

- Both indexes are used to query
- eg: F1_IDX is used to search 1 and F2_IDX is used to search 4
- F1 = 1 gets all the rows where it satisfies
- F2 = 4 gets all the rows where it satisfies
- Then an intersection is done,

![Untitled](How%20do%20database%20optimizers%20decide%20to%20use%20indexes%203435e6ea10964751879401f03331b503/Untitled%201.png)

the resultset rows are 7 and 10.

Do we need to use both indexes?????

- If the data that we know we will get is very little, it doesn't make sense to use both indexes. We can use a single index and get the data from the heap, then use somekind of filter or recheck.
- If the data that we know we will get is very large, it doesn't make sense to use any index at all. We can just go directly to the heap and get the data from there.

## Case 2

- Optimizer decides to favor one index over another
- Only one index is used
- eg: F1_IDX to search for 1, rowIds are collected and the table is accessed and then filtered for F2=4 is filtered

![Untitled](How%20do%20database%20optimizers%20decide%20to%20use%20indexes%203435e6ea10964751879401f03331b503/Untitled%202.png)

## Case 3

- No indexes are used. Does a full table scan
- eg: Database decides the search will yield so many rows that its going to be faster to do a full table scan.

![Untitled](How%20do%20database%20optimizers%20decide%20to%20use%20indexes%203435e6ea10964751879401f03331b503/Untitled%203.png)