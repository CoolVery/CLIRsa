package dto

import (
	"github.com/CoolVery/CLIRsa.git/internal/errors/apperr"
)

//Класс для комбинаторных запросов к модели
type DtoRsaUser struct {
	Id           *int `json:"id"`
	RsaNum       *string `json:"rsa_num"`
	RsaKey       *string `json:"rsa_key"`
	UserNameFull *string `json:"user_name_full"`
	UserLogin    *string `json:"user_login"`
}
//Функция проверки полей модели
func (dtoRU DtoRsaUser) Validate() error {
	if dtoRU.RsaKey == nil && dtoRU.RsaNum == nil && dtoRU.UserLogin == nil && dtoRU.UserNameFull == nil && dtoRU.Id == nil {
		return apperr.ErrNotFilter
	}
	if dtoRU.RsaKey != nil && len(*dtoRU.RsaKey) != 40 {
		return apperr.ErrRsaLen
	}
	return nil
}

