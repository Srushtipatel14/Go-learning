package main

import "fmt"

func main() {
	var num int
	fmt.Scan((&num))
	var code int = 65
	for i := 0; i <= num-1; i++ {

		// for j := num - 1; j > i; j-- {
		// 	fmt.Print(" ")
		// }
		for j := 0; j <= i; j++ {
			val := code + num - 1 + j - i
			fmt.Printf("%c", val)
		}
		// for j := i - 1; j >= 0; j-- {
		// 	fmt.Printf("%c", code+j)
		// }
		fmt.Print("\n")
	}

	// for i := num - 1; i >= 0; i-- {
	// 	// for j := 0; j < num-i-1; j++ {
	// 	// 	fmt.Print(" ")
	// 	// }
	// 	// for j := i; j >= 0; j-- {
	// 	// 	fmt.Print("*")
	// 	// }
	// 	for j := i - 1; j >= 0; j-- {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Print("\n")
	// }
}
