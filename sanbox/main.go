package main

// Интерфейсы: получение информации о объекте
type Item interface {
	GetTitle() string
	GetAuthor() string
	GetYear() int
}

type Borrowable interface {
	Borrow() bool
	Return() bool
	IsBorrowed() bool
}

type Periodical interface {
	isPeriodical() bool
}

// Структура и методы для книги
type Book struct {
	title    string
	author   string
	year     int
	borrowed bool
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b Book) IsBorrowed() bool {
	return b.borrowed
}

// Borrow

func (b *Book) Borrow() bool {
	if b.borrowed {
		return false
	}
	b.borrowed = true
	return true
}

// Return
func (b *Book) Return() bool {
	if !b.borrowed {
		return false
	}
	b.borrowed = false
	return true
}
