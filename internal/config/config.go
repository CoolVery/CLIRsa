package config

import (
	"fmt"
	"os"
)
//Структура конфигурации
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}
//Функция чтения и записи в конфигурацию значений из .env
func Load() Config {
	return Config{
		Host: getEnv("DB_HOST", "localhost"),
		Port: getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USERNAME", "postgrea"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName: getEnv("DB_NAME", "postgrea"),
	}
}
//Функция создания строки подключения в базе данных PostgreSQL
func (c Config) ConnectToPostgres() string {
	return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.Host, c.Port, c.Username, c.Password, c.DBName,
    )
}
//Функция чтения данных из .env
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}