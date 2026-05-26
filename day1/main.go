package main

import (
	"fmt"
)

func val() func() int {
	count := 0

	return func() int {
		count = count + 1
		return count
	}
}

func main() {
	result := val()
	fmt.Println(result())
	fmt.Println(result())
}
