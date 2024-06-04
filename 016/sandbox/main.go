package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background() // создаем контекст с базовыми значениями
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // создаем контекст с таймаутом
	d := time.Now().Add(time.Second) // создаем время, которое будет добавлено к текущему времени
	ctx, cancel := context.WithDeadline(ctx, d) // создаем контекст с дедлайном 
	defer cancel()         // отмена контекста
	wg := sync.WaitGroup{} // создаем группу ожидания
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1) // добавляем горутину в группу ожидания
		go func() {
			defer wg.Done() // указываем, что горутина завершила работу
			for {
				select {
				case <-ctx.Done(): // если контекст отменен, то выходим из цикла
					fmt.Println("стопгорутина: ", i)
					return
				default:
					fmt.Println("сложныевычислениягорутины: ", i)
				}
			}
		}()
	}
	wg.Add(1) // добавляем горутину в группу ожидания
	go func() {
		defer wg.Done() // указываем, что горутина завершила работу
		fmt.Println("стоп!")
		cancel()
	}()
	wg.Wait() // ждем завершения всех горутин
}
