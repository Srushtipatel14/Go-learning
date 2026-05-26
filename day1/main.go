package main

import (
	"fmt"
)

func main() {
	age := 15
	if age >= 18 {
		fmt.Println("you are adult")
	} else if age < 10 && age >= 0 {
		fmt.Println("you are child")
	} else {
		fmt.Println("you are teenager")
	}
}
