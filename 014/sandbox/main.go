package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 1; i < 5; i++ {
		go func(id int) {
			fmt.Printf("goroutine %d start working\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("goroutine %d end working\n", id)
		}(i)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("happy end")
}
