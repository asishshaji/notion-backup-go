# Strings

- An immutable sequence of characters, where characters are Unicode points.
- Earlier characters were represented using ASCII, using 1 byte for representation.
- Couldn't represent foreign characters
- Unicode representation was introduced
    - Variable length encoding
    - Golang uses the utf-8 encoding scheme, ie it can use 1 byte up to 4 bytes for storing characters in memory
- Go stores string physically in memory as a sequence of bytes, 1 byte to 4 bytes for each character.
- Logically as a sequence of runes (Unicode points)

byte ↔ uint8 

rune ↔ int32

```go
s := "Énemy"
fmt.Printf("%T %[1]v %d\n", s, len(s))
fmt.Printf("%T %[1]v\n", []rune(s))
fmt.Printf("%T %[1]v\n", []byte(s))

string Énemy 6
[]int32 [201 110 101 109 121]
[]uint8 [195 137 110 101 109 121]
```

In rune conversion, each of the values in the array is the Unicode representation of the characters.

here 201 is stored in memory as two bytes, and É takes 2 bytes to store in memory physically.

***When we print the length of the string we get 6.*** 

From the declaration, we know that a string is a `byte` sequence wrapper. 

```go
type _string struct {
	elements *byte // underlying bytes
	len      int   // number of bytes
}
```

![Untitled](Strings%20e6ee53a1f2604cffadd5a43c357e2bb0/Untitled.png)

If we do 

h = “Cat “

h += “asish”

New memory will be allocated for this and the previous data will be copied, and then the concatenated value is added.

In a conversion from a rune slice to a string, each slice element (a rune value) will be UTF-8 encoded as from one to four bytes and stored in the result string. 

When a byte slice is converted to a string, the underlying byte sequence of the result string is also just a deep copy of the byte slice. A memory allocation is needed to store the deep copy in each of such conversions. The reason why a deep copy is essential is slice elements are mutable but the bytes stored in strings are immutable, so a byte slice and a string can't share byte elements.

- for-range on a string will iterate the Unicode code points(as `rune` values), instead of bytes, in a string.