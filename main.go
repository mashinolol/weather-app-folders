package main

import (
	"context"
	"log"
	"net/http"

	"weather-app/config"
	"weather-app/handler"
	"weather-app/repository"
	"weather-app/service"

	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализация конфигурации
	cfg := config.NewConfig()

	// Подключение к MongoDB
	dbClient, err := config.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer dbClient.Disconnect(context.TODO())

	// Инициализация репозитория, сервиса и хендлера
	weatherRepo := repository.NewWeatherRepository(dbClient, cfg.DatabaseName)
	weatherService := service.NewWeatherService(weatherRepo, cfg.BaseURL, cfg.APIKey)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	// Настройка маршрутов
	http.HandleFunc("/weather", weatherHandler.HandleWeather)

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
