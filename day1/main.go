package main

import (
	"fmt"
)

func main() {
	var a = make([]int, 0, 4)

	a = append(a, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)

	fmt.Println(len(a), cap(a))
}
