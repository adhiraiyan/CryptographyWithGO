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




--- 

## Week 3: Threads in GO




--- 

## Week 4: Synchronized Communication




--- 

## My Notes





---
## Definitions

- Hiding latency: is executing a task while another task is waiting.