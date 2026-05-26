package main

import (
	"fmt"
)

func printVal(val *int) {
	*val = 20
	fmt.Println(*val)
}

func main() {
	val := 200
	printVal(&val)
	fmt.Println(val)
}
