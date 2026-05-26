package main

import (
	"fmt"
)

func main() {
	age := 1
	switch {
	case age >= 18:
		{
			fmt.Println("you are adult")
		}
	case age >= 13 && age < 18:
		{
			fmt.Println("you are teenager")
		}
	case age < 13:
		{
			fmt.Println("you are child")
		}
	}
}
