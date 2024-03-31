# Concurrency in Go: Katherine Cox Buday

**Amdahl’s law** describes a way in which to model the potential performance gains from implementing the solution to a problem in a parallel manner. Simply put, it states that the gains are bounded by how much of the program must be written in a sequential manner.

[[Amdahl’s Law](https://learn.saylor.org/mod/page/view.php?id=27029)](../../System%20Design%20and%20Architecture%2006b337409f31496590be5398016faab0/Performance%20f021e38997ed4ecd8e8a971da093db09/Parallel%20Request%207bd9d477ce644b86973fd832fd5bc6f0.md) 

For example, imagine you were writing a program that was largely GUI based: a user is presented with an interface, clicks on some buttons, and stuff happens. This type of program is bounded by one very large sequential portion of the pipeline: human interaction. No matter how many cores you make available to this program, it will always be bounded by how quickly the user can interact with the interface.

### Amdahl’s law (Chatgpt)

Amdahl's Law is a fundamental principle in computer architecture and parallel computing that quantifies the potential speedup of a program when you parallelize it. It was formulated by Gene Amdahl in 1967 and is often used to illustrate the limits of how much a program can benefit from parallel processing. Amdahl's Law can be expressed mathematically as follows:

![Untitled](Concurrency%20in%20Go%20Katherine%20Cox%20Buday%2003ab336e56e14e89be33df132c7561b9/Untitled.png)

Where:

- **Speedup** represents the improvement in execution time or performance when running a program in parallel compared to running it sequentially. It tells you how much faster the parallel version is.
- **Fraction\_Sequential** is the portion of the program that must be executed sequentially or cannot be parallelized. It is expressed as a decimal between 0 and 1. This part of the program cannot benefit from multiple processors.
- **Number\_of\_Processors** is the number of processors or computing units available for parallel execution.

The key insight from Amdahl's Law is that the ***speedup of a program is fundamentally limited by the fraction of the program that cannot be parallelized.*** No matter how many processors you add, if a significant portion of the program is inherently sequential, there will be a limit to how much you can speed up the entire program.

For example, if 90% of a program can be parallelized (Fraction\_Sequential = 0.1), Amdahl's Law shows that even with an infinite number of processors, the maximum speedup you can achieve is 10x (1 / 0.1). In practice, it's rare to have a program that can be perfectly parallelized, so the actual speedup achieved is often less than this theoretical maximum.

***Amdahl's Law is crucial for making decisions about when and how to parallelize programs.*** It highlights the importance of identifying and optimizing the sequential parts of a program to achieve meaningful speedup when parallelizing. It also emphasizes that the returns on parallelization diminish as you add more processors, especially if a large portion of the program remains sequential. Therefore, it's essential to strike a balance between the effort required to parallelize a program and the potential speedup that can be achieved.

[Fundamentals of concurrency](Concurrency%20in%20Go%20Katherine%20Cox%20Buday%2003ab336e56e14e89be33df132c7561b9/Fundamentals%20of%20concurrency%20098dc77b7bde496ca435e2d31e90a9ca.md)

[Communicating Sequential Process](Concurrency%20in%20Go%20Katherine%20Cox%20Buday%2003ab336e56e14e89be33df132c7561b9/Communicating%20Sequential%20Process%2007addbdd08804bb381e88da917fec539.md)