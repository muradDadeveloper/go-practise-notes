# Go by Example: Closures

Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

This function `intSeq` returns another function, which we define anonymously in the body of `intSeq`. The returned function _closes over_ the variable `i` to form a closure.

We call `intSeq`, assigning the result (a function) to `nextInt`. This function value captures its own `i` value, which will be updated each time we call `nextInt`.

See the effect of the closure by calling `nextInt` a few times. To confirm that the state is unique to that particular function, create and test a new one.
 
```go
package main
import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}
```

Output:

```Bash
$ go run closures.go
1
2
3
1
```