# Build Your Own Database From Scratch | Build
Your Own Database From Scratch

URL: https://build-your-own.org/database/
Category: Database, Design, General, System Design

![https://build-your-own.org/database/img/book_byodb_banner.jpg](https://build-your-own.org/database/img/book_byodb_banner.jpg)

---

## Introduction

Databases are not black boxes. Understand them by building your own from scratch!

This book contains a walk-through of a minimal persistent database implementation. The implementation is incremental. We start with a B-Tree, then add a new concept with each chapter, and eventually go from a simple KV to a mini relational DB.

Although the book is short and the implementation is minimal, it covers three important topics:

1. **Persistence.** How not to lose or corrupt your data. Recovering from a crash.
2. **Indexing.** Efficiently querying and manipulating your data. (B-tree).
3. **Concurrency.** How to handle multiple (large number of) clients. And transactions.

Sample code written in Golang.

## Contents

### Part I: Simple KV Store

*(Self-contained, free-to-read web version.)*

1. [Introduction](https://build-your-own.org/database/00a_overview.html)
2. [Files vs Databases](https://build-your-own.org/database/01_files.html)
3. [Indexing](https://build-your-own.org/database/02_indexing.html)
4. [B-Tree: The Ideas](https://build-your-own.org/database/03_btree_intro.html)
5. [B-Tree: The Practice (Part I)](https://build-your-own.org/database/04_btree_code_1.html)
6. [B-Tree: The Practice (Part II)](https://build-your-own.org/database/05_btree_code_2.html)
7. [Persist to Disk](https://build-your-own.org/database/06_btree_disk.html)
8. [Free List: Reusing Pages](https://build-your-own.org/database/07_free_list.html)

### Part II: Mini Relational DB

*(Included in the [ebook and paperback](https://build-your-own.org/database/90_end.html) editions.)*

1. Rows and Columns
2. Range Query
3. Secondary Index
4. Atomic Transactions
5. Concurrent Readers and Writers
6. Query Language: Parser
7. Query Language: Execution

**The book is available for [purchase](https://build-your-own.org/database/90_end.html).**

![Build%20Your%20Own%20Database%20From%20Scratch%20Build%20Your%20Ow%20c00058bb03e64937999577bf57cf5908/book_byodb_cover_1024.jpg](Build%20Your%20Own%20Database%20From%20Scratch%20Build%20Your%20Ow%20c00058bb03e64937999577bf57cf5908/book_byodb_cover_1024.jpg)