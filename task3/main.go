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
	a, b := 0, 1
	return func() int {
		res := a
		a, b = b, a+b
		return res
	}
}
