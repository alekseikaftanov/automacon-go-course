package main

import "fmt"

func main() {

	var divident, dividor int
	fmt.Println("Введите делимое")
	fmt.Scanf("%d", &divident)

	fmt.Println("Введите делитель")
	fmt.Scanf("%d", &dividor)

	result := divident / dividor
	remainder := divident % dividor

	fmt.Printf("Результат: %d, остаток от деления %d, тип результата: %T\n", result, remainder, result)

}
