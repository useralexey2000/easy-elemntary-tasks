package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	palindrome(n)
}

// Task 4
// Input 1234437
func palindrome(n int) {
	s := strconv.Itoa(n)
	found := false
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			found = true
			start, end := i-1, i
			for j, k := start, end; j > 0 && k < len(s); j, k = j-1, k+1 {
				if s[j] == s[k] {
					start, end = j, k
					continue
				}
				res, _ := strconv.Atoi(s[start : end+1])
				fmt.Println(res)
				break
			}
		}
	}
	if !found {
		fmt.Println(0)
	}
}
