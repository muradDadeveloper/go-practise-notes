# Declaring a Constant

- Here is an example of declaring a constant in Go:

```go
package main  
import ("fmt")  
  
const PI = 3.14  
  
func main() {  
  fmt.Println(PI)  
}
```

## Constant Rules

- Constant names follow the same naming rules as [variables](https://www.w3schools.com/go/go_variable_naming_rules.php)
- Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
- Constants can be declared both inside and outside of a function

## Constant Types

There are two types of constants:

- Typed constants
- Untyped constants

## Typed Constants

- Typed constants are declared with a defined type:

```go
package main  
import ("fmt")  
  
const A int = 1  
  
func main() {  
  fmt.Println(A)  
}
```

## Untyped Constants

- Untyped constants are declared without a type:

```go
package main  
import ("fmt")  
  
const A = 1  
  
func main() {  
  fmt.Println(A)  
}
```

## Multiple Constants Declaration

- Multiple constants can be grouped together into a block for readability:

```go
package main  
import ("fmt")  
  
const (  
  A int = 1  
  B = 3.14  
  C = "Hi!"  
)  
  
func main() {  
  fmt.Println(A)  
  fmt.Println(B)  
  fmt.Println(C)  
}
```