# Go Functions

## Create a Function

To create (often referred to as declare) a function, do the following:

- Use the `func` keyword.
- Specify a name for the function, followed by parentheses ().
- Finally, add code that defines what the function should do, inside curly braces {}.

## Call a Function

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

In the example below, we create a function named "myMessage()". The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function. The function outputs "I just got executed!". To call the function, just write its name followed by two parentheses ():

```go
package main  
import ("fmt")  
  
func myMessage() {  
  fmt.Println("I just got executed!")  
}  
  
func main() {  
  myMessage() // call the function  
}
```

## Function With Parameter Example

The following example has a function with one parameter (`fname`) of type `string`. When the familyName() function is called, we also pass along a name (e.g. Liam), and the name is used inside the function, which outputs several different first names, but an equal last name:

```go
package main  
import ("fmt")  
  
func familyName(fname string) {  
  fmt.Println("Hello", fname, "Refsnes")  
}  
  
func main() {  
  familyName("Liam")  
  familyName("Jenny")  
  familyName("Anja")  
}
```

## Return Values

If you want the function to return a value, you need to define the data type of the return value (such as `int`, `string`, etc), and also use the `return` keyword inside the function:

Here, `myFunction()` receives two integers (`x` and `y`) and returns their addition (`x + y`) as integer (`int`):

```go
package main  
import ("fmt")  
  
func myFunction(x int, y int) int {  
  return x + y  
}  
  
func main() {  
  fmt.Println(myFunction(1, 2))  
}
```

## Named Return Values

In Go, you can name the return values of a function. Here, we name the return value as `result` (of type `int`), and return the value with a naked return (means that we use the `return` statement without specifying the variable name):

```go
package main  
import ("fmt")  
  
func myFunction(x int, y int) (result int) {  
  result = x + y  
  return  
}  
  
func main() {  
  fmt.Println(myFunction(1, 2))  
}
```

## Multiple Return Values

Go functions can also return multiple values.

```go
package main  
import ("fmt")  
  
func myFunction(x int, y string) (result int, txt1 string) {  
  result = x + x  
  txt1 = y + " World!"  
  return  
}  
  
func main() {  
  fmt.Println(myFunction(5, "Hello"))  
}
```

