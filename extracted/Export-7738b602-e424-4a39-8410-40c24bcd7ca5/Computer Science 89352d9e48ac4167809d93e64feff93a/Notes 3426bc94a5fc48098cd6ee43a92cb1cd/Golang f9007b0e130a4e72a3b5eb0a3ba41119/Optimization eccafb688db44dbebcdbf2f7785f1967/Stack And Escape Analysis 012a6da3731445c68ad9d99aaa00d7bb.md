# Stack And Escape Analysis

main.go:3:6: can inline main
main.go:4:14: make([]int, 10000, 10000) escapes to heap

## Escape Analysis

Not all values are capable of being allocated on the stack. ***One principle condition to allocate a value part on stack is the value part will be only used in one goroutine during its life time.*** 

If the compiler detects a value is used by more than one goroutine or it’s unable to make sure whether or not the value part is used by only one goroutine, then it will let the value part be allocated on heap. ie esacpes to the heap.

It’s often very complicated to find all the values which are to be allocated to the heap during the compile time. And a more powerful escape analysis implementation will increase compilation time.

So at run time, some value parts will be allocated on heap even if they could be safely be allocated on the stack.

If the size of value part is too large, then the compiler will still let the value part allocated to the heap.

The basic escape analysis units are functions. O***nly the local values will be escape analyzed, all the package-level variables are allocated on the heap.***

-m compiler option is used to show escape analysis result.

```go
func Heap() {
	var (
		a = 1
		b = false
		c = make(chan struct{})
	)

	go func() {
		if b {
			a++
		}
		close(c)
	}()

	<-c
	println(a, b)
}

/home/asish/Learning/golang/heap/escape.go:10:5: can inline Heap.func1
/home/asish/Learning/golang/heap/escape.go:5:3: moved to heap: a
/home/asish/Learning/golang/heap/escape.go:10:5: func literal escapes to heap
```

The direct part of c is allocated on the stack, the indirect parts of channels are always allocated on heap, so in escape messages for channel indirect parts are always omitted.

Here `b` is allocated on the stack, a is allocated to the heap, why???

The reason is that the escape analysis module is so smart that it detects the variable `b` is never modified, and it understands it’s a good idea to use a copy of the variable b in the new goroutine.

```go
func Heap() {
	var (
		a = 1
		b = false
		c = make(chan struct{})
	)

	go func() {
		if b {
			a++
		}
		close(c)
	}()

	b = !b
	<-c
	println(a, b)
}

/home/asish/Learning/golang/heap/escape.go:10:5: can inline Heap.func1
/home/asish/Learning/golang/heap/escape.go:5:3: moved to heap: a
/home/asish/Learning/golang/heap/escape.go:6:3: moved to heap: b
/home/asish/Learning/golang/heap/escape.go:10:5: func literal escapes to heap
```

Now that I modify the value of b, it’s now moved to the heap.

## Stack frames of function calls

At compile time, the compiler calculates the maximum possible stack size  of a function that can be used during runtime. The calculated size is called the stack frame size of the function.

-S compiler option is used.

Initial default stack size is 2KiB. Default max stacksize is 1GB on 64bit systems and 250MB on 32-bit systems.

For some goroutines, as functions go deeper and deeper, more and more function stack frames are needed. The stack would need to grow on demand. And stacks should also support shrinkage, once the function call depth becomes shallower.

## Situations where variables escapes

- ***Local variables declared in a loop will escape to heap if it’s referenced by a value out of the loop.***
    
    ```go
    package main 
    
    func main(){
    	var x *int 
    	for {
    		var n = 1
    		x = &n
    		break
    	}
    	_ = x
    }
    ```
    
    - The reason is there might be many coexisting instances of n if its containing loop needs many steps to finish.
    - The number of instances is often hard or impossible to determine at compile time.
- ***Value parts referenced by an argument will escape to heap if the argument is passed to interface method calls.***
    
    ```go
    type I interface {
    	M(*int)
    }
    
    type T struct{}
    
    func (T) M(*int) {}
    
    var t T
    var i I = t
    
    func Interface() {
    	var x int
    	t.M(&x)
    	var y int
    	i.M(&y)
    }
    ```
    
    - Here y escapes to the heap.
    - It often impossible or too expensive for compilers to determine the dynamic value of an interface, so compiler often gives up to do so for most cases.
    - Potentially, the concrete method of the interface value could pass its argument to some other goroutines.
    - So for safety, the compiler lets the value part referenced by the arguments escape to heap.

---

Global variables, local variables with a large memory footprint, and local variables that cannot be reclaimed immediately after a function call are all stored inside the heap.

---

```bash
func demoFunction() *int {
	var data int = 11
	return &data
}
$> go tool compile -m
main.go:3:6: can inline demoFunction
main.go:8:6: can inline main
main.go:9:20: inlining call to demoFunction
main.go:4:6: moved to heap: data
```

```bash

func main() {
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}$> go tool compile -m
main.go:3:6: can inline main
main.go:4:14: make([]int, 10000, 10000) escapes to heap
```

A pointer to a variable inside the function is returned, if the pointer to it was not returned the variable would be stored inside the stack frame of the function. Here since the pointer to the variable is returned, the compiler will think that the variable is used elsewhere after the function ends, so it will keep it in heap.

Insufficient stack space causes the variable to escape