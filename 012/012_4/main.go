// В представленном ниже коде возникает паника.
// Нужно понять где и объяснить почему.
// Также нужно предложить не менее 3-х вариантов решения проблемы
// с объяснением причины устранения ошибки

package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}
type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	return d.voice
}
func main() {
	var d = &Duck{
		voice: "Кря-кря", // Здесь инициализируем значением структуры
	}
	fmt.Println(d) // Здесь видим, что переменная инициализируется как nil, что и вызывает панику
	song, err := Sing(d)
	if err != nil { // Проверка на nil не работает должным образом
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}
func Sing(b Bird) (string, error) {

	fmt.Println(b)
	if b != nil {
		return "", errors.New("Ошибка пения!") // Исправляем код
	}
	return b.Sing(), nil
}
