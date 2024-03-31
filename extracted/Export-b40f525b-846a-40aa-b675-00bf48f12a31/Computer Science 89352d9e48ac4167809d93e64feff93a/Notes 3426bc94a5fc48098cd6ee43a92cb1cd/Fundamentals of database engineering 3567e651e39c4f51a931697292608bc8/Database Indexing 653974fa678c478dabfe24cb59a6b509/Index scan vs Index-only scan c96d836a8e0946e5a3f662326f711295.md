# Index scan vs Index-only scan

Owner: Asish Shaji Thomas
Last edited time: August 10, 2023 7:20 AM

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled.png)

Here there is no index on the id field. So it does a sequential scan on the table.

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled%201.png)

We create an index for id on the grades table, a btree.

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled%202.png)

Now you repeat the select query, it uses an index scan. It’s because you've created a new index for the id field. You search through the btree and you get the page location where the row lives and then you go to the heap and find that page. If you were selecting only id then **index-only** scan would be used.

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled%203.png)

How do I make this more useful????

You can use non-keys.

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled%204.png)

Now the name is also stored in the index tree.

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled%205.png)

Now the scanner says index-only scan.

![Untitled](Index%20scan%20vs%20Index-only%20scan%20c96d836a8e0946e5a3f662326f711295/Untitled%206.png)

This is an index scan because g is not in the index.

But increasing the size of the index is a very very bad idea.