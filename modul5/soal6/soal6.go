package main

import (
	"fmt"
)

func nilai(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * nilai(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan angka x dan y: ")
	fmt.Scan(&x, &y)

	result := nilai(x, y)
	fmt.Println("Hasil:", result)
}
