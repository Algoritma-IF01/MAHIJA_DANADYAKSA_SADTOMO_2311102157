package main

import (
	"fmt"
)

func operasi_F(k float64, hasil *float64) {
	up := (4*k + 2) * (4*k + 2)
	down := (4*k + 1) * (4*k + 3)
	*hasil *= up / down  
}

func operasi_akar2(k int, hasil *float64) {
	*hasil = 1.0 
	for i := 0; i <= k; i++ {
		operasi_F(float64(i), hasil)
	}
}

func main() {
	var k int
	var hasil float64

	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	operasi_akar2(k, &hasil)
	fmt.Printf("Nilai hampiran akar 2 = %.10f\n", hasil)
}