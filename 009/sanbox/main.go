package main

import "fmt"

// func getNumber() {

// 	var num int
// 	var userInput string
// 	fmt.Println("Enter value from 1 to 10")
// 	fmt.Scanf("%s", &userInput)

// 	num, err := strconv.Atoi(userInput)

// 	if err != nil {
// 		fmt.Printf("Value is incorrect")
// 		return
// 	} else if num < 0 {
// 		fmt.Printf("Number is incorrect, use number more than 0")
// 		return
// 	} else if num > 10 {
// 		fmt.Printf("Number is incorrect, use number less than 10")
// 		return
// 	} else {
// 		fmt.Printf("Value %d is correct", num)
// 	}

// }

// func valid(v string) error {

// 	num, err := strconv.Atoi(v)
// 	if err != nil {
// 		return fmt.Errorf("Value is incorrect:  %s", v)
// 	}
// 	if num < 0 {
// 		return fmt.Errorf("Введённое значение меньше 0: %d", num)
// 	}
// 	if num > 10 {
// 		return fmt.Errorf("Введённое значение больше 10: %d", num)
// 	}
// 	return nil
// }

// func getNumber() {

// 	var inputValue string

// 	fmt.Println("Enter number from 1 to 10")
// 	fmt.Scanf("%s\n", &inputValue)

// 	if err := valid(inputValue); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Printf("Значение %s корректно.", inputValue)
// 	}

// }

func main() {
	// getNumber()

	const myConst = 10

	var a int = (1 + 1) * myConst

	fmt.Print("Hello world!")
	fmt.Print(a)

}
