# sync.RWMutex

[Go Concurrency Visually Explained – sync.RWMutex](https://blog.stackademic.com/go-concurrency-visually-explained-sync-rwmutex-622fc04278cb)

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled.png)

Partier comes and acquires an exclusive lock on TV. 

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%201.png)

Now another guy called Stringer comes along, and he wants exclusive access to the TV using `tv.Lock()`, but Partier already have an exclusive lock on the TV, so Stringer has to wait till Partier unlocks the exclusive lock. So after Partier finishes the critical section, he releases the lock. Now Stringer can move from the queue and get exclusive access over the TV.

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%202.png)

Stringer has an exclusive lock over the TV, now Candier wants to watch the TV with Stringer (he doesn’t want to change the channel, i.e. change the state of the TV, he just wants to watch it, i.e. Read, using an RLock). But since Stringer, has an “exclusive” access over the TV, he doesn't want Candier to watch TV with him. So Candier is put to wait till Stringer gives up his “exclusive” access.

Similarly, Swimmer also wants to watch TV with Stringer, and you get the same situation as above, Swimmer is also asked to wait in the queue.

They will need to wait till Stringer gives up his exclusive access to the TV.

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%203.png)

Stringer gives up his exclusive access, and now all the guys in the queue can RLock on the TV. Now there is no exclusive access, there is only RLock.

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%204.png)

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%205.png)

Now a new guy comes along, Xavior. He wants exclusive access to the TV(He wants to change the TV). Since Swimmer, Candier and Stringer are watching the TV, Xavior won’t get exclusive access to the TV. He will be put in the queue. And he has to wait till all guys unlock their rLock on the TV.

Another interesting thing to note here is if another person tries to get a read lock, i.e. he wants to watch the TV with them, in the earlier situation, he could join with the boys. Now he cannot, because there is already an exclusive lock in the queue.

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%206.png)

After some time, the guys got tired, and they left the room(i.e. unlocked the read locks). Now Xavior can get an exclusive lock over the TV. And Guy has to wait till Xavior is done with the exclusive lock.

![Untitled](sync%20RWMutex%203bf14bc6d8bc4567a08df845d393e069/Untitled%207.png)

## Key takeaways

- When there is an exclusive lock over a critical section, all the read locks are queued.
- When there are read locks on a critical section, all the exclusive locks are queued. The read locks after an exclusive locks are also queued.