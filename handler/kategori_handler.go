package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"
)

type KategoriHandler struct {
	Repo *repository.KategoriRepository
}

// Menampilkan daftar kategori dalam bentuk tabel
func (h *KategoriHandler) ListKategori() {
	kategoris, err := h.Repo.GetAll()
	if err != nil {
		fmt.Println("Gagal mengambil data kategori.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNama\tDeskripsi")
	for _, k := range kategoris {
		fmt.Fprintf(w, "%d\t%s\t%s\n", k.ID, k.Nama, k.Deskripsi)
	}
	w.Flush()
}

// Menambah kategori baru
func (h *KategoriHandler) AddKategori() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Kategori: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama == "" {
		fmt.Println("Nama kategori tidak boleh kosong.")
		return
	}

	// Cek duplikat nama kategori
	kategoris, err := h.Repo.GetAll()
	if err != nil {
		fmt.Println("Gagal mengambil data kategori.")
		return
	}
	for _, k := range kategoris {
		if strings.EqualFold(k.Nama, nama) {
			fmt.Println("Nama kategori sudah ada, tidak boleh duplikat.")
			return
		}
	}

	fmt.Print("Deskripsi: ")
	deskripsi, _ := reader.ReadString('\n')
	deskripsi = strings.TrimSpace(deskripsi)

	kategori := &model.Kategori{
		Nama:      nama,
		Deskripsi: deskripsi,
	}
	err = h.Repo.Create(kategori)
	if err != nil {
		fmt.Println("Gagal menambah kategori:", err)
		return
	}
	fmt.Println("Kategori berhasil ditambahkan dengan ID:", kategori.ID)
}

// Melihat detail kategori berdasarkan ID
func (h *KategoriHandler) DetailKategori() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Kategori: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	kategori, err := h.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Kategori tidak ditemukan.")
		return
	}

	fmt.Println("ID        :", kategori.ID)
	fmt.Println("Nama      :", kategori.Nama)
	fmt.Println("Deskripsi :", kategori.Deskripsi)
}

// Edit kategori
func (h *KategoriHandler) EditKategori() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Kategori yang akan diedit: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	kategori, err := h.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Kategori tidak ditemukan.")
		return
	}

	fmt.Printf("Nama lama (%s), masukkan nama baru (kosongkan jika tidak ingin mengubah): ", kategori.Nama)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama != "" {
		kategori.Nama = nama
	}

	fmt.Printf("Deskripsi lama (%s), masukkan deskripsi baru (kosongkan jika tidak ingin mengubah): ", kategori.Deskripsi)
	deskripsi, _ := reader.ReadString('\n')
	deskripsi = strings.TrimSpace(deskripsi)
	if deskripsi != "" {
		kategori.Deskripsi = deskripsi
	}

	err = h.Repo.Update(kategori)
	if err != nil {
		fmt.Println("Gagal mengedit kategori:", err)
		return
	}
	fmt.Println("Kategori berhasil diupdate.")
}

// Hapus kategori
func (h *KategoriHandler) DeleteKategori() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Kategori yang akan dihapus: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	err = h.Repo.Delete(id)
	if err != nil {
		fmt.Println("Gagal menghapus kategori:", err)
		return
	}
	fmt.Println("Kategori berhasil dihapus.")
}

func NewKategoriCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kategori",
		Short: "Kelola kategori barang",
		// Add Run or subcommands here
	}
	return cmd
}
