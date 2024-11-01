package main

import (
	"fmt"
)

func perkalian(n, m int) []int {
	
	if m == 0 {
		return []int{0}
	}

	if m == 1 {
		return []int{n}
	}

	result := perkalian(n, m-1)
	result = append(result, n) 
	return result
}

func main() {
	var n, m int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&m)
	fmt.Print("Masukkan nilai m: ")
	fmt.Scan(&n)

	steps := perkalian(n, m)
	fmt.Printf("Proses: %v\n", steps)
	fmt.Printf("Hasil dari %d x %d = %d\n", m, n, m*n)
}
