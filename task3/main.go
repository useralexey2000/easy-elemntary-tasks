package main

import "fmt"

func main() {
	f := fib()
	for {
		fmt.Scanln()
		fmt.Print(f())
	}
}

// 0, 1, 1, 2, 3, 5
// Task3
func fib() func() int {
	n := 0
	arr := []int{0, 1}
	return func() int {
		n++
		if n == 0 || n == 1 {
			return arr[n-1]
		}
		next := arr[n-2] + arr[n-1]
		arr = append(arr, next)
		return arr[n-1]
	}
}
