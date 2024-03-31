# Assembler

[A Quick Guide to Go's Assembler - The Go Programming Language](https://go.dev/doc/asm)

[Chapter I: Go Assembly](https://cmc.gitbook.io/go-internals/chapter-i-go-assembly)

Go offers pseudo-assembly. The compiler outputs an abstract, portable form of assembly that doesn’t actually map to any real hardware.

This pseudo-assembly is then used to generate machine specific instructions for the targeted hardware. 

```go
//go:noinline
func add(a, b int32) (int32, bool) { return a + b, true }
func main() { add(10, 32) }
```

```bash
GOOS=linux GOARCH=amd64 go tool compile -S direct_topfunc_call.go

0x0000 TEXT        "".add(SB), NOSPLIT, $0-16
  0x0000 FUNCDATA    $0, gclocals·f207267fbf96a0178e8758c6e3e0ce28(SB)
  0x0000 FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x0000 MOVL        "".b+12(SP), AX
  0x0004 MOVL        "".a+8(SP), CX
  0x0008 ADDL        CX, AX
  0x000a MOVL        AX, "".~r2+16(SP)
  0x000e MOVB        $1, "".~r3+20(SP)
  0x0013 RET

0x0000 TEXT        "".main(SB), $24-0
  ;; ...omitted stack-split prologue...
  0x000f SUBQ        $24, SP
  0x0013 MOVQ        BP, 16(SP)
  0x0018 LEAQ        16(SP), BP
  0x001d FUNCDATA    $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x001d FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  0x001d MOVQ        $137438953482, AX
  0x0027 MOVQ        AX, (SP)
  0x002b PCDATA        $0, $0
  0x002b CALL        "".add(SB)
  0x0030 MOVQ        16(SP), BP
  0x0035 ADDQ        $24, SP
  0x0039 RET
  ;; ...omitted stack-split epilogue...
```

### Dissecting `add`

0x0000 TEXT "".add(SB), NOSPLIT, $0-16

- Ox0000 Offset of current instruction, relative to the start of the function
- TEXT “”.add
    - The TEXT declares “”.add symbol as part of the .text section and indicates that the instructions that follow are the body of the function. “” will be replaced by the name of the current package at link-time, ie “”.add → main.add once linked into our final binary
    - SB is the virtual register that holds the “static-base” pointer, ie the address of the beginning of the address-space of our program.
        - “”.add(SB) declares that our symbol is located at some constant offset from the start of our address-space.

```go
;; Declare global function symbol "".add (actually main.add once linked)
;; Do not insert stack-split preamble
;; 0 bytes of stack-frame, 16 bytes of arguments passed in
;; func add(a, b int32) (int32, bool)
0x0000 TEXT    "".add(SB), NOSPLIT, $0-16
  ;; ...omitted FUNCDATA stuff...
  0x0000 MOVL    "".b+12(SP), AX        ;; move second Long-word (4B) argument from caller's stack-frame into AX
  0x0004 MOVL    "".a+8(SP), CX        ;; move first Long-word (4B) argument from caller's stack-frame into CX
  0x0008 ADDL    CX, AX            ;; compute AX=CX+AX
  0x000a MOVL    AX, "".~r2+16(SP)   ;; move addition result (AX) into caller's stack-frame
  0x000e MOVB    $1, "".~r3+20(SP)   ;; move `true` boolean (constant) into caller's stack-frame
  0x0013 RET                ;; jump to return address stored at 0(SP)

   |    +-------------------------+ <-- 32(SP)              
   |    |                         |                         
 G |    |                         |                         
 R |    |                         |                         
 O |    | main.main's saved       |                         
 W |    |     frame-pointer (BP)  |                         
 S |    |-------------------------| <-- 24(SP)              
   |    |      [alignment]        |                         
 D |    | "".~r3 (bool) = 1/true  | <-- 21(SP)              
 O |    |-------------------------| <-- 20(SP)              
 W |    |                         |                         
 N |    | "".~r2 (int32) = 42     |                         
 W |    |-------------------------| <-- 16(SP)              
 A |    |                         |                         
 R |    | "".b (int32) = 32       |                         
 D |    |-------------------------| <-- 12(SP)              
 S |    |                         |                         
   |    | "".a (int32) = 10       |                         
   |    |-------------------------| <-- 8(SP)               
   |    |                         |                         
   |    |                         |                         
   |    |                         |                         
 \ | /  | return address to       |                         
  \|/   |     main.main + 0x30    |                         
   -    +-------------------------+ <-- 0(SP) (TOP OF STACK)

(diagram made with https://textik.com)
```

---

[Golang UK Conference 2016 - Michael Munday - Dropping Down Go Functions in Assembly](https://www.youtube.com/watch?v=9jpnFmJr2PE)

![Untitled](Assembler%20c69e75348f0a4d54827dcaaa5e0f3b92/Untitled.png)

![Untitled](Assembler%20c69e75348f0a4d54827dcaaa5e0f3b92/Untitled%201.png)