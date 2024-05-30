package main

import (
	"fmt"
	"time"
)

/*
Задача 14.1 Необходимо создать дочернюю горутину,
которая выведет в stdout Привет из дочерней горутины!
*/

func main() {
	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена")
}
