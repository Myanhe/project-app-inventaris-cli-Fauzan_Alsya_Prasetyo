-- Tabel kategori
CREATE TABLE IF NOT EXISTS kategori (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL UNIQUE,
    deskripsi TEXT
);

-- Tabel barang
CREATE TABLE IF NOT EXISTS barang (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    harga NUMERIC(15,2) NOT NULL,
    tanggal_beli DATE NOT NULL,
    kategori_id INTEGER NOT NULL REFERENCES kategori(id) ON DELETE CASCADE,
    total_penggunaan INTEGER DEFAULT 0
);

-- Index untuk pencarian nama barang
CREATE INDEX IF NOT EXISTS idx_barang_nama ON barang(nama);