# Go Arrays

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

## Declare an Array

In Go, there are two ways to declare an array:

1. With the `var` keyword
2. With the `:=` sign

## Array Examples

```go
package main  
import ("fmt")  
  
func main() {  
  var arr1 = [3]int{1,2,3}  
  arr2 := [5]int{4,5,6,7,8}  
  
  fmt.Println(arr1)  
  fmt.Println(arr2)  
}
```

Output:
```Bash
[1 2 3]
[4 5 6 7 8]
```


This example declares two arrays (arr1 and arr2) with inferred lengths:

```go
package main  
import ("fmt")  
  
func main() {  
  var arr1 = [...]int{1,2,3}  
  arr2 := [...]int{4,5,6,7,8}  
  
  fmt.Println(arr1)  
  fmt.Println(arr2)  
}
```

Output:
```Bash
[1 2 3]
[4 5 6 7 8]
```

This example declares an array of strings:

```go
package main  
import ("fmt")  
  
func main() {  
  var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}  
  fmt.Print(cars)  
}
```

Output:
```Bash
[Volvo BMW Ford Mazda]
```

## Access Elements of an Array

You can access a specific array element by referring to the index number.
In Go, array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

This example shows how to access the first and third elements in the prices array:

```go
package main  
import ("fmt")  
  
func main() {  
  prices := [3]int{10,20,30}  
  
  fmt.Println(prices[0])  
  fmt.Println(prices[2])  
}
```

Output:

```Bash
10
30
```

## Change Elements of an Array

You can also change the value of a specific array element by referring to the index number.
This example shows how to change the value of the third element in the prices array:

```go
package main  
import ("fmt")  
  
func main() {  
  prices := [3]int{10,20,30}  
  
  prices[2] = 50  
  fmt.Println(prices)  
}
```

Output:

```bash
[10 20 50]
```

## Array Initialization

If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

**Tip:** The default value for int is 0, and the default value for string is "".

```go
package main  
import ("fmt")  
  
func main() {  
  arr1 := [5]int{} //not initialized  
  arr2 := [5]int{1,2} //partially initialized  
  arr3 := [5]int{1,2,3,4,5} //fully initialized  
  
  fmt.Println(arr1)  
  fmt.Println(arr2)  
  fmt.Println(arr3)  
}
```

Output: 

```Bash
[0 0 0 0 0]
[1 2 0 0 0]
[1 2 3 4 5]
```

## Initialize Only Specific Elements

It is possible to initialize only specific elements in an array. This example initializes only the second and third elements of the array:

```go
package main  
import ("fmt")  
  
func main() {  
  arr1 := [5]int{1:10,2:40}  
  
  fmt.Println(arr1)  
}
```

Output:

```Bash
[ 10 40 0 0]
```

#### Example Explained

 The array above has 5 elements.

- `1:10` means: assign `10` to array index `1` (second element).
- `2:40` means: assign `40` to array index `2` (third element).

## Find the Length of an Array

The `len()` function is used to find the length of an array:

```go
package main  
import ("fmt")  
  
func main() {  
  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}  
  arr2 := [...]int{1,2,3,4,5,6}  
  
  fmt.Println(len(arr1))  
  fmt.Println(len(arr2))  
}
```

```Bash
4
6
```