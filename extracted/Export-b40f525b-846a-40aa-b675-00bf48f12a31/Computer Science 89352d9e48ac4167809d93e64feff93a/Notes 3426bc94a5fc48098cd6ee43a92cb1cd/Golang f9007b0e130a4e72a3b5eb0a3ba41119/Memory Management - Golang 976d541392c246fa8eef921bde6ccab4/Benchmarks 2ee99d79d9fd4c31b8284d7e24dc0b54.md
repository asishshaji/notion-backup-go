# Benchmarks

[How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

There can be multiple solutions to a problem. During a **benchmark**, execution statistics are gathered (the computation time, the number of allocations, and the number of function calls). Using these statistics we can choose a solution.

## How to run a benchmark?

```bash
go test -bench=.
```

### bench.go

```go
package basic

import (
	"bytes"
	"strings"
)

func ConcatenateBuffer(first string, second string) string {
	var buffer bytes.Buffer
	buffer.WriteString(first)
	buffer.WriteString(second)
	return buffer.String()
}

func ConcatenateJoin(first string, second string) string {
	return strings.Join([]string{first, second}, "")
}
```

### bench_test.go

```go
package basic

import "testing"

var result string

func BenchmarkConcatenateBuffer(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateBuffer("test2", "test3")
	}
	result = s
}

func BenchmarkConcatenateJoin(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = ConcatenateJoin("test2", "test3")
	}
	result = s
}
```

```go
// standard library
// src/testing/benchmark.go (v1.11.4)
type BenchmarkResult struct {
    N         int           // The number of iterations.
    T         time.Duration // The total time taken.
    Bytes     int64         // Bytes processed in one iteration.
    MemAllocs uint64        // The total number of memory allocations.
    MemBytes  uint64        // The total number of bytes allocated.
}

```

`go test -bench Join` Runs the benchmark functions that contain the string `Join`. BenchmarkConcatenateJoin will run the other won’t. 

`go test -bench Join -benchtime 5s` benchmark’s execution time.

## How to read benchmark results?

`go test -bench . -benchmem`

![Untitled](Benchmarks%202ee99d79d9fd4c31b8284d7e24dc0b54/Untitled.png)

![Untitled](Benchmarks%202ee99d79d9fd4c31b8284d7e24dc0b54/Untitled%201.png)

The first part is two Go env variables `GOOS` and `GOARCH`. 

The ***number of cores*** is appended to the function name, `-cpu` can be used to change the number of cores

***Iterations*** represent the number of times the for loop(for loop inside the benchmark function) has run to obtain the statistics. We can increase the number of iterations by using `benchtime` flag to increase the duration.

***Nanoseconds per operation*** gives an idea on average how your function runs. `ConcatenateBuffer` function takes an average of 55.97 ns to run. `ConcatenateJoin` function takes on average 33.63 ns to run. The fastest function in the context of our benchmark is `ConcatenateJoin`. 

***Number of bytes allocated per operation*** is present only if you add the flag `-benchmem`. This gives the memory consumption of your benchmark functions. If your focus is to improve memory usage, then you should focus on these statistics.

***Number of allocations per operation*** represents the average number of memory allocations per run.

## Detect memory allocations

```go
package main

// inports

func main() {
    basic.ConcatenateBuffer("first", "second")
    basic.ConcatenateJoin("first", "second")
}
```

```bash
go build -o allocDetect main.go
GODEBUG=allocfreetrace=1 ./allocDetect &>> trace.log
```

`GODEBUG` is an env variable that accepts a list of key-value pairs. We ask the runtime to generate a stack trace for each allocation and free.

```bash
cat trace.log | grep -n basic/bench.go
1091:   /home/runner/FlamboyantDangerousNewsaggregator/basic/bench.go:9 +0x4f fp=0xc0000a0f28 sp=0xc0000a0ed8 pc=0x45d98f
1108:   /home/runner/FlamboyantDangerousNewsaggregator/basic/bench.go:12 +0x99 fp=0xc0000a0f28 sp=0xc0000a0ed8 pc=0x45d9d9
1129:   /home/runner/FlamboyantDangerousNewsaggregator/basic/bench.go:16
```

![Untitled](Benchmarks%202ee99d79d9fd4c31b8284d7e24dc0b54/Untitled%202.png)

9: `var buffer bytes.Buffer`

12: `buffer.String()`