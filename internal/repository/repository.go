package repository

import (
	"context"
	"database/sql"

	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
)

//Интерфейс для работы с моделью RsaUser
type RsaUser interface {
	CreateRsaUser(ctx context.Context , newRsaUser RsaUser) error
	GetRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (RsaUser, error)
	UpdateRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (bool, error)
	DeleteRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (bool, error)
}
//Структура репозитория, которая будет хранить интерфейсы для работы со всеми моделями
type Repository struct {
	RsaUser
}
//Функция нового репозитория, которая принимает подключение к БД и вызывает функцию создания структуры, которая реализуется RsaUser
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		RsaUser: NewRsaUserRep(db),
	}
}