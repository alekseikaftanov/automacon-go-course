package main

import (
	"fmt"
)

// 5.1 Необходимо создать указатель на строковое значение

func stringAdress(s string) {
	var pointer *string = &s
	fmt.Println("String pointer: ", pointer)
}

// 5.2 Необходимо создать переменную. В консоль вывести её значение и адрес

func valueAndAddress() {
	var a int = 42
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
}

// 5.3 Необходимо создать переменную и указатель на неё. Необходимо через указатель изменить значение переменной

func changeValue() {
	myVal := 42
	myValPointer := &myVal
	*myValPointer = 21
	fmt.Println("Value of myVal: ", myVal)
}

// 5.4 Создать группу переменных, в консоль вывести их значения и адреса. Проанализировать, как отличаются адреса и почему

func groupVariables() {

	myVar1 := "John Doe"
	myVar2 := 24
	MyVar3 := true
	myVar4 := 42.23

	fmt.Println("Value of myVar1: ", myVar1)
	fmt.Println("Address of myVar1: ", &myVar1)
	fmt.Println("Value of myVar2: ", myVar2)
	fmt.Println("Address of myVar2: ", &myVar2)
	fmt.Println("Value of myVar3: ", MyVar3)
	fmt.Println("Address of myVar3: ", &MyVar3)
	fmt.Println("Value of myVar4: ", myVar4)
	fmt.Println("Address of myVar4: ", &myVar4)
}

/* 5.5 Необходимо создать функцию change, которая принимает параметр и изменяет его значение.
В функции main необходимо создать локальную переменную и вызвать change таким образом,
чтобы она изменила значение локальной переменной.*/

func change(a *int) {
	*a = 21
	fmt.Println("Value of a: ", *a)
}

/* 5.6 Необходимо создать пользовательский тип square на базе int.
Далее необходимо создать переменную s типа square и значением 25.
Необходимо вывести значение переменной в консоль*/

func createType() {
	type square int
	s := square(25)
	fmt.Println(s)
}

/* 5.7 Необходимо создать пользовательский тип square на базе int.
Далее необходимо создать переменную s типа square и значением 30.
Значение переменной s необходимо увеличить на 15 и вывести
результат в консоль. */

func increaseSquare() {
	type square int
	s := square(30)
	s += 15
	fmt.Println("Square: ", s)
}

/* 5.8 Необходимо создать пользовательский тип square на базе int. Далее необходимо создать переменную s типа square и значением 34.
Значение переменной s необходимо увеличить на 10 и вывести
результат в консоль. Тип square при выводе в консоль должен
автоматически дописывать м², то есть результат должен
выглядеть: 44 м².*/

func printSquareMeters() {
	type square int
	s := square(34)
	s += 10
	fmt.Println(s, "м²")
}

func main() {
	stringAdress("Hello, Go(5.1)!") // 5.1
	fmt.Println("-----")
	valueAndAddress() // 5.2
	fmt.Println("-----")
	changeValue() // 5.3
	fmt.Println("-----")
	groupVariables() // 5.4
	fmt.Println("-----")
	a := 42
	change(&a) // 5.5
	fmt.Println("-----")
	createType() // 5.6
	fmt.Println("-----")
	increaseSquare() // 5.7
	fmt.Println("-----")
	printSquareMeters() // 5.8
}
