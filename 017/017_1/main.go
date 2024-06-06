package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

/*
Задача 17.1 Необходимо код примера 1 изменить так, чтобы tcp-сервер
обрабатывал подключения параллельно
*/

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	go func() { // Для решения задачи обернули функцию из примера 1 в горутину
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go handleConn(conn)
		}
	}()
	log.Println("завершение работы")
}

func handleConn(c net.Conn) {
	defer c.Close()
	var i int
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Println(err)
			return
		}
		i++
		time.Sleep(time.Second)
	}
}
