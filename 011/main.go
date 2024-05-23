package main

import (
	"errors"
	"fmt"
)

/*Задача 11.1
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1».
*/

func someFunction() error {
	err := someOtherFunction()

	if err != nil {
		return fmt.Errorf("Ошибка 2: %w", err)
	}
	return nil
}

func someOtherFunction() error {

	return errors.New("Ошибка 1")
}

/*
Задача 11.2
Нужно создать, используя оборачивание, ошибку
«ошибка3:ошибка2:ошибка1». Из созданной цепочки ошибок нужно
получить ошибку «ошибка2:ошибка1» и вывести в stdout.
*/

func main() {
	// Задача 11.1:
	err := someFunction()
	if err != nil {
		fmt.Println("Ошибка 3:", err)
	}

}
