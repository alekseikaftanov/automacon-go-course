package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Задача 16.1 Нужно запустить 5 горутин и остановить, используя контекст с отменой.

func main() {
	// Создаем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Инициализируем WaitGroup
	var wg sync.WaitGroup

	// Запускаем 5 горутин
	for i := 0; i < 5; i++ {
		wg.Add(1) // Добавляем 1 горутину в WaitGroup
		go func(i int) {
			defer wg.Done() // Указываем, что горутина завершила работу
			for {
				select {
				case <-ctx.Done(): // Проверяем, не завершилась ли горутина
					fmt.Printf("goroutine %d is stopped\n", i)
					return
				default: // Если горутина не завершилась, выполняем работу
					fmt.Printf("goroutine %d is running\n", i)
					time.Sleep(500 * time.Millisecond) // Симуляция работы
				}
			}
		}(i)
	}

	// Выводим сообщение, когда все горутины запущены
	fmt.Println("All goroutines are running")
	time.Sleep(1 * time.Second)
	fmt.Println("Stopping all goroutines")

	cancel()
	wg.Wait()

	fmt.Println("All goroutines are stopped")
}
