package main

import (
	"fmt"
)

func main() {

	var (
		h    int
		w    int
		chip string
	)
	fmt.Scanf("%d", &h)
	fmt.Scanf("%d", &w)
	fmt.Scanf("%s", &chip)

	chessboard(h, w, chip)
}

// Task 0
func chessboard(h, w int, chip string) {
	a, b := " ", chip
	for i := 0; i < h; i++ {
		a, b = b, a
		for j := 0; j < w; j++ {
			fmt.Print(a, b)
		}
		fmt.Println()
	}
}
