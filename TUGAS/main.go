package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var pilihanMenu int

	fmt.Println("")
	fmt.Println("-=-=-=- Welcome to Jordan's W3Schools Practice -=-=-=-")
	fmt.Println("")
	fmt.Println("Menu:")
	fmt.Println("1. Syntax")
    fmt.Println("2. Comments 1")
    fmt.Println("3. Comments 2")
    fmt.Println("4. Variables 1")
    fmt.Println("5. Variables 2")
    fmt.Println("6. Variables 3")
    fmt.Println("7. Variables 4")
    fmt.Println("8. Data Types")
    fmt.Println("9. Array 1")
    fmt.Println("10. Array 2")
    fmt.Println("11. Array 3")
    fmt.Println("12. Operator 1")
    fmt.Println("13. Operator 2")
    fmt.Println("14. Operator 3")
    fmt.Println("15. Conditions 1")
    fmt.Println("16. Conditions 2")
    fmt.Println("17. Conditions 3")
    fmt.Println("18. Conditions 4")
    fmt.Println("19. Switch 1")
    fmt.Println("20. Switch 2")
    fmt.Println("21. Loops 1")
    fmt.Println("22. Loops 2")
    fmt.Println("23. Loops 3")
    fmt.Println("24. Loops 4")
    fmt.Println("25. Functions 1")
    fmt.Println("26. Functions 2")
    fmt.Println("27. Functions 3")
    fmt.Println("28. Functions 4")

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihanMenu)

	reader := bufio.NewReader(os.Stdin) // Membuat reader baru dimana Stdin adalah input dari terminal
	reader.ReadString('\n') // Membuang karakter newline

	fmt.Println("")

	switch pilihanMenu {
    case 1:
        syntax1()
    case 2:
        comments1()
    case 3:
        comments2()
    case 4:
        variables1()
    case 5:
        variables2()
    case 6:
        variables3()
    case 7:
        variables4()
    case 8:
        dataTypes1()
    case 9:
        arrays1()
    case 10:
        arrays2()
    case 11:
        arrays3()
    case 12:
        operators1()
    case 13:
        operators2()
    case 14:
        operators3()
    case 15:
        conditions1()
    case 16:
        conditions2()
    case 17:
        conditions3()
    case 18:
        conditions4()
    case 19:
        switch1()
    case 20:
        switch2()
    case 21:
        loops1()
    case 22:
        loops2()
    case 23:
        loops3()
    case 24:
        loops4()
    case 25:
        functions1()
    case 26:
        functions2()
        functions2()
    case 27:
        functions3("John")
    case 28:
        functions4((3))
	default:
		fmt.Println("Huh? What number did you just type?")
	}
}