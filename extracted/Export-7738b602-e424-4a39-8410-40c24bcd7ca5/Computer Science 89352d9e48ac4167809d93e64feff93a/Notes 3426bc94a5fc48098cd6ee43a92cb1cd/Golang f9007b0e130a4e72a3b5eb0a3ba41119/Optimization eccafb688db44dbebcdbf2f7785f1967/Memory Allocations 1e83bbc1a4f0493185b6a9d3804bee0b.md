# Memory Allocations

Memory allocation operations will consume some CPU resources, more the allocation more the CPU consumption. In programming we should try to avoid unnecessary memory allocations to get better code execution performances.

Go runtime might find a memory block on the stack or on the heap. More the memory blocks on the heap, the more the pressure on the garbage collection.

Allocating memory on the stack also has costs, but they are mosly cheaper

The following operations will make at least one allocation

- declaring variables
- calling new function
- calling make function
- modifying slices and maps with composite literals
- converting int to strings
- concatenate strings by using +
- converting strings to byte and vice versa
- strings to rune slices
- box values into interfaces (converting non-interface values into interfaces)
- append elements to a slice and the capacity of slice is not large enough
- put new entries into map and the underlying array of the map is not large enough to store the new entries.

In most of the cases, the Go compiler makes some special code optimizations so that some of the above listed operations will not make allocations.

## Memory wasting caused by allocated memory blocks larger than needed

[Some memory block size](https://github.com/golang/go/blob/master/src/runtime/sizeclasses.go) classes are predefined. 8,16,24,32,48,64,…..

For memory blocks larger than 32,768 bytes, each of them is always composed of multiple memory pages. The memory page size used by the run time is 8192 bytes

So, 

- to allocate a heap memory block the value whose size is in the range [33, 48], the size of block is 48. 15 bytes is wasted if the value size is 33.
    
    ```go
    var t *[5]int64
    
    func Benchmark_f(b *testing.B) {
    	for i := 0; i < b.N; i++ {
    		t = &[5]int64{} // 5 * 8 = 40
    	}
    }
    
    Benchmark_f-8   	49204881	        23.77 ns/op	      48 B/op	       1 allocs/op
    ```
    
    The value is 48.
    
- For memory blocks larger than 32,768 bytes, each of them is always composed of multiple memory pages. The heap in Golang is composed of a block of memory, represented as `heapArena` in runtime, each block is 64MB in size and consists of many pages, with one page occupying 8KB of space (at 64bit).
- to create a byte slice with 32,769 elements on heap. So about 8191 bytes will be wasted here.
    - 32, 768 = 8192 * 4
    - so 32,768 bytes can be stored
    - here we need one more, so another 8182 bytes block is needed, and thus the wastage

## Avoid unnecessary allocations by allocating enough in advance

Consider a situation in which we are appending to a slice. If the append function needs to called multiple times, it’s best to ensure that the base slice has a large enough capacity, to avoid several unnecessary allocations in the whole pushing process.

```go
func MergeWithOneLoop(data ...[]int) []int {
	var r []int
	for _, s := range data {
		r = append(r, s...)
	}
	return r
}

func MergeWithTwoLoops(data ...[]int) []int {
	n := 0
	for _, s := range data {
		n += len(s)
	}

	r := make([]int, 0, n)
	for _, s := range data {
		r = append(r, s...)
	}
	return r
}

Benchmark_MergeWithOneLoop-8    	 7739151	       152.6 ns/op	     352 B/op	       4 allocs/op
Benchmark_MergeWithTwoLoops-8   	18837558	        61.98 ns/op	     176 B/op	       1 allocs/op
```

By calculating the size of array before hand we reduced the allocation to 1 allocs/op 

## Save memory and reduce allocations by combining memory blocks

Sometimes, we could allocate one large memory block to carry many value parts instead of allocating a small memory block for each value part.

```go
type Book struct {
	Title  string
	Author string
	Pages  int
}

func CreateBooksInOneBlock(n int) []*Book {
	books := make([]Book, n)
	pBooks := make([]*Book, n)
	for i := 0; i < n; i++ {
		pBooks[i] = &books[i]
	}

	return pBooks
}

func CreateBooksInMultiple(n int) []*Book {
	pBooks := make([]*Book, n)
	for i := 0; i < n; i++ {
		pBooks[i] = new(Book)
	}
	return pBooks
}

Benchmark_CreateBooksInOneBlock-8   	 1329960	       899.0 ns/op	    4992 B/op	       2 allocs/op
Benchmark_CreateBooksInMultiple-8   	  307483	      4022 ns/op	    5696 B/op	     101 allocs/op
```

From these results,

- spends much less CPU time
    - Two allocations are better than 101 allocations
- consumes a bit less memory
    - When the size of a small value doesnt exactly match any memory block classes, then a bit larger memory block than needed will be allocated for the small value, if small value is created individually.
    - The size of Book type is 40 bytes, but 48 is closer, so allocating each book wastes 8 bytes.
    
    4992 = [4096](https://github.com/golang/go/blob/master/src/runtime/sizeclasses.go) + 896 
    
    5696 = 896 + 48 * 100
    

The second part may not always make sense, if the value if N = 820, it needsto allocate a memory block of min size 32,800(820 * 40), which is larger than the largest small memory block 32,768. So the memory block needs on more memory page. 8192 * 5 = 40,960

40,960 - 32,800 = 8160 bytes are wasted.

Despite it sometimes wastes more memory, generally speaking, allocating many small value parts on one large memory block is comparatively better than allocating each of them on a separated memory block. This is especially true when the life times of the small value parts are almost the same, in which case allocating many small value parts on one large memory block could often effectively avoid memory fragmentation.

***Use value cache pool to avoid some allocations.***