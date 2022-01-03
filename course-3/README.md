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

- Common operation of a channel is to iteratively read from a channel:
    
    ```go
    for i := range c {  // one iteration each time will read a new value its received
        fmt.Println(i)  // this will continue until a close happens
    }
    ```

- If you read `c` from a `range` construct, the sender needs to `close(c)` the channel so the receiver knows to quit the loop.

- Reading from multiple channels received from multiple goroutines. Multiple channels may be used to receive from multiple sources. For example lets say there are two channels and there is a `func` that needs data from both the channels to do a product calculation:

    ```go
    a := <- c1
    b := <- c2
    fmt.Println(a*b)
    ```

- Sometimes you may want to read from just one channel, you may have a choice of which data to use. For example first come first served. If say, we don't know which channel the data is going to come from, we can use a select statement so we aren't waiting for data from a channel that isn't there yet. `select` allows you to wait for a data from a set of channels.

    ```go
    select {
        case a = <- c1:
            fmt.Println(a)
        case b = <- c2:
            fmt.Println(b)
    }
    ```

- By using select, we can both block on sending and receiving data to/from a channel:

    ```go
    select {
        case a = <- inchan:
            fmt.Println("received a")
        case outchan <- b:
            fmt.Println("sent b")
    }
    ```

- One common use of a select is to have a separate abort channel.

    ```go
    for {
        select {
            case a = <- c:
                fmt.Println(a)
            case <- abort:  // we aren't focusing on what data is coming in cause its from abort
                return
        }
    }
    ```

- We may also want a default operation to avoid blocking

    ```go
    select {
        case a = <- c1:
            fmt.Println(a)
        case b = <- c2:
            fmt.Println(b)
        default:  // this will not wait to get data from the two channels
            fmt.Println("nope")
    }
    ```

- Sharing variables concurrently can cause problems. Two goroutines writing to the same shared variable can interfere with each other. A thread program is said to be concurrency safe if it can be executed or invoked concurrently with other goroutines. Below is an example of variable sharing. Here, two goroutines write to `i`. `i` should be equal to 2 but this doesn't always happen. This is because of possible interleavings. Initial it may seem like there is no problem and it may not change the value of `i`. 

    ```go
    var i int = 0
    var wg sync.WaitGroup

    func inc() {
        i = i+1
        wg.Done()
    }

    func main() {
        wg.Add(2)
        go inc()
        go inc()
        wg.Wait()
        fmt.Println(i)
    }
    ```

- Concurrency is at the machine code level and interleavings happen here. The source code instructions can partition or chop up things right in the middle. `i = i + 1` might be three machine instructions (read i, increment, write i). This can be where things mess up. Interleaving machine instructions causes unexpected problems.

- Correct Sharing:
    - Don't let 2 goroutines write to a shared variable at the same time.
    - Need to restrict possible interleavings.
    - Access to shared variables cannot be interleaved.
    - This can be done using _mutual exclusion_. Code segments in different goroutines which can't execute concurrently. Writing to shared variables should be mutually exclusive.

- `sync.mutex` ensures mutual exclusion using a binary semaphore. If the flag is up, the shared variable is in use and you aren't allowed to use it. The flag being down means the shared variable is available and you can put the flag up and use the variable.

- `sync.mutex` methods:
    - `Lock()` method puts the flag up (shared variable in use). If another goroutine tries to call lock, then the lock will block the second goroutine until the flag is put down.
    - `Unlock()` gets called when we are done using shared variable. This method puts the flag down. So one of the waiting routines can start using the variable. If the programmer doesn't specify the unlock, a deadlock may occur. Below is an example on how we can modify the incrementing problem:

        ```go
        var i int = 0
        var mut sync.Mutex

        func inc() {
            mut.Lock()
            i = i + 1
            mut.Unlock()
        }
        ```

- Synchronization initialization must happen once and must happen before everything else. This can be tricky when multiple goroutines run in parallel. So, to guarantee one initialization happens before the goroutines happen is:
    1. To initialize before goroutines
    2. `once.Do(f)`, where function `f` is executed only one time even if its called in multiple goroutines. All calls to `once.Do(f)` will be blocked until the first one returns. This will ensure that initialization executes first.

        ```go
        var on sync.Once
        
        // setup should execute only once
        func setup() {
            fmt.Println("Init")
        }

        func dostuff() {
            on.Do(setup)
            fmt.Println("Hello")
            wg.Done()
        }

        var wg sync.WaitGroup
        
        func main() {
            wg.Add(2)
            go dostuff()  // Init, Hello
            go dostuff()  // Hello
            wg.Wait()
        }
        ```

- Deadlock comes from synchronization dependencies. Synchronization causes the execution of different goroutines to depend on each other. Circular dependencies cause all involved goroutines to block. For example G1 waits for G2 and G2 waits for G1, this is called _deadlock_. For example:

    ```go
    func dostuff(c1 chan int, c2 chan int) {
        <- c1  // read from first channel
        c2 <- 1  // write to second channel
        wg.Done()
    }

    func main() {
        ch1 := make(chan int)
        ch2 := make(chan int)
        wg.Add(2)
        go dostuff(ch1, ch2)  // Each goroutine blocked on channel read
        go dostuff(ch2, ch1)
        wg.Wait()
    }
    ```

