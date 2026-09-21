package service

import (
	"context"

	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
	"github.com/CoolVery/CLIRsa.git/internal/repository"
)

//Интерфейс для работы с моделью RsaUser
type RsaUser interface {
	CreateRsaUser(ctx context.Context , newRsaUser models.User) (int, error)
	GetRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (*models.User, error)
	UpdateRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (bool, error)
	DeleteRsaUser(ctx context.Context, idUser int) (bool, error)
}
//Общий сервис
type Service struct {
	RsaUser
}
//Метод создания нового сервиса
func NewService(rep repository.Repository) *Service {
	return &Service{
		RsaUser: NewRsaUserServ(rep.RsaUser),
	}
}

