# Getting Started with GO

--- 

## Week 1


- While GO being compiled language, it has features from interpreted language like Garbage Collection (automatic memory management)

- GO is weakly object oriented but have fewer features. Object oriented programming organizes your code through _encapsulation_ and group together data and functions which are related. These can be thought of as user-defined type. For example ints have data (the number) and functions (+, -, *, etc). GO doesn't use the term _class_, it uses _structs_ with associated methods. GO doesn't have inheritance, constructors, generics. 

- _Concurrency_ is the management of multiple tasks at the same time. It doesn't have to be executing at the same time, but they are alive at the same time. Concurrent programming enables parallelism. It helps in the management, communication and synchronization between tasks. GO includes concurrency primitives such as Goroutines (represent concurrent tasks), Channels (are used to communicate between tasks) Select (enables task synchronization)

- Workspace structure:
    - src: contains source code files
    - pkg: contains packages (libraries)
    - bin: contains executables

- GO Tools:
    - go build: compiles the program
    - go doc: prints documentation for a package
    - go fmt: formats source code files
    - go get: downloads packages and installs them
    - go list: list all installed packages
    - go run: compiles .go files and runs the executables
    - go test: runs test using files ending in "_test.go"

- Most basic declaration: `var x int` (keyword name type). 

- After defining a type: `type Celsius float64`, it can be used to declare variables: `var temp Celsius`. Two ways to initialize the declaration 1: `var x int = 100` and 2: `var x int, x = 100`. You can also perform declaration and initialization together with := operator: `x := 100` (here the type is inferred), this can only be done inside a function.


---

## Week 2

- Pointers: a pointer is an address to data in memory. __&__ operator returns the address of a variable/function and __*__ operator returns data at an address (dereferencing). 

    ```go
    var x int = 1
    var y int 
    var ip *int  // ip is pointer to int
    ip = &x  // ip now points to x 
    y = *ip  // y is now 1
    ```

- `new()` function creates a variable and returns a pointer to the variable. This is an alternate way to create a variable. Variable is initialized to zero by default

    ```go
    ptr := new(int)
    *ptr = 3
    ```

- Stacks: is an area of memory dedicated to function calls, for example local variables. These are deallocated after function completes in other languages like C but GO tweaks this a bit.

    ```go
    func f() {
        var x = 1  // stack
        fmt.Printf("%d", x)
    }

    func g() {
        fmt.Printf("%d", x)
    }
    ```

- Heap: is a persistent region of memory. This needs to explicitly deallocated in other languages like C, but in GO it tweaks this a bit. 

    ```go
    var x = 1  // heap
    func f() {
        fmt.Printf("%d", x)
    }

    func g() {
        fmt.Printf("%d", x)
    ```

- Garbage Collection in GO keeps tracks of pointers and only when all the references are gone, the memory is deallocated. GO also figures out what goes into stack and heap and garbage collects appropriately in the background. This is the tweak I mention above. The downside is a minor performance issue.

- Type Conversion:
    ```go
    var x int32 = 1 
    var y int16 = 2 
    /*
    If you want to assign the value of y to x, you can't simply do x=y
    since the types are different
    */

    x = int32(y)  // converts y into int32 and assigns to x
    ```

- iota can be used to generate a set of related but distinct constants. The constants must be different but _actual value is not important_. Its like an enumerated type in other languages like C.

    ```go
    type Grades int 
    const (
        A Grades = iota 
        B
        C
        D
        E
        F
    )
    ```

- For loops: 
    
    ```go
    for <init>; <cond>; <update> {
        <statements>
    }

    // Three common forms
    for i:=0; i<10; i++ {
        fmt.Printf("Hi")
    }

    i = 0 
    for i<10 {
        fmt.Printf("Hi")
        i++
    }

    for {
        fmt.Printf("Hi")
    }
    ```

- Switch/Case:

    ```go
    switch x {
        case 1: 
            fmt.Printf("Case 1")
        case 2: 
            fmt.Printf("Case 2")
        default: 
            fmt.Printf("Default")
    }
    ```

- Tagless Switch (case contains a boolean expression to evaluate and whichever condition becomes true will get evaluated first):

    ```go
    switch {
        case x>1: 
            fmt.Printf("Case 1")
        case x<-1: 
            fmt.Printf("Case 2")
        default: 
            fmt.Printf("Default")
    }
    ```

- Break and continue:

    ```go
    i := 0
    for i < 10 {
        i++
        if i == 5 {break}  // 1, 2, 3, 4, 5
        fmt.Printf("Hi")
    }
    
    for i < 10 {
        i++
        if i == 5 {continue} // 1, 2, 3, 4, 6, 7, 8, 9, 10
        fmt.Printf("Hi")
    }
    ```

- scan, takes a pointer as an argument, typed data is then written to pointer. It returns the number of scanned items and error if any:

    ```go
    var appleNum int 
    
    fmt.Printf("Number of apples?")

    num, err := fmt.Scan(&appleNum)
    fmt.Printf(appleNum)
    ```


--- 

# Week 3




--- 

# Week 4



---
### Definitions

- inheritance: is a concept that acquires the properties from one class to other classes.

- constructors: is a special method that is used to initialize objects.

- generics: is a class or interface that is parameterized over types, meaning that a type can be assigned by performing generic type invocation.

- Variable Scope: the places in code where a variable can be accessed.

- Blocks: a sequence of declarations and statements within matching brackets, __{}__.

- Lexical Scoping: GO is lexically scoped using blocks.

- 