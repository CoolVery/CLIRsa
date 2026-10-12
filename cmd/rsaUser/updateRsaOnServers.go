package cmd

import (
	"github.com/CoolVery/CLIRsa.git/internal/models"
	"github.com/spf13/cobra"
)

func UpdateRsaOnServers() *cobra.Command {
	var serachRsaUser models.User

	updateRsaKey := &cobra.Command{
        Use:   "updateRsaKey",
        Short: "Обновление RSA ключа на серверах",
		Long: "Команда которая находит блок пользователя и меняет его RSA-ключ на всех серверах",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
    }
	updateRsaKey.Flags().StringVarP(&serachRsaUser.UserLogin, "login-user", "l", "", "Добавить в условие поиска Логин пользователя")
	updateRsaKey.Flags().StringVarP(&serachRsaUser.RsaKey, "rsa-key-user", "k", "", "Добавить в условие поиска RSA ключ пользователя")


	return updateRsaKey
}


