// Не изменяя структуры Xml, Csv и функцию main необходимо
// доработать следующий код так, чтобы в консоли увидели:
// Данные в формате xml
// Данные в формате csv

package main

import "fmt"

// Добавляем интерфейс Formatter
type Formatter interface {
	Format()
}

//
type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}
func Report(x Formatter) {
	x.Format()
}

func ReportCsv(c Formatter) {
	c.Format()
}
func main() {
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}
