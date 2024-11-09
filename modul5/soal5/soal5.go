package main

import (
	"fmt"
)

func printOdd(n, current int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Print(current, " ")
	}

	printOdd(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&N)

	printOdd(N, 1)
	fmt.Println()
}
