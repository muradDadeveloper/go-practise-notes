# Declare Variables 

<p>There are three types of variables in Go: </p>

* int
* float32
* string
* bool

## Declaring Variables 

* <p>With the <code>var </code> keyword </p>
* <p>With the <code>:= </code> sign </p>

## Variable Declaration With Initial Value
<p>If the value of a variable is known from the start, you can declare the variable and assign a value to it on one line:</p>

```go
package main  
import ("fmt")  
  
func main() {  
  var student1 string = "John" _//type is string_  
  var student2 = "Jane" _//type is inferred_  
  x := 2 _//type is inferred_  
  
  fmt.Println(student1)  
  fmt.Println(student2)  
  fmt.Println(x)  
}
```

## Variable Declaration Without Initial Value

<p>In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:</p>

```go
package main  
import ("fmt")  
  
func main() {  
  var a string  
  var b int  
  var c bool  
  
  fmt.Println(a)  
  fmt.Println(b)  
  fmt.Println(c)  
}
```

## Difference Between var and :=

| var                                                                  | :=                                                                                                      |
| -------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| Can be used **inside** and **outside** of functions                  | Can only be used **inside** functions                                                                   |
| Variable declaration and value assignment **can be done separately** | Variable declaration and value assignment **cannot be done separately** (must be done in the same line) |

## Go Multiple Variable Declaration

<p>In Go, it is possible to declare multiple variables on the same line.</p>

```go
package main  
import ("fmt")  
  
func main() {  
  var a, b, c, d int = 1, 3, 5, 7  
  
  fmt.Println(a)  
  fmt.Println(b)  
  fmt.Println(c)  
  fmt.Println(d)  
}
```

## Go Variable Naming Rules

- A variable name must start with a letter or an underscore character (_)
- A variable name cannot start with a digit
- A variable name can only contain alpha-numeric characters and underscores (`a-z, A-Z`, `0-9`, and `_` )
- Variable names are case-sensitive (age, Age and AGE are three different variables)
- There is no limit on the length of the variable name
- A variable name cannot contain spaces
- The variable name cannot be any Go keywords

## Multi-Word Variable Names

* Camel Case: myVariableName = "John"
* Pascal Case: MyVariableName = "John"
* Snake Case: my_variable_name = "John"