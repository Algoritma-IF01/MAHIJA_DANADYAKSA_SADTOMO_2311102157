package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

// Fungsi untuk mengisi array dengan input dari user
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Teks: ")
	fmt.Scanln(&input)

	*n = len(input)
	for i := 0; i < *n && i < NMAX; i++ {
		(*t)[i] = rune(input[i])
	}
}

// Fungsi untuk mencetak array sebagai string
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

// Fungsi untuk membalikkan isi array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Fungsi untuk memeriksa apakah array merupakan palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Mengisi array dengan input dari user
	isiArray(&tab, &m)

	// Menampilkan teks asli
	fmt.Print("Reverse: ")
	cetakArray(tab, m)

	// Memeriksa apakah teks tersebut palindrom
	if palindrom(tab, m) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}

}
