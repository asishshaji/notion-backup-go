# Mutex and Lock Internals in Golang

URL: https://blog.stackademic.com/mutex-internals-in-golang-1624749f35a6
Category: Golang

![https://miro.medium.com/v2/resize:fit:1152/1*J3nG_zYNaeBVsu6LRPZlMQ.png](https://miro.medium.com/v2/resize:fit:1152/1*J3nG_zYNaeBVsu6LRPZlMQ.png)

---

Photo by [Chinmay B](https://unsplash.com/@geekgunda?utm_source=medium&utm_medium=referral) on [Unsplash](https://unsplash.com/?utm_source=medium&utm_medium=referral)

[Mutex%20and%20Lock%20Internals%20in%20Golang%202229edcfb3264fe4a5ac25d2a78bdcba/0DYOSsNufwc3INAr8](Mutex%20and%20Lock%20Internals%20in%20Golang%202229edcfb3264fe4a5ac25d2a78bdcba/0DYOSsNufwc3INAr8)

In modern software development, the concept of concurrency pertains to a program’s ability to execute multiple tasks or operations simultaneously. This enables efficient utilization of system resources and enhanced performance. A standout feature of the Go programming language (Golang) is its exceptional support for concurrency. However, concurrent systems introduce challenges that require careful handling. One such challenge arises when multiple resources attempt to access and modify shared resources concurrently, leading to unpredictable and erroneous behavior.

## How to safeguard critical sections:

To tackle this issue, the concept of locks was introduced. Locks are used to safeguard critical sections of code where shared resources are accessed. They ensure that only one thread can access the protected resource at any given time.

In Golang, we leverage goroutines, which are lightweight threads. The scheduling of these goroutines is managed by the Go runtime scheduler. While delving deeply into the intricacies of the Go runtime scheduler is beyond the scope of this article, you can refer to ***[this](https://medium.com/@ravikumar19997/exploring-concurrency-threads-and-goroutines-88edbf3422df)*** article for a more depth understanding of the Go runtime scheduler.

In this article, we’ll start by exploring the fundamentals of locks. We’ll discuss how to create locks when needed. After that, we’ll delve into various locking strategies. Finally, we’ll take a deep dive into the Go mutex. I kindly ask you all to stay with me until the end of the article.

**Lock In Golang:**

In Go the unit of concurrency is goroutine, and goroutines are conceptually very similar to regular threads. You use them just like you use threads in a language like Java or C++, but the big difference is that they are user space. They are run and managed entirely by the Go runtime and not by the operating system. They’re created and scheduled by the runtime.

Like with regular threads, when you have two goroutines accessing a shared memory location, that access needs to be synchronized. You can synchronize it using Go’s lock implementation, the implementation we’ll be looking at today, the ***sync.Mutex***. The Mutex is a blocking non-recursive lock, so, there are no try acquire semantics. If a goroutine tries to acquire a lock and it can’t get it, it will block.

Goroutines Accessing the same shared resource

![Mutex%20and%20Lock%20Internals%20in%20Golang%202229edcfb3264fe4a5ac25d2a78bdcba/1J3nG_zYNaeBVsu6LRPZlMQ.png](Mutex%20and%20Lock%20Internals%20in%20Golang%202229edcfb3264fe4a5ac25d2a78bdcba/1J3nG_zYNaeBVsu6LRPZlMQ.png)

Enough of discussion if we have to build a lock from the ground up, let’s explore the essential steps involved in crafting an effective locking mechanism.

## Lock Internals and Locking Mechanism

Let’s think we have two processes in which: T1 is a reader retrieving tasks from a queue and T2 is a writer adding tasks to the queue. To prevent simultaneous access, a locking mechanism is essential. Below is the code for both of these processes:

```
// Thread 1 which is reading the tasks
func reader() {
 //read a task
 t := tasks.get()
 //Do something with the tasks
}

// Thread 2 which is writing the tasks
func writer() {
 //write to tasks
 tasks.put(t)
}
```

Let’s build a locking mechanism that ensures T1 and T2 never access the tasks queue at the same time.

**Initial Attempt: Using a Flag**

An initial approach involves ***using a flag*** to indicate whether the task queue is in use or free. When a thread accesses the queue, it sets the flag and proceeds. Code for this approach looks like:

```
var flag int
var tasks Tasks

func reader() {
 for {
  if flag == 0 {
   //Set flag which indicates flag is being set so some thread already processing.
   flag++
   t := tasks.get()
   //Unset flag which means tasks is not accessed by any thread
   flag --
   return
  }
  //else keep looping
 }
}

func writer() {
 for {
  if flag == 0 {
   //Set flag which indicates flag is being set so some thread already processing.
   flag++
   tasks.put(t)
   //Unset flag which means tasks is not accessed by any thread
   flag --
   return
  }
  //else keep looping
 }
}
```

**How do we do this?**

We use a flag to keep track of whether the task queue is available (free) or being used. If it’s available, the flag is set to 0. If it’s being used, the flag is set to 1.

How does the reader and writer code look like:

1. Check if the flag is 0 (task queue is free) and if not keep looping.
2. If it’s 0, set the flag to 1 to indicate the queue is being used.
3. Do the necessary task.
4. Set the flag back to 0 when done.

*Is this method effective? Can we finish and leave?*

No, this wouldn’t work. The first problem is with this line, ***this flag++.*** This is a read-modify-write instruction, so, the variable is read from memory into a local CPU register, it’s modified, here it’s incremented, and then it’s written back to memory.

Our one instruction actually leads to two memory accesses. Why is this an issue? The problem is that these memory accesses from different threads can get mixed up. Even if ***Thread 1’s “flag++”*** starts before Thread 2’s read, it’s possible that Thread 2 still sees the old value of the flag. So, Thread 2 might see the flag as 0, which is a problem.

This situation is called a “non-atomic operation.” It means that other processors can see the changes it’s making before it’s finished. Operations become non-atomic when they involve multiple small steps behind the scenes, like working with big data or using multiple memory accesses. They can also be non-atomic when they appear as one step but actually involve several memory actions, as we saw in the earlier example.

Our flag is not going to work because it’s not atomic fine why cannot we build something which is atomic? This sounds useful. Can we use this to solve our conundrum? Yes, we can.

**Second Attempt: Using Atomic operation**

We’ve found the solution: the atomic compare-and-swap instruction. It does what we need. This instruction updates a variable under certain conditions — it changes it only if it’s already a specific value. Importantly, this read-modify-write process happens all at once, which is exactly what we require.

Let’s apply this atomic compare-and-swap to build our lock. Code for this approach looks like:

```
var flag int
var tasks Tasks

func reader() {
 for {
  if atomic.CompareAndSwap(&flag, 0, 1) { //Set flag to 1 succeeded
   t := tasks.get()
   atomic.Store(&flag, 0) //Unset flag to 0
   return
  }
  //else keep looping
 }
}

func writer() {
 for {
  if atomic.CompareAndSwap(&flag, 0, 1) { //Set flag to 1 succeeded
   tasks.put(t)
   atomic.Store(&flag, 0) //Unset flag to 0
   return
  }
  //else keep looping
 }
}
```

How does our reader and writer code look like:

1. Try to CAS flag from 0 to 1. So, if the flag is indeed zero, it will set the flag to one in a single, uninterrupted step.
2. If CAS instruction fails to set the flag (maybe because another process changed it before we could), we won’t give up. We’ll simply try again. We’ll keep looping until we successfully set the flag to one
3. Do the necessary task.
4. Atomically restored flag value to 0.

*It makes sense, does this work? Actually, it does work.*

This is a simplified implementation of a spinlock, and the spinlocks are used extensively all through the Linux kernel.

***Should we ship it?***

No. There is one big problem, and the big problem is that we have this thread just spinning, just busy waiting. This is wasteful, it burns CPU cycles, and it takes precious CPU time away from other threads that could be running.

It would be really nice if, when our thread did this compare-and-swap and it failed, ***rather than spinning, we could just put it away.*** If we could just put it to sleep and resume it when the value flag has changed and it can try again, that sounds nice. Lucky for us, the operating system gives us a way to do just that.

***How to do this to put our thread to sleep and wake up when needed:***

In Linux, you have these [futexes](https://man7.org/linux/man-pages/man2/futex.2.html), and futexes provide both an interface and a mechanism for user space, for programs to request the kernel to sleep and wake threads. The interface piece of this is the futex system call, and the mechanism is kernel-managed wait queues.

**How Futexes Work:**

[Lock on Linux - A Deep Dive](Lock%20on%20Linux%20-%20A%20Deep%20Dive%20eae787a11ae041f68e7d6876e2bd3167.md) 

Now we are going to extend our flag variable instead of having only two values. It will start storing three values:

The “flag” variable can be ***0 (free), 1 (locked), or 2 (waiter present).***

So how does our reader and writer processor function extend:

```go
func reader() {
 for {
  if atomic.CompareAndSwap(&flag, 0, 1) { //Set flag to 1 succeeded
   t := tasks.get()
   atomic.Store(&flag, 0) //Unset flag to 0
   return
  }
  // CAS failed, set flag to sleeping set flag to 2 (there’s a waiter)
  v := atomic.Xchg(&flag, 2)
  // and go to sleep.
  futex(&flag, FUTEX_WAIT, ...)
 }
}
```

Above code is performing below operations:

- It tries an atomic action to switch flag from 0 to 1.
- If successful perform the operation and store backs flag to 0 and then check if there is any waiters trigger that waiter.
- If unsuccessful (another thread set it to 1), it marks itself as a waiter (flag = 2).
- It then uses the futex system call to tell the kernel to put it to sleep until the flag changes.
- When woken, it retries the atomic action.

Switching to kernel land, let’s see what the kernel does. The kernel needs to do two things:

- The kernel stores information about the waiting thread — To store this information kernel maintains the user-space address. This is stored away in a wait queue entry that stores the thread and the key.
- The thread is put to sleep.
- When a thread finishes and sets the flag to 0, it uses a futex call to wake up waiting for threads.
- The kernel identifies the right waiting threads and schedules them to run.

How Kernel Maintains Waiting Queue

This is pretty nice, this is really convenient. This implementation we saw was an extremely simplified Futex implementation. What they give us is this nice, lightweight primitive to build synchronization constructs like locks. ***In fact, the pthread Mutex that you use in C and Java, they all use variants of this futex.***

At this point, it’s worth asking, ***if it makes sense to sleep rather than spin. Really, it comes down to an amortized cost argument.*** It makes sense to spin if the lock duration, the time for which we’re going to hold a lock is going to be short, but the trade-off is, while spinning, we’re not doing any useful work, we’re burning CPU cycles. At some point, it makes sense to pay the cost of that thread contact switch, to put the thread to sleep and run other threads so we can do useful work.

So, we need to build something which spins for some iteration before going to sleep and this is the same implementation provided by hybrid futexes.

***Hybrid futexes*** introduce a clever approach to thread synchronization. When a thread tries to lock using a compare-and-swap operation and fails, it doesn’t immediately go to sleep. Instead, it first spins a small number of times to see if the situation changes. Only if the spinning doesn’t result in acquiring the lock does the thread then resort to suspending itself using a futex-like mechanism. This approach combines the efficiency of spinning with the effectiveness of sleeping.

***At this point, we’ve evolved from spinlocks to futexes to hybrid futexes. Are we done now?***

Well, we could be. If we were talking about a language like Java or C++ that has a native threading model, we would be.

What about a language that uses user-space threads like Go with its goroutines? Can we do even better? Let’s see how locks/mutexes being implemented in golang.

## Golang Mutex Implementation

As discussed in the previous section we were blocking the kernel thread and putting them in to wait queue due to which OS has to put the extra cost of thread switching.

But Golang has goroutines which provide us with an opportunity to do something clever. This gives us an opportunity to block the goroutine rather than the underlying operating system thread. If a goroutine tries to acquire a Mutex and it can’t, we can totally put the goroutine to sleep without putting the underlying thread to sleep so we don’t have to pay that expensive thread-switching cost. This is clever and this is exactly what the Go runtime does.

Let’s see how this works. In our program, when the compare-and-swap fails, the Go runtime is going to add the goroutine to a wait queue, except, in this case, the wait queue is managed in user space. The wait queue looks very similar to the futex wait queue. We have a hash table and we have a wait queue per hash bucket. Once the goroutine has been added to the wait queue, the Go runtime reschedules the goroutine by calling into the Go scheduler. The goroutine is suspended, but the underlying thread keeps running, and the underlying thread just picks up other goroutines to run instead. When the writer comes along, it finishes its thing, The Go runtime walks the wait queue, finds the first goroutine to run that was waiting on the flag variable, and then reschedules it onto an operating system thread so the goroutine can run.

This is clever, we’ve found a way to avoid that heavy thread context switch cost. At this point, we could almost be done, but we’re not, and we’re not for a couple of reasons. There are a couple of problems with this implementation, and they both come down to this — they come down to the fact that once a goroutine is woken up, so a goroutine tries to acquire a lock, it failed, it was put on a wait queue. Once it’s resumed, it has to compete, it has to do that compare and swap, it has to compete with other incoming goroutines that are also trying to do the compare and swap. It’s totally possible it’s going to lose again, and it’s going to be put to sleep again. In fact, not only is it possible, it’s very likely this will happen because the other goroutines, the incoming goroutines, they’re already scheduled onto threads, they’re already running on the CPU, versus the goroutine that we just woke up, there’s a scheduling delay before it’s put onto a thread and run. It’s likely that it will lose.

To avoid the above problem goroutines operate in two modes:

1. Normal Mode
2. Starvation Mode

Let’s think a scenario in which two goroutines trying to acquire a lock:

```
func main() {
 var mutex sync.Mutex
 var data int

//Writer goroutines -- G1
 go func() {
  mutex.Lock()
  data = 42
  mutex.Unlock()
 }()

 // Reader goroutines -- G2
 for i := 0; i < 3; i++ {
  go func() {
   mutex.Lock()
   fmt.Println("Reader: Data read:", data)
   mutex.Unlock()
  }()
 }

 // Allow time for the goroutines to complete
 time.Sleep(time.Second)
}
```

**Normal Mode:**

In this mode, waiting goroutines form a line, like in a queue. However, when a waiting goroutine gets the signal to proceed, it doesn’t immediately own the lock. Instead, it competes with new goroutines trying to get the lock. But new goroutines have a head start as they are already on the CPU, so the waiting one often loses. To give waiting goroutines a chance, when they wake up, they’re moved to the front of the line. If a waiting goroutine fails to get the lock multiple times and it takes more than 1 millisecond, the system switches to a different mode.

**Lock acquiring in normal mode:**

**Lock release in normal mode:**

Once it releases the lock so at that time there could be some waiting goroutines and some goroutines are already trying to acquire lock. So, once lock gets released all goroutines compete each other for lock acquiring.

Now the problem with normal mode there is always competition between waiting goroutines and running goroutines for lock acquiring and waiting goroutines has a fair chance of losing. So our lock switches from normal mode to starvation mode after spinning goroutines for 1ms.

For this with the ***mutexLocked flag***, we are maintaining mode also ***mutexStarving*** which specify in which mode our mutex is running

**Starvation Mode:**

In this mode, the lock is directly given from the one releasing it to the waiting goroutine at the front of the line. New goroutines don’t spin or try to grab the lock. Instead, they patiently join the end of the line.

Lock acquired in starvation mode:

Park Goroutines in FIFO order in waiting queue.

Lock release in starvation mode:

While releasing the lock and we are operating in starvation mode we see for processing goroutine if it’s the last one in line or if it waited for less than 1 millisecond, the system goes back to Normal Mode.

***Why the Two Modes?***

Normal Mode is usually faster, allowing a goroutine to get the lock multiple times quickly. Starvation Mode helps avoid extreme cases where a goroutine is stuck waiting forever. So, the combination of these modes ensures that things stay both efficient and fair.

We’ve discussed how locks work and how other languages approach them. With Go, things get smarter. Go uses its own way of managing threads, called goroutines. This makes locks more efficient. Instead of blocking the whole thread, only the part that needs it waits, and the system switches different modes to make it more efficient and effective.

More information about mutex can be explored at the ***sync*** package in ***mutex.go*** file.

We haven’t covered the types of mutexes in Go and when to use them. If you’d like to delve into that topic, please comment below, and I’ll be happy to provide an explanation of the different types of mutexes in Go and the scenarios in which each type is most appropriate to use.