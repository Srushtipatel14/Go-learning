package main

import (
	"fmt"
)

type val struct {
	x int
	y int
	z string
}

func main() {
	p := val{x: 1, y: 2, z: "ten"}
	q := val{x: 1, y: 2, z: "ten"}
	q.x = 10

	fmt.Println(p, q)

}
