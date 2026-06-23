package main

import "fmt"

func main() {
	var num int
	fmt.Scan((&num))

	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(num-j, " ")
		}
		for j := num - 1; j > i; j-- {
			fmt.Print(num-i, " ")
		}
		for j := num - 1; j > i; j-- {
			fmt.Print(num-i, " ")
		}
		for j := i; j >= 0; j-- {
			fmt.Print(num-j, " ")
		}
		fmt.Print("\n")
	}
	for i := 0; i < num-1; i++ {

		for j := num - 1; j > i; j-- {
			fmt.Print(j+1, " ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print(i+2, " ")
		}
		for j := i; j >= 0; j-- {
			fmt.Print(i+2, " ")
		}
		for j := num - 1; j > i; j-- {
			fmt.Print(num-j+i+1, " ")
		}
		fmt.Print("\n")
	}

}
