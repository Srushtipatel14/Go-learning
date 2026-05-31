package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("id", id)
}

func main() {
	for i := 0; i < 10; i++ {
		go task(i)
	}

	time.Sleep(time.Second * 2)
}
