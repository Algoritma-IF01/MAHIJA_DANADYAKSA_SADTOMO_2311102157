package main

import (
	"fmt"
)

func main() {
	var r float64
	var luas, volume float64

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&r)
	luas = 4 * 3.14 * r * r
	volume = (4.0 / 3.0) * 3.14 * r * r * r
	fmt.Println("Bola dengan jejari ", r, " memiliki volume ", volume, " dan luas kulit ", luas)
}