package service

import (
	"context"
	"log"

	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
	"github.com/CoolVery/CLIRsa.git/internal/repository"
)

//Сервис для модели RsaUser
type RsaUserSer struct {
	rep repository.RsaUser
}

func NewRsaUserServ(rep repository.RsaUser) *RsaUserSer {
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
func (ruSer *RsaUserSer) GetRsaUser(ctx context.Context, idUser int, rsaUserArgs models.User) (*models.User, error) {
	var rsaUserDTO dto.DtoRsaUser

	if rsaUserArgs.UserLogin != "" {
		rsaUserDTO.UserLogin = &rsaUserArgs.UserLogin
	}
	if rsaUserArgs.RsaKey != "" {
		rsaUserDTO.RsaKey = &rsaUserArgs.RsaKey
	}
	if rsaUserArgs.RsaNum != "" {
		rsaUserDTO.RsaNum = &rsaUserArgs.RsaNum
	}
	log.Printf("-- %s",rsaUserArgs.UserNameFull)
	if rsaUserArgs.UserNameFull != "" {
		rsaUserDTO.UserNameFull = &rsaUserArgs.UserNameFull
	}
	if rsaUserArgs.UserLogin != "" {
		rsaUserDTO.UserLogin = &rsaUserArgs.UserLogin
	}

	getUser, err := ruSer.rep.GetRsaUser(ctx, idUser, rsaUserDTO)
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