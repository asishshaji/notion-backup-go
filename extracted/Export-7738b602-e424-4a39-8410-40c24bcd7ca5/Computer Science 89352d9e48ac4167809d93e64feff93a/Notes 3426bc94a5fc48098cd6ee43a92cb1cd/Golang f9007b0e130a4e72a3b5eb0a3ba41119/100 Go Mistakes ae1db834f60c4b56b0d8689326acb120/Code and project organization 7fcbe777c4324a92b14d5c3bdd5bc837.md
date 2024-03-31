# Code and project organization

## **Variable shadowing**

![Wrong](Code%20and%20project%20organization%207fcbe777c4324a92b14d5c3bdd5bc837/Untitled.png)

Wrong

![Correct](Code%20and%20project%20organization%207fcbe777c4324a92b14d5c3bdd5bc837/Untitled%201.png)

Correct

`client` variable is declared outside, `client` variable is declared again in the local scope, when you use the client variable outside, the value is `nil`

## Nested Loops

> **Align the happy path to the left; you should quickly be able to scan down one column to see the expected execution flow. (AVOID NESTING)**
> 

![Untitled](Code%20and%20project%20organization%207fcbe777c4324a92b14d5c3bdd5bc837/Untitled%202.png)

If any *If *****block returns, we should omit the *else* block in all cases. It’s more readable.

![Wrong](Code%20and%20project%20organization%207fcbe777c4324a92b14d5c3bdd5bc837/Untitled%203.png)

Wrong

![Correct](Code%20and%20project%20organization%207fcbe777c4324a92b14d5c3bdd5bc837/Untitled%204.png)

Correct

> Reduce the number of nested blocks, aligning the happy path on the left, and return early as possible
> 

## Init functions

It takes no arguments and returns no result. 

```go
package main

import "fmt"

var a = func() int {
    fmt.Println("var")
    return 0
}()

func init() {
    fmt.Println("init")
}

func main() {
		fmt.Println("main")
}

var
init
main
```

The init function is executed when a package is initialized, ie, when importing a package.

We can define multiple init functions per package. When we do, the execution order of the init function inside the package is based on the source file’s alphabetical order.

a.go → init()

b.go → init(), both belonging to the same package, a.go init will be executed first.

The same file can have multiple init functions, they are executed in the order they are defined.

![Untitled](Code%20and%20project%20organization%207fcbe777c4324a92b14d5c3bdd5bc837/Untitled%205.png)

Doing this will call the `init` method of the foo package, cannot call another method.

`init` method cannot be invoked directly.

***DON’T EVER OPEN A DATABASE CONNECTION INSIDE AN INIT METHOD.***

- since init doesn't return anything, you now have to account for bad error handling.
- if we add tests to this file, the init method will be executed first, we might not be what we want.

## Getters and Setters

The getter for a field *balance* should be`.Balance()`, the setter should be `SetBalance`.

## Interfaces

> Abstractions should be discovered, not created.
>