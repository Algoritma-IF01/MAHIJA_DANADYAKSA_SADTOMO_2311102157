package main

import (
	"fmt"
)

func myfunction(fname string) {
	fmt.Printf("%v Doe", fname)
}

func main() {
	myfunction("John")
}
