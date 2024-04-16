package main

import (
	"fmt"
)

// 6.1

type Contract struct {
	ID           int
	Number, Date string
}

/*Задача 6.2  сполями:
ID int, Number string, Date string. Далее создать экземпляр структуры со значениями
полей: ID=1, Number=«#000A101\t01», Date=«2024-01-31».
Вконсольнужновывестиструктурутакимобразом, чтобыданныеотображалисьввиде: {ID:1 Number:#000A10101 Date:2024-01-31}Задача

6.3 сполями: ID int, Number string, Date string. Далее создать экземпляр структуры со значениями полей: ID=1, Number=«#000A\n101», Date=«2024-01-31». При29
передачи экземпляра структуры в fmt.Println в   консоли должно отображаться: ДоговорNo #000A\n101 от 2024-01-31Задача

6.4Необходимоубратьповторяющийсякод - поля Addss и Phone изструктур:

Послепроведения
рефакторинга
строка fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone) должна выводить
в консоль адрес и телефон пользователя и сотрудника соответственно*/

// 6.2 Необходимо объявить локальную структуру contract

func printStruct() {
	type Contract struct {
		ID           int
		Number, Date string
	}
	c := Contract{ID: 1, Number: "«#000A101\t01»", Date: "2024-01-31"}
	fmt.Println("ID: ", c.ID, "Number: ", c.Number, "Date: ", c.Date)
}

//6.3 Необходимо объявить глобальную структуру contract

type Contract2 struct {
	ID           int
	Number, Date string
}

// 6.4

type user struct {
	ID      int
	Name    string
	Address string
	Phone   string
}
type employee struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

func main() {
	c := Contract{ID: 1, Number: "«#000A101\t01»", Date: "2024-01-31"}

	fmt.Println("ID: ", c.ID, "Number: ", c.Number, "Date: ", c.Date) //6.1
	printStruct()
	c2 := Contract2{ID: 1, Number: "«#000A\n101»", Date: "2024-01-31"} // 6.2
	fmt.Println("Договор No ", c2.Number, "от: ", c2.Date)

	u := user{ID: 1, Name: "Jhon Doe", Address: "New Hempshire", Phone: "790123232"}
	e := employee{ID: 1, Name: "Jhon Doe", Address: "New Hempshire", Phone: "790123232"}
	fmt.Println(u.Address, u.Phone, e.Address, e.Phone)

}
