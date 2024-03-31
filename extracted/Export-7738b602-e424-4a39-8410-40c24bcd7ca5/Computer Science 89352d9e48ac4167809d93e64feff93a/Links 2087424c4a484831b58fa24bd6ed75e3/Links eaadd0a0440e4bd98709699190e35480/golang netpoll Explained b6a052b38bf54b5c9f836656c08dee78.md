# golang netpoll Explained

URL: https://www.sobyte.net/post/2021-09/golang-netpoll/
Category: Golang, Network

![https://www.sobyte.net/android-chrome-192x192.png](https://www.sobyte.net/android-chrome-192x192.png)

---

Computer io models are distinguished into a variety of, currently the most used is also nio, epoll, select.

Combining different scenarios with different io models is the right solution.

## Network io in golang

golang is naturally suited for concurrency, why? One is the lightweight concurrency, and the other is the abstraction of complex io, which simplifies the process.

For example, if we simply access an http service, a few simple lines of code will do the trick:

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-0-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-0-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-0-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-0-4 | 

tr := &recordingTransport{}
client := &Client{Transport: tr}
url := "http://dummy.faketld/"
client.Get(url) // Note: doesn't hit network |
| --- | --- |

So what optimizations does golang make for Io? What about the ability to achieve such a simple switch?

## groutinue Scheduling for io events

We assume here that you already have some knowledge of groutinue scheduling.

We know that in go, each process is bound to a virtual machine, and in the machine, it has a g0 that traverses its own queue locally to get g or gets g from the global queue.

![golang%20netpoll%20Explained%20b6a052b38bf54b5c9f836656c08dee78/b6526812bc5647129a9639cb1f734e8e.png](golang%20netpoll%20Explained%20b6a052b38bf54b5c9f836656c08dee78/b6526812bc5647129a9639cb1f734e8e.png)

We also know that when g is running, g hands over execution to g0 for rescheduling, so how does g hand back events to g0 in io events? This is when it comes to our main character today —-netpoll.

## netpoll

The go language uses the I/O multiplexing model in the network poller to handle I/O operations, but it does not choose the most common system call, `select`. `select` can also provide I/O multiplexing capabilities, but there are more limitations to its use.

- Limited listening capability - can only listen to a maximum of 1024 file descriptors, which can be changed by manually modifying the limit, but at a relatively high cost in all respects.
- High memory copy overhead - a larger data structure needs to be maintained to store the file descriptors, which needs to be copied to the kernel.
- time complexity - after returning the number of ready events, all file descriptors need to be traversed.

golang officially encapsulates a network event poll in a unified way, independent of platform, providing a specific implementation for epoll/kqueue/port/AIX/Windows.

- `src/runtime/netpoll_epoll.go`
- `src/runtime/netpoll_kqueue.go`
- `src/runtime/netpoll_solaris.go`
- `src/runtime/netpoll_windows.go`
- `src/runtime/netpoll_aix.go`
- `src/runtime/netpoll_fake.go`

These modules, which implement the same functionality on different platforms, form a common tree structure. When the compiler compiles a Go language program, it selects specific branches in the tree for compilation based on the target platform

The methods that must be implemented are

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-1-8 | 

​netpollinit 初始化网络轮询器，通过 `sync.Once` 和 `netpollInited` 变量保证函数只会调用一次
​netpollopen 监听文件描述符上的边缘触发事件，创建事件并加入监听poll_runtime_pollOpen函数，这个函数将用户态协程的pollDesc信息写入到epoll所在的单独线程，从而实现用户态和内核态的关联。
​netpoll  轮询网络并返回一组已经准备就绪的 Goroutine，传入的参数会决定它的行为：
  - 如果参数小于0，阻塞等待文件就绪
  - 如果参数等于0，非阻塞轮询
  - 如果参数大于0，阻塞定期轮询
​netpollBreak 唤醒网络轮询器，例如：计时器向前修改时间时会通过该函数中断网络轮询器
​netpollIsPollDescriptor  判断文件描述符是否被轮询器使用 |
| --- | --- |

There are 2 important structures in netpoll.

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-8https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-9https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-10https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-11https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-12https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-13https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-14https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-15https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-16https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-17https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-18https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-19https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-20https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-21https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-22https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-23https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-24https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-25https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-26https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-27https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-28https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-29https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-30https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-31https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-32https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-33https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-34https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-35https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-36https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-2-37 | 

//pollCache  
//pollDesc
type pollDesc struct {
	link *pollDesc // in pollcache, protected by pollcache.lock
	// The lock protects pollOpen, pollSetDeadline, pollUnblock and deadlineimpl operations.
	// This fully covers seq, rt and wt variables. fd is constant throughout the PollDesc lifetime.
	// pollReset, pollWait, pollWaitCanceled and runtime·netpollready (IO readiness notification)
	// proceed w/o taking the lock. So closing, everr, rg, rd, wg and wd are manipulated
	// in a lock-free way by all operations.
	// NOTE(dvyukov): the following code uses uintptr to store *g (rg/wg),
	// that will blow up when GC starts moving objects.
	lock    mutex // protects the following fields
	fd      uintptr
	closing bool
	everr   bool      // marks event scanning error happened
	user    uint32    // user settable cookie
	rseq    uintptr   // protects from stale read timers
	rg      uintptr   // pdReady, pdWait, G waiting for read or nil
	rt      timer     // read deadline timer (set if rt.f != nil)
	rd      int64     // read deadline
	wseq    uintptr   // protects from stale write timers
	wg      uintptr   // pdReady, pdWait, G waiting for write or nil
	wt      timer     // write deadline timer
	wd      int64     // write deadline
	self    *pollDesc // storage for indirect interface. See (*pollDesc).makeArg.
}
type pollCache struct {
	lock  mutex
	first *pollDesc
	// PollDesc objects must be type-stable,
	// because we can get ready notification from epoll/kqueue
	// after the descriptor is closed/reused.
	// Stale notifications are detected using seq variable,
	// seq is incremented when deadlines are changed or descript |
| --- | --- |
- `rseq` and `wseq` - indicate that the file descriptor is reused or the timer is reset.
- `rg` and `wg` - indicate binary semaphores, possibly `pdReady`, `pdWait`, goroutine waiting for the file descriptor to become readable or writable, and `nil`.
- `rd` and `wd` - deadlines for waiting for file descriptors to become readable or writable.
- `rt` and `wt` - timers for waiting for file descriptors.

golang on io time to do a lot of unified encapsulation under runtime/netpoll (actually called under the internal/poll package), and then through the internal package to the runtime package to call, the internal package also encapsulates a pollDesc object of the same name, but is a pointer (there is a detail about the internal is that this package can not be called externally).

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-3-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-3-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-3-3 | 

type pollDesc struct {
	runtimeCtx uintptr
} |
| --- | --- |

In fact, it is ultimately a call under the runtime, but encapsulates some easy-to-use methods, such as read, write, and do some abstraction.

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-8https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-9https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-10https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-11https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-4-12 | 

func runtime_pollServerInit()  //初始化
func runtime_pollOpen(fd uintptr) (uintptr, int)  //打开
func runtime_pollClose(ctx uintptr)   //关闭
func runtime_pollWait(ctx uintptr, mode int) int //等待
func runtime_pollWaitCanceled(ctx uintptr, mode int) int  //等待并（失败时）退出
func runtime_pollReset(ctx uintptr, mode int) int  //重置状态,复用
func runtime_pollSetDeadline(ctx uintptr, d int64, mode int) //设置读/写超时时间
func runtime_pollUnblock(ctx uintptr)  // 解锁 
func runtime_isPollServerDescriptor(fd uintptr) bool  
// 这里的ctx实际上是一个io fd，不是上下文
// mod 是 r 或者 w  ,io事件毕竟只有有这两种
// d 意义和time.d差不多，就是关于时间的 |
| --- | --- |

The specific implementation of these methods is under runtime, so let’s pick a few important ones.

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-8https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-9https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-10https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-11https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-12https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-13https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-14https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-15https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-16https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-17https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-18https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-19https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-20https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-21https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-22https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-23https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-5-24 | 

//将就绪好得io事件，写入就绪的grotion对列
// netpollready is called by the platform-specific netpoll function.
// It declares that the fd associated with pd is ready for I/O.
// The toRun argument is used to build a list of goroutines to return
// from netpoll. The mode argument is 'r', 'w', or 'r'+'w' to indicate
// whether the fd is ready for reading or writing or both.
//
// This may run while the world is stopped, so write barriers are not allowed.
//go:nowritebarrier
func netpollready(toRun *gList, pd *pollDesc, mode int32) {
	var rg, wg *g
	if mode == 'r' || mode == 'r'+'w' {
		rg = netpollunblock(pd, 'r', true)
	}
	if mode == 'w' || mode == 'r'+'w' {
		wg = netpollunblock(pd, 'w', true)
	}
	if rg != nil {
		toRun.push(rg)
	}
	if wg != nil {
		toRun.push(wg)
	}
} |
| --- | --- |

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-8https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-9https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-10https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-11https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-12https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-13https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-14https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-15https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-16https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-17https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-18https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-19https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-20https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-21https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-22https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-23https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-24https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-25https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-26https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-27https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-28https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-29https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-30https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-31https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-32https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-33https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-34https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-35https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-36https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-37https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-38https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-39https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-40https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-41https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-42https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-43https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-44https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-45https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-46https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-47https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-48https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-49https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-50https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-51https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-52https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-53https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-54https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-55https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-56https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-57https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-58https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-59https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-60https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-61https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-62https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-63https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-64https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-65https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-66https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-67https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-6-68 | 

//轮询时调用的方法，如果io就绪了返回ok，如果没就绪，返回flase
// returns true if IO is ready, or false if timedout or closed
// waitio - wait only for completed IO, ignore errors
func netpollblock(pd *pollDesc, mode int32, waitio bool) bool {
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}
	// set the gpp semaphore to pdWait
	for {
		old := *gpp
		if old == pdReady {
			*gpp = 0
			return true
		}
		if old != 0 {
			throw("runtime: double wait")
		}
		if atomic.Casuintptr(gpp, 0, pdWait) {
			break
		}
	}
	// need to recheck error states after setting gpp to pdWait
	// this is necessary because runtime_pollUnblock/runtime_pollSetDeadline/deadlineimpl
	// do the opposite: store to closing/rd/wd, membarrier, load of rg/wg
	if waitio || netpollcheckerr(pd, mode) == 0 {
	  //gopark是很重要得一个方法，本质上是让出当前协程执行权，一般是返回到g0让g0重新调度
		gopark(netpollblockcommit, unsafe.Pointer(gpp), waitReasonIOWait, traceEvGoBlockNet, 5)
	}
	// be careful to not lose concurrent pdReady notification
	old := atomic.Xchguintptr(gpp, 0)
	if old > pdWait {
		throw("runtime: corrupted polldesc")
	}
	return old == pdReady
}
//获取到当前io所在的协程，如果协程已关闭，直接返回nil
func netpollunblock(pd *pollDesc, mode int32, ioready bool) *g {
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}
	for {
		old := *gpp
		if old == pdReady {
			return nil
		}
		if old == 0 && !ioready {
			// Only set pdReady for ioready. runtime_pollWait
			// will check for timeout/cancel before waiting.
			return nil
		}
		var new uintptr
		if ioready {
			new = pdReady
		}
		if atomic.Casuintptr(gpp, old, new) {
			if old == pdWait {
				old = 0
			}
			return (*g)(unsafe.Pointer(old))
		}
	}
} |
| --- | --- |

