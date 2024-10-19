package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hitungSkor(soal *int, skor *int, waktu []int) {
	*soal = 0
	*skor = 0
	for _, w := range waktu {
		if w < 301 {
			*soal++
			*skor += w
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pemenang string
	var maxSoal, minWaktu int
	maxSoal = 0
	minWaktu = 999999

	for {
		fmt.Println("Masukkan nama dan waktu pengerjaan soal (ketik 'Selesai' untuk mengakhiri input):")
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		data := strings.Fields(input)
		if len(data) != 9 {
			fmt.Println("Input tidak valid, masukkan nama dan 8 nilai waktu pengerjaan soal.")
			continue
		}

		nama := data[0]
		waktu := make([]int, 8)
		for i := 0; i < 8; i++ {
			fmt.Sscanf(data[i+1], "%d", &waktu[i])
		}

		var soal, skor int
		hitungSkor(&soal, &skor, waktu)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}

	if pemenang != "" {
		fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta.")
	}
}
