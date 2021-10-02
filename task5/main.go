package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		min int
		max int
	)
	fmt.Scanf("%d", &min)
	fmt.Scanf("%d", &max)
	luckyTicket(min, max)
}

// Task 5

//Input  Min: 120123
//       Max: 320320
func luckyTicket(min, max int) {
	easy, hard := 0, 0

	for i := min; i <= max; i++ {
		bt := strconv.Itoa(i)

		right := (bt[0] ^ 48) + (bt[1] ^ 48) + (bt[2] ^ 48)
		left := (bt[3] ^ 48) + (bt[4] ^ 48) + (bt[5] ^ 48)

		if right == left {
			easy++
		}
		var even byte
		var odd byte
		for j := 0; j < len(bt); j++ {
			if bt[j]&1 == 1 {
				odd += (bt[j] ^ 48)
			} else {
				even += (bt[j] ^ 48)
			}
		}

		if even == odd {
			hard++
		}

	}
	fmt.Println("EasyFormula: ", easy)
	fmt.Println("HardFormula: ", hard)
}
