package service

import (
	"context"

	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
	"github.com/CoolVery/CLIRsa.git/internal/repository"
)
//Сервис для модели RsaUser
type RsaUserSer struct {
	rep *repository.RsaUserRep
}

func NewRsaUserServ(rep *repository.RsaUserRep) *RsaUserSer {
	return &RsaUserSer{rep: rep}
}
//Реализация методов для работы с моделью
func (ruSer *RsaUserSer) CreateRsaUser(ctx context.Context , newRsaUser models.User) (int, error) {
	idUser, err := ruSer.rep.CreateRsaUser(ctx, newRsaUser)
	if err != nil {
		return 0, err
	}
	return idUser, nil
}
func (ruSer *RsaUserSer) GetRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (*models.User, error) {
	getUser, err := ruSer.rep.GetRsaUser(ctx, idUser, rsaUserArgs)
	if err != nil {
		return nil, err
	}
	return getUser, nil
}
func (ruSer *RsaUserSer) UpdateRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (bool, error) {
	isUpdated, err := ruSer.rep.UpdateRsaUser(ctx, idUser, rsaUserArgs)
	if err != nil {
		return isUpdated, err
	}
	return isUpdated, nil
}
func (ruSer *RsaUserSer) DeleteRsaUser(ctx context.Context, idUser int) (bool, error) {
	isDeleted, err :=ruSer.rep.DeleteRsaUser(ctx, idUser)
	if err != nil {
		return isDeleted, err 
	}
	return isDeleted, nil
}