package main

import (
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/db"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/handler"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"

	"github.com/spf13/cobra"
)

func main() {
	database := db.Connect()
	kategoriRepo := &repository.KategoriRepository{DB: database}

	rootCmd := &cobra.Command{Use: "inventaris"}
	rootCmd.AddCommand(handler.NewKategoriCmd(kategoriRepo))
	rootCmd.Execute()
}