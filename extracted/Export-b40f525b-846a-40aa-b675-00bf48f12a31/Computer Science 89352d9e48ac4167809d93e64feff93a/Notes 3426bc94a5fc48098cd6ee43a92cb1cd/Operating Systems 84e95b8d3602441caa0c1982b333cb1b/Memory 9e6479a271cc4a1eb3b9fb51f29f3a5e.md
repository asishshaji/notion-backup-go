# Memory

RAM = static (SRAM) and dynamic (DRAM)

SRAM is faster and more expensive. DRAM is used for the main memory plus the frame buffer of a graphics system. SRAM stores each bit in a bistable memory cell. Each cell is implemented with a [6-transistor circuit](https://www.iue.tuwien.ac.at/phd/entner/node34.html). This circuit can stay indefinitely in either of two voltage configurations, or states as long as power is supplied. This is in contrast to DRAM where periodic refreshes are necessary.

SRAM provides higher performance benefits because SRAM does not require periodic refreshments like DRAM. Due to the bistable nature of SRAM. Even when a disturbance, such as electrical noise, perturbs the voltages, the circuit will return to the stable value when the disturbance is removed.

DRAM = Dynamic RAM

DRAM cell consist of a single capactor and a transistor which stores one bit as charge(30 * 10^-15 farads). Unlike SRAM, DRAM is very senstive to disturbances. When the capacitor voltage is disturbed, it will never recover. Exposure to light rays will change the voltages of the capacitor. Sensors of digital cameras are essentially an array of DRAM cells.

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled.png)

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%201.png)

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%202.png)

The wordline is connected to the transistors gate and bitline is connected to the channel of the transistor.

Applying a voltage to the wordline turns on the transistor, while it’s on electron can flow through the channel, thus connecting the capacitor to the bitline. This can allow us to either charge or discharge the capacitor.

When the voltage applied to wordline is disconnected, the gate is closed, thus the capcitor can store the data.

One thing to note is the transistor is very small, and thus electrons will escape through the transistor. That’s why DRAM requires refresh periodically.

Various current leakages cause the capacitor to lose its memory(10 to 100 ms, this retention time is actually good), so the memory system must periodically refresh every bit by reading it and rewriting it.

In conclusion SRAM will persist the memory, DRAM needs to refresh the memory periodically.

![image-1696132593125.jpg8138106242047248259.jpg](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/image-1696132593125.jpg8138106242047248259.jpg)

## Conventional DRAMs

The cells in a DRAM chip are partitioned into d supercells, each with the ability to store w bits. So DRAM stores d * w bits of information. The supercells are organised into 2-d array with r rows and c columns, (r,c) = d. Each supercell has an address of the form (i,j).

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%203.png)

Here d = 16, w = 8

There total of 16 supercells, each supercell have a capacity of 8 bits.

Information from and to the chip flows through the 2 pins, one 2 bit and the other 8 bit. 2 bit pin is used for sending the address of the supercell the memory controller wants to access to the DRAM. 8 bit pin is used to read and write data into the supercell.

The memory controller first sends the row number (row access strobe) to the chip, the chip then copies the contents of that row to the interal row buffer. Then the mem controller sends the column number (column access strobe). DRAM responds by copying the supercell data of supercell (i,j) from the internal row buffer and sends to the memory controller.

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%204.png)

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%205.png)

One reason to keep DRAMs as 2-d array is to reduce the number of address pins. If our 128 bit DRAM was organized as linear array, addresses from 0 - 15, we will need 4 pins.

Only disadvantage is the access time is increased, you send i, then a copy to the internal row buffer, then you send the column number and then atlast you get the data.

[How does Computer Memory Work? 💻🛠](https://www.youtube.com/watch?v=7J7X7aZvMXQ)

## Memory Modules

DRAM chips are packaged into memory modules, they are plugged into memory slots on the motherboard. A stick of a RAM is called Dual Inline Memory Module (DIMM).

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%206.png)

## Memory Read Transaction

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%207.png)

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%208.png)

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%209.png)

IO bridge acts as an interface between system bus and memory bus. It converts system bus electrical signals into memory bus electrical signals.

I/O bridge is also connected to various I/O devices.

```nasm
movq A, %rax
```

Circuitry on the CPU chip called bus interface initiates a read transaction on the bus. First the CPU places the address A on the system bus. The I/O bridge passes the signal along to the memory bus. Main memory reads the address A from the bus, retrivies the data from the DRAM, and writes the data into the memory bus.

## Locality

Programs with good locality tend to reference data items that are near other recently referenced data items or that were recently referenced themselves. 

*Temporal Locality:* memory location that is referenced once is likely to be referenced again multiple times in the near future.

*Spatial Locality:* If a memory location is referenced once, then the program is likely to reference a nearby memory location in the near by future.

Programs with good locality run faster than programs with poor locality.

- Programs that repeatedly reference the same variables enjoy good temporal locality.
- For programs with stride-k reference patterns, the smaller the stride, the better the spatial locality. Programs with stride-1 reference patterns have good spatial locality. Programs that hop around memory with large strides have poor spatial locality.

![Untitled](Memory%209e6479a271cc4a1eb3b9fb51f29f3a5e/Untitled%2010.png)

- Loops have good temporal and spatial locality with respect to instruction feteches. The smaller the loop body and the greater the number of loop iterations, the better the locality.