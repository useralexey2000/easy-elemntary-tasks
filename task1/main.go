package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//  input 30,-1,-6,90,-6,22,52,123,2,35,6
	var s string
	fmt.Scanf("%s", &s)
	countPosEven(s)
}

// Task 1
func countPosEven(str string) {
	strs := strings.Split(str, ",")
	ints := make([]int, len(strs))

	for i, s := range strs {
		ints[i], _ = strconv.Atoi(s)
	}

	count := 0
	for _, v := range ints {
		if v > 0 && v&1 == 0 {
			count++
		}
	}
	fmt.Println(count)
}
