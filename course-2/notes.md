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
