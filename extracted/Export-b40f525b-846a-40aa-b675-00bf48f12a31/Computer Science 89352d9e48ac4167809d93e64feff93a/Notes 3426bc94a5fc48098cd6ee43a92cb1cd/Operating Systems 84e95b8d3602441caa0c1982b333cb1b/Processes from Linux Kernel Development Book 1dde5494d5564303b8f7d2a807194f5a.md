# Processes from Linux Kernel Development Book

[OSCA Academy](https://www.youtube.com/@oscaacademy5386)

[CPU Cores VS Threads Explained](https://www.youtube.com/watch?v=hwTYDQ0zZOw)

[what's the difference between processes, threads, and io multiplexing?](https://www.youtube.com/watch?v=85T_ZaT8EUI)

The kernel schedules individual threads, not processes. Processes can have multiple threads. 

Linux doesn't differentiate between threads and processes. To Linux, a thread is just a special kind of process.

***Another name for a process is a task. The Linux kernel internally refers to processes as tasks.***

***TCB - task_struct***

Linux uses kernel-level threading. 

![Untitled](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled.png)

---

The kernel stores the list of processes in a circular doubly linked list called the ***task list***. Each element in the task list is a process descriptor of type *`struct task_struct`*. 

`task_struct` is a large struct, 1.7 KB on a 32-bit machine. Considering this struct contains all the information that the kernel has and needs about a process, the size is justifiable. 

![Untitled](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%201.png)

## Allocating the Process Descriptor

The `task_struct` is allocated via the slab allocator to provide object reuse. Before the 2.6 kernel series, `task_struct` is stored at the end of the kernel stack of each process. 

Now the process descriptor is created dynamically via the slab allocator, `struct thread_info`, which also lives at the bottom of the stack(for stacks that grow down) and at the top of the stack(for stacks that grow up)

 

```c
struct thread_info {
			struct task_struct *task;
			struct exec_domain *exec_domain;
			__u32 flags;
			__u32 status;
			__u32 cpu;
			int preempt_count;
			mm_segment_t addr_limit;
			struct restart_block restart_block;
			void *sysenter_return;
			int uaccess_err;
};
```

![Untitled](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%202.png)

Each task’s `thread_info` structure is allocated at the end of its stack. The task element of the structure is a pointer to the task’s actual `task_struct`.

## Storing the Process Descriptor

Some architectures save a pointer to the `task_struct` of the currently running process in a register, enabling for efficient access. 

Other architectures, make use of the fact that `struct thread_info` is stored on the kernel stack to calculate the location of `thread_info` and subsequently the `task_struct`.

---

All processes in linux are descendants of the `init` process which has PID 1. Every process has exactly one parent. The relationship is stored in the task_struct, which contains a pointer to the parent of the process and a list of children named `children`.

![Getting the parent of a task](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%203.png)

Getting the parent of a task

![Iterating over the children of a task](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%204.png)

Iterating over the children of a task

## Process Creation

`fork()` and `exec()`. 

- The fork creates a child process that is a copy of the current task.
- The exec loads a new executable into the address space and begins executing it. It’s often used after fork when you want to execute a different program within the child process. It loads the new program from the file system and initializes the code, data, and stack segments.

# The Linux Implementation of Threads

On single-core machines threads enable concurrent programming and, on multiple processor systems, true parallelism.

Linux has a different approach when it comes to threads, Linux implements all threads as standard processes. The Linux kernel does not provide any special scheduling semantics or data structures to represent threads. Instead, a thread is merely a process that shares certain resources with other processes. 

Each thread has a unique `task_struct` and appears to the kernel as a normal process - threads just happen to share resources, such as an address space, with other processes. 

## Kernel Threads

It is useful for the kernel to perform some operations in the background. The main difference with a user-space thread is that a kernel space thread doesn't have an address space. 

Linux delegates several tasks to kernel threads, most notably the flush tasks and ksoftirqd task. 

Kernel threads are created on system boot by other kernel threads. A kernel thread can be created only by another kernel thread.

```bash
ps -ef // lists all the kernel threads
```

The kernel handles this by forking all new kernel threads off of the kthreadd kernel process.

---

---

# Process Scheduling

***Preemptive Scheduling***: The scheduler decides when a process is to stop running and a new process is to begin running. The act of involuntarily suspending a running process is called *preemption.* Runs for a timeslice, it prevents a process from monopolizing the processor. In modern systems, the time slice is calculated dynamically.

**Cooperative Scheduling:** The process doesn't stop running until it voluntarily decides to do so.- The process needs to decide when to stop. The scheduler doesn't make the decisions here. Processes can monopolize the processor.

Most of the OSs use ***Preemptive Scheduling*** for obvious reasons. 

---

[https://medium.com/geekculture/linux-cpu-context-switch-deep-dive-764bfdae4f01#:~:text=“CPU context switching” is to,to by the program counter](https://medium.com/geekculture/linux-cpu-context-switch-deep-dive-764bfdae4f01#:~:text=%E2%80%9CCPU%20context%20switching%E2%80%9D%20is%20to,to%20by%20the%20program%20counter).

![Untitled](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%205.png)

![Untitled](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%206.png)

![Untitled](Processes%20from%20Linux%20Kernel%20Development%20Book%201dde5494d5564303b8f7d2a807194f5a/Untitled%207.png)

---