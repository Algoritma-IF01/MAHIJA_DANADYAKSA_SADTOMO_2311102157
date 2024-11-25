package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rumahkerabat(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah daerah (n): ")
	nInput, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nInput))

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan banyak rumah di daerah %d: ", i+1)
		mInput, _ := reader.ReadString('\n')
		m, _ := strconv.Atoi(strings.TrimSpace(mInput))

		fmt.Printf("Masukkan nomor rumah di daerah %d (dipisahkan dengan spasi): ", i+1)
		housesInput, _ := reader.ReadString('\n')
		housesStr := strings.Fields(housesInput)

		houses := make([]int, m)
		for j := 0; j < m; j++ {
			houses[j], _ = strconv.Atoi(housesStr[j])
		}

		rumahkerabat(houses)

		fmt.Printf("Hasil daerah %d: ", i+1)
		for _, num := range houses {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