Think about.

1. a, b two co-processes, b io blocking, finished, has not obtained the scheduling rights, what will happen.
2. a, b two co-processes, b io blocking, 2s time out, but a has been occupying the execution rights, b has not obtained the scheduling rights, 5s before obtaining to, b to the use of the end has timed out, this time is time out or not time out

So set the time out, not necessarily the real io waiting, may not get the right to execute.

## How is the read event triggered?

Because write io is an active operation for us, how does read perform the operation? This is a passive state

First we understand a structure. golang identifies all network events and file reads and writes with fd (located under the internal package).

```go
// FD is a file descriptor. The net and os packages use this type as a
// field of a larger type representing a network connection or OS file.
type FD struct {
	// Lock sysfd and serialize access to Read and Write methods.
	fdmu fdMutex
	// System file descriptor. Immutable until Close.
	Sysfd int
	// I/O poller.
	pd pollDesc
	// Writev cache.
	iovecs *[]syscall.Iovec
	// Semaphore signaled when file is closed.
	csema uint32
	// Non-zero if this file has been set to blocking mode.
	isBlocking uint32
	// Whether this is a streaming descriptor, as opposed to a
	// packet-based descriptor like a UDP socket. Immutable.
	IsStream bool
	// Whether a zero byte read indicates EOF. This is false for a
	// message based socket connection.
	ZeroReadIsEOF bool
	// Whether this is a file rather than a network socket.
	isFile bool
}
```

