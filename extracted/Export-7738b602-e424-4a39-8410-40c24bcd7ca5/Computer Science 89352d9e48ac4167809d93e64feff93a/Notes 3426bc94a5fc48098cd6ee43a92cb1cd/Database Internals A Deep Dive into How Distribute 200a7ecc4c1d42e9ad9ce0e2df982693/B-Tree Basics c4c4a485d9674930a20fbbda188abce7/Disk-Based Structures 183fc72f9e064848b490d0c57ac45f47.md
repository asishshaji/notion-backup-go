# Disk-Based Structures

Some data structures are better suited to be used on disk and some work better in memory. Not every data structure that satisfies space and time complexity requirements can be effectively used for on-disk storage. 

Data structures used in databases have to be adapted to account for persistent medium limitations.

## Hard Disk Drives

On spinning disks, seeks increases costs of random reads because they require disk rotation and mechanical head movements to position the read/write head to the desired location. Seeking is the expensive part, reading or writing contiguous bytes is relatively cheaper.

The smallest transfer unit of a spinning drive is a sector, so when some operation is performed, at least an entire sector can be read or written. Sector sizes: 512 bytes to 4Kb.

Head positing is the most expensive part of an operation on an HDD. That’s why *sequential I/O*(reading and writing contiguous memory segments from disk) **is best for HDD.

## Solid State Drives

These don't have moving parts, so there’s no disk that spins or head that has to be positioned for the read.

SSD is built of memory cells. connected into strings, strings combined into arrays, arrays combined into pages, and pages combined into blocks.

cell: one or multiple bits of data

pages: 2 - 16Kb

blocks: 64 - 512 pages

Blocks are organized into planes and, finally, planes are placed on a die. SSDs can have one or more dies. 

![Untitled](Disk-Based%20Structures%20183fc72f9e064848b490d0c57ac45f47/Untitled.png)

The smallest unit that can be written(programmed) or read is a page.