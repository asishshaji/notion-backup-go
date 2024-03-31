# Type System Overview

![Untitled](Type%20System%20Overview%20370fe32d05764fdba99019d3e0470f2a/Untitled.png)

- Registers store binary data, composed of ones and zeros, but their interpretation can vary widely.
- Without a type system, operations on registers might produce results that don't make sense.
- A type system specifies valid operations for each data type, ensuring meaningful computations.
- It guides developers to perform operations in the intended manner.
- The goal is to enforce correctness and prevent unintended consequences in computations.

### Creating new types

- type definition
    
    ```jsx
    type Myint int
    ```
    
    Myint and int two separate types. The new defined type and the source has the same underlying type, and their values can be converted to each other.
    
- type alias
    
    ```jsx
    type Myint = int
    ```
    
    Myint and int are exactly the same type, not used much. Useful for refactoring code.
    

In Go, types can be categorised as either named or unnamed, also known as composite types. Let's explore each category:

1. **Named Types**:
    - **Basic Types**: These are types like `int`, `float64`, `string`, etc., which have predefined behaviour and are built into the language. They are named because they have a specific identifier associated with them.
    - **Struct Types**: When you define a struct type, you give it a name, such as `type MyStruct struct {...}`. The name `MyStruct` identifies the type.
    - **Interface Types**: Similarly, when you define an interface type, you give it a name, such as `type MyInterface interface {...}`. The name `MyInterface` identifies the interface type.
    - **Pointer Types**: When you define a pointer to a type, such as `type MyPointer *MyStruct`, the pointer type is also named (`MyPointer`), even though it's composed of another type.
    - **Function Types**: When you define a function type, such as `type MyFunc func(int) string`, the function type is named (`MyFunc`).
    
    ```jsx
    type (
    	MyInt int
    	Age   int
    	Text  string
    )
    
    ```
    
2. **Unnamed Types (Composite Types)**:
    - **Array Types**: When you define an array type, such as `type MyArray [5]int`, the type is unnamed (`[5]int`). Arrays have a specific size and element type but lack a distinct name.
    - **Slice Types**: Similarly, when you define a slice type, such as `type MySlice []int`, the type is unnamed (`[]int`). Slices are dynamically sized views into arrays but don't have a separate name beyond their composition.
    - **Map Types**: Map types, defined like `type MyMap map[string]int`, are also unnamed (`map[string]int`). Maps represent a collection of key-value pairs but don't have a distinct name beyond their composition.
    - **Channel Types**: Channel types, defined like `type MyChannel chan int`, are also unnamed (`chan int`). Channels represent communication channels for goroutines but don't have a distinct name beyond their composition.
    
    ```jsx
    type StringArray [5]string
    type StringSlice []string
    
    ```
    

Unnamed types are composite types composed of other types but don't have an explicit identifier beyond their definition. They are typically used when the specific type itself is less important than its structure or composition.

"Unnamed" reflects the fact that these types lack an explicit identifier beyond their composition and are defined as part of a larger type declaration.

## Underlying types

- Each type has an underlying type
- for builtin types the underlying type is the builtin type itself
    
    ```jsx
    // The underlying types of the following ones are both int.
    type (
    	MyInt int
    	Age   MyInt
    )
    ```
    
- underlying type of an unnamed type is the composite type itself
    
    ```jsx
    // The following new types have different underlying types.
    type (
    	IntSlice   []int   // underlying type is []int
    	MyIntSlice []MyInt // underlying type is []MyInt
    	AgeSlice   []Age   // underlying type is []Age
    )
    ```
    

- in type declaration, the newly declared type and source type both has the same underlying type.
    - How to find the underlying type.
        - if a builtin type or unnamed type is encountered, the tracing must be stopped.
    
    ```jsx
    // The underlying types of []Age, Ages, and AgeSlice
    // are all the unnamed type []Age.
    type Ages AgeSlice
    ```
    

```
MyInt → int
Age → MyInt → int
IntSlice → []int
MyIntSlice → []MyInt →[]int
AgeSlice → []Age →[]MyInt →[]int
Ages → AgeSlice → []Age →[]MyInt →[]int
```

[A Closer Look at Go (golang) Type System.](https://medium.com/@ankur_anand/a-closer-look-at-go-golang-type-system-3058a51d1615)

## Named types

- int, float, string, bool etc are predeclared
- predeclared types are called named types
- any type created using type declaration is a named type

## Unnamed types

- Composite types - array, struct, pointer, function, interfaces, slice, map, and channel types are all unnamed types

```jsx
[]string // unnamed type
map[string]string // unnamed type
[10]int // unnamed type
```

## Underlying types

![Untitled](Type%20System%20Overview%20370fe32d05764fdba99019d3e0470f2a/Untitled%201.png)

- Each type has an underlying type
- For basic types it is the type itself
- Otherwise, T’s underlying type is one which T refers to in its type declaration
    - **3 and 8:** We have a predeclared type of `string` , so the underlying type will be `T` itself — `string`
    - **5 and 7:** We have a `type literal`, so the underlying type will be `T` itself — `map[string]int` and `N` pointer. **Note** these type literals are also `unnamed type`
    - **4, 6, and 10:** `T`'s underlying type is the underlying type which `T` refers to in its [type declaration](https://golang.org/ref/spec#Type_declarations). For example,`B` refers to `A,` hence B is a string.

# Assignability

 

![Untitled](Type%20System%20Overview%20370fe32d05764fdba99019d3e0470f2a/Untitled%202.png)

![Untitled](Type%20System%20Overview%20370fe32d05764fdba99019d3e0470f2a/Untitled%203.png)

i and j both have the same underlying type, but cannot be assigned directly. `i` is a named type, `j` is a named type of aint

![Untitled](Type%20System%20Overview%20370fe32d05764fdba99019d3e0470f2a/Untitled%204.png)

This works, because m and mM are both unnamed types, and the underlying type of both are the same.