We see that the pollDesc associated with the fd calls the various platform io events implemented inside the runtime package through the pollDesc.

When we do a read operation (here is the code capture)

```go
for {
	n, err := ignoringEINTRIO(syscall.Read, fd.Sysfd, p)
	if err != nil {
		n = 0
		if err == syscall.EAGAIN && fd.pd.pollable() {
			if err = fd.pd.waitRead(fd.isFile); err == nil {
				continue
			}
		}
	}
	err = fd.eofError(n, err)
	return n, err
	}
```

Will block the call waiteRead method, the method is mainly called inside the runtime_pollWait.

```go

func poll_runtime_pollWait(pd *pollDesc, mode int) int {
	errcode := netpollcheckerr(pd, int32(mode))
	if errcode != pollNoError {
		return errcode
	}
	// As for now only Solaris, illumos, and AIX use level-triggered IO.
	if GOOS == "solaris" || GOOS == "illumos" || GOOS == "aix" {
		netpollarm(pd, mode)
	}
	for !netpollblock(pd, int32(mode), false) {
		errcode = netpollcheckerr(pd, int32(mode))
		if errcode != pollNoError {
			return errcode
		}
		// Can happen if timeout has fired and unblocked us,
		// but before we had a chance to run, timeout has been reset.
		// Pretend it has not happened and retry.
	}
	return pollNoError
}
```

