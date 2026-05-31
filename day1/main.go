package main

import "fmt"

func printSlice[t int | string | bool](s []t) {

	for _, v := range s {
		fmt.Printf("%v ", v)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	strings := []string{"Hello", "World", "Go", "Programming"}
	printSlice(numbers)
	printSlice(strings)
}
