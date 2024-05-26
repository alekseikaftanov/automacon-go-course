// Следующий код программы должен вывести в консоль сообщения:
// Утка - Умею летать!
// Утка - Умею плавать!
// Воробей - Умею летать!

package main

import "fmt"

type Bird interface {
	Fly()
	Swim()
}
type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}
func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}
func (s Sparrow) Swim() {
	// Фикстивный метод, реализующий интерфейс, но ничего не возвращающий в случае плавающего воробья
}

func main() {
	var d, s Bird
	d = Duck{}
	Do(d)
	s = Sparrow{}
	Do(s)
}
func Do(b Bird) {
	b.Fly()
	b.Swim()
}
