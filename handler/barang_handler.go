package handler

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"

	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/model"
	"project-app-inventaris-cli-Fauzan_Alsya_Prasetyo/repository"
)

type BarangHandler struct {
	Repo         *repository.BarangRepository
	KategoriRepo *repository.KategoriRepository
}

// Menampilkan daftar barang dalam bentuk tabel
func (h *BarangHandler) ListBarang() {
	barangs, err := h.Repo.GetAll()
	if err != nil {
		fmt.Println("Gagal mengambil data barang.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNama\tHarga\tTanggal Beli\tKategori ID\tTotal Penggunaan (hari)")
	for _, b := range barangs {
		fmt.Fprintf(w, "%d\t%s\t%.2f\t%s\t%d\t%d\n",
			b.ID, b.Nama, b.Harga, b.TanggalBeli.Format("2006-01-02"), b.KategoriID, b.TotalPenggunaan)
	}
	w.Flush()
}

// Menambah barang baru
func (h *BarangHandler) AddBarang() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Barang: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama == "" {
		fmt.Println("Nama barang tidak boleh kosong.")
		return
	}

	fmt.Print("Harga: ")
	hargaStr, _ := reader.ReadString('\n')
	hargaStr = strings.TrimSpace(hargaStr)
	harga, err := strconv.ParseFloat(hargaStr, 64)
	if err != nil {
		fmt.Println("Harga tidak valid.")
		return
	}

	fmt.Print("Tanggal Beli (YYYY-MM-DD): ")
	tanggalStr, _ := reader.ReadString('\n')
	tanggalStr = strings.TrimSpace(tanggalStr)
	tanggal, err := time.Parse("2006-01-02", tanggalStr)
	if err != nil {
		fmt.Println("Format tanggal salah.")
		return
	}

	fmt.Print("ID Kategori: ")
	kategoriIDStr, _ := reader.ReadString('\n')
	kategoriIDStr = strings.TrimSpace(kategoriIDStr)
	kategoriID, err := strconv.Atoi(kategoriIDStr)
	if err != nil {
		fmt.Println("ID kategori tidak valid.")
		return
	}

	// Validasi kategori
	_, err = h.KategoriRepo.GetByID(kategoriID)
	if err != nil {
		fmt.Println("Kategori tidak ditemukan.")
		return
	}

	barang := &model.Barang{
		Nama:            nama,
		Harga:           harga,
		TanggalBeli:     tanggal,
		KategoriID:      kategoriID,
		TotalPenggunaan: 0,
	}
	err = h.Repo.Create(barang)
	if err != nil {
		fmt.Println("Gagal menambah barang:", err)
		return
	}
	fmt.Println("Barang berhasil ditambahkan dengan ID:", barang.ID)
}

// Melihat detail barang berdasarkan ID
func (h *BarangHandler) DetailBarang() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Barang: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	barang, err := h.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	fmt.Println("ID               :", barang.ID)
	fmt.Println("Nama             :", barang.Nama)
	fmt.Printf("Harga            : %.2f\n", barang.Harga)
	fmt.Println("Tanggal Beli     :", barang.TanggalBeli.Format("2006-01-02"))
	fmt.Println("Kategori ID      :", barang.KategoriID)
	fmt.Println("Total Penggunaan :", barang.TotalPenggunaan, "hari")
}

// Edit barang
func (h *BarangHandler) EditBarang() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Barang yang akan diedit: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	barang, err := h.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	fmt.Printf("Nama lama (%s), masukkan nama baru (kosongkan jika tidak ingin mengubah): ", barang.Nama)
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	if nama != "" {
		barang.Nama = nama
	}

	fmt.Printf("Harga lama (%.2f), masukkan harga baru (kosongkan jika tidak ingin mengubah): ", barang.Harga)
	hargaStr, _ := reader.ReadString('\n')
	hargaStr = strings.TrimSpace(hargaStr)
	if hargaStr != "" {
		harga, err := strconv.ParseFloat(hargaStr, 64)
		if err == nil {
			barang.Harga = harga
		}
	}

	fmt.Printf("Tanggal Beli lama (%s), masukkan tanggal baru (YYYY-MM-DD, kosongkan jika tidak ingin mengubah): ", barang.TanggalBeli.Format("2006-01-02"))
	tanggalStr, _ := reader.ReadString('\n')
	tanggalStr = strings.TrimSpace(tanggalStr)
	if tanggalStr != "" {
		tanggal, err := time.Parse("2006-01-02", tanggalStr)
		if err == nil {
			barang.TanggalBeli = tanggal
		}
	}

	fmt.Printf("Kategori ID lama (%d), masukkan kategori ID baru (kosongkan jika tidak ingin mengubah): ", barang.KategoriID)
	kategoriIDStr, _ := reader.ReadString('\n')
	kategoriIDStr = strings.TrimSpace(kategoriIDStr)
	if kategoriIDStr != "" {
		kategoriID, err := strconv.Atoi(kategoriIDStr)
		if err == nil {
			// Validasi kategori
			_, err = h.KategoriRepo.GetByID(kategoriID)
			if err == nil {
				barang.KategoriID = kategoriID
			}
		}
	}

	fmt.Printf("Total Penggunaan lama (%d), masukkan total penggunaan baru (kosongkan jika tidak ingin mengubah): ", barang.TotalPenggunaan)
	totalStr, _ := reader.ReadString('\n')
	totalStr = strings.TrimSpace(totalStr)
	if totalStr != "" {
		total, err := strconv.Atoi(totalStr)
		if err == nil {
			barang.TotalPenggunaan = total
		}
	}

	err = h.Repo.Update(barang)
	if err != nil {
		fmt.Println("Gagal mengedit barang:", err)
		return
	}
	fmt.Println("Barang berhasil diupdate.")
}

