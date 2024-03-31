# Process vs Thread

## Process

- A set of instructions.
- Has an isolated memory.
- Has a PID
- Scheduled in the CPU
    - Suppose there is a 4-core machine and 10 processes. These processes are distributed among these processes.
    - When any of the processes tries to access value from memory, it will be kicked out. Whenever the process does something non-cpu related, it will get kicked out. eg : reading large file from disk. Why??? Because there other starving processes waiting for the CPU space for executing those instructions.

## Thread

- Light Weight Process (LWP)

![Untitled](Process%20vs%20Thread%20a19600731e2244f99d8e1d60317b14fc/Untitled.png)

- A set of instructions
- shares memory with parent process
    - Some competing happens between process and threads.
- has a PID
- scheduled in CPU

---

## How many process is too many?

- #core = #process 4 core - 4 process. Each process gets a dedicated core
- You can technically have any number of process
- lets take 100 processes, these processes are fighting for 4 core
    - So you to take in account of context switching
    - At a time 96 processes are waiting