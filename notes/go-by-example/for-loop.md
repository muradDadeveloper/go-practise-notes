# Go by Example: For

`for` is Go’s only looping construct. Here are some basic types of `for` loops.

The most basic type, with a single condition. A classic initial/condition/after `for` loop. Another way of accomplishing the basic “do this N times” iteration is `range` over an integer.

`for` without a condition will loop repeatedly until you `break` out of the loop or `return` from the enclosing function. You can also `continue` to the next iteration of the loop.

```go
package main
import "fmt"

func main() {
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    for i := range 3 {
        fmt.Println("range", i)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```

Output:

```Bash
$ go run for.go
1
2
3
0
1
2
range 0
range 1
range 2
loop
1
3
5
```