package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku []Buku

// Prosedur DaftarkanBuku
func tambahbuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		*pustaka = append(*pustaka, buku)
	}
}

// Prosedur CetakTerfavorit
func favoritBuku(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku di pustaka.")
		return
	}
	maxRating := pustaka[0].rating
	for _, buku := range pustaka {
		if buku.rating > maxRating {
			maxRating = buku.rating
		}
	}
	fmt.Println("Buku dengan rating tertinggi:")
	for _, buku := range pustaka {
		if buku.rating == maxRating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
		}
	}
}

// Prosedur UrutBuku (menggunakan insertion sort)
func UrutBuku(pustaka *DaftarBuku) {
	n := len(*pustaka)
	for i := 1; i < n; i++ {
		key := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < key.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

// Prosedur Cetak5Terbaru
func ratingTertinggiBuku(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < len(pustaka); i++ {
		buku := pustaka[i]
		fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
	}
}

// Prosedur CariBuku
func CariBuku(pustaka DaftarBuku, r int) {
	found := false
	fmt.Printf("Buku dengan rating %d:\n", r)
	for _, buku := range pustaka {
		if buku.rating == r {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	tambahbuku(&pustaka, n)
	fmt.Println("\nData Buku Terfavorit:")
	favoritBuku(pustaka)

	UrutBuku(&pustaka)
	fmt.Println("\n5 Buku dengan Rating Tertinggi setelah diurutkan:")
	ratingTertinggiBuku(pustaka)

	fmt.Print("\nMasukkan rating yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
