# Consistency

- Consistency in data: if the data is consistent in disk with the data model
    - **Defined by**
    - the user
    - Referential integrity (foreign keys)
    - Atomicity
    - Isolation
    
    eg : 
    
    ![Untitled](Consistency%207a2ec6b53bca48719751f912619fa7ac/Untitled.png)
    
    There is referential integrity between the two tables. The number of likes in the first table should match the number of likes in the second table. The data needs to be ***consistent***.
    
    Another example is if you see a picture id in the second table which doesn't exist in the first table.
    
    ---
    
- Consistency in reads: the reading of data can become inconsistent when you have multiple instances running.
    
    eg : 
    

![Untitled](Consistency%207a2ec6b53bca48719751f912619fa7ac/Untitled%201.png)

We update field x, and if we read the field x it should give the same value that was written before.

![Untitled](Consistency%207a2ec6b53bca48719751f912619fa7ac/Untitled%202.png)