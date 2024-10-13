package main

import (
	"fmt"
)

func BiayaPos(beratPaket int) (int, int, int) {

	biayaDasarKg := 10000
	totalBiaya := 0

	kg := beratPaket / 1000
	gr := beratPaket % 1000

	biayaPerKg := kg * biayaDasarKg
	totalBiaya += biayaPerKg

	biayaTambahan := 0
	if kg < 10 {
		if gr > 0 && gr < 500 {

			biayaTambahan = gr * 15
		} else if gr >= 500 {

			biayaTambahan = gr * 5
		}
		totalBiaya += biayaTambahan
	}

	return kg, gr, totalBiaya
}

func main() {
	var beratPaket int

	for {

		fmt.Print("Masukkan berat paket (gram): ")
		fmt.Scan(&beratPaket)

		if beratPaket == 0 {
			fmt.Println("Program selesai.")
			break
		}

		kg, gr, totalBiaya := BiayaPos(beratPaket)

		fmt.Printf("Berat parsel (gram): %d\n", beratPaket)
		fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kg*10000, totalBiaya-(kg*10000))
		fmt.Printf("Total biaya: Rp. %d\n\n", totalBiaya)
	}
}