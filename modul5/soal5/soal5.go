package main

import (
	"fmt"
)

func cetakbilangan(n, current int) {
	if current > n {
		return
	}

	if current%2 != 0 {
		fmt.Print(current, " ")
	}

	cetakbilangan(n, current+1)
}

func main() {
	var N int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&N)

	cetakbilangan(N, 1)
	fmt.Println()
}
