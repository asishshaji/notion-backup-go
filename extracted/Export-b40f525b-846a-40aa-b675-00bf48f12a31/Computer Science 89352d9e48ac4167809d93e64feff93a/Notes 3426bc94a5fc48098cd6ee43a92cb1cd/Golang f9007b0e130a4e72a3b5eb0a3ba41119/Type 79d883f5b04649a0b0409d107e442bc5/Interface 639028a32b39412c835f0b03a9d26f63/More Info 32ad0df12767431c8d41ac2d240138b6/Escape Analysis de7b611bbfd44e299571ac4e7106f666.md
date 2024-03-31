# Escape Analysis

[](https://github.com/akutz/go-interface-values/blob/main/docs/03-escape-analysis/01-overview.md)

Escape analysis applies only to reference types.

The compile time process Go uses to determine whether memory is dynamically managed on the heap or can be allocated on the stack is known as escape analysis. Escape analysis walks a program’s abstract syntax tree to build a graph of all the variables encountered.

```
./docs/03-escape-analysis/examples/ex1/main.go:23:2: leaking param: username
./docs/03-escape-analysis/examples/ex1/main.go:24:2: password does not escape
./docs/03-escape-analysis/examples/ex1/main.go:25:2: leaking param: token to result ~r0 level=0
./docs/03-escape-analysis/examples/ex1/main.go:31:10: "newLoginToken" escapes to heap
./docs/03-escape-analysis/examples/ex1/main.go:42:3: moved to heap: username1
./docs/03-escape-analysis/examples/ex1/main.go:45:3: moved to heap: cookieJar
./docs/03-escape-analysis/examples/ex1/main.go:43:18: new(string) escapes to heap
```

***leak, escape, moved.***

# Leak

[](https://github.com/akutz/go-interface-values/blob/main/docs/03-escape-analysis/02-leak.md)

Escape analysis inspects variables with the potential to escape when they are used. If that potential exists, the variables are marked as ***leaking***.

### Criteria for leaking

- The variable must be a function parameter
- The variable must be a reference type, e.g.: channels, interfaces, maps, pointers, slices

If the above criteria are met, then a parameter will leak if:

- The variable is returned from the same function
- is assigned to a sink outside of the stack frame to which the variable belongs.

Two types of leaks

- Leaking to a sink
- Leaking to a result

### Leaking to a sink

If a function’s parameter is a reference type and the function assigns the parameter to a variable outside of the function, the variable is leaking. 

```go
var sink *int32

func recordID(id *int32) { // leaking param: id
	sink = id
}
```

The function recordID leaks the parameter id to the package level field sink. 

### Leaking to result

The reference parameter is returned from a function.

```go
func validateID(id *int32) *int32 { // leaking param: id to result ~r1 level=0
	return id
}
```

A leak doesn’t necessarily mean that the variable will escape to the heap.

### Leaking without escape

```go
func main() {
	var id1 int32 = 4096
	if validateID(&id1) == nil {
		os.Exit(1)
	}

	var id2 *int32 = new(int32) // new(int32) does not escape
	*id2 = 4096
	validID := validateID(id2)
	if validID == nil {
		os.Exit(1)
	}
}

./docs/03-escape-analysis/examples/ex2/main.go:22:17: leaking param: id to result ~r0 level=0
./docs/03-escape-analysis/examples/ex2/main.go:32:22: new(int32) does not escape
```

- The `id` parameter for the `validateID` function is leaking to the result because the function returns the incoming parameter and thus there is potential for the value of `id` to outlive its stack frame.
- The value `id1` is not even mentioned because it is a value type and escape analysis only applies to reference types. While a pointer is passed to `validateID`, the Go compiler optimized the pointer to the stack.
- The value `id2` is a reference type, but it does not escape. `validID` and id2 are on the same stack frame, thus it doesn't outlive its stack frame. Therefore `id2` doesn't escape to the heap.

## Escape

### Criteria

- The variable must be a reference type, e.g.: channels, interfaces, maps, pointers, slices
- A value type stored in an interface value can also escape to the heap.

If the above criteria are met, then a parameter will escape if it outlives its current stack frame. That usually happens when 

- The variable is sent to a function that assigns the variable to a sink outside the stack frame.
- The function where the variable is declared assigns it to a sink outside the stack frame.

## Move

### Criteria

- The variable must be a value type
- A reference to the variable is assigned to a location outside of the local stack frame.

### *A value **escapes** if it is initially a reference type or the operation is storing a value type in an interface*

### *A value **moves** if it is initially a value type and the operation is not storing a value type in an interface.*