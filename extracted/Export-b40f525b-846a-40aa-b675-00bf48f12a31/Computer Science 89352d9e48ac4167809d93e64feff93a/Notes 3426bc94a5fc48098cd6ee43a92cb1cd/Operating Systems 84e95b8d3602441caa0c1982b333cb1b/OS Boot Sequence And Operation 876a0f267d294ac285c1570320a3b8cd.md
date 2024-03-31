# OS Boot Sequence And Operation

CPU should execute instructions from some start address (stored in Flash ROM). When your CPU starts some predefined memory address will be loaded into the program counter, and the CPU will start executing. Instead of regular memory, the memory at this location is generally memory mapped to some chip on the board (ROM). 

![Untitled](OS%20Boot%20Sequence%20And%20Operation%20876a0f267d294ac285c1570320a3b8cd/Untitled.png)

The code that executes is called the BIOS. It finds and initialises the storage device and loads the first sector(block of data). This sector will contain the bootloader. It’s responsible for loading the OS.

Then the OS will boot. It will load the drivers. And at the end, it goes into the init() function. 

![Untitled](OS%20Boot%20Sequence%20And%20Operation%20876a0f267d294ac285c1570320a3b8cd/Untitled%201.png)