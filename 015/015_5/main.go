package main

import (
	"fmt"
	"sync"
)

/*
Задача 15.5
Необходимо реализовать интерфейс
type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

Речь про температуру окружающей среды. ReadTemp читает температуру,
ChangeTemp изменяет температуру. Код должен быть потокобезопасным, т.е.
при работе температуры нескольких параллельных горутин
не должно возникать состояние гонки.
*/

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type SafeMeteo struct {
	mu   sync.Mutex
	temp string
}

func (sm *SafeMeteo) ReadTemp() string {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.temp
}

func (sm *SafeMeteo) ChangeTemp(v string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.temp = v
}

func main() {
	meteo := SafeMeteo{temp: "20"} // Понять зачем &
	var wg sync.WaitGroup

	// Запускаем несколько горутин для чтения и изменения температуры
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Горутина %d читает температуру %s\n", id, meteo.ReadTemp())
		}(i)
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			newTemp := fmt.Sprintf("%d", +id)
			fmt.Printf("Горутина %d изменяет температуру на %s\n", id, newTemp)
		}(i)
	}

	wg.Wait()
	fmt.Println("Финальная температура:", meteo.ReadTemp())
}
