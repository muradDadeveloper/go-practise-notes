package main

import (
	"fmt"
)

func myfunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {

	fmt.Println(myfunction(5, "Hello"))

}
