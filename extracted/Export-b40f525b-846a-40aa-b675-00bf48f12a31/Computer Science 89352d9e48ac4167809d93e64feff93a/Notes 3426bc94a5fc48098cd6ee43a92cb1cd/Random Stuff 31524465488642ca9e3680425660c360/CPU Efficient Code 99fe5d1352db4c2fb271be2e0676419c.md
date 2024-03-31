# CPU Efficient Code

[Operating System — How to Write CPU-Efficient Code](https://medium.com/@tonylixu/operating-system-how-to-write-cpu-efficient-code-c656dbb18e26)

[Why software developers should care about CPU caches](https://medium.com/software-design/why-software-developers-should-care-about-cpu-caches-8da04355bb8a)

During the running of a program, the data from memory goes to shared L3 cache, then to L2→L1. Only after reaching cache, the CPU can access the data.

![Untitled](CPU%20Efficient%20Code%2099fe5d1352db4c2fb271be2e0676419c/Untitled.png)

# Structure of CPU Cache

- When CPU is reading data from main memory or from cache levels, it doesn’t read a single byte but instead reads a block of bytes, usually 64 bytes. We call this block of bytes as cache lines.

- Consists of cache line
    - basic units of data retrieval by the CPU from memory
    - one cache line is about 64 bytes

- CPU caches pull the data from the memory in chunks rather than individual data elements.
- These chunks are referred to as cache lines within the CPU cache.
- When CPU reads data, it accesses the Cache first, regardless of the data is present in cache or not.
- Only when the data is not found in the cache does the CPU access the memory, transfer data from memory→cache, and then CPU read the data from the CPU cache.

![Untitled](CPU%20Efficient%20Code%2099fe5d1352db4c2fb271be2e0676419c/Untitled%201.png)

---

[Cache Thrashing and False Sharing](https://medium.com/@ali.gelenler/cache-trashing-and-false-sharing-ce044d131fc0)

If CPUs need to work on variables shared on the same cache line then it will create cache ***thrashing*** and will cause ***false sharing***.

- There are two threads operating on a multi-core system.
- They access a shared variable, X, which is read by both thread 1 and thread 2.
- When both threads read the variable, X, it is stored in their CPU cache lines.
- If thread 1 updates the variable, X, its cache line must be invalidated.
- This invalidation also affects thread 2's cache line.
- This behavior is expected and is called ***true sharing***.
- Making the variable **`volatile`** can force the invalidation.
- Without **`volatile`**, cache line invalidation might not always occur.

---

- If an additional variable, Y, shares the same cache line as X.
- Thread 1 operates on X, while thread 2 is only interested in Y.
- Despite thread 2's interest solely in Y, both threads' cache lines contain both X and Y.
- When thread 1 updates X, it invalidates the cache line.
- As a result, thread 2 is also instructed to invalidate the cache line, even though it wasn't interested in X.
- This phenomenon is called false sharing.

# How is data accesses from cache

- How does the CPU know if the data it needs to access is in the cache?
- And where in cache it is stored?

Direct Mapped Cache (DMC) → type of cache memory that maps each block of main memory maps to exactly one cache line.

To find the cache line, where the data is stored, the cache use a direct mapping of bits from the memory address.

## How does it work??

- memory address → the tag, the cache line number (index), and the byte offset within the cache line

![Untitled](CPU%20Efficient%20Code%2099fe5d1352db4c2fb271be2e0676419c/Untitled%202.png)

- Using the cache line number you get the cache line
- The tag is compared with the tag from the requested memory address
- If there is a match, it indicates a cache hit
- If cache is hit, then the data is loaded from the cache using the byte offset.
- If the cache miss, it retrieves the data from the main memory and stored in the cache. The CPU then reads from the cache.