--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- PostgreSQL database dump complete
--

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