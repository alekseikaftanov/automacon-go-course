package main

import (
	"first"
	"fmt"
)

func main() {
	msg := fmt.Sprintf("Hello, my %s package!", first.Hello())
	fmt.Println(msg)
}
