# <h1 align="center">Laporan Praktikum Modul 2</h1>
<p align="center">Mahija Danadyaksa Sadtomo_2311102157</p>

## A. Hello World

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var greetings string = "Selamat datang di dunia go!"
	var a, b int

	fmt.Println(greetings)
	fmt.Scanln(&a, &b)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}
```
![hello world!](assets/p1.png)

## B. Hipotenusa

```go
package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var hipotenusa bool

	fmt.Print("Masukkan nilai A: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai B: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan nilai C: ")
	fmt.Scanln(&c)
	hipotenusa = (c * c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa segitiga a, b, c: ", hipotenusa)
}
```
![hello world!](assets/p2.png)

## C. Latihan 1

```go
package main

import (
	"fmt"
)

func main() {
	var(
	satu, dua, tiga string
	temp string
	)
	fmt.Print("Masukkan input String: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukkan input String: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukkan input String: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal: ", satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir: ", satu + " " + dua + " " + tiga)
}
```
![hello world!](assets/p3.png)

## D. Latihan 2

```go
package main

import (
	"fmt"
)

func main() {
	var tahun int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)

	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				fmt.Println("Tahun kabisat")
			} else {
				fmt.Println("Bukan tahun kabisat")
			}
		} else {
			fmt.Println("Tahun kabisat")
		}
	} else {
		fmt.Println("Bukan tahun kabisat")
	}
}
```
![hello world!](assets/p4.png)
