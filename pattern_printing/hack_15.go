package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		for k := 0; k < i; k++ {
			fmt.Print(" ")
		}

		for j := 0; j < n-i; j++ {
			fmt.Print("#")
		}

		fmt.Print("\n")
	}
}
