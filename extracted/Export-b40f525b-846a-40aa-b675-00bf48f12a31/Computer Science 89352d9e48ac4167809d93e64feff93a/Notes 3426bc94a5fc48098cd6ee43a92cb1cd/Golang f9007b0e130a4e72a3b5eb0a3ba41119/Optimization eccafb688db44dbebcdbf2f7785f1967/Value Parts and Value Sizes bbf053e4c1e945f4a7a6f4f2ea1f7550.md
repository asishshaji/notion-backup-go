# Value Parts and Value Sizes

In Go, value of some kinds of types always contains only one part, whereas a value of other kinds of types might contain more than one part. 

Direct part and indirect part. The direct part references the indirect part.

![one part](Value%20Parts%20and%20Value%20Sizes%20bbf053e4c1e945f4a7a6f4f2ea1f7550/Untitled.png)

one part

![multiple part](Value%20Parts%20and%20Value%20Sizes%20bbf053e4c1e945f4a7a6f4f2ea1f7550/Untitled%201.png)

multiple part

When assigning/copying a value, only the direct part of the value is copied. The indirect parts are usually shared.

So when it comes to defining a size of a type, only the direct parts are considered.

Internally, a slice uses a pointer (on the direct part) to reference all its elements (on the indirect part). The len and cap of a slice are stored on the direct part of the slice.

A string references all its containing bytes (on the indirect part), the length of a string is stored on the direct part of the string as an int value.

## Copying costs

To achieve high code execution performance, we should avoid

- copying a large quantityof large-size values in a loop
- copying very large size arrays and structs

In practice, we can view struct types with no more than 4 native word size fields and array types with no more than 4 native word size elements as small-size values.

### Value copy scenarios

**EXAMPLE 1 :** Lets benchmark ranging over an array and taking the sum. This can be achieved in multiple ways, by passing the whole array, by passing a pointer, by passing a slice.

```jsx
//go:noinline
func Sum_RangeArray(a [N]int) (r int) {
	for _, v := range a {
		r += v
	}
	return
}

//go:noinline
func Sum_RangeArrayPtr1(a *[N]int) (r int) {
	for _, v := range *a {
		r += v
	}
	return
}

//go:noinline
func Sum_RangeArrayPtr2(a *[N]int) (r int) {
	for _, v := range a {
		r += v
	}
	return
}

//go:inline
func Sum_RangeSlice(a []int) (r int) {
	for _, v := range a {
		r += v
	}
	return
}

cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
Benchmark_Sum_RangeArray-8       	   13458	     88853 ns/op
Benchmark_Sum_RangeArrayPtr1-8   	   16028	     74745 ns/op
Benchmark_Sum_RangeArrayPtr2-8   	   15824	     72071 ns/op
Benchmark_Sum_RangeSlice-8       	   23794	     50381 ns/op
```

The benchmark results shows that slice is faster. Let’s understand why.

In RangeArray the whole array is copied twice. Once during the function call and once during the range operation. So when the size of the array increases so does the cpu cycles. 

For Pointer part1, we copy the array only once, in the range operation. During the function call we pass the pointer.

```jsx
for _, v := range *a {
		r += v
	}
```

For pointer part2, we dont copy the array ever, both the function call and the range operation take a pointer to the array, which makes it much faster than part1.

For the slice operation also no copying is used.

**EXAMPLE 2: Testing the performance of various loops**

```go
//go:noinline
func Sum_PlainForLoop(s []Element) (r int64) {
	for i := 0; i < len(s); i++ {
		r += s[i][0]
	}
	return
}

//go:noinline
func Sum_OneIterationVar(s []Element) (r int64) {
	for i := range s {
		r += s[i][0]
	}
	return
}

//go:noinline
func Sum_UseSecondIterationVar(s []Element) (r int64) {
	for _, v := range s {
		r += v[0]
	}
	return
}

cpu: Intel(R) Core(TM) i5-9300H CPU @ 2.40GHz
Benchmark_PlainForLoop-8            	 2290671	       514.0 ns/op
Benchmark_OneIterationVar-8         	 2314377	       519.2 ns/op
Benchmark_UseSecondIterationVar-8   	  509379	      2275 ns/op
```

Use second iteration var takes much more time, because during every iteration the value is copied to the variable v(int64 is not a small type). Other two loops are much faster because they avoid copying.

**EXAMPLE 3: Sum using various loops**

```go
type S struct {
	a, b, c, d, e int
}

//go:noinline
func sum_UseSecondIterationVar(s []S) int {
	var sum int
	for _, v := range s {
		sum += v.c
		sum += v.d
		sum += v.e
	}
	return sum
}

//go:noinline
func sum_OneIterationVar_index(s []S) int {
	var sum int
	for i := range s {
		sum += s[i].c
		sum += s[i].d
		sum += s[i].e
	}
	return sum
}

//go:noinline
func sum_OneIterationVar_Ptr(s []S) int {
	var sum int
	for i := range s {
		v := &s[i]
		sum += v.c
		sum += v.d
		sum += v.e
	}
	return sum
}

Benchmark_UseSecondIterationVarSum-8   	  832596	      1411 ns/op
Benchmark_OneIterationVar_IndexSum-8   	 1571810	       761.0 ns/op
Benchmark_OneIterationVar_PtrSum-8     	 1568836	       762.1 ns/op
```

Benchmark_UseSecondIterationVarSum took more time because of value copy during every opeation.

One intresting thing to note is, the copying cost is significant if the value we are copying is a large sized value. Here it’s a large sized value, that is why there is huge performance difference in these functions.

If I reduce the number of fields in the struct, ie make the struct small sized, the performance difference is almost zero. They copying cost is very small in this case.

```
Benchmark_UseSecondIterationVarSum-8   	 1622338	       736.4 ns/op
Benchmark_OneIterationVar_IndexSum-8   	 1618110	       737.5 ns/op
Benchmark_OneIterationVar_PtrSum-8     	 1623066	       737.6 ns/op
```