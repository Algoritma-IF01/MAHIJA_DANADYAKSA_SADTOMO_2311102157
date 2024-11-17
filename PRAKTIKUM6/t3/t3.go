package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int) {
	// inisialisasi bMin dan bMax dengan nilai pertama array
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// perulangan untuk mencari nilai bMin dan bMax
	for i := 1; i < n; i++ { // Menggunakan batas n untuk iterasi
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}


func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	// menggunakan batas n untuk iterasi
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var arrBerat arrBalita
	var bMin, bMax float64
	var n int

	// input jumlah data
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scanln(&n)

	// validasi jumlah data
	if n <= 0 || n > len(arrBerat) {
		fmt.Println("Jumlah data tidak valid")
		return
	}

	// input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scanln(&arrBerat[i])
	}

	// hitung dan tampilkan nilai minimum, maksimum, dan rerata
	hitungMinMax(arrBerat, &bMin, &bMax, n)

	// tampilkan hasil
	fmt.Println("Berat balita minimum: ", bMin, "kg")
	fmt.Println("Berat balita maksimum: ", bMax, "kg")
	fmt.Println("Rerata berat balita: ", rerata(arrBerat, n), "kg")
}
