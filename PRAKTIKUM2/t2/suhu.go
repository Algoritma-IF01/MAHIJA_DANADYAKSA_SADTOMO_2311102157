package main

import (
	"fmt"
)

func main() {
	var celcius, fahrenheit, kelvin, reamur float64

	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scanln(&celcius)
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = celcius + 273.15
	reamur = celcius * 4 / 5
	fmt.Println("Suhu dalam fahrenheit: ", fahrenheit)
	fmt.Println("Suhu dalam kelvin: ", kelvin)
	fmt.Println("Suhu dalam reamur: ", reamur)
}