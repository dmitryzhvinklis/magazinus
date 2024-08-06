package main

import (
	"db_service/internal/database"
	"db_service/internal/logging"
	"log"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger := logging.InitLogger()
	defer func() {
		if err := logger.Sync(); err != nil {
			zap.L().Error("Ошибка при синхронизации логгера", zap.Error(err))
		}
	}()

	logger.Info("Запуск приложения...")

	db, err := database.Connect()
	if err != nil {
		logger.Fatal("Не удалось подключиться к базе данных", zap.Error(err))
	}

	logger.Info("Успешное подключение к базе данных")

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			logger.Error("Ошибка при получении sql.DB", zap.Error(err))
			return
		}
		if err := sqlDB.Close(); err != nil {
			logger.Error("Ошибка при закрытии соединения с базой данных", zap.Error(err))
		} else {
			logger.Info("Соединение с базой данных закрыто")
		}
	}()

	logger.Info("Сервис базы данных работает")

	for {
		log.Println("Application is running...")
		time.Sleep(10 * time.Second)
	}
}
