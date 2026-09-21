package repository

import (
	"context"
	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
	"github.com/jmoiron/sqlx"
)

//Интерфейс для работы с моделью RsaUser
type RsaUser interface {
	CreateRsaUser(ctx context.Context , newRsaUser models.User) (int, error)
	GetRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (*models.User, error)
	UpdateRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (bool, error)
	DeleteRsaUser(ctx context.Context, idUser int) (bool, error)
}
//Структура репозитория, которая будет хранить интерфейсы для работы со всеми моделями
type Repository struct {
	RsaUser
}
//Функция нового репозитория, которая принимает подключение к БД и вызывает функцию создания структуры, которая реализуется RsaUser
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		RsaUser: NewRsaUserRep(db),
	}
}