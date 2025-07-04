package main

import (
	"fmt"
	"os"

	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/db"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/handler"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"

	"github.com/spf13/cobra"
)

func main() {
	dsn := "host=localhost port=5433 user=postgres password=c4k3 dbname=inventaris sslmode=disable" // TODO: replace with your actual connection string
	database, err := db.Connect(dsn)
	if err != nil {
		fmt.Println("Gagal koneksi database:", err)
		os.Exit(1)
	}

	// Inisialisasi repository
	repository.NewKategoriRepository(database)
	repository.NewBarangRepository(database)

	// Inisialisasi handler
	// barangHandler := &handler.BarangHandler{Repo: barangRepo, KategoriRepo: kategoriRepo}

	// Root command
	rootCmd := &cobra.Command{Use: "inventaris"}

	// Tambahkan command kategori dan barang
	rootCmd.AddCommand(handler.NewKategoriCmd())
	rootCmd.AddCommand(handler.NewBarangCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		os.Exit(1)
	}
}
