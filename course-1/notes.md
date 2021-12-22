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
















### Definitions

- inheritance: is a concept that acquires the properties from one class to other classes

- constructors: is a special method that is used to initialize objects.

- generics: generic type is a class or interface that is parameterized over types, meaning that a type can be assigned by performing generic type invocation,
