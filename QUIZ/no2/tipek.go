package main

import (
	"fmt"
)

func main() {
	jumlahTebak, jumlahCendol, jumlahLamang := 0, 0, 0
	var n int
	fmt.Print("Masukkan kartu: ")
	fmt.Scan(&n)
	
	fmt.Print("Masukkan angka pada kartu: ", "\n")
	for i := 0; i < n; i++ {
		var kartu int
		fmt.Scan(&kartu)


		if tebak(kartu) {
			fmt.Println("Es Tebak")
			jumlahTebak++
		} else if cendol(kartu) {
			fmt.Println("Es Cendol")
			jumlahCendol++
		} else {
			fmt.Println("Lamang")
			jumlahLamang++
		}
	}

	fmt.Printf("Es Tebak: %d\n", jumlahTebak)
	fmt.Printf("Es Cendol: %d\n", jumlahCendol)
	fmt.Printf("Lamang: %d\n", jumlahLamang)
}


func tebak(kartu int) bool {
	digit := kartu % 10
	if digit%2 == 0 { 
		return false
	}

	for kartu > 0 {
		if kartu%10 != digit || kartu%2 == 0 { 
			return false
		}
		kartu /= 10
	}
	return true
}


func cendol(kartu int) bool {
	for kartu > 0 {
		if kartu%2 != 0 { 
			return false
		}
		kartu /= 10
	}
	return true
}