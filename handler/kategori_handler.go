package handler

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"
)

func NewKategoriCmd(repo *repository.KategoriRepository) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kategori",
		Short: "Manajemen kategori barang",
	}

	cmd.AddCommand(
		listKategoriCmd(repo),
		addKategoriCmd(repo),
		getKategoriCmd(repo),
		editKategoriCmd(repo),
		deleteKategoriCmd(repo),
	)
	return cmd
}

func listKategoriCmd(repo *repository.KategoriRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Tampilkan daftar kategori",
		Run: func(cmd *cobra.Command, args []string) {
			kategoris, err := repo.GetAll()
			if err != nil {
				fmt.Println("Gagal mengambil data:", err)
				return
			}
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "ID\tNama\tDeskripsi")
			for _, k := range kategoris {
				fmt.Fprintf(w, "%d\t%s\t%s\n", k.ID, k.Nama, k.Deskripsi)
			}
			w.Flush()
		},
	}
}

func addKategoriCmd(repo *repository.KategoriRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "add [nama] [deskripsi]",
		Short: "Tambah kategori baru",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			nama, deskripsi := args[0], args[1]
			if nama == "" {
				fmt.Println("Nama tidak boleh kosong")
				return
			}
			err := repo.Create(nama, deskripsi)
			if err != nil {
				fmt.Println("Gagal menambah kategori:", err)
			} else {
				fmt.Println("Kategori berhasil ditambah")
			}
		},
	}
}

func getKategoriCmd(repo *repository.KategoriRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "get [id]",
		Short: "Lihat detail kategori",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			k, err := repo.GetByID(id)
			if err != nil {
				fmt.Println("Kategori tidak ditemukan:", err)
				return
			}
			fmt.Printf("ID: %d\nNama: %s\nDeskripsi: %s\n", k.ID, k.Nama, k.Deskripsi)
		},
	}
}

func editKategoriCmd(repo *repository.KategoriRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "edit [id] [nama] [deskripsi]",
		Short: "Edit kategori",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			nama, deskripsi := args[1], args[2]
			if nama == "" {
				fmt.Println("Nama tidak boleh kosong")
				return
			}
			err := repo.Update(id, nama, deskripsi)
			if err != nil {
				fmt.Println("Gagal edit kategori:", err)
			} else {
				fmt.Println("Kategori berhasil diupdate")
			}
		},
	}
}

func deleteKategoriCmd(repo *repository.KategoriRepository) *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Hapus kategori",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			err := repo.Delete(id)
			if err != nil {
				fmt.Println("Gagal hapus kategori:", err)
			} else {
				fmt.Println("Kategori berhasil dihapus")
			}
		},
	}
}