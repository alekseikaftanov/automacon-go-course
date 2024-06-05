package main

import (
	"010/10_1/first"
	"010/10_3/second"
	"fmt"
)

/*
Задача 10.3
Необходимо создать и установить пакет second. Пакет должен иметь функцию Hello,
которая возвращает строку «Hello, Louis!».
Далее написать программу, которая будет вызывать функцию Hello пакета
first задачи 8.1 и функцию Hello пакета second. Оба результата должны
выводиться в консоль
*/

func main() {
	fmt.Println(first.Hello())
	fmt.Println(second.Hello())
}
