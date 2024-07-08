# Установите переменную окружения для версии Go


# Используйте базовый образ Golang с указанной версией
FROM golang:1.20 as builder

# Установите рабочую директорию внутри контейнера
WORKDIR /app

# Скопируйте go.mod и go.sum файлы
COPY go.mod go.sum ./

# Установите зависимости
RUN go mod download

# Скопируйте остальные файлы приложения
COPY . .

# Соберите приложение
RUN go build -o /app/library-app ./cmd

# Проверьте, что файл был создан
RUN ls -l /app/library-app

# Используйте минимальный образ для запуска приложения
FROM alpine:latest

# Установите рабочую директорию внутри контейнера
WORKDIR /root/

# Скопируйте скомпилированное приложение из этапа сборки
COPY --from=builder /app/library-app .

# Установите права на выполнение
RUN chmod +x /root/library-app

# Проверьте, что файл был скопирован
RUN ls -l /root/library-app

# Запустите приложение
CMD ["./library-app"]