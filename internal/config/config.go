package config

import (
	"fmt"
	"os"
)
//Структура конфигурации
type Config struct {
	host     string
	port     string
	userName string
	password string
	dbName   string
}
//Функция чтения и записи в конфигурацию значений из .env
func Load() Config {
	return Config{
		host: getEnv("DB_HOST", "localhost"),
		port: getEnv("DB_PORT", "5432"),
		userName: getEnv("DB_USERNAME", "postgrea"),
		password: getEnv("DB_PASSWORD", ""),
		dbName: getEnv("DB_NAME", "postgrea"),
	}
}
//Функция создания строки подключения в базе данных PostgreSQL
func (c Config) ConnectToPostgres() string {
	return fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        c.host, c.port, c.userName, c.password, c.dbName,
    )
}
//Функция чтения данных из .env
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}