// Hapus barang
func (h *BarangHandler) DeleteBarang() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Barang yang akan dihapus: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	err = h.Repo.Delete(id)
	if err != nil {
		fmt.Println("Gagal menghapus barang:", err)
		return
	}
	fmt.Println("Barang berhasil dihapus.")
}

// Menampilkan barang yang sudah digunakan lebih dari 100 hari
func (h *BarangHandler) ListBarangPerluDiganti() {
	barangs, err := h.Repo.GetAll()
	if err != nil {
		fmt.Println("Gagal mengambil data barang.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNama\tHari Penggunaan\tTanggal Beli\tKategori ID")
	found := false
	for _, b := range barangs {
		// Hitung hari penggunaan berdasarkan tanggal beli jika total_penggunaan belum diupdate
		hariPenggunaan := b.TotalPenggunaan
		if hariPenggunaan == 0 {
			hariPenggunaan = int(time.Since(b.TanggalBeli).Hours() / 24)
		}
		if hariPenggunaan > 100 {
			found = true
			fmt.Fprintf(w, "%d\t%s\t%d\t%s\t%d\n",
				b.ID, b.Nama, hariPenggunaan, b.TanggalBeli.Format("2006-01-02"), b.KategoriID)
		}
	}
	w.Flush()
	if !found {
		fmt.Println("Tidak ada barang yang perlu diganti (lebih dari 100 hari).")
	}
}

// Fungsi menghitung depresiasi saldo menurun 20% per tahun
func hitungDepresiasiSaldoMenurun(harga float64, tahun int, rate float64) float64 {
	nilai := harga
	for i := 0; i < tahun; i++ {
		nilai = nilai * (1 - rate)
	}
	return nilai
}

// Menampilkan total nilai investasi dan depresiasi seluruh barang
func (h *BarangHandler) LaporanInvestasiDepresiasi() {
	barangs, err := h.Repo.GetAll()
	if err != nil {
		fmt.Println("Gagal mengambil data barang.")
		return
	}

	totalInvestasi := 0.0
	totalSetelahDepresiasi := 0.0
	now := time.Now()

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNama\tHarga Awal\tTahun\tNilai Setelah Depresiasi")
	for _, b := range barangs {
		tahun := int(math.Floor(now.Sub(b.TanggalBeli).Hours() / 24 / 365))
		if tahun < 0 {
			tahun = 0
		}
		nilaiDepresiasi := hitungDepresiasiSaldoMenurun(b.Harga, tahun, 0.2)
		fmt.Fprintf(w, "%d\t%s\t%.2f\t%d\t%.2f\n", b.ID, b.Nama, b.Harga, tahun, nilaiDepresiasi)
		totalInvestasi += b.Harga
		totalSetelahDepresiasi += nilaiDepresiasi
	}
	w.Flush()
	fmt.Printf("\nTotal Investasi Awal: Rp%.2f\n", totalInvestasi)
	fmt.Printf("Total Setelah Depresiasi: Rp%.2f\n", totalSetelahDepresiasi)
}

// Menampilkan nilai investasi dan depresiasi untuk barang tertentu berdasarkan ID
func (h *BarangHandler) LaporanDepresiasiBarangByID() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan ID Barang: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID tidak valid.")
		return
	}

	barang, err := h.Repo.GetByID(id)
	if err != nil {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	now := time.Now()
	tahun := int(math.Floor(now.Sub(barang.TanggalBeli).Hours() / 24 / 365))
	if tahun < 0 {
		tahun = 0
	}
	nilaiDepresiasi := hitungDepresiasiSaldoMenurun(barang.Harga, tahun, 0.2)

	fmt.Println("ID               :", barang.ID)
	fmt.Println("Nama             :", barang.Nama)
	fmt.Printf("Harga Awal       : Rp%.2f\n", barang.Harga)
	fmt.Println("Tahun            :", tahun)
	fmt.Printf("Nilai Setelah Depresiasi: Rp%.2f\n", nilaiDepresiasi)
}

// Fitur pencarian barang berdasarkan nama (keyword)
func (h *BarangHandler) CariBarangByNama() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan keyword nama barang: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	if keyword == "" {
		fmt.Println("Keyword tidak boleh kosong.")
		return
	}

	barangs, err := h.Repo.GetAll()
	if err != nil {
		fmt.Println("Gagal mengambil data barang.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNama\tHarga\tTanggal Beli\tKategori ID\tTotal Penggunaan (hari)")
	found := false
	for _, b := range barangs {
		if strings.Contains(strings.ToLower(b.Nama), strings.ToLower(keyword)) {
			found = true
			fmt.Fprintf(w, "%d\t%s\t%.2f\t%s\t%d\t%d\n",
				b.ID, b.Nama, b.Harga, b.TanggalBeli.Format("2006-01-02"), b.KategoriID, b.TotalPenggunaan)
		}
	}
	w.Flush()
	if !found {
		fmt.Println("Barang dengan keyword tersebut tidak ditemukan.")
	}
}

func NewBarangCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "barang",
		Short: "Kelola barang inventaris",
		// Add Run or subcommands here
	}
	return cmd
}
