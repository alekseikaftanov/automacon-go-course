package main

import (
	"fmt"
)

func bankPercent() {
	var ammount int
	fmt.Print("Введите сумму ")
	fmt.Scanf("%d", &ammount)

	var percent int
	fmt.Print("Введите банковский процент ")
	fmt.Scanf("%d", &percent)

	var qurency string
	fmt.Print("Введите валюту ")
	fmt.Scanf("%s", &qurency)

	var monthIncome, yearIncome int // Переменные объявлены здесь, чтобы они были видны во всей функции main()

	if qurency == "Тенге" || qurency == "Тэнге" {
		monthIncome = (ammount * percent) / 100 * 85 / 100 / 12 // Убрал 0.85, умножил на 85 и разделил на 100
		yearIncome = (ammount * percent) / 100 * 85 / 100       // Убрал 0.85, умножил на 85 и разделил на 100
		fmt.Printf("Если положить вклад в банк Казахстана под %d процентов, то в месяц будет %d %s, а в год %d %s.", percent, monthIncome, qurency, yearIncome, qurency)
	} else {
		monthIncome = (ammount * percent) / 100 / 12
		yearIncome = (ammount * percent) / 100
		fmt.Printf("Если положить вклад в банк под %d процентов, то в месяц будет %d %s, а в год %d %s.", percent, monthIncome, qurency, yearIncome, qurency)
	}
}

// func main() {
// 	bankPercent()
// }
