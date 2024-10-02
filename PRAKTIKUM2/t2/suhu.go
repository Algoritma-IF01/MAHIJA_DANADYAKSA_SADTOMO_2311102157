package main

import (
	"fmt"
)

func main() {
	var celcius, fahrenheit, kelvin float64

	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scanln(&celcius)
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = celcius + 273.15
	fmt.Println("Suhu dalam fahrenheit: ", fahrenheit)
	fmt.Println("Suhu dalam kelvin: ", kelvin)
}