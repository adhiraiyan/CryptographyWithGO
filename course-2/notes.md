# Functions, Methods and Interfaces in GO

--- 

## Week 1: Functions and Organization

- Function Example:

    ```go
    func addTwo (x, y int) (int, int, int) {
        return x, y, x + y
    }

    func main() {
        x, y, z := addTwo(2, 3)
    }
    ```

- GO call by reference:

    ```go
    func foo (y *int) {
        *y = *y + 1
    }

    func main() {
        x := 2
        foo(&x)
        fmt.Printf(x)  // 3
    }
    ```

- Passing arrays and slice:

    ```go
    func foo(x [3]int) int {
        return x[0]
    }

    func main() {
        a := [3]int{1, 2, 3}
        fmt.Printf(foo(a))
    }
    ```

- Instead of passing array pointers like so:

    ```go
    func foo(x *[3]int) {
        (*x)[0] = (*x)[0] + 1
    }

    func main() {
        a := [3]int{1, 2, 3}
        foo(&a)
        fmt.Printf(a)
    }
    ```

- You can use a slice, since slices contain a pointer to the array, passing a slice copies the pointer instead of the whole array. 

    ```go
    func foo(sli []int) {
        sli[0] = sli[0] + 1
    }

    func main() {
        a := []int{1, 2, 3}
        foo(a)
        fmt.Printf(a)
    }
    ```


--- 

## Week 2: Function Types

- Functions can be treated like other types.
    - variables can be declared with a function type.
    - can be created dynamically.
    - can be passed as arguments and returned as values.
    - can be store in data structures.

    ```go
    var funcVar func (int) int func incFn (x int) int {
        return x + 1
    }

    func main() {
        funcVar = incFn
        fmt.Printf(funcVar(1))  // 2
    }
    ```

- Functions can be passed to another function as an argument.

    ```go
    func applyIt (afunct func (int) int, val int) int {
        return afunct(val)
    }

    func incFn(x int) int {return x + 1}
    func decFn(x int) int {return x - 1}

    func main() {
        fmt.Println(applyIt(incFn, 2))  // 3
        fmt.Println(applyIt(decFn, 2))  // 1
    }
    ```

- Anonymous Functions.

    ```go
    func applyIt (afunct func (int) int, val int) int {
        return afunct(val)
    }

    func main() {
        v := applyIt(func (x int) int {return x + 1}, 2)
        fmt.Println(v)  // 3
    }

    ```

- Functions as return values.

    ```go
    // Origin location is passed as an argument, origin is built into the returned function
    func MakeDistOrigin(o_x, o_y float64) func (float64, float64) float64 {
        fn := func(x, y, float64) float64 {
            math.Sqrt(math.Pow(x - o_x, 2) + math.Pow(y - o_y, 2) + )
        }
        return fn
    }

    func main() {
        Dist1 := MakeDistOrigin(0, 0)
        Dist2 := MakeDistOrigin(2, 2)
        fmt.Println(Dist1(2, 2))
        fmt.Println(Dist1(2, 2))
    }
    ```

- Variadic Function: Use `...` to specify a function can take a variable number of arguments.

    ```go
    func getMax(vals ...int) int {
        maxV := -1
        for _, v := range vals {
            if v > maxV {
                maxV = v
            }
        }
        return maxV
    }

    func main() {
        fmt.Println(getMax(1, 3, 6, 4))
        
        // Variadic Slice argument
        vslice := []int{1, 3, 6, 7}
        fmt.Println(getMax(vslice...))
    }
    ```

- Deferred Function: call can be deferred until the surrounding function completes. Typically used for cleanup activities. Arguments of a deferred call are evaluated immediately.

    ```go
    func main() {
        defer fmt.Println("Bye")
        fmt.Println("Hello")
    }
    ```


--- 

## Week 3: Object Orientation in GO





--- 

## Week 4: Interfaces for Abstraction



--- 

## My Notes





---
## Definitions

- Call by value: passed arguments are copied to parameters. Therefore modifying parameters has no effect outside of the function.

- Call by reference: is not built in GO, all you have to do is pass a pointer.

- Debugging tip: When you have a bug, its either the function is wrong, written wrong or the data that's passed to it is wrong. To make sure the data you get is as it should be you do data test, then function test to verify if a function is working as it should be and finally integration test to make sure when the data is passed around, it behaves as it should be.
