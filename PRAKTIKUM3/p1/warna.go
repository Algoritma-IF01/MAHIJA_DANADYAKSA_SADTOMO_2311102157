package main

import (
	"fmt"
)


func main() {
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	for i := 1; i<=5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d\n", i)
		fmt.Printf("Masukkan warna 1: ")
		fmt.Scanln(&warna1)
		fmt.Printf("Masukkan warna 2: ")
		fmt.Scanln(&warna2)
		fmt.Printf("Masukkan warna 3: ")
		fmt.Scanln(&warna3)
		fmt.Printf("Masukkan warna 4: ")
		fmt.Scanln(&warna4)

		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] || warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {
			hasil = false
			break
		}
	} 
	 //tampilkan hasil
	if hasil {
		fmt.Println("Selamat, anda berhasil!", hasil)
	} else {
		fmt.Println("Anda gagal!")
	}
}
