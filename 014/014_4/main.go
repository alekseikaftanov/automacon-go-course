package main

/*
Если запустить следующий код программы:
package main
import "fmt"
func main() {
ch := make(chan int)
stop := make(chan struct{})
go func() {
<-ch
stop <- struct{}{}
}()
<-stop
fmt.Println("happy end")
}
возникнет ошибка блокировки:
fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:
main.main()
/go/src/course/main.go:12 +0x9c
goroutine 18 [chan receive]:
main.main.func1()
/go/src/course/main.go:9 +0x2c
created by main.main in goroutine 1
/go/src/course/main.go:8 +0x90
exit status 2

Не меняя логику дочерней горутины, нужно дописать
логику главной горутины так, чтобы ошибки блокировки не возникало.
Вместо ошибки в консоль должна выводится фраза «happy end».
Так как канал ch является «особым» , то в него ЗАПРЕЩАЕТСЯ писать значения.
*/

import "fmt"

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()

	/*
		Чтобы избежать deadlock, необходимо гарантировать, что
		дочерняя горутина завершит свою работу и отправит сигнал в канал stop.
		Поскольку мы не можем записывать значения в канал ch, можно просто
		закрыть канал ch. Когда канал закрыт, операция чтения из него завершится,
		и дочерняя горутина сможет продолжить выполнение.
	*/
	close(ch) // закрывает канал ch. Когда канал закрыт, все текущие и будущие операции чтения из канала будут немедленно завершаться.
	<-stop
	fmt.Println("happy end")
}
