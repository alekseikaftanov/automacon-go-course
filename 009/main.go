package main

import "fmt"

/*
Задача 9.1
Необходимо написать функцию fruitMarket, которая будет принимать
название фруктов, а возвращать их количество. Например,
«апельсины» - 5. Сама функция должна создать карту: апельсин = 5,
яблоки = 3, сливы = 1, груши = 0. Если запрашиваемых фруктов нет в
карте, то в консоль должно выводится сообщение об отсутствии.
*/

func fruitMarket(fruit string) {

	var fruits = map[string]int{
		"Oranges": 5,
		"Apples":  3,
		"Plums":   1,
		"Pears":   0,
	}

	if quantity, ok := fruits[fruit]; ok {
		fmt.Println(quantity)
	} else {
		fmt.Println("There is no such fruit!")
	}
}

/*
Задача 9.2
Необходимо создать срез целых чисел: 1, 2, 3. Далее создать 4
вложенных цикла. Каждый цикл должен в консоль выводить текущее
значение среза. Причём внутренние срезы должны содержать отступы
для облегчения визуального восприятия. Внутренний срез на ключе 1
должен остановить все циклы, начиная со второго цикла.
*/

/*
Задача 9.3
Необходимо написать функцию checkFood, которая принимает в
качестве параметра название еды, а в консоль выводит «это фрукт»,
«это овощ». Для проверки необходимо использовать оператор выбора.
Причём проверке подлежат: "груша", "яблоко", "апельсин", "тыква",
"огурец", "помидор". Если checkFood получит название еды, которое не
входит в этот список, то в консоли должно появиться сообщение: «что-
то странное…». Для вывода информации в консоль необходимо
использовать только fmt.Println, но не более трёх раз.
*/

func checkFood(food string) {
	switch food {
	case "Pear", "Apple", "Orange":
		fmt.Printf("%s is fruit", food)

	case "Pumpkin", "Cucumber", "Tomato":
		fmt.Printf("%s is vegetable", food)

	default:
		fmt.Println("Something went wrong")
	}
}

// Нужно понять почему эта версия не работает так, как нужно.

// func checkFood() {
// 	foods := []string{"Pear", "Apple", "Orange", "Pumpkin", "Cucumber", "Tomato"}

// 	var userInput string
// 	fmt.Println("Please fill your food name:")
// 	_, err := fmt.Scanf("%s", &userInput)
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	for _, food := range foods {

// 		if food == userInput && (food == "Pear" || food == "Apple" || food == "Orange") {
// 			fmt.Printf("%s is fruit", userInput)
// 			return
// 		}
// 		for _, food := range foods {
// 			if food == "Pumpkin" || food == "Cucumber" || food == "Tomato" {
// 				fmt.Println("This is veg")
// 				return
// 			}
// 		}

// 	}
// 	fmt.Println("Something went wrong")
// }

func main() {
	fruitMarket("Plums") // Задача 9.1
	// loop() // Задача 9.2
	checkFood("Apple")
	checkFood("Pumpkin")
	checkFood("Heizenberg") // Задача 9.3
}