- Golang runtime automatically detects when all goroutines are deadlocked. It cannot however detect when a subset of goroutines are deadlocked.

- Dining Philosophers Issue:
    - Each chopstick is a mutex
    - Each philosopher is associated with a goroutine and two chopsticks
    - Naive Code:

    ```go
    type ChopS struct {
        sync.Mutex
    }

    type Philo struct {
        leftCS, rightCS *ChopS
    }

    func (p Philo) eat() {
        for {
            // interleaving where each philosopher can grab their left chopstick
            p.leftCS.Lock()
            p.rightCS.Lock()

            fmt.Println("eating")

            p.rightCS.Unlock()
            p.leftCS.Unlock()
        }
    }

    func main() {
        CSticks := make([] *ChopS, 5)  // initialize chopsticks
        for i:=0; i<5; i++ {
            CSticks[i] = new(ChopS)
        }

        philos := make([]*Philo, 5)  // initialize philosophers
        for i:=0; i<5; i++ {
            philos[i] = &Philo{
                // this violates picking up lowest numbered chopsticks first
                CSticks[i], CSticks[(i+1)%5]
            }
        }
        
        // start dining
        for i:=0; i<5; i++ {
            go philos[i].eat()
        }
    }
    ```

- In the above code, deadlock can happen if all the philosophers grab their left chopstick. Then if it tries to lock the right, its not possible. One way to fix this is dijkstra's way. He proposed that each philosopher picks up lowest numbered chopstick first. For example, Philosopher 4 picks up chopstick 0 before chopstick 4, he blocks allowing philosopher 3 to eat, no deadlock, but philosopher 4 may starve.


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


### Dining Philosophers Problem

- The Golang runtime scheduler has feature to manages all the goroutines that are created and need processor time. The scheduler binds operating system's threads to logical processors in order to execute the goroutines. By sitting on top of the operating system, scheduler controls everything related to which goroutines are running on which logical processors at any given time.

- Communicating Sequential Processes, or CSP for short, is used to describe how systems that feature multiple concurrent models should interact with one another. It typically relies heavily on using channels as a medium for passing messages between two or more concurrent processes, and is the underlying mantra of Golang.

- __Concurrency__ - Concurrency is about to handle numerous tasks at once. This means that you are working to manage numerous tasks done at once in a given period of time. However, you will only be doing a single task at a time. This tends to happen in programs where one task is waiting and the program determines to drive another task in the idle time. It is an aspect of the problem domain — where your program needs to handle numerous simultaneous events.

- __Parallelism__ - Parallelism is about doing lots of tasks at once. This means that even if we have two tasks, they are continuously working without any breaks in between them. It is an aspect of the solution domain — where you want to make your program faster by processing different portions of the problem in parallel.

- A concurrent program has multiple logical threads of control. These threads may or may not run in parallel. A parallel program potentially runs more quickly than a sequential program by executing different parts of the computation simultaneously (in parallel). It may or may not have more than one logical thread of control.

#### Illustration of the Dining Philosophers Problem

- Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent philosophers. Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on their left as they become available, but cannot start eating before getting both forks. Eating is not limited by the remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed. The problem is how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve; i.e., each can forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to eat or think.


    ```go
    package main

    import (
        "hash/fnv"
        "log"
        "math/rand"
        "os"
        "sync"
        "time"
    )

    // Number of philosophers is simply the length of this list.
    var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

    const hunger = 3                // Number of times each philosopher eats
    const think = time.Second / 100 // Mean think time
    const eat = time.Second / 100   // Mean eat time

    var fmt = log.New(os.Stdout, "", 0)

    var dining sync.WaitGroup

    func diningProblem(phName string, dominantHand, otherHand *sync.Mutex) {
        fmt.Println(phName, "Seated")
        h := fnv.New64a()
        h.Write([]byte(phName))
        rg := rand.New(rand.NewSource(int64(h.Sum64())))
        rSleep := func(t time.Duration) {
            time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
        }
        for h := hunger; h > 0; h-- {
            fmt.Println(phName, "Hungry")
            dominantHand.Lock() // pick up forks
            otherHand.Lock()
            fmt.Println(phName, "Eating")
            rSleep(eat)
            dominantHand.Unlock() // put down forks
            otherHand.Unlock()
            fmt.Println(phName, "Thinking")
            rSleep(think)
        }
        fmt.Println(phName, "Satisfied")
        dining.Done()
        fmt.Println(phName, "Left the table")
    }

    func main() {
        fmt.Println("Table empty")
        dining.Add(5)
        fork0 := &sync.Mutex{}
        forkLeft := fork0
        for i := 1; i < len(ph); i++ {
            forkRight := &sync.Mutex{}
            go diningProblem(ph[i], forkLeft, forkRight)
            forkLeft = forkRight
        }
        go diningProblem(ph[0], fork0, forkLeft)
        dining.Wait() // wait for philosphers to finish
        fmt.Println("Table empty")
    }
    ```


---
## Definitions

- Hiding latency: is executing a task while another task is waiting.

- A process is an instance of a running program.

- Context switch: control flow changes from one process to another.