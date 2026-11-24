
# Program Kasir Sederhana

## Deskripsi
Program Kasir Sederhana merupakan program berbasis
command line yang dibuat menggunakan bahasa pemrograman Go (Golang).

Program ini digunakan untuk menghitung total belanja
berdasarkan harga satuan dan jumlah barang yang dibeli,
serta menghitung uang kembalian berdasarkan jumlah uang
yang dibayarkan oleh pembeli.

## Fitur
- Memasukkan nama barang.
- Memasukkan harga satuan barang.
- Memasukkan jumlah barang.
- Memasukkan jumlah uang yang dibayarkan.
- Menghitung total belanja.
- Menghitung uang kembalian.

## Teknologi
- Bahasa Pemrograman: Go (Golang)
- Input dan Output: Package fmt

## Cara Menjalankan

1. Pastikan Go sudah terinstal di komputer.
2. Clone repository ini:

   ```bash
   git clone URL_REPOSITORY
   ```

3. Masuk ke folder project:

   ```bash
   cd nama-folder
   ```

4. Jalankan program:

   ```bash
   go run main.go
   ```

## Contoh Penggunaan

```text
Masukkan nama barang: Buku
Masukkan harga satuan: 5000
Masukkan jumlah barang: 3
Masukkan jumlah uang yang dibayarkan: 20000

Nama barang: Buku
Total belanja: 15000
Uang kembalian: 5000
```

## Rumus Perhitungan

Total Belanja:

harga satuan × jumlah barang

Uang Kembalian:

uang dibayarkan - total belanja

## Catatan
Program ini merupakan project sederhana untuk
mempelajari dasar-dasar pemrograman Go, khususnya
penggunaan variabel, input, output, dan operasi aritmatika.
