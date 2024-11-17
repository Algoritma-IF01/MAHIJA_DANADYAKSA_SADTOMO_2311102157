package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung total berat setiap wadah dan rata-rata berat per wadah
func calculateContainerWeightsAndAverage(x, y int, weights []float64) ([]float64, float64) {
	// Hitung jumlah wadah
	numContainers := int(math.Ceil(float64(x) / float64(y)))
	containerWeights := make([]float64, numContainers)

	// Hitung total berat per wadah
	for i, weight := range weights {
		containerWeights[i/y] += weight
	}

	// Hitung rata-rata berat
	total := 0.0
	for _, weight := range containerWeights {
		total += weight
	}
	average := total / float64(numContainers)

	return containerWeights, average
}

func main() {
	var x, y int
	fmt.Print("Masukkan banyak ikan yang akan dijual: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan banyak ikan yang akan dimasukkan ke dalam wadah: ")
	fmt.Scan(&y)

	// Validasi input
	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan harus antara 1 dan 1000, dan kapasitas wadah harus lebih dari 0.")
		return
	}

	weights := make([]float64, x)
	fmt.Printf("Masukkan berat %d ikan (dipisahkan dengan spasi): ", x)
	for i := 0; i < x; i++ {
		if _, err := fmt.Scan(&weights[i]); err != nil || weights[i] < 0 {
			fmt.Println("Berat ikan tidak valid. Masukkan angka positif.")
			return
		}
	}

	// Panggil fungsi untuk menghitung total berat dan rata-rata
	containerWeights, totalAverage := calculateContainerWeightsAndAverage(x, y, weights)

	// Cetak hasil
	fmt.Println("\nTotal berat di setiap wadah:")
	for i, weight := range containerWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, weight)
	}

	fmt.Printf("\nRata-rata berat per wadah: %.2f\n", totalAverage)
}
