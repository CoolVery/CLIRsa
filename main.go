/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/CoolVery/CLIRsa.git/cmd"
	"github.com/CoolVery/CLIRsa.git/internal/config"
	"github.com/CoolVery/CLIRsa.git/internal/repository"
	"github.com/CoolVery/CLIRsa.git/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//Чтение конфигурации из .env
	if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("no .env file, using system env")
    }
	//Создание конфига для бд
	cfg := config.Load()
	//Подключение к бд
	db, err := sqlx.Connect("postgres", cfg.ConnectToPostgres())
	if err != nil {
		log.Fatalf("db err %v", err)
	}
	defer db.Close()
	//Создание репозитория
	rsaUserRep := repository.NewRepository(db)
	//Создание сервиса
	rsaUserServ := service.NewService(*rsaUserRep)
	//Перелаем сервис для команд с работой RsaUser
	cmd.Execute(*rsaUserServ)
}
