package main

import (
	"fmt"
)

func faktor(a int) []int {
	var faktor []int
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func prima(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for {
		var bil int
		fmt.Print("Bilangan: ")
		fmt.Scanln(&bil)
		
		fmt.Println("Faktor: ", faktor(bil))
		fmt.Println("Prima: ", prima(bil))
		fmt.Println()
		
		if bil == -1 {
			fmt.Println("Program Selesai")
			break
		}
	}
}