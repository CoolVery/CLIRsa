package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/models/dto"
	"github.com/CoolVery/CLIRsa.git/internal/models/tables"
	"github.com/jmoiron/sqlx"
)

//Структура Репозитория модели RsaUser
type RsaUserRep struct {
	db *sqlx.DB
}
//Функция создания нового Репозитория модели RsaUser
func NewRsaUserRep(db *sqlx.DB) *RsaUserRep {
	return &RsaUserRep{
		db: db,
	}
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

func (ruRep *RsaUserRep) GetRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (*models.User, error) {
	var getUser models.User
	
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argInt := 1

	if rsaUserArgs.RsaKey != nil {
		setValues = append(setValues, fmt.Sprintf("AND rsa_key=$%d", argInt))
		args = append(args, *rsaUserArgs.RsaKey)
		argInt++
	}

	if rsaUserArgs.RsaNum != nil {
		setValues = append(setValues, fmt.Sprintf("AND rsa_num=$%d", argInt))
		args = append(args, *rsaUserArgs.RsaNum)
		argInt++
	}

	if rsaUserArgs.UserLogin != nil {
		setValues = append(setValues, fmt.Sprintf("AND user_login=$%d", argInt))
		args = append(args, *rsaUserArgs.UserLogin)
		argInt++
	}

	if rsaUserArgs.UserNameFull != nil {
		setValues = append(setValues, fmt.Sprintf("AND user_name_full=$%d", argInt))
		args = append(args, *rsaUserArgs.UserNameFull)
		argInt++
	}
	if len(setValues) == 0 {
		return nil, errors.New("no fields to update provided")
	}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$%d %s`, tables.UsersTable, argInt, setValues)
	args = append(args, idUser)


	err := ruRep.db.GetContext(ctx, &getUser, query, args...)
	if err != nil {
		return nil, err
	}
	return &getUser, nil
}

func (ruRep *RsaUserRep) UpdateRsaUser(ctx context.Context, idUser int, rsaUserArgs dto.DtoRsaUser) (bool, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argInt := 1

	if rsaUserArgs.RsaKey != nil {
		setValues = append(setValues, fmt.Sprintf("rsa_key=$%d", argInt))
		args = append(args, *rsaUserArgs.RsaKey)
		argInt++
	}

	if rsaUserArgs.RsaNum != nil {
		setValues = append(setValues, fmt.Sprintf("rsa_num=$%d", argInt))
		args = append(args, *rsaUserArgs.RsaNum)
		argInt++
	}

	if rsaUserArgs.UserLogin != nil {
		setValues = append(setValues, fmt.Sprintf("user_login=$%d", argInt))
		args = append(args, *rsaUserArgs.UserLogin)
		argInt++
	}

	if rsaUserArgs.UserNameFull != nil {
		setValues = append(setValues, fmt.Sprintf("user_name_full=$%d", argInt))
		args = append(args, *rsaUserArgs.UserNameFull)
		argInt++
	}
	if len(setValues) == 0 {
		return false, errors.New("no fields to update provided")
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE users SET %s WHERE user_id=$%d`, setQuery, argInt)
	args = append(args, idUser)


	_, err := ruRep.db.ExecContext(ctx, query, args...)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ruRep *RsaUserRep) DeleteRsaUser(ctx context.Context, idUser int) (bool, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", tables.UsersTable)
	_, err := ruRep.db.ExecContext(ctx, query, idUser)
	if err != nil {
		return false, err
	}
	return true, nil
}