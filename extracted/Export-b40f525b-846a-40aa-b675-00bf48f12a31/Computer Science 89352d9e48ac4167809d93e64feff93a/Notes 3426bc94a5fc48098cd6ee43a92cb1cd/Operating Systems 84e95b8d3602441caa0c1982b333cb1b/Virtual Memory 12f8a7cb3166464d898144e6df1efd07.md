# Virtual Memory

The main memory of a computer system is organised as an array of M contiguous byte-size cells. Each byte has a unique physical address.

The most natural way for a CPU to access memory would be to use physical addresses. Early PCs used physical addressing, and systems such as digital signal processors, embedded microcontrollers, and Cray supercomputers continue to do so.

Modern processors use a form of addressing known as virtual addressing.

![Untitled](Virtual%20Memory%2012f8a7cb3166464d898144e6df1efd07/Untitled.png)

With virtual addressing, the CPU accesses main memory by generating a virtual address, which is converted to the appropriate physical address before sent to the main memory. This conversion is called *address translation*. Dedicated hardware on the CPU chip called the memory managment unit translates virtual addresses on the fly using a map called page map.

![Untitled](Virtual%20Memory%2012f8a7cb3166464d898144e6df1efd07/Untitled%201.png)

Each byte of main memory has a virtual address chosen from the virtual address space, and a physical address chosen from the physical address space.

## Virtual Memory as a Tool for Caching

Both virtual and physical address spaces are divided into fixed sized blocks called pages.

Typical page size is 4kb - 16kb

Virtual address = virtual page number + offset

Physical address = physical page number + offset 

![Untitled](Virtual%20Memory%2012f8a7cb3166464d898144e6df1efd07/Untitled%202.png)

The key idea is MMU manages pages not individual memory addresses(Why waste resource by translating every addresses). Think about of the principle of locality.

The contents of disk are cached in main memory. The data on the disk(lower level) is partitioned into blocks that serve as the transfer units between the disk and the main memory (upper level) (cache hierrarchy). VMs system handle this by partitioning the virtual memory into fixed size blocks called virtual pages. Each virtual page is P = 2^p bytes. Physical memory is also partitioned into physical pages, also P bytes in size. (page frames)

At any point in time, the set of virtual pages is partitioned into three disjoint subsets.

- Unallocated: pages that have not yet been allocated by the VM system. They dont have any data associated with them, and thus donot occupy any space on the disl
- Cached: allocated pages that are currently cached in physical memory
- Uncached: allocated pages that are not cached in physical memory.

![Untitled](Virtual%20Memory%2012f8a7cb3166464d898144e6df1efd07/Untitled%203.png)

SRAM cache misses are not very expensive, because the data will be served from the DRAM.

DRAM cache misses are expensive, because the data will be served from the disk.

Best video I watched about virtual memory.

[Virtual Memory](https://www.youtube.com/watch?v=pRjokFrxt_I&ab_channel=John'sBasement)