Presentasi Program Penggunaan Komputer Warnet (Bahasa Go)

📌 1. Modularitas (Menggunakan Prosedur dan Fungsi)
Program ini disusun secara modular dengan membagi setiap fitur ke dalam fungsi-fungsi terpisah seperti:

•	tampilkanMenu() | untuk menampilkan menu utama.

•	tambahData() | untuk memasukkan data baru.

•	tampilkanData() | untuk melihat seluruh data.

•	updateData() | untuk mengedit data.

•	hapusData() | untuk menghapus data.

•	cariData() | untuk pencarian data.

•	urutkanDataBiaya() | untuk mengurutkan data berdasarkan biaya.

•	resetData() | untuk menghapus semua data.
________________________________________
📌 2. Array Statis dan Tipe Bentukan
Program menggunakan:

•	Array statis: var data [maxData]Penggunaan untuk menyimpan data dengan kapasitas maksimum 100.

•	Tipe bentukan: type Penggunaan struct { ... } untuk merepresentasikan satu entri penggunaan komputer.
________________________________________
📌 3. Fitur CRUD Lengkap
Program mendukung fitur:

•	Create: dengan fungsi tambahData().

•	Read: dengan fungsi tampilkanData().

•	Update: dengan fungsi updateData(), memungkinkan pengguna mengedit nama dan nomor komputer.

•	Delete: dengan fungsi hapusData(), menghapus data berdasarkan nama pengguna.
________________________________________
📌 4. Pencarian dengan Sequential Search
Pencarian dilakukan menggunakan sequential search dalam fungsi cariIndex(nama string), yang digunakan juga oleh updateData(), hapusData(), dan cariData().
________________________________________
📌 5. Pengurutan dengan Selection Sort
Pengurutan data berdasarkan biaya penggunaan komputer menggunakan algoritma selection sort:

•	urutkanDataBiaya(true) untuk pengurutan menaik.

•	urutkanDataBiaya(false) untuk pengurutan menurun.
________________________________________
📌 6. Perhitungan Sederhana
Program menghitung durasi penggunaan komputer dan biaya yang harus dibayar berdasarkan durasi dengan tarif Rp15.000 per jam:
biaya := int(jam * tarifPerJam)
Jika waktu penggunaan kurang dari 1 jam, tetap dikenakan biaya minimum 1 jam.
________________________________________
📌 7. Fitur Opsional: Akun/Registrasi
Fitur akun tidak diterapkan karena sifat program hanya fokus pada pencatatan penggunaan komputer harian secara langsung.
________________________________________
📌 8. Penggunaan Break/Continue
Program tidak menggunakan break atau continue secara langsung, dan hanya menggunakan struktur kontrol standar seperti if, for, dan switch.
________________________________________
📌 9. Variabel Global
Variabel global yang digunakan hanya untuk:

•	data [maxData]Penggunaan (array statis)

•	jumlahData int (penghitung jumlah data yang valid)

Sesuai dengan aturan, variabel global dibatasi hanya untuk array dan tipe bentukan.
________________________________________
✅ Kesimpulan
Program ini sudah sesuai dengan seluruh ketentuan yang diberikan:

•	Modular

•	Menggunakan array statis dan struct

•	Memiliki fitur CRUD

•	Pencarian & pengurutan sesuai algoritma dasar

•	Ada proses perhitungan biaya

•	Tidak melanggar batasan penggunaan variabel global dan struktur kontrol
