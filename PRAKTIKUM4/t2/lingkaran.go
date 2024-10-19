package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi untuk memeriksa apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2 float64 // Pusat dan radius lingkaran 1 dan 2
	var x, y float64                      // Titik sembarang

	// Input nilai untuk lingkaran 1 (pusat dan radius)
	fmt.Println("Masukkan pusat (cx1, cy1) dan radius (r1) untuk lingkaran 1:")
	fmt.Scanf("%f %f %f", &cx1, &cy1, &r1)

	// Input nilai untuk lingkaran 2 (pusat dan radius)
	fmt.Println("Masukkan pusat (cx2, cy2) dan radius (r2) untuk lingkaran 2:")
	fmt.Scanf("%f %f %f", &cx2, &cy2, &r2)

	// Input nilai untuk titik sembarang (x, y)
	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scanf("%f %f", &x, &y)

	// Memeriksa apakah titik berada di dalam lingkaran 1 atau 2
	dalamLingkaran1 := diDalam(cx1, cy1, r1, x, y)
	dalamLingkaran2 := diDalam(cx2, cy2, r2, x, y)

	// Menentukan posisi titik
	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
