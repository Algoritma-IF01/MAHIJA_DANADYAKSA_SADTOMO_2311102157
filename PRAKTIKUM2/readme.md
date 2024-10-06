package main

import (
	"fmt"
)

func main() {
	var integers [5]int
	var chars [3]rune

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &integers[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%c", rune(integers[i]))
	}
	fmt.Println()

	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i]+3)
	}
	fmt.Println()
}
