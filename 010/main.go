package main

import (
	"fmt"
	first "myfirst/010"
)

func main() {
	msg := fmt.Sprintf("Hello, my %s package!", first.Hello())
	fmt.Println(msg)
}
