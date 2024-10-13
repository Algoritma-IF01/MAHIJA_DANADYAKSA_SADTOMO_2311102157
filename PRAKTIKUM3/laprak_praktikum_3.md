# <h1 align="center">Laporan Praktikum Modul 1 - Hello World</h1>
<p align="center">Mahija Danadyaksa Sadtomo_2311102157</p>

## A. Warna Kimia

```go
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
```
![hello world!](assets/p1.png)

## B. Pita Bunga

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner:= bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input
		} else {
			pita = pita + " - " + input
		}
		bungaCount++
	}

	fmt.Println("Pita bunga: ", pita)
	fmt.Println("Jumlah bunga: ", bungaCount)
}
```
![hello world!](assets/p2.png)


## C. Berat Karung (Perulangan no. 3)

```go
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
```
![hello world!](assets/t1.png)

## D. Fungsi K (Perulangan no. 4)

```go

package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(k)
func f(k float64) float64 {
	numerator := math.Pow((4*k + 2), 2)
	denominator := (4*k + 1) * (4*k + 3)
	return numerator / denominator
}

// Fungsi untuk menghitung akar dari 2 dengan jumlah iterasi K
func sqrt2(k int) float64 {
	result := 1.0
	for i := 0; i <= k; i++ {
		result *= f(float64(i))
	}
	return result
}

func main() {
	var K int

	
		// Input nilai K
		fmt.Print("Nilai K = ")
		fmt.Scan(&K)

		// // Menghitung f(K)
		fk := f(float64(K))
		fmt.Printf("Nilai f(K) = %.10f\n\n", fk)

		for i := 1; i <= 3; i++ {

		fmt.Print("Nilai K = ")
		fmt.Scan(&K)
		// Menghitung aproksimasi sqrt(2) dengan K iterasi
		approxSqrt2 := sqrt2(K)
		fmt.Printf("Nilai akar 2 = %.10f\n\n", approxSqrt2)
	}

	fmt.Println("Proses selesai.")
}
```
![hello world!](assets/t2.png)

## E. Biaya Pos (Percabangan no. 1)

```go
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
```
![hello world!](assets/t3.png)

## F. Nilai Akhir Mata Kuliah (Percabangan no. 2)

```go
package main

import (
	"fmt"
)

func main () {
	var nam float64
	var nmk string
	
	for {
		if nam == -1 {
			fmt.Println("Program Selesai")
			break
		}
	fmt.Print("Nilai Akhir Mata Kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {	
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else
	{
		nmk = "E"
	}
	fmt.Println("Nilai Akhir Mata Kuliah: ", nmk)
}
}
```
Jawablah pertanyaan-pertanyaan berikut:
1. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesual spesifikasi soal? <br/>
![hello world!](assets/t4(1).png)
Jawab: <br/>
Tidak sesuai, karena output dari program tersebut adalah D seharusnya "A" karena nam > 80.

2. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya! <br/>
Jawab: <br/>
![hello world!](assets/t4(2).png)
Kesalahannya terdapat pada logika if statements. Karena if dievaluasi secara independen, ketika nilai nam lebih dari satu kondisi, kondisi terakhir yang dievaluasi akan mengeset nilai nmk. Misal, user menginputkan nilai 80.1, kondisi itu menjadi benar untuk "B", "BC", "C", dan kondisi terakhir "D". Oleh karena itu nmk nya menjadi "D". <br/>
Untuk mengatasi hal itu perlu menggunakan else if sehingga ketika suatu kondisi benar, kondisi lainnya tidak akan diperiksa lagi. <br/>

3. Perbaiki program tersebut! Ujilah dengan masulkanı: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah "A", "B", dan D <br/>
Jawab: <br/>
![hello world!](assets/t4(3).png)

## G. Bilangan Prima dan Faktor b(Percabangan no. 3)

```go
package main

import (
	"fmt"
)

func faktor(a int) []int {
	var faktor []int
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func prima(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; i < a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for {
		var bil int
		fmt.Print("Bilangan: ")
		fmt.Scanln(&bil)
		
		fmt.Println("Faktor: ", faktor(bil))
		fmt.Println("Prima: ", prima(bil))
		fmt.Println()
		
		if bil == -1 {
			fmt.Println("Program Selesai")
			break
		}
	}
}
```
![hello world!](assets/t5.png)