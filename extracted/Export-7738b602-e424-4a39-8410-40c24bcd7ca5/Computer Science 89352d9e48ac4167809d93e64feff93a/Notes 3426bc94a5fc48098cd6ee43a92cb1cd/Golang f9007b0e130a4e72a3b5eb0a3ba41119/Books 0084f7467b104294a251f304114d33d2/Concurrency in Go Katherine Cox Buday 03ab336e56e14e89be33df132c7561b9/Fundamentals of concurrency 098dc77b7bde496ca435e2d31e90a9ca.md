# Fundamentals of concurrency

## Data Race

Data race happens when two or more operations must happen in an order, but the code has been written in such a way that the order is not maintained. The program is written in such a way that the order of execution is not guarrented. 

```go
var data int
go func() {
  data++
}()
if data == 0 {
fmt.Printf("the value is %v.\n", data)
}
```

Most of the time data race happens because the developer thinks in a sequential manner. They assume that because the lines of code falls before another that it will run first. 

Using locks to synchronise memory is often a bad idea, calling lock() function can make our program slow. Every time we lock something, we are imposing a delay.

## Deadlock

### Code

![Untitled](Fundamentals%20of%20concurrency%20098dc77b7bde496ca435e2d31e90a9ca/Untitled.png)

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Bank struct {
	mu      sync.Mutex
	Balance int
	Owner   string
}

func (b *Bank) TransferTo(wg *sync.WaitGroup, to *Bank, amount int) {
	defer wg.Done()
	fmt.Printf("%s in locking their account\n", b.Owner)
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Balance -= amount

	time.Sleep(2 * time.Second)

	fmt.Printf("%s is locking %s's account\n", b.Owner, to.Owner)
	to.mu.Lock()
	to.mu.Lock()
	defer to.mu.Unlock()
	to.Balance += amount

	fmt.Printf("Money transfered from %s to %s : amount : %d\n", b.Owner, to.Owner, amount)
}

func Init() {
	alice := Bank{
		Balance: 1000,
		Owner:   "Alice",
	}
	bob := Bank{
		Balance: 2000,
		Owner:   "Bob",
	}

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go bob.TransferTo(&wg, &alice, rand.Intn(200))
		go alice.TransferTo(&wg, &bob, rand.Intn(200))
	}
	wg.Wait()
}
```

![deadlock.png](Fundamentals%20of%20concurrency%20098dc77b7bde496ca435e2d31e90a9ca/deadlock.png)

1, 2 := process A locks ‘a’ and process B locks ‘b’.

4 = B tries to access ‘a’, since a is already locked, process B is put to sleep. Inorder for process B to wake up

var ‘a’ must be unlocked.

3 = process A tries to lock ‘b’. But ‘b’ is already locked by Process B. So by locking var ‘b’, process A is put to sleep. 

The program will never recover without outside intervention.

```go
type Value struct {
	val int
	mu sync.Mutex
}

var wg sync.Waitgroup
printSum := function(a, b * Value) {
	def wg.Done()
	
	a.mu.Lock()
	defer a.mu.Unlock()

	time.Sleep(time.Second)

	b.mu.Lock()
	defer b.mu.Unlock()

	fmt.Println(a.val + b.val)
}

var a, b Value
wg.Add(2)
go printSum(&a,&b)
go printSum(&b,&a)
wg.Wait()
```

## Livelock

a more “smarter” deadlock.

A livelock is a situation in concurrent programming where two or more threads or processes are actively trying to resolve a resource conflict, but they end up stuck in a repetitive, non-productive cycle, where none of them makes any progress. In essence, it's a scenario where the threads are constantly reacting to each other's actions, but they are unable to complete their tasks. Unlike a deadlock where threads are completely blocked, in a livelock, threads are not blocked but are not making any real progress either.

[Deadlock, Livelock, Race condition and Starvation](https://pradeesh-kumar.medium.com/deadlock-livelock-race-condition-and-starvation-c225018bbae6)

![Untitled](Fundamentals%20of%20concurrency%20098dc77b7bde496ca435e2d31e90a9ca/Untitled%201.png)

- T1 has acquired the lock on the file and waiting for the printer to print the file, but the printer has been locked by T2.
- T2 has acquired the lock on the printer and waiting for the file, but the file has been locked by T2.
- After some time, at the same time, both threads realize that they will not get the lock, it releases their lock.

![Untitled](Fundamentals%20of%20concurrency%20098dc77b7bde496ca435e2d31e90a9ca/Untitled%202.png)

- Now it's possible that earlier, T1 will acquire the lock on the Printer and T2 will acquire the lock on the file.
- T1 wants to acquire the file, which has been locked by T2. T2 wants to acquire the printer, which has been locked by T1.

![Untitled](Fundamentals%20of%20concurrency%20098dc77b7bde496ca435e2d31e90a9ca/Untitled%203.png)

• This will continue to exchange the resources forever. Both threads will never acquire the file and printer at the same time. This situation is called Livelock.

[](https://github.com/kat-co/concurrency-in-go-src/blob/master/notes/dead-writing/livelocks/livelock-example.go)

We are not exactly stuck, but we are not making any progress.