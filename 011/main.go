package main

import (
	"errors"
	"fmt"
)

/*Задача 11.1
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1».
*/

func createThreeErrors() error {
	// Создаем первую ошибку
	err1 := errors.New("ошибка1")
	// Оборачиваем ее в ошибку 2
	err2 := fmt.Errorf("ошибка2:%w", err1)
	// Оборачиваем ее в ошибку 3
	err3 := fmt.Errorf("ошибка2:%w", err2)
	// Отображаем ошибку
	fmt.Println(err3)
	return err3
}

/*
Задача 11.2
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Из созданной цепочки ошибок нужно
получить ошибку «ошибка2:ошибка1» и вывести в stdout.
*/

func createAndExtractErrors() {
	// Создаем первую ошибку
	err1 := errors.New("ошибка 1")
	// Оборачиваем ее в ошибку 2
	err2 := fmt.Errorf("ошибка 2: %w", err1)
	// Оборачиваем ее в ошибку 3
	err3 := fmt.Errorf("ошибка 3: %w", err2)

	// Выводим полную цепочку ошибок
	fmt.Println("Полная цепочка ошибок:", err3)

	// Разворачиваем ошибки до ошибки2:ошибка1
	extractedErr := errors.Unwrap(err3)
	if extractedErr != nil {
		fmt.Println("Result:", extractedErr)
	}
}

/*
11.3
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Не используя unwrap, нужно определить была ли ошибка «ошибка1»
*/

func isErrorHappens() {
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2: %w", err1)
	err3 := fmt.Errorf("ошибка2: %w", err2)
	targetError := err1
	if errors.Is(err3, targetError) {
		fmt.Println("ошибка 1 есть в цепочке ошибок")
	} else {
		fmt.Println("ошибка 1 не присутсвует в цепочке ошибок")
	}
}

/*
11.5
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Также нужно создать свою ошибку в виде структуры myFirstError,
которая обязательно должна иметь метод Error() string.
Необходимо убедиться, что в созданной цепочке ошибок не было ошибки типа myFirstError
*/

type myFirstError struct {
	Message string
}

func (e *myFirstError) Error() string {
	return e.Message
}

func myFirstErrorFunc() {
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2: %w", err1)
	err3 := fmt.Errorf("ошибка2: %w", err2)

	targetError := &myFirstError{Message: "Моя ошибка"}

	if !errors.As(err3, &targetError) {
		fmt.Println("Ошибка отсутвует в цепочке ошибок")
	} else {
		fmt.Println("Ошибка отсутвует в цепочке ошибок")
	}

}

func main() {

	// createThreeErrors() // Задача 11.1
	// createAndExtractErrors() // Задача 11.2
	// isErrorHappens() // Задача 11.3
	myFirstErrorFunc() // Задача 11.4
}
