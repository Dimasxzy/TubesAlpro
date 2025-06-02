🧠 Tujuan Program
Program ini digunakan untuk mencatat, mengelola, dan menganalisis penggunaan komputer warnet atau laboratorium dengan berbagai fitur seperti:

•	Menambah, mengedit, menghapus, dan mencari data pengguna

•	Menghitung durasi dan biaya penggunaan

•	Menyediakan laporan harian, rata-rata durasi, dan komputer aktif
________________________________________
📦 Struktur Program
1. Import Package
import (
    "fmt"
    "time"
)

•	fmt: Untuk input/output (misal: fmt.Println, fmt.Scanln)

•	time: Untuk manipulasi waktu (time.Parse, time.Duration, dll.)
________________________________________
2. Struct Penggunaan
type Penggunaan struct {

   Nama     string

   Komputer int

   Mulai    string

   Selesai  string

   Durasi   string

   Biaya    int
}
Struktur data untuk menyimpan satu entri penggunaan komputer:

•	Nama: Nama pengguna

•	Komputer: Nomor komputer

•	Mulai, Selesai: Jam mulai dan selesai (dalam format string)

•	Durasi: Lama pemakaian (string hasil dari Duration.String())

•	Biaya: Total biaya penggunaan
________________________________________
3. Konstanta & Variabel Global
const (
    tarifPerJam = 15000
    maxData     = 100
)
var (
    data       [maxData]Penggunaan
    jumlahData int
)

•	tarifPerJam: Biaya per jam komputer

•	data: Array berukuran maksimal 100 untuk menyimpan data penggunaan

•	jumlahData: Jumlah data yang sedang tersimpan
________________________________________
4. Fungsi main
Fungsi utama program:

•	Menampilkan menu secara berulang

•	Menangani pilihan menu dari user menggunakan switch
________________________________________
🧭 Fungsi-Fungsi Menu
✅ tampilkanMenu()
Menampilkan daftar menu pilihan.
________________________________________
✅ tambahData(format string)

•	Meminta input nama, nomor komputer, waktu mulai dan selesai

•	Menghitung durasi dan biaya berdasarkan tarif

•	Menyimpan data ke array data

Catatan: Jika waktu selesai < mulai, dianggap lewat tengah malam (+24 jam).
________________________________________
✅ tampilkanData()
Menampilkan semua data penggunaan komputer.
________________________________________
✅ cariIndex(nama string) int
Mencari indeks array berdasarkan nama pengguna. Return -1 jika tidak ditemukan.
________________________________________
✅ updateData()

•	Cari data berdasarkan nama

•	Jika ditemukan, minta input nama & komputer baru

•	Tidak memperbarui waktu atau biaya (bisa ditambahkan nanti)
________________________________________
✅ hapusData()

•	Cari data berdasarkan nama

•	Hapus dengan cara menggeser data setelahnya ke kiri
________________________________________
✅ cariData()

•	Mencari dan menampilkan data berdasarkan nama
________________________________________
✅ urutkanDataBiaya(naik bool)

•	Mengurutkan data berdasarkan biaya (naik atau turun)

•	Menggunakan algoritma Selection Sort
________________________________________
✅ resetData()

•	Menghapus semua data dengan mengatur jumlahData = 0
________________________________________
✅ laporanHarian()

•	Menampilkan total pengguna & total pendapatan hari itu

•	Mengakumulasi semua Biaya dari array data
________________________________________
✅ rataRataDurasi()

•	Menghitung total durasi semua pengguna

•	Membagi dengan jumlahData untuk mendapatkan rata-rata durasi
________________________________________
✅ daftarKomputerAktif()

•	Menampilkan daftar komputer yang digunakan hari ini

•	Menggunakan map[int]bool untuk memastikan unik
________________________________________
📝 Contoh Output
Jika pengguna menambahkan 2 data:

1. Budi | PC #1 | 09:00 - 10:30 (1h30m0s) | Rp22500

2. Siti | PC #2 | 08:15 - 09:15 (1h0m0s) | Rp15000

Laporan harian:
Total Pengguna: 2
Total Pendapatan: Rp37500
________________________________________
🧠 Catatan Tambahan

•	Durasi dihitung dengan time.Sub, dan dikonversi ke jam untuk hitung biaya.

•	Jika hasil biaya = 0 (durasi < 1 jam), tetap dikenakan minimum 1 jam.

•	Fungsi rataRataDurasi() memakai time.ParseDuration untuk mengubah string ke Duration lagi.
