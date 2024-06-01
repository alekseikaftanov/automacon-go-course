package main

import "fmt"

// // Интерфейсы: получение информации о объекте
// type Item interface {
// 	GetTitle() string
// 	GetAuthor() string
// 	GetYear() int
// }

// type Borrowable interface {
// 	Borrow() bool
// 	Return() bool
// 	IsBorrowed() bool
// }

// type Periodical interface {
// 	isPeriodical() bool
// }

// // Структура и методы для книги
// type Book struct {
// 	title    string
// 	author   string
// 	year     int
// 	borrowed bool
// }

// func (b *Book) GetTitle() string {
// 	return b.title
// }

// func (b *Book) GetAuthor() string {
// 	return b.author
// }

// func (b *Book) GetYear() int {
// 	return b.year
// }

// func (b Book) IsBorrowed() bool {
// 	return b.borrowed
// }

// Borrow

// func (b *Book) Borrow() bool {
// 	if b.borrowed {
// 		return false
// 	}
// 	b.borrowed = true
// 	return true
// }

// // Return
// func (b *Book) Return() bool {
// 	if !b.borrowed {
// 		return false
// 	}
// 	b.borrowed = false
// 	return true
// }

func main() {
	a := []int{0, 1, 2, 3, 4}
	fmt.Printf("a=%v, len=%d, cap=%d\n ", a, len(a), cap(a)) //a=[0 1 2 3 4], len=5, cap=5

	b := a[1:2]
	fmt.Printf("b=%v, len=%d, cap=%d\n ", b, len(b), cap(b))

	c := append(a, 5)
	fmt.Printf("c=%v, len=%d, cap=%d\n ", c, len(c), cap(c))

	b = append(b, 100)

	fmt.Printf("a=%v, len=%d, cap=%d\n ", a, len(a), cap(a))
}
