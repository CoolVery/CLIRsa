package cmd

import (
	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/CoolVery/CLIRsa.git/internal/service"
	"github.com/spf13/cobra"
)

func GetUserInfo(rsaUserServ service.Service) *cobra.Command {
	var serachRsaUser models.User

	getRsaUser := &cobra.Command{
        Use:   "getUser",
        Short: "Получение пользователя",
		Long: "Команда для получения информации о пользователе по параметрам",
		RunE: func(cmd *cobra.Command, args []string) error {
			getUser, err := rsaUserServ.GetRsaUser(cmd.Context(), serachRsaUser.UserId, serachRsaUser)
			if err != nil {
				return err
			}
			getUser.Print()
			return nil
		},
    }

	getRsaUser.Flags().IntVarP(&serachRsaUser.UserId, "id-user", "i", -1, "Поиск происходит по идентификатору")
	getRsaUser.Flags().StringVarP(&serachRsaUser.UserNameFull, "full-name-user", "f", "", "Добавить в условие поиска ФИО пользователя")
	getRsaUser.Flags().StringVarP(&serachRsaUser.UserLogin, "login-user", "l", "", "Добавить в условие поиска Логин пользователя")
	getRsaUser.Flags().StringVarP(&serachRsaUser.RsaKey, "rsa-key-user", "k", "", "Добавить в условие поиска RSA ключ пользователя")
	getRsaUser.Flags().StringVarP(&serachRsaUser.RsaNum, "rsa-num-user", "n", "", "Добавить в условие поиска Номер RSA пользователя")


	return getRsaUser
}


