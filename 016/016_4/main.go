package main

import (
	"context"
	"fmt"
)

/*
Задача 16.4 Нужно создать контекст со значениями
some key1: some value1
some key2: some value2
Контекст следует передать в функцию, которая выведет значения
some key1 и some key2 в stdout.
*/

type contextKey string

const ( // Константы для ключей
	key1 contextKey = "some key1"
	key2 contextKey = "some key2"
)

func printValues(ctx context.Context) {
	val1 := ctx.Value(key1) // Получаем значения по ключу
	val2 := ctx.Value(key2)
	fmt.Printf("some key1: %v\n", val1) // Выводим значения
	fmt.Printf("some key2: %v\n", val2)
}

func main() {
	ctx := context.WithValue(context.Background(), key1, "some value1") // Создаем контекст с ключами и значениями
	ctx = context.WithValue(ctx, key2, "some value2")
	printValues(ctx) // Выводим значения
}
