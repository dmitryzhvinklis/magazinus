package database

import (
	"context"
	"db_service/internal/logging"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config содержит конфигурацию базы данных
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

// LoadConfig загружает конфигурацию из YAML файла
func LoadConfig() (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile("configs/database/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	// Переопределяем значения переменными окружения, если они заданы
	if host := os.Getenv("DATABASE_HOST"); host != "" {
		config.Database.Host = host
	}
	if port := os.Getenv("DATABASE_PORT"); port != "" {
		portInt, err := strconv.Atoi(port)
		if err == nil {
			config.Database.Port = portInt
		}
	}
	if user := os.Getenv("DATABASE_USER"); user != "" {
		config.Database.User = user
	}
	if password := os.Getenv("DATABASE_PASSWORD"); password != "" {
		config.Database.Password = password
	}
	if dbname := os.Getenv("DATABASE_NAME"); dbname != "" {
		config.Database.DBName = dbname
	}

	return config, nil
}

// CustomLogger реализует интерфейс gorm.Logger для использования zap
type CustomLogger struct {
	Logger *logging.Logger
}

func (cl *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	return &CustomLogger{Logger: cl.Logger}
}

func (cl *CustomLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	cl.Logger.Sugar().Infof(msg, args...)
}

func (cl *CustomLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	cl.Logger.Sugar().Warnf(msg, args...)
}

func (cl *CustomLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	cl.Logger.Sugar().Errorf(msg, args...)
}

func (cl *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if err != nil {
		cl.Logger.Sugar().Errorf("SQL Error: %s", err)
	}
	cl.Logger.Sugar().Infof("SQL Trace: %s [rows: %d]", sql, rows)
}

// Connect устанавливает соединение с базой данных и выполняет миграцию
func Connect() (*gorm.DB, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName, config.Database.SSLMode)

	// Настройка логгера
	logger := logging.InitLogger()
	gormLogger := &CustomLogger{Logger: logger}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	err = Migrate(db)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to the database successfully")
	return db, nil
}