Here is mainly controlled by netpollblock,netpollblock method we have said above, when the io is not yet ready, directly release the current execution rights, otherwise it is already read and write the io event, directly read the operation can be.

## Summary

Overall flow listenStream -> bind&listen&init -> pollDesc.Init -> poll_runtime_pollOpen -> runtime.netpollopen - > epollctl(EPOLL_CTL_ADD)

Draw a diagram to understand more easily, of course, I was lazy to find the diagram

![golang%20netpoll%20Explained%20b6a052b38bf54b5c9f836656c08dee78/b068bfd61f21470c868696b75bf6f456.png](golang%20netpoll%20Explained%20b6a052b38bf54b5c9f836656c08dee78/b068bfd61f21470c868696b75bf6f456.png)

When g io events are encountered in golang, they are encapsulated in a unified way, first establishing a system event (this article focuses on epoll), then giving up the cpu (gopark), and then executing other g’s in a concurrent scheduling process. when the g io event is completed, it will interact with epoll to see if it is ready (epoll ready list), and if it is ready, pop will take out a g and execute it down the line. (Actually, there is some logic in the pop ready list, such as delayed processing)

runtime/proc.go.

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-8https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-9https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-10https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-11https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-12https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-13https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-14https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-15https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-16https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-17https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-18https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-19https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-20https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-21https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-22https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-23https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-24https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-25https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-26https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-27https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-28https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-29https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-30https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-31https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-10-32 | 

