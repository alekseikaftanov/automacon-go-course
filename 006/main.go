package main

import (
	"fmt"
)

// 6.1

type Contract struct {
	ID           int
	Number, Date string
}

func main() {
	c := Contract{ID: 1, Number: "«#000A101\t01»", Date: "2024-01-31"}
	fmt.Println("ID: ", c.ID, "Number: ", c.Number, "Date: ", c.Date)

}
