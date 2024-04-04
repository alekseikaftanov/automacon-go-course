package main

import "fmt"

// Задача 3.1 и 3.2 Необходимо создать глобальную константу и вывести её значение в консоль

func showConst() {
	const myConst = 10
	fmt.Println(myConst)
}

// Задача 3.3. Необходимо создать ukj,fkmye. константу и затенить ее значение

const myConst = 10

func changeConstantValue() {

	const myConst = 20
	fmt.Println(myConst)

}

// Задача 3.4. и 3.5 Создание группы глобальных констант

func showConstantsGroup() {
	const (
		myConst1 = 1
		myConst2 = 2
		myConst3 = 3
		myConst4 = 4
		myConst5 = 5
	)
	fmt.Println(myConst1, myConst2, myConst3, myConst4, myConst5)
}

// Задача 3.6. Приведение типизированной константы к значению float

func changeConstantType() {
	const n int = 5
	m := 3.4 + float64(n)
	fmt.Println(m)
}

// Задача 3.7. Генератор Iota

func iotaGenerator() {
	const (
		myConst1 = 1 << iota * 2
		myConst2
		myConst3
		myConst4
	)

	fmt.Println(myConst1, myConst2, myConst3, myConst4)

}

func main() {
	showConst()           // Задача 3.1 и 3.2
	changeConstantValue() // Задача 3.1
	showConstantsGroup()  // Задача 3.4
	changeConstantType()  // Задача 3.5 и 3.6
	iotaGenerator()       // Задача 3.7
}
