package main

import "fmt"

func main() {
	a := 5

	for i := 1; i <= a; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
