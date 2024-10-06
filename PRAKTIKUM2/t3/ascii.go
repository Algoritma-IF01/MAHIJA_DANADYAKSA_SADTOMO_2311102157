package main

import "fmt"

func main() {

	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	
	var kata string
	fmt.Scan(&kata)
	
	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)
	
	if len(kata) == 3 {
		fmt.Printf("%c%c%c\n", kata[0]+1, kata[1]+1, kata[2]+1)
	} else {
		return
	}
}