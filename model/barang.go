package model

import "time"

type Barang struct {
	ID              int
	Nama            string
	Harga           float64
	TanggalBeli     time.Time
	KategoriID      int
	TotalPenggunaan int
}