// Finds a runnable goroutine to execute.
// Tries to steal from other P's, get g from local or global queue, poll network.
func findrunnable() (gp *g, inheritTime bool) {
	_g_ := getg()
	// The conditions here and in handoffp must agree: if
	// findrunnable would return a G to run, handoffp must start
	// an M.
top:
	_p_ := _g_.m.p.ptr()
	//......
	// Poll network.
	// This netpoll is only an optimization before we resort to stealing.
	// We can safely skip it if there are no waiters or a thread is blocked
	// in netpoll already. If there is any kind of logical race with that
	// blocked thread (e.g. it has already returned from netpoll, but does
	// not set lastpoll yet), this thread will do blocking netpoll below
	// anyway.
	if netpollinited() && atomic.Load(&netpollWaiters) > 0 && atomic.Load64(&sched.lastpoll) != 0 {
		if list := netpoll(0); !list.empty() { // non-blocking
			gp := list.pop()
			injectglist(&list)
			casgstatus(gp, _Gwaiting, _Grunnable)
			if trace.enabled {
				traceGoUnpark(gp, 0)
			}
			return gp, false
		}
	}
	//......
} |
| --- | --- |

Also in sysmon, netpoll is scheduled.

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-8https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-9https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-10https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-11https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-12https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-13https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-14https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-15https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-16https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-17https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-18https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-19https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-20https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-21https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-22https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-23https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-24https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-25https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-26https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-27https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-28https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-11-29 | 

// Always runs without a P, so write barriers are not allowed.
//
//go:nowritebarrierrec
func sysmon() {
	lock(&sched.lock)
	sched.nmsys++
	checkdead()
	unlock(&sched.lock)
	//......
	// poll network if not polled for more than 10ms
	lastpoll := int64(atomic.Load64(&sched.lastpoll))
	if netpollinited() && lastpoll != 0 && lastpoll+10*1000*1000 < now {
		atomic.Cas64(&sched.lastpoll, uint64(lastpoll), uint64(now))
		list := netpoll(0) // non-blocking - returns list of goroutines
		if !list.empty() {
			// Need to decrement number of idle locked M's
			// (pretending that one more is running) before injectglist.
			// Otherwise it can lead to the following situation:
			// injectglist grabs all P's but before it starts M's to run the P's,
			// another M returns from syscall, finishes running its G,
			// observes that there is no work to do and no other running M's
			// and reports deadlock.
			incidlelocked(-1)
			injectglist(&list)
			incidlelocked(1)
		}
	}
	//......
} |
| --- | --- |

## Remarks

### epoll

epoll is a separate thread maintained by the system kernel, not by go itself

### Constants

FD_CLOEXEC is used to set the close-on-exec status criteria of the file. This, emm, is quite difficult to understand.

### gc

pollDesc is maintained by pollCache and is not monitored by GC (persistentalloc method allocation), so in the normal case about io operations, we must perform a manual shutdown to clean up the reference objects in epoll (specific implementation in poll_runtime_Semrelease).

| 

https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-1https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-2https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-3https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-4https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-5https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-6https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-7https://www.sobyte.net/post/2021-09/golang-netpoll/#hl-12-8 | 

// Must be in non-GC memory because can be referenced
// only from epoll/kqueue internals.
mem := persistentalloc(n*pdSize, 0, &memstats.other_sys)
for i := uintptr(0); i < n; i++ {
	pd := (*pollDesc)(add(mem, i*pdSize))
	pd.link = c.first
	c.first = pd
} |
| --- | --- |

### sysmon

Go’s standard library provides a thread to monitor your application and help you (find) any bottlenecks your application may encounter. This thread is called sysmon, the system monitor. In the GMP model, this (special) thread is not linked to any P, which means that the scheduler does not take it into account, and is therefore always running.

![golang%20netpoll%20Explained%20b6a052b38bf54b5c9f836656c08dee78/7ebb38990f984a3ba32ee88f3d43c0da.png](golang%20netpoll%20Explained%20b6a052b38bf54b5c9f836656c08dee78/7ebb38990f984a3ba32ee88f3d43c0da.png)

The sysmon thread has a wide range of roles, mainly in the following areas:

- Timers (timers) created by the application. The sysmon thread looks at timers that should be running but are still waiting for execution time. In this case, Go will look at the list of idle M and P timers so that it can run them as fast as possible.
- **Network pollers and system calls. It will run goroutines that are blocked during network operations.**
- Garbage collector (if it has not been running for a long time). If the garbage collector has not run for two minutes, sysmon will force a round of garbage collection (GC).
- Preemption of long-running goroutines. Any goroutine that runs for more than 10 milliseconds will be preempted, leaving the running time for other goroutines.