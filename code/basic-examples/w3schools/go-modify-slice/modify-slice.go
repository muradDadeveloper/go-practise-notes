package main

import (
	"fmt"
)

func main() {

	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice3 := []int{1, 2, 3}
	myslice4 := []int{4, 5, 6}
	myslice5 := append(myslice3, myslice4...)

	fmt.Printf("myslice3=%v\n", myslice5)
	fmt.Printf("length=%d\n", len(myslice5))
	fmt.Printf("capacity=%d\n", cap(myslice5))

	arr1 := [6]int{9, 10, 11, 12, 13, 14}
	myslice6 := arr1[1:5]
	fmt.Printf("myslice1 = %v\n", myslice6)
	fmt.Printf("length = %d\n", len(myslice6))
	fmt.Printf("capacity = %d\n", cap(myslice6))

	myslice7 := arr1[1:3]
	fmt.Printf("myslice1 = %v\n", myslice7)
	fmt.Printf("length = %d\n", len(myslice7))
	fmt.Printf("capacity = %d\n", cap(myslice7))

	myslice8 := append(myslice7, 20, 21, 22, 23)
	fmt.Printf("myslice8 = %v\n", myslice8)
	fmt.Printf("length = %d\n", len(myslice8))
	fmt.Printf("capacity = %d\n", cap(myslice8))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}
