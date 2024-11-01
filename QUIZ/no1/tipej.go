package main

import "fmt"

func main() {
  var input int
  fmt.Print("Masukkan angka: ")
  fmt.Scan(&input)

  digitCount := 0
  temp := input
  for temp > 0 {
    digitCount++
    temp /= 10
  }

  mid := (digitCount + 1) / 2
  leftPart := input / pow(10, digitCount-mid)
  rightPart := input % pow(10, digitCount-mid)

  fmt.Print("Bilangan kiri: ")
  fmt.Println(leftPart)
  fmt.Print("Bilangan kanan: ")
  fmt.Println(rightPart)
  fmt.Print("Hasil penjumlahan bilangan kanan dan kiri: ")
  fmt.Println(leftPart + rightPart)
}

func pow(base, exp int) int {
  result := 1
  for exp > 0 {
    result *= base
    exp--
  }
  return result
}


