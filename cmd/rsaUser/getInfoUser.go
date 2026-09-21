package cmd

import (
	"fmt"

	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/service"
	"github.com/spf13/cobra"
)

func GetUserInfo(rsaUserServ service.Service) *cobra.Command {
	var serachRsaUser models.User

	rsaUser := &cobra.Command{
        Use:   "user",
        Short: "Получение пользователя",
		Long: "Команда для получения информации о пользователи по его идентификатору и параметрам",
		RunE: func(cmd *cobra.Command, args []string) error {

			getUser, err := rsaUserServ.GetRsaUser(cmd.Context(), serachRsaUser.UserId, serachRsaUser)
			if err != nil {
				return err
			}

			fmt.Println(getUser)
			return nil
		},
    }

	rsaUser.Flags().IntVarP(&serachRsaUser.UserId, "id-user", "i", -1, "Поиск происходит по идентификатору")
	rsaUser.Flags().StringVarP(&serachRsaUser.UserNameFull, "full-name-user", "f", "", "Добавить в условие поиска ФИО пользователя")
	rsaUser.Flags().StringVarP(&serachRsaUser.UserLogin, "login-user", "l", "", "Добавить в условие поиска Логин пользователя")
	rsaUser.Flags().StringVarP(&serachRsaUser.RsaKey, "rsa-key-user", "k", "", "Добавить в условие поиска RSA ключ пользователя")
	rsaUser.Flags().StringVarP(&serachRsaUser.RsaNum, "rsa-num-user", "n", "", "Добавить в условие поиска Номер RSA пользователя")


	_ = rsaUser.MarkFlagRequired("id-user")
	return rsaUser
}


