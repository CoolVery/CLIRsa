package db

import (
	"database/sql"
	"time"
)

//Названия таблиц
const (
	usersTable = "users"
)
//Функция подключения к базе данных
func ConnectPostgres(url string) (*sql.DB, error) {
	//Открываем соединение
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	//Задаем параметры
	//Максимальное времени жизни соединения
	db.SetConnMaxLifetime(time.Minute * 5)
	//Максимальное количество соединений, хранящихся в пуле
	db.SetMaxIdleConns(5)
	//Максимальное количество активных соединений
	db.SetMaxOpenConns(10)
	//Если пинг к бд не прошел, то вернем err
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}