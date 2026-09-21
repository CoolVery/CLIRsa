/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/CoolVery/CLIRsa.git/internal/service"
	"github.com/spf13/cobra"
)

// rootCmd - корневая команда
var rootCmd = &cobra.Command{
	Use:   "CLIRsa",
	Short: "A brief description of your application",
	Long: `Главная команда`,
	Run: func(cmd *cobra.Command, args []string) { 
		
	},
}
//Добавление остальных подкоманд
func Execute(rsaUserServ service.Service) {
	rootCmd.AddCommand(GetUserInfo(rsaUserServ))
	err := rootCmd.Execute()
	
	if err != nil {
		os.Exit(1)
	}
}



