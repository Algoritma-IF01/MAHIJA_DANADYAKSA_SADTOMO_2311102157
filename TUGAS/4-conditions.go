package main

import (
	"fmt"
)

func conditions4() {
	var x = 50
	var y = 50

	if x == y {
		fmt.Print("1")
	} else if x > y {
		fmt.Print("2")
	} else {
		fmt.Print("3")
	}
}
