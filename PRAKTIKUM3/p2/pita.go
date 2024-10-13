package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner:= bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita = pita + " - " + input
		}
		bungaCount++
	}

	fmt.Println("Pita bunga: ", pita)
	fmt.Println("Jumlah bunga: ", bungaCount)
}
