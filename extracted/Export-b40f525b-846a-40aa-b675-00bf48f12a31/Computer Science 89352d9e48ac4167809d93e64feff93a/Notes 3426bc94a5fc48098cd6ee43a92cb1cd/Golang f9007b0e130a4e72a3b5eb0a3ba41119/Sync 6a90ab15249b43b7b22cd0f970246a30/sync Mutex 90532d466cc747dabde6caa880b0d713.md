# sync.Mutex

Using this lock an exclusive right over a critical section can be acquired. When another goroutine tries to access the lock, it won't be able to acquire the lock. That goroutine can wait for the lock to be released(active waiting) or go somewhere and come back later(passive waiting).?????

The former is called “spinning”, they hold the CPU resource, increasing their chance of acquiring the lock when it becomes available, without the overhead of expensive context switching. 

However, when the mutex is not likely to become available anytime soon, continuing to hold CPU resources will diminish the chances of other goroutines getting their share of CPU time.

![Untitled](sync%20Mutex%2090532d466cc747dabde6caa880b0d713/Untitled.png)

Golang allows arriving goroutine to spin for a while. If it cannot acquire the mutex within a specified time slice, it will sleep to give other goroutines a fair chance to run.

![Untitled](sync%20Mutex%2090532d466cc747dabde6caa880b0d713/Untitled%201.png)

## Mutex Fairness

[](https://github.com/golang/go/blob/go1.21.0/src/sync/mutex.go#L74C5-L74C80)

Now the new guy has a big chance of acquiring the lock, if the lock is released before the new guy is put to sleep.

This is called normal mode.

If the new guy fails to acquire the lock within a time slice, it is pushed to the waiting queue. Then if a new guy comes he now has more advantage in acquiring the lock than the guys in the queue.

This introduces an unfairness for goroutines for accessing the mutex.

If a waiter at the front of the wait queue fails to acquire the lock for more than 1 ms, the run time switches the mutex to starvation mode.

In starvation mode ownership of the mutex is directly handed off from the unlocking goroutine to the waiter at the front of the queue. New arriving goroutines don't try to acquire the mutex even if it appears to be unlocked, and don't try to spin. Instead, they queue themselves at the tail of the wait queue.

If a waiter receives ownership of the mutex and sees that either 

- It is the last waiter in the queue
- It waited for less than 1 ms

It switches the mutex back to normal mode. 

[Go Concurrency Visually Explained – sync.Mutex](https://blog.stackademic.com/go-concurrency-visually-explained-sync-mutex-a72fe533b8f4)