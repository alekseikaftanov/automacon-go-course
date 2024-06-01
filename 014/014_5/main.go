package main

import (
	"fmt"
	"time"
)

/*
Следующий код программы запускает излишне «болтливую» горутину:

package main
import (
"fmt"
"time"
)
func main() {
ch := make(chan int)
stop := make(chan struct{}, 2)
go func() {
OUT:
for {
select {
case <-stop:
break OUT
case v, ok := <-ch:
if !ok {
break OUT
}
fmt.Println(v)
default:
continue
}
}
fmt.Println("завершение работы горутины_1")
}()
go func() {
var i int
OUT:
for {
i++
select {
case <-stop:
break OUT
default:
time.Sleep(time.Second)
if ch == nil {
continue
}
ch <- i
}
}
fmt.Println("завершение работы горутины_2")
}()
time.Sleep(5 * time.Second)
stop <- struct{}{}
stop <- struct{}{}
time.Sleep(time.Second)
fmt.Println("завершение работы главной горутины")
}
которая в консоль выводит счёт. Горутина настолько утомительна,
что в коде поставлено ограничение по времени выполнения программы.
Не меняя логику дочерних горутин, необходимо дописать логику так,
чтобы «болтушка» не смогла вывести в консоль свой утомительный счёт.
Вместо пустой болтовни должны увидеть:
завершение работы горутины_1
завершение работы горутины_2
завершение работы главной горутины

Причём логика главной горутины должна отработать в полном объёме,
без изменения длительности временных задержек.
*/

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 3)
	stop <- struct{}{}
	stop <- struct{}{}
	go func() {
	OUT:
		for {
			select {
			case <-stop: //Если получено значение из канала stop, выполнение переходит на метку OUT, что приводит к завершению цикла.
				break OUT
			case v, ok := <-ch: // Если получено значение из канала ch, оно присваивается переменной v, а переменная ok указывает на успешность операции. Если ok равно false (канал закрыт), происходит выход из цикла.
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default: // Если ни один из других каналов не готов, выполнение продолжается с continue.
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	ch = nil
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	fmt.Println("завершение работы главной горутины")
}
