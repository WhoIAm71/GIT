package main

import "fmt"

func main() {
	var namaBarang string
	var hargaSatuan, jumlahBarang, uangDibayar int

	fmt.Print("Masukkan nama barang: ")
	fmt.Scan(&namaBarang)

	fmt.Print("Masukkan harga satuan: ")
	fmt.Scan(&hargaSatuan)

	fmt.Print("Masukkan jumlah barang: ")
	fmt.Scan(&jumlahBarang)

	fmt.Print("Masukkan jumlah uang yang dibayarkan: ")
	fmt.Scan(&uangDibayar)

	totalBelanja := hargaSatuan * jumlahBarang
	uangKembalian := uangDibayar - totalBelanja

	fmt.Printf("Nama barang: %s\n", namaBarang)
	fmt.Printf("Total belanja: %d\n", totalBelanja)
	fmt.Printf("Uang kembalian: %d\n", uangKembalian)
}
