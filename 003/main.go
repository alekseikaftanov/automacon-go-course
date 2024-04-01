package main

import "fmt"

// Задача 3.1. Необходимо создать глобальную константу и вывести её значение в консоль

// const myConst = 10

// func main() {
// 	fmt.Println(myConst)
// }

// Задача 3.2. Необходимо создать локальную константу и вывести её значение в консоль

// func main() {
// 	const myConst = 10
// 	fmt.Println(myConst)
// }

// Задача 3.3. Необходимо создать ukj,fkmye. константу и затенить ее значение

// const myConst = 10

// func main() {
// 	const myConst = 20
// 	fmt.Println(myConst)
// }

// Задача 3.4. Создание группы глобальных констант

// const (
// 	myConst1 = 1
// 	myConst2 = 2
// 	myConst3 = 3
// 	myConst4 = 4
// 	myConst5 = 5
// )

// func main() {
// 	fmt.Println(myConst1, myConst2, myConst3, myConst4, myConst5)
// }

// Задача 3.5. Создание группы локальных констант

// func main() {
// 	const (
// 		myConst1 = 1
// 		myConst2 = 2
// 		myConst3 = 3
// 		myConst4 = 4
// 		myConst5 = 5
// 	)
// 	fmt.Println(myConst1, myConst2, myConst3, myConst4, myConst5)
// }

// Задача 3.6. Приведение типизированной константы к значению float

// func main() {
// 	const n int = 5
// 	m := 3.4 + float64(n)
// 	fmt.Println(m)
// }

// Задача 3.7. Генератор Iota

func main() {

	const (
		myConst1 = 1 << iota * 2
		myConst2
		myConst3
		myConst4
	)

	fmt.Println(myConst1, myConst2, myConst3, myConst4)

}
