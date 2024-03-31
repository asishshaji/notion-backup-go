# Explanation using code

```go
/* line 19 */ func ex1() {
/* line 20 */     var x int64
/* line 21 */     var y interface{}
/* line 22 */     x = 2
/* line 23 */     y = x
/* line 24 */     _ = y
/* line 25 */ }

	ex1.go:20		0x454c4e		48c7042400000000	       MOVQ $0x0, 0(SP)	
  ex1.go:21		0x454c56		440f117c2410		           MOVUPS X15, 0x10(SP)	
  ex1.go:22		0x454c5c		48c7042402000000	       MOVQ $0x2, 0(SP)	
  ex1.go:23		0x454c64		48c744240802000000	      MOVQ $0x2, 0x8(SP)	
  ex1.go:23		0x454c6d		488d056c410000		        LEAQ 0x416c(IP), AX	
  ex1.go:23		0x454c74		4889442410		             MOVQ AX, 0x10(SP)	
  ex1.go:23		0x454c79		488d442408		             LEAQ 0x8(SP), AX	
  ex1.go:23		0x454c7e		4889442418		             MOVQ AX, 0x18(SP)
```

**ex1.go:20		0x454c4e		48c7042400000000	        MOVQ $0x0, 0(SP)**

- **0x454c4e:** program counter in hexadecimal
- **48c7042400000000:** executable instruction formatted as hexadecimal
- **MOVQ $0x0, 0(SP): MOVQ SRC, DST**
    - **Q** in MOVQ stands for quadword.
    - MOVQ is used for copying 8 bytes.	****
    - **$0x0:** $ indicates SRC is not a memory address, but a literal value
    - The value to copy is 0x0, or the integer value 0.
    - ie the value of x is set to zero
    - set the 8 bytes from the start of SP, ie 0 → 7
- **0(SP)**
    - 0 indicates the offset of zero bytes from some address
    - SP, stack pointer, which points to the top of the current call stack frame.
    - 0(SP) can be translated as zero bytes from the top of the current stack frame. ****
    
    ![Untitled](Explanation%20using%20code%200b615461aee540d0b6b026592bf555e5/Untitled.png)
    

**ex1.go:21		0x454c56		440f117c2410		           MOVUPS X15, 0x10(SP)**	

- var y interface{}
- we know that an interface is two word size. Therefore we need 16 bytes.
- MOVUPS copies an unaligned, packed, single-precision floating point value.
- X15 is a special register and holds the zero value

![Untitled](Explanation%20using%20code%200b615461aee540d0b6b026592bf555e5/Untitled%201.png)

- [https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md](https://github.com/golang/go/blob/master/src/cmd/compile/abi-internal.md)
- Because MOVUPS is copying 16 bytes of data and the X15 register is 0, this instruction is essentially reserving 16 bytes on the stack starting at DST.
- **0x10(SP):** 16 bytes from the top of the current stack frame.

**ex1.go:22	0x454c5c	48c7042402000000	MOVQ $0x2, 0(SP)**

- x = 2
- Copies the literal value 2 to the 8bytes starting from the current stack pointer.

![Untitled](Explanation%20using%20code%200b615461aee540d0b6b026592bf555e5/Untitled%202.png)

**ex1.go:23	0x454c64	48c744240802000000	MOVQ $0x2, 0x8(SP)**

- y = x
- Moves the literal value 2 to the 8 bytes starting from 0x8(SP)
- The Go compiler was able to determine that the only value ever assigned to y would be an int64, and so an extra eight bytes were allocated on the stack in order to store the int64 value assigned to y.

![Untitled](Explanation%20using%20code%200b615461aee540d0b6b026592bf555e5/Untitled%203.png)

**ex1.go:23	0x454c6d	488d056c410000	LEAQ 0x416c(IP), AX**

- y = x
- LEAQ
    - Load Effective Address, Q quad word
    - Copies the addresses around
    - LEAQ 0x416c(IP), AX stores the address of the next CPI instruction in register AX.
    - Ultimately what is stored in AX is the address of type.int64, a global value that specifies the internal type for an int64.
    - Basically loads the type of value that you want to store in the interface.

**ex1.go:23	0x454c74	4889442410	MOVQ AX, 0x10(SP)**

- y = x
- Now since the interface is two words, we are now going to write the type of the interface using the assembly.
- This assigns the address of the global value type.int64 to the interface’s first uintptr, the one that points to the underlying type.

![Untitled](Explanation%20using%20code%200b615461aee540d0b6b026592bf555e5/Untitled%204.png)

**ex1.go:23	0x454c79	488d442408	LEAQ 0x8(SP), AX**

- y = x
- Loads the address of 0x8(SP) to AX, why???? for the value ptr inside the interface to refer to.

**ex1.go:23	0x454c7e	4889442418	MOVQ AX, 0x18(SP)**

- y = x
- Sets the address to the value pointer of the interface

![Untitled](Explanation%20using%20code%200b615461aee540d0b6b026592bf555e5/Untitled%205.png)