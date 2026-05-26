package main

import (
	"fmt"
)

func add(num ...int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum

}

func main() {
	result := add(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
