package cmd

import (
	"fmt"
	"os"

	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/db"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/handler"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "inventaris",
	Short: "Aplikasi CLI Inventaris Barang",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Terjadi kesalahan:", err)
		os.Exit(1)
	}
}

func init() {
	// Inisialisasi koneksi DB
	connStr := "your_connection_string_here" // Ganti dengan string koneksi yang sesuai
	conn, err := db.Connect(connStr)
	if err != nil {
		fmt.Println("Gagal koneksi database:", err)
		os.Exit(1)
	}

	// Repository & Handler
	kategoriRepo := repository.NewKategoriRepository(conn)
	barangRepo := repository.NewBarangRepository(conn)
	kategoriHandler := &handler.KategoriHandler{Repo: kategoriRepo}
	barangHandler := &handler.BarangHandler{Repo: barangRepo, KategoriRepo: kategoriRepo}

	// Kategori Command
	var kategoriCmd = &cobra.Command{
		Use:   "kategori",
		Short: "Manajemen kategori barang",
	}
	var kategoriListCmd = &cobra.Command{
		Use:   "list",
		Short: "Tampilkan daftar kategori",
		Run: func(cmd *cobra.Command, args []string) {
			kategoriHandler.ListKategori()
		},
	}
	var kategoriAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Tambah kategori baru",
		Run: func(cmd *cobra.Command, args []string) {
			kategoriHandler.AddKategori()
		},
	}
	var kategoriDetailCmd = &cobra.Command{
		Use:   "detail",
		Short: "Lihat detail kategori",
		Run: func(cmd *cobra.Command, args []string) {
			kategoriHandler.DetailKategori()
		},
	}
	var kategoriEditCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit kategori",
		Run: func(cmd *cobra.Command, args []string) {
			kategoriHandler.EditKategori()
		},
	}
	var kategoriDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Hapus kategori",
		Run: func(cmd *cobra.Command, args []string) {
			kategoriHandler.DeleteKategori()
		},
	}
	kategoriCmd.AddCommand(kategoriListCmd, kategoriAddCmd, kategoriDetailCmd, kategoriEditCmd, kategoriDeleteCmd)
	rootCmd.AddCommand(kategoriCmd)

	// Barang Command
	var barangCmd = &cobra.Command{
		Use:   "barang",
		Short: "Manajemen barang inventaris",
	}
	var barangListCmd = &cobra.Command{
		Use:   "list",
		Short: "Tampilkan daftar barang",
		Run: func(cmd *cobra.Command, args []string) {
			barangHandler.ListBarang()
		},
	}
	var barangAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Tambah barang baru",
		Run: func(cmd *cobra.Command, args []string) {
			barangHandler.AddBarang()
		},
	}
	var barangDetailCmd = &cobra.Command{
		Use:   "detail",
		Short: "Lihat detail barang",
		Run: func(cmd *cobra.Command, args []string) {
			barangHandler.DetailBarang()
		},
	}
	var barangEditCmd = &cobra.Command{
		Use:   "edit",
		Short: "Edit barang",
		Run: func(cmd *cobra.Command, args []string) {
			barangHandler.EditBarang()
		},
	}
	var barangDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Hapus barang",
		Run: func(cmd *cobra.Command, args []string) {
			barangHandler.DeleteBarang()
		},
	}
	barangCmd.AddCommand(barangListCmd, barangAddCmd, barangDetailCmd, barangEditCmd, barangDeleteCmd)
	rootCmd.AddCommand(barangCmd)
}
