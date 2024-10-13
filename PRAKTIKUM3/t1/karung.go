package main

import (
	"fmt"
	"math"
)

func main() {
	var kiri, kanan float64

	for {

		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)

		if kiri < 0 || kanan < 0 ||   kiri+kanan > 150 {
			fmt.Println("Proses selesai.")
			break
		}
		
		selisih := math.Abs(kiri - kanan)
		oleng := selisih >= 9

			fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", oleng)
	}
}
