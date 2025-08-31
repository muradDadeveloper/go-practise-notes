## Append Elements To a Slice

You can append elements to the end of a slice using the `append()`function:

```go
package main  
import ("fmt")  
  
func main() {  
  myslice1 := []int{1, 2, 3, 4, 5, 6}  
  fmt.Printf("myslice1 = %v\n", myslice1)  
  fmt.Printf("length = %d\n", len(myslice1))  
  fmt.Printf("capacity = %d\n", cap(myslice1))  
  
  myslice1 = append(myslice1, 20, 21)  
  fmt.Printf("myslice1 = %v\n", myslice1)  
  fmt.Printf("length = %d\n", len(myslice1))  
  fmt.Printf("capacity = %d\n", cap(myslice1))  
}
```

## Append One Slice To Another Slice

To append all the elements of one slice to another slice, use the `append()`function:

```go
package main  
import ("fmt")  
  
func main() {  
  myslice1 := []int{1,2,3}  
  myslice2 := []int{4,5,6}  
  myslice3 := append(myslice1, myslice2...)  
  
  fmt.Printf("myslice3=%v\n", myslice3)  
  fmt.Printf("length=%d\n", len(myslice3))  
  fmt.Printf("capacity=%d\n", cap(myslice3))  
}
```

## Change The Length of a Slice

Unlike arrays, it is possible to change the length of a slice.

```go
package main  
import ("fmt")  
  
func main() {  
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array  
  myslice1 := arr1[1:5] // Slice array  
  fmt.Printf("myslice1 = %v\n", myslice1)  
  fmt.Printf("length = %d\n", len(myslice1))  
  fmt.Printf("capacity = %d\n", cap(myslice1))  
  
  myslice1 = arr1[1:3] // Change length by re-slicing the array  
  fmt.Printf("myslice1 = %v\n", myslice1)  
  fmt.Printf("length = %d\n", len(myslice1))  
  fmt.Printf("capacity = %d\n", cap(myslice1))  
  
  myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items  
  fmt.Printf("myslice1 = %v\n", myslice1)  
  fmt.Printf("length = %d\n", len(myslice1))  
  fmt.Printf("capacity = %d\n", cap(myslice1))  
}
```

## Memory Efficiency

 When using slices, Go loads all the underlying elements into the memory.
If the array is large and you need only a few elements, it is better to copy those elements using the `copy()` function.

The `copy()` function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. The `copy()` function takes in two slices _dest_ and _src_, and copies data from _src_ to _dest_. It returns the number of elements copied.

```Go
package main  
import ("fmt")  
  
func main() {  
  numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}  
  // Original slice  
  fmt.Printf("numbers = %v\n", numbers)  
  fmt.Printf("length = %d\n", len(numbers))  
  fmt.Printf("capacity = %d\n", cap(numbers))  
  
  // Create copy with only needed numbers  
  neededNumbers := numbers[:len(numbers)-10]  
  numbersCopy := make([]int, len(neededNumbers))  
  copy(numbersCopy, neededNumbers)  
  
  fmt.Printf("numbersCopy = %v\n", numbersCopy)  
  fmt.Printf("length = %d\n", len(numbersCopy))  
  fmt.Printf("capacity = %d\n", cap(numbersCopy))  
}
```