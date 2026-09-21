package cmd

import (
	"github.com/CoolVery/CLIRsa.git/internal/service"
	"github.com/spf13/cobra"
)

func GetUserInfo(rsaUserServ service.Service) *cobra.Command {
	rsaUser := &cobra.Command{
        Use:   "user",
        Short: "Получение пользователя",
    }

	return rsaUser
}