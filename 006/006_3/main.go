package main

import "fmt"

type Contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	first := Contract{1, "#000A\n101", "2024-01-31"}

	fmt.Println("Договор №", first.Number, "от", first.Date)

}
