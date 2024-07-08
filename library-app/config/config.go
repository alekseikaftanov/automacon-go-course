package config

import (
	"errors"
	"os"
)

type Config struct { // Структура для хранения конфигурации
	DatabaseURL   string // URL для подключения к базе данных
	ServerAddress string // Адрес сервера
}

func LoadConfig() (*Config, error) { // Функция для загрузки конфигурации
	cfg := &Config{ // Создание экземпляра структуры Config
		DatabaseURL:   os.Getenv("DATABASE_URL"), // Получение URL для подключения к базе данных
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}

	// Проверка на обязательные параметры
	if cfg.DatabaseURL == "" || cfg.ServerAddress == "" { // Проверка на наличие URL и адреса сервера
		return nil, errors.New("необходимо задать переменные окружения DATABASE_URL и SERVER_ADDRESS")
	}

	return cfg, nil
}
