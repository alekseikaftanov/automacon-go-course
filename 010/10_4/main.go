/*
Задача 10.4
Необходимо создать модуль и залить в публичный репозиторий на https://github.com/. Модуль должен
иметь функцию Hello. Используя теги git,
необходимо создать пару версий модуля: v1.0.0 и v1.1.0.
Разные версии должны выводить разные сообщения:
«Hello, v1.0.0» и «Hello, v1.1.0» соответсвенно.
 Написать программу, которая будет одновременно
 использовать логику версий v1.0.0 и v1.1.0.
Результат должен выводиться в консоль.
*/

// main.go
package main

import (
	"fmt"

	helloV1 "github.com/alekseikaftanov/hello-module v1.0.0"
	helloV2 "github.com/alekseikaftanov/hello-module v1.1.0"
)

func main() {
	fmt.Println("Using version v1.0.0:")
	helloV1.Hello()
	fmt.Println("Using version v1.1.0:")
	helloV2.Hello()
}
