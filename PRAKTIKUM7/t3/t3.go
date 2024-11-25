package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func penghitunganjarak(arr []int) string {
	if len(arr) < 2 {
		return "Data berjarak tidak tetap"
	}

	diff := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return "Data berjarak tidak tetap"
		}
	}

	return fmt.Sprintf("Data berjarak %d", diff)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan sekumpulan bilangan bulat (akhiri dengan bilangan negatif):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	inputs := strings.Fields(input)

	var numbers []int
	for _, val := range inputs {
		num, _ := strconv.Atoi(val)
		if num < 0 {
			break
		}
		numbers = append(numbers, num)
	}

	insertionSort(numbers)

	status := penghitunganjarak(numbers)

	fmt.Println("Hasil setelah pengurutan:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
	fmt.Println(status)
}
