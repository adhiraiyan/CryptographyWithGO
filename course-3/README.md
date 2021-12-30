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

- Creating a goroutine:
    - One goroutine is created automatically to execute the `main()`.
    - Other goroutines are created using the `go` keyword.

        ```go
        a = 1
        go foo()  // New go routine crated for foo(), main goroutine doesn't block
        a = 2
        ```
    
- A goroutine exits when its code is complete. When the main goroutine is complete, all other goroutines exit. A goroutine may not complete its execution because main completes early.

- In the code below, we expect both to be printed, we just don't know which ones' going to get printed first. But the below code only prints "main routine" because the main finishes before the new goroutine.

    ```go
    func main() {
        go fmt.Printf("New Routine")
        fmt.Printf("Main routine")
    }
    ```

- This is unwanted behavior, regardless of the order, we want both to be printed, so what we can do is add a delay in the main routine, but this isn't what you should do, because we are making assumptions about the timing (the delay of 100ms) would be enough to execute the new goroutine, about the go scheduler and the OS scheduler. What we need is formal synchronization constructs.

    ```go
    func main() {
        go fmt.Printf("New Routine")
        time.Sleep(100 * time.Millisecond)  // add a delay in the main routine
        fmt.Printf("Main routine")
    }
    ```

- Basic Synchronization is when multiple threads agree on the timing of the event by using global events. Synchronization is used to restrict bad interleavings.

- The Synch package contains functions to synchronize between goroutines. The `synch.WaitGroup` forces a goroutine to wait for other goroutines. It contains an internal counter that increments or decrements for each goroutine added or completed and the waiting goroutine can't complete until the counter reaches 0.

    ```go
    func foo(wg *sync.WaitGroup) {
        fmt.Printf("New Routine")
        wg.Done()  // Signal to the wait group that the task is complete
    }

    func main() {
        var wg sync.WaitGroup
        wg.Add(1)  // Create a WaitGroup of 1
        go foo(&wg)  // Call the concurrent goroutine
        wg.Wait()  // Ask the WaitGroup to wait till the task completes
        fmt.Printf("Main Routine")  // executes the code after wait receives done
    }
    ```

- Goroutines usually work together to perform a bigger task, they are not completely independent, they often need to send data to collaborate. For example lets say you want to find the product of 4 integers, you can accomplish this by making 2 goroutines each multiplying a pair and then the main goroutine multiples the 2 results. In order to do this, we need to send data from the main routine to the two subroutines. Then the results need to come back to the main routine. This is a basic and simplest type of communication that share data at the start and the end, but this doesn't have to be the only way, you can have the data be sent in the middle of the routine. 

- Communication, transfer of data between routines are done using channels. Channels are typed. We can use `make()` to create a channel: `c := make(chan int)`. We can then send and receive data using the `<-` operator. To send data on a channel: `c <- 3` and to receive data from a channel: `x := <- c`.

    ```go
    func prod(v1 int, v2 int, c chan int) {
        c <- v1 * v2
    }

    func main() {
        c := make(chan int)
        go prod(1, 2, c)
        go prod(3, 4, c)

        a := <- c
        b := <- c
        fmt.Println(a * b)
    }
    ```

- Unbuffered channels cannot hold data in transit and by default it is unbuffered. What this means is sending data across a channel is blocked until data is received and receiving is blocked until data is sent. Since a task waits for one to complete, channels also does synchronization, this synchronization is built in. Blocking is the same as waiting for communication.

- Channels can contain a limited number of objects, the default size of a buffered channel is zero. But you can make an buffered channel with some size. The capacity is the number of objects it can hold in transit. You can define this as an optional argument to make: `c := make(chan int, 3)`. It will then only do blocks if buffer is full while sending or empty while receiving. Blocking is generally a bad thing since it reduces concurrency.

- The main use of buffering is that the sender and receiver don't need to operate at exactly the same speed. Incase of over producing or consuming, the buffer can be a place where things can just get filled up, a speed mismatch is acceptable.


--- 

## Week 4: Synchronized Communication




--- 

## My Notes

### Parallel Bubble Sort

- The sorting methods can be divided into two classes by the complexity of the algorithms used. The complexity of a sorting algorithm is generally written in the big-O notation and it is expressed based on the size of sets the algorithm is run against.

- The two classes of sorting algorithms are O(n^2), which includes the bubble, insertion, selection, shell sorts and O(n log n), which includes the heap, merge, quick sorts. 

- The sequential version of the bubble sort is the oldest, the simplest and the slowest sorting algorithm in use having a complexity level of O(n^2).

- The parallel version of bubble sort can be obtained if we use the odd-even transposition method  that implies the existence of n phases, each requiring n/2 compare and exchanges. In the first phase, called odd phase, the elements having odd indexes are compared with the neighbors from the right and the values are swapped when necessary. In the even phase, the elements having even indexes are compared with the elements from the right and the exchanges are performed only if necessary.


### Parallel Merge Sort in Go

- Psuedocode:

    ```
    mergesort (s []int) {
        if len(s) > 1 {
            middle := len(s) / 2

            // Create two sub-tasks
            createTask -> mergesort(s[:middle])
            createTask -> mergesort([middle:])

            merge(s, middle)
        }
    } 
    ```

- Two main strategies for context switching: 
    - Work Sharing: When a processor generates new threads, it attempts to migrate some of them to the other processor with the hopes of them being utilized by the idle/underutilized processors.
    - Work-stealing: An underutilized processor actively looks for other processor's threads and 'steal' some.

- GO's scheduler is based on the second strategy. The main benefit of this strategy is to reduce the frequency of goroutines migration (from one thread/CPU core to another). The migration is only considered when a thread is idle (no more goroutines to handle or every goroutine is blocked due to an I/O call for example) and this is going to have a positive impact performance-wise.




---
## Definitions

- Hiding latency: is executing a task while another task is waiting.

- A process is an instance of a running program.

- Context switch: control flow changes from one process to another.