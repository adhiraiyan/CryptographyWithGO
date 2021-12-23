# Getting Started with GO

--- 

## Week 1: Getting started with GO


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

## Week 2: Basic Data Types

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

## Week 3: Composite Data Types

- Arrays: is a fixed length series of elements of a chosen type. Elements are initialized to zero value unlike in C. `...` can be used to infer the size of the array.

    ```go
    var x [5] int
    x[0] = 2
    fmt.Printf("x[1]")  // 0

    x := [...]int{1, 2, 3, 4}
    ```

- Iterating through an array:

    ```go
    x := [3]int {1, 2, 3}

    // range returns two values, index and element at index
    for i, v range x {
        fmt.Printf("ind %d, val %d", i, v)
    }
    ```

- Slice: is a window on an underlying array. Slices can have variable size, up to the whole array. Every slice has three properties:
    1. Pointer: indicates the start of the slice
    2. Length: is the number of elements in the slice, `len()`
    3. Capacity: is maximum number of elements, `cap()`

    ```go
    arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
    s1 := arr[1:3]  // "b", "c", len=2, cap=6
    s2 := arr[2:5]  // "c", "d", "e", len=3, cap=5
    ```

- To create a slice literal: `s1 := []int{1, 2, 3}`, note that `[]` is empty, this is how the compiler knows this is a slice.

- `make()` can be used to create a slice (and array), this can be done in two way:
    1. `sli = make([]int, 10)`, 2 args: type, length/capacity
    2. `sli = make([]int, 10, 15)`, 3 args: type, length, capacity

- size of a slice can be increased using `append()` by adding elements to the end of a slice. If you reach the size of the underlying array, the underlying arrays size will expand.

- Maps in GO are implementation of a hash table. Can be created using `make()`:

    ```go
    var idMap map[string]int  // string: key type, int: value type
    idMap = make(map[string]int)

    // can also be defined as a literal map
    idMap := map[string]int {
        "joe": 123
    }

    // adding or modifying the key value pair
    idMap["jane"] = 456
    ```

- Two value assignment tests for existence of the key:

    ```go
    id, p := idMap["joe"]  // id is value, p is presence of key
    ```

- Iterating through a map:

    ```go
    for key, val := range idMap {
        fmt.Println(key, val)
    }
    ```

- Structs groups together objects of arbitrary type, similar to from C language.

    ```go
    type struct Person {
        // each property is a field
        name string
        addr string
        phone string
    }
    var p1 Person

    // to access and assign values of the property
    p1.name = "joe"
    x = p1.addr

    p2 := new(Person)  // initializes fields to zero

    p3 := Person(name: "joe", addr: "a st.", phone: "123")  // can initialize using a struct literal
    ```


--- 

## Week 4: Protocols and Formats

- RFCs (Request for Comments) is a definitions of internet protocols and formats. Example  protocols:
    - HTML: Hypertext Markup Language, 1866
    - URI: Uniform Resource Identifier, 3986
    - HTTP: Hypertext Transfer Protocol, 2616

- JSON Marshalling- `Marshal()` returns JSON representation as `[]byte`: 

    ```go
    type struct Person {
        name string
        addr string
        phone string
    }

    p1 := Person(name: "joe", addr: "a st.", phone: "123")
    barr, err := json.Marshal(p1)
    ```

- JSON Unmarshalling - `Unmarshal()` converts a JSON `[]byte` into a GO object.

    ```go
    var p2 Person
    err := json.Unmarshal(barr, &p2)
    ```

- Basic File operations:
    - Open: get handle for access
    - Read: read bytes into `[]byte`
    - Write: write `[]byte` into file
    - Close: release handle
    - Seek: move read/write head

- `io/ioutil` package has basic functions for file operations. It has implicit open/close, but this may cause problems if the files are large:

    ```go
    dat, e := ioutil.ReadFile("test.txt")  // dat is []byte filled with contents of entire file

    err := ioutil.WriteFile("outfile.txt", dat, 0777)  // filename, object, permission
    ```

- OS Package File access:
    - `os.Open()` opens a file
    - `os.Close()` closes a file
    - `os.Read()` reads from a file into a `[]byte`
    - `os.Write()` writes a `[]byte` into a file

    ```go
    f, err := os.Open("dt.txt")
    barr := make([]byte, 10)
    nb, err := f.Read(barr)  // returns # of bytes read
    f.Close()

    f, err := os.Create("outfile.txt")
    barr := []byte{1, 2, 3}
    nb, err := f.Write(barr)  // writes any unicode sequence
    nb, err := f.WriteString("Hi")  // writes a string
    ```


--- 

## My Notes





---
## Definitions

- inheritance: is a concept that acquires the properties from one class to other classes.

- constructors: is a special method that is used to initialize objects.

- generics: is a class or interface that is parameterized over types, meaning that a type can be assigned by performing generic type invocation.

- Variable Scope: the places in code where a variable can be accessed.

- Blocks: a sequence of declarations and statements within matching brackets, __{}__.

- Lexical Scoping: GO is lexically scoped using blocks.

- Array literal is an array pre-defined with values, for example `var x [5] int = [5]{1, 2, 3, 4, 5}`

- Hash function is used to compute the slot for a key.

- JSON Marshalling is generating JSON representation from a GO object.