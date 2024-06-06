package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

/*
Задача 17.2 Измените код примера 4 так, чтобы логировались
абсолютно все запросы, даже которые не прошли авторизацию.
*/

func main() {
	mux := http.NewServeMux()                                                        // Создаем мукс
	mux.HandleFunc("/hello", hello)                                                  // Добавляем обработчик
	l := log.New(os.Stdout, "", 0)                                                   // Создаем логгер
	logHandler := logMiddleware(l)                                                   // Создаем обертку для логирования
	httpServer := &http.Server{Addr: ":8080", Handler: authHandler(logHandler(mux))} // Создаем сервер
	if err := httpServer.ListenAndServe(); err != nil {                              // Запускаем сервер
		log.Fatalln(fmt.Errorf("не удалось запустить сервер: %w", err)) // Логируем ошибку, если сервер не запустился
	}
}

// Обработчик для маршрута "/hello"
func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"       // Сообщение, которое будет отправлено клиенту
	log.Println("resp:", msg) // Логируем сообщение
	fmt.Fprint(res, msg)      // Отправляем сообщение в ответе
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			log.Println("Unauthorized request:", r.URL) // Логируем неавторизованный запрос
			return
		}
		h.ServeHTTP(w, r) // Если авторизация прошла, передаем запрос дальше
	})
}

func logMiddleware(logger *log.Logger) func(h http.Handler) http.Handler { // Используем переданный логгер
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println("url:", r.URL) // Логируем URL запроса
			h.ServeHTTP(w, r)             // Передаем запрос дальше
		})
	}
}
