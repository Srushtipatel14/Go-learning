package main

import (
	"fmt"
)

func main() {
	val := make(map[string]interface{})

	val["one"] = 1
	val["two"] = 2.7
	val["three"] = true
	val["four"] = "four"
	val["five"] = []string{"five", "six", "seven"}
	val["six"]=map[string]interface{}{
		"eight": 8,
	}
	fmt.Println(val)

}
