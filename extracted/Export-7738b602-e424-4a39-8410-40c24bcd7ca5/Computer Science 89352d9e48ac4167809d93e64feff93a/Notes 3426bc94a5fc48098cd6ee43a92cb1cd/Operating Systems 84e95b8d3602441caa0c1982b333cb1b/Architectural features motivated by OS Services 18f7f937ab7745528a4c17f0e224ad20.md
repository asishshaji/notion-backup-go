# Architectural features motivated by OS Services

## Protection

There are two types of assembly instructions regular `MOV, ADD` & sensitive or privileged, which are only executed by the kernel `HLT`. 

Regular users are restricted to not 

- address I/O directly
- use instructions that manipulate the state of memory `page table pointers, TLB...`
- set mode bits that determines user or kernel mode
- disable and enable interrupts
- halt machine

In kernel mode, OS can do all the items mentioned above. There is a bit that identifies mode of execution `user or kernel mode`. It’s status bit in protected processor register. This mode bit is stored in a register called **Program Status Word (PSW)** register. 

[The Operating System](https://medium.com/swlh/the-operating-system-d2fe86842971)

Apart from these two privilege levels, there are also some intermediate levels `ring privilege`;these comes handy in virtualisation. 

[What Are Rings in Operating Systems? | Baeldung on Computer Science](https://www.baeldung.com/cs/os-rings)

![Untitled](Architectural%20features%20motivated%20by%20OS%20Services%2018f7f937ab7745528a4c17f0e224ad20/Untitled.png)

## Crossing Protection Boundaries

System call is an OS procedure call that executes privileged instructions, eg : I/O  

[Crossing Protection Boundaries](https://courses.cs.washington.edu/courses/cse451/98sp/Lectures/2-arch/tsld006.htm)

- When a system call is triggered; it issues a trap, which asks the kernel to take over
- Each trap has a trap handler registered.
- The mode in PSW is set to 0; indicating the execution is now in kernel mode
- The stat of the user program is saved so it can restore control to user process after completing the system call.

![Untitled](Architectural%20features%20motivated%20by%20OS%20Services%2018f7f937ab7745528a4c17f0e224ad20/Untitled%201.png)

Kernel uses the IDT table to define the different functions to be executed when an interrupt happens. The Interrupt Descriptor Table (IDT) is a data structure used by the x86 architecture to implement an interrupt vector table. The IDT is used by the processor to determine the correct response to interrupts and exceptions. 

The **interrupt** **vector table**, often abbreviated to **IVT** or simply **IV**, is an array of pointers to functions, associated by the CPU to handle specific *exceptions*, such as faults, system service requests from the application, and interrupt requests from peripherals.

[IDT and interrupts](https://samypesse.gitbook.io/how-to-create-an-operating-system/chapter-7)

[Interrupt Vector Table - OSDev Wiki](https://wiki.osdev.org/Interrupt_Vector_Table)

### The keyboard example

When the user pressed a key on the keyboard, the keyboard controller will signal an interrupt to the Interrupt Controller. If the interrupt is not masked, the controller will signal the interrupt to the processor, the processor will execute a routine to manage the interrupt (key pressed or key released), this routine could, for example, get the pressed key from the keyboard controller and print the key to the screen. Once the character processing routine is completed, the interrupted job can be resumed.

The [PIC](http://en.wikipedia.org/wiki/Programmable_Interrupt_Controller) (Programmable interrupt controller)is a device that is used to combine several sources of interrupt onto one or more CPU lines, while allowing priority levels to be assigned to its interrupt outputs. When the device has multiple interrupt outputs to assert, it asserts them in the order of their relative priority

[Interrupts — The Linux Kernel  documentation](https://linux-kernel-labs.github.io/refs/heads/master/lectures/interrupts.html)

## Memory Protection

A simple technique is to use a base and limit registers, which marks the starting and end of the memory allocated for a process.

![Untitled](Architectural%20features%20motivated%20by%20OS%20Services%2018f7f937ab7745528a4c17f0e224ad20/Untitled%202.png)

![https://www.computersciencejunction.in/wp-content/uploads/2017/04/memory_protection.png](https://www.computersciencejunction.in/wp-content/uploads/2017/04/memory_protection.png)

[Memory Protection in Operating Systems - GeeksforGeeks](https://www.geeksforgeeks.org/memory-protection-in-operating-systems/)

CPU checks each user reference (instruction and data address); ensuring it fall between the base and limit register values.