package main

import (
	"fmt"
)

func main() {
	var a, b string = "Hello", "World"
	fmt.Print(a)
	fmt.Print(b)

	var c, d string = "Hello", "World"
	fmt.Print(c, "\n")
	fmt.Print(d, "\n")

	var e, f string = "Hello", "World"
	fmt.Print(e, "\n", f)

	var g, h string = "Hello", "World"
	fmt.Print(g, "\n", h)

	var i, j string = "Hello", "World"
	fmt.Print(i, " ", j)

	var k, l = 10, 20
	fmt.Print(k, l)

	var m, n string = "Hello", "World"
	fmt.Println(m, n)

	var o string = "Hello"
	var p int = 15

	fmt.Printf("i has value: %v and type: %T\n", o, o)
	fmt.Printf("i has value: %v and type: %T", p, p)

}
