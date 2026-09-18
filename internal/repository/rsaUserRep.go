package repository

import (
	"context"
	"database/sql"

	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
)
//Структура Репозитория модели RsaUser
type RsaUserRep struct {
	db *sql.DB
}
//Функция создания нового Репозитория модели RsaUser
func NewRsaUserRep(db *sql.DB) *RsaUserRep {
	return &RsaUserRep{
		db: db,
	}
}
//Реализация всех функций интерфейса RsaUser
func (ruRep *RsaUserRep) CreateRsaUser(ctx context.Context , newRsaUser RsaUser) error {
	return nil
}

func (ruRep *RsaUserRep) GetRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (RsaUser, error) {
	return nil, nil
}

func (ruRep *RsaUserRep) UpdateRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (bool, error) {
	return false, nil
}

func (ruRep *RsaUserRep) DeleteRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (bool, error) {
	return true, nil
}