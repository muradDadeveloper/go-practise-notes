# Go Output Functions

Go has three function is output text: 
- <code>Print()</code>
- <code>Println()</code>
- <code>Printf()</code>

## The Print() Function

The `Print()` function prints its arguments with their default format.

```go
package main  
import ("fmt")  
  
func main() {  
  var i,j string = "Hello","World"  
  
  fmt.Print(i)  
  fmt.Print(j)  
}
```

Output:
```Bash
Hello   
World
```


It is also possible to only use one `Print()` for printing multiple variables.

```go
package main  
import ("fmt")  
  
func main() {  
  var i,j string = "Hello","World"  
  
  fmt.Print(i, "\n",j)  
}
```

Output:
```Bash
Hello
World
```

If we want to add a space between string arguments, we need to use " ":
```go
package main  
import ("fmt")  
  
func main() {  
  var i,j string = "Hello","World"  
  
  fmt.Print(i, " ", j)  
}
```

Output:
```Bash
Hello World
```

`Print()` inserts a space between the arguments if **neither** are strings:

```go
package main  
import ("fmt")  
  
func main() {  
  var i,j = 10,20  
  
  fmt.Print(i,j)  
}
```

Output: 
```Bash
10 20
```


## The Println() Function

The `Println()` function is similar to `Print()` with the difference that a whitespace is added between the arguments, and a newline is added at the end:

```go
package main  
import ("fmt")  
  
func main() {  
  var i,j string = "Hello","World"  
  
  fmt.Println(i,j)  
}
```

Output:

```Bash
Hello World
```

## The Printf() Function

The `Printf()` function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

- `%v` is used to print the **value** of the arguments
- `%T` is used to print the **type** of the arguments

```go
package main  
import ("fmt")  
  
func main() {  
  var i string = "Hello"  
  var j int = 15  
  
  fmt.Printf("i has value: %v and type: %T\n", i, i)  
  fmt.Printf("j has value: %v and type: %T", j, j)  
}
```

Output:

```bash
i has value: Hello and type: string
j has value: 15 and type: int
```