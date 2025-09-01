package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "Year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	var c = make(map[string]string)
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"

	var d = make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)

	var e = make(map[string]string)
	var f map[string]string

	fmt.Println(e == nil)
	fmt.Println(f == nil)

	var g = make(map[string]string)
	g["brand"] = "Ford"
	g["model"] = "Mustang"
	g["year"] = "1964"

	fmt.Println(g)
	delete(g, "year")

	fmt.Println(g)

	var h = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := h["brand"]
	val2, ok2 := h["color"]
	val3, ok3 := h["day"]
	_, ok4 := h["model"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	i := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for j, k := range i {
		fmt.Printf("%v: %v, ", j, k)
	}

	l := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var m []string
	m = append(m, "one", "two", "three", "four")

	for n, o := range l {
		fmt.Printf("%v : %v, ", n, o)
	}

	fmt.Println()

	for _, element := range m {
		fmt.Printf("%v: %v,", element, l[element])
	}
}
