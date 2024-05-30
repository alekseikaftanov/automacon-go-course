package main

import "fmt"

/*
Необходимо создать канал с буфером 4, записать в него 2 значения:
«Привет», «буферизованный канал!».
Далее необходимо прочитать все значения из канала и вывести в stdout
*/

func main() {
	ch := make(chan string, 8)
	ch <- "Привет"
	ch <- "буферизованный канал!"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
