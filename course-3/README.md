# Concurrency in GO

--- 

## Week 1: Why Use Concurrency?

- Concurrency is built into the language in GO.

- Parallel Execution is when two programs execute at exactly the same time. Generally, one core runs one instruction at a time, so to have two instructions or more run at the same time, you either need multi-core processors or multiple CPUs.

- Von Neumann Bottleneck: the time the CPU takes to read and write to memory. Cache, on memory chip was a solution to this.

- Power Wall: Increasing the density of transistors increased the power use and high power use increases temperature.

- Dennard Scaling: Voltage should scale with transistor size. This would keep power consumption and temperature low, but the problem is voltage can't go too low because it must stay above threshold voltage and there is leakage power.

- Concurrent Execution is not necessarily execution as same as parallel execution. Concurrent start and end times overlap, whereas parallel execute at exactly the same time. So why use concurrent execution? Parallel tasks must be executed on different hardware, whereas concurrent tasks may be executed on the same hardware since only one task actually executes at a time, and mapping from tasks to hardware is not directly controlled by the programmer in parallel execution, atleast not in GO.

- In concurrent programming, programmer determines which tasks can be executed in parallel and what will be executed is determined by the operating system, mapping to hardware and the go runtime scheduler.

- Concurrency improves performance even without parallelism, multiple processors or cores.


--- 

## Week 2: Concurrency basics 

- Operating system schedules processes for execution. It gives the illusion of parallel execution.

- One downside of processes is that the context switching time may be long since it takes time to read data from and write to memory. To overcome this, a thread was created, threads share some context and many threads can exist in one process. OS schedules threads rather than processes.

- Goroutines are like a thread in GO. Many Goroutines execute withing a single OS thread. The process of scheduling goroutines inside an OS thread is done by the GO runtime scheduler. GO runtime scheduler is like a little OS inside a single OS thread. Since all the goroutines are running inside a thread, unless you create a logical process (is mapped to a thread) you can't achieve true parallelism.

- With concurrent code, its difficult to debug since the overall state of the machine is not deterministic.

- Order of execution between tasks isn't known, meaning the instructions can be interleaved in different ways. To make sure the code works as it should, we must consider all possibilities of interleavings.

- Race conditions is a problem where the outcome of the program depends on interleavings.Since interleavings are non-deterministic, the outcome will become non-deterministic. Races occur due to communication between tasks, for example say you have two tasks that needs to communicate information about a shared variable and say one writes or reads before the other should and communicate the results at different times, this will lead to different outcomes on different executions.

- Threads are largely independent, they mostly don't have to communicate with each other.


--- 

## Week 3: Threads in GO




--- 

## Week 4: Synchronized Communication




--- 

## My Notes





---
## Definitions

- Hiding latency: is executing a task while another task is waiting.

- A process is an instance of a running program.

- Context switch: control flow changes from one process to another.