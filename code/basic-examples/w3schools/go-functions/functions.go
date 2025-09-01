package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func main() {
	myMessage()
	familyName("John")
}
