# Go by Example: Multiple Return Values

The `(int, int)` in this function signature shows that the function returns 2 `int`s.

Here we use the 2 different return values from the call with _multiple assignment_. If you only want a subset of the returned values, use the blank identifier `_`. Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.
 
```go
package main
import "fmt"

func vals() (int, int) {
    return 3, 7
}

func main() {

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}
```

Output:

```Bash
$ go run multiple-return-values.go
3
7
7
```