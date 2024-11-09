package main

import (
	"fmt"
)

func printSequence(n, current int) {
	if current > n {
		return
	}

	fmt.Print(current, " ")

	if current > 1 {
		printSequence(n, current-1)
	}

	fmt.Print(current, " ")
}

func main() {
	var N int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&N)

	printSequence(N, N)
	fmt.Println()
}
