package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
	"github.com/CoolVery/CLIRsa.git/internal/models/tables"
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
func rre() {

}
//Реализация всех функций интерфейса RsaUser
func (ruRep *RsaUserRep) CreateRsaUser(ctx context.Context , newRsaUser models.User) (int, error) {
	//Формируем строку запроса
	query := fmt.Sprintf(`
	INSERT INTO %s (rsa_num, rsa_key, user_name_full, user_login)
	VALUES ($1, $2, $3, $4)
	RETURNING user_id
	`, tables.UsersTable)
	
	var newId int
	//Формируем строку-запрос в SQL
	row := ruRep.db.QueryRowContext(ctx, query, newRsaUser.RsaNum, newRsaUser.RsaKey, newRsaUser.UserNameFull, newRsaUser.UserLogin)
	//Выполняем запрос
	if err := row.Scan(&newId); err != nil {
		return 0, err
	}

	return newId, nil
}

func (ruRep *RsaUserRep) GetRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (RsaUser, error) {
	return nil, nil
}

func (ruRep *RsaUserRep) UpdateRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (bool, error) {
	
	// setValues := make([]string, 0)
	// args := make([]interface{}, 0)
	// argInt := 1

	// if dtoRU.RsaKey != nil {
	// 	setValues = append(setValues, fmt.Sprintf("rsa_key=$%d", argInt))
	// 	args = append(args, *dtoRU.RsaKey)
	// 	argInt++
	// }

	// if dtoRU.RsaNum != nil {
	// 	setValues = append(setValues, fmt.Sprintf("rsa_num=$%d", argInt))
	// 	args = append(args, *dtoRU.RsaNum)
	// 	argInt++
	// }

	// if dtoRU.UserLogin != nil {
	// 	setValues = append(setValues, fmt.Sprintf("user_login=$%d", argInt))
	// 	args = append(args, *dtoRU.UserLogin)
	// 	argInt++
	// }

	// if dtoRU.UserNameFull != nil {
	// 	setValues = append(setValues, fmt.Sprintf("user_name_full=$%d", argInt))
	// 	args = append(args, *dtoRU.UserNameFull)
	// 	argInt++
	// }

	// if len(setValues) == 0 {
	// 	return "", errors.New("no fields to update provided")
	// }

	// setQuery := strings.Join(setValues, ", ")
	// query := fmt.Sprintf(`UPDATE users SET %s WHERE user_id=$%d`, setQuery, argInt)
	// args = append(args, userId)


	// _, err := us.db.ExecContext(ctx, query, args...)
	return false, nil
}

func (ruRep *RsaUserRep) DeleteRsaUser(ctx context.Context, rsaUserArgs dto.DtoRsaUser) (bool, error) {
	return true, nil
}