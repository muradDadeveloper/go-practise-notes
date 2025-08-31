## Go Slices

Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

In Go, there are several ways to create a slice:

- Using the []_datatype_{_values_} format
- Create a slice from an array
- Using the make() function

## Create a Slice With []_datatype_{_values_}

A common way of declaring a slice is like this:

```go
myslice := []int{}
```

The code above declares an empty slice of 0 length and 0 capacity.
To initialize the slice during declaration, use this:

```Go
myslice := []int{1,2,3}
```

The code above declares a slice of integers of length 3 and also the capacity of 3.
In Go, there are two functions that can be used to return the length and capacity of a slice:
- `len()` function - returns the length of the slice (the number of elements in the slice)
- `cap()` function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

```go
package main  
import ("fmt")  
  
func main() {  
  myslice1 := []int{}  
  fmt.Println(len(myslice1))  
  fmt.Println(cap(myslice1))  
  fmt.Println(myslice1)  
  
  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}  
  fmt.Println(len(myslice2))  
  fmt.Println(cap(myslice2))  
  fmt.Println(myslice2)  
}
```

Output:
```Bash
0
0
[]
4 
4
[Go Slices Are Powerful]
```

## Create a Slice From an Array

You can create a slice by slicing an array:

```go
var myarray = [length]datatype{values} // An array  
myslice := myarray[start:end] // A slice made from the array
```

This example shows how to create a slice from an array:

```Go
package main  
import ("fmt")  
  
func main() {  
  arr1 := [6]int{10, 11, 12, 13, 14,15}  
  myslice := arr1[2:4]  
  
  fmt.Printf("myslice = %v\n", myslice)  
  fmt.Printf("length = %d\n", len(myslice))  
  fmt.Printf("capacity = %d\n", cap(myslice))  
}
```

Output: 

```Go
myslice = [12 13]
length = 2
capacity = 4
```

## Create a Slice With The make() Function

The `make()` function can also be used to create a slice.

```go
package main  
import ("fmt")  
  
func main() {  
  myslice1 := make([]int, 5, 10)  
  fmt.Printf("myslice1 = %v\n", myslice1)  
  fmt.Printf("length = %d\n", len(myslice1))  
  fmt.Printf("capacity = %d\n", cap(myslice1))  
  
  // with omitted capacity  
  myslice2 := make([]int, 5)  
  fmt.Printf("myslice2 = %v\n", myslice2)  
  fmt.Printf("length = %d\n", len(myslice2))  
  fmt.Printf("capacity = %d\n", cap(myslice2))  
}
```

Output:

```bash
myslice1 = [0 0 0 0 0]
length = 5
capacity = 10
myslice2 = [0 0 0 0 0]
length = 5
capacity = 5
```

