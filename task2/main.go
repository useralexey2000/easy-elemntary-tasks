package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	cardNum(s)
}

// Input example: 4539 1488 0343 6467
// Task2
func cardNum(card string) {
	card = strings.Trim(card, " ")
	arr := strings.Split(card, " ")
	str := strings.Join(arr, "")
	if len(str) != 4*4 {
		fmt.Printf("not valid, %s - not correct lenght of card %d\n", str, len(str))
		return
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("not valid, %v - not a number\n", i)
		return
	}
	res := strings.Join(arr[:3], " ")
	fmt.Println(res)

}
