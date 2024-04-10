package main

import "fmt"

// 4.1. Функция Hello, GO

func hello() {
	fmt.Println("Hello, GO (4.1)!")
}

//4.2 Анонимная функция

var helloAgain = func() string {

	fmt.Println("Hello, anonymous GO (4.2)!")
	return "Hello, anonymous GO!"
}

/*4.3 Необходимо создать и вызвать функцию «hello», которая в качестве параметра
принимает и вызывает анонимную функцию. Анонимная функция должна выводить в stdout фразу «Hello, Go!».*/

func helloWithParam(anonFunc func()) {
	anonFunc()
}

/* 4.4 Необходимо создать функцию «hello», которая возвращает анонимную функцию.
Необходимо вызвать анонимную функцию, которая в stdout выводит «Hello, Go!». */

func helloAnonimus() func() {
	return func() {
		fmt.Println("Hello, Go (4.4)!")
	}
}

/* 4.5 Необходимо создать и вызвать функцию «hello», которая выводит «Hello, Go!» в stdout.
Также, используя defer необходимо вывести фразу «завершение работы».*/

func helloDeffer() {
	defer fmt.Println("завершение работы (4.5)")
	fmt.Println("Hello, Go (4.5)!")
}

// 4.6 Вывести первые 23 числа Фибоначчи, не используя циклы и максимум один оператор if

func printFibonacci(n int) {
	fibonacciHelper(n, 0, 1)
}

func fibonacciHelper(n, a, b int) {
	if n == 0 {
		return
	}
	fmt.Println(a)
	fibonacciHelper(n-1, b, a+b)
}

func main() {
	hello()      // 4.1.
	helloAgain() // 4.2
	helloWithParam(func() {
		fmt.Println("Hello, Go(4.3)!")
	}) // 4.3
	helloFunc := helloAnonimus()
	helloFunc()        // 4.4
	helloDeffer()      // 4.5
	printFibonacci(23) // 4.6
}
