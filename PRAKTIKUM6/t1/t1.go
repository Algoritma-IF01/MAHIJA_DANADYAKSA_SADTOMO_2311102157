package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
		return
	}

	weight := make([]float64, N)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&weight[i])
	}

	minWeight, maxWweight := weight[0], weight[0]

	for _, w := range weight[1:] {
		if w < minWeight {
			minWeight = w
		}
		if w > maxWweight {
			maxWweight = w
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxWweight)
}
