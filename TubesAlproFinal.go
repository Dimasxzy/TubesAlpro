package main

import (
    "fmt"
    "time"
)

type Penggunaan struct {
    Nama     string
    Komputer int
    Mulai    string
    Selesai  string
    Durasi   string
    Biaya    int
}

const (
    tarifPerJam = 15000
    maxData     = 100
)

var (
    data       [maxData]Penggunaan
    jumlahData int
)

func main() {
    format := "15:04"

    for {
        tampilkanMenu()

        var pilihan int
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            tambahData(format)
        case 2:
            tampilkanData()
        case 3:
            updateData()
        case 4:
            hapusData()
        case 5:
            cariData()
        case 6:
            urutkanDataBiaya(true)
        case 7:
            urutkanDataBiaya(false)
        case 8:
            resetData()
        case 9:
            fmt.Println("👋 Terima kasih, program selesai.")
            return
        default:
            fmt.Println("❗ Pilihan tidak tersedia. Silakan coba lagi.\n")
        }
    }
}

func tampilkanMenu() {
    fmt.Println("\n╔════════════════════════════════════════╗")
    fmt.Println("║ 💻  SISTEM PENGGUNAAN KOMPUTER         ║")
    fmt.Println("╠════════════════════════════════════════╣")
    fmt.Println("║ 1. ➕ Tambah Data                      ║")
    fmt.Println("║ 2. 📄 Lihat Data                       ║")
    fmt.Println("║ 3. ✏️  Edit Data                       ║")
    fmt.Println("║ 4. 🗑️  Hapus Data                      ║")
    fmt.Println("║ 5. 🔍 Cari Data (Nama)                 ║")
    fmt.Println("║ 6. 📈 Urutkan Biaya (Naik)             ║")
    fmt.Println("║ 7. 📉 Urutkan Biaya (Turun)            ║")
    fmt.Println("║ 8. 🔄 Reset Data                       ║")
    fmt.Println("║ 9. ❌ Keluar Program                   ║")
    fmt.Println("╚════════════════════════════════════════╝")
    fmt.Print("Pilih menu: ")
}

func tambahData(format string) {
    if jumlahData >= maxData {
        fmt.Println("❌ Data penuh, tidak dapat menambah lagi.\n")
        return
    }

    var nama string
    var komputer int
    var mulai, selesai string

    fmt.Print("Nama: ")
    fmt.Scanln(&nama)

    fmt.Print("Nomor Komputer: ")
    fmt.Scanln(&komputer)

    fmt.Print("Waktu Mulai (HH:MM): ")
    fmt.Scanln(&mulai)
    fmt.Print("Waktu Selesai (HH:MM): ")
    fmt.Scanln(&selesai)

    tMulai, err1 := time.Parse(format, mulai)
    tSelesai, err2 := time.Parse(format, selesai)

    if err1 != nil || err2 != nil {
        fmt.Println("❌ Format waktu salah!\n")
        return
    }

    if tSelesai.Before(tMulai) {
        tSelesai = tSelesai.Add(24 * time.Hour)
    }

    durasi := tSelesai.Sub(tMulai)
    jam := durasi.Hours()

    biaya := int(jam * tarifPerJam)
    if biaya == 0 {
        biaya = tarifPerJam
    }

    data[jumlahData] = Penggunaan{
        Nama:     nama,
        Komputer: komputer,
        Mulai:    mulai,
        Selesai:  selesai,
        Durasi:   durasi.String(),
        Biaya:    biaya,
    }
    jumlahData++
    fmt.Println("✅ Data berhasil ditambahkan!\n")
}

func tampilkanData() {
    if jumlahData == 0 {
        fmt.Println("📂 Belum ada data penggunaan.\n")
        return
    }
    fmt.Println("🗂️  === Daftar Penggunaan Komputer ===")
    for i := 0; i < jumlahData; i++ {
        d := data[i]
        fmt.Printf("%d. %s | PC #%d | %s - %s (%s) | Rp%d\n",
            i+1, d.Nama, d.Komputer, d.Mulai, d.Selesai, d.Durasi, d.Biaya)
    }
}

func cariIndex(nama string) int {
    for i := 0; i < jumlahData; i++ {
        if data[i].Nama == nama {
            return i
        }
    }
    return -1
}

func updateData() {
    var nama string
    fmt.Print("Masukkan nama yang ingin diubah: ")
    fmt.Scanln(&nama)
    idx := cariIndex(nama)
    if idx == -1 {
        fmt.Println("❌ Data tidak ditemukan.")
        return
    }
    fmt.Print("Nama baru: ")
    fmt.Scanln(&data[idx].Nama)
    fmt.Print("Nomor Komputer baru: ")
    fmt.Scanln(&data[idx].Komputer)
    fmt.Println("✅ Data berhasil diperbarui!\n")
}

func hapusData() {
    var nama string
    fmt.Print("Masukkan nama yang ingin dihapus: ")
    fmt.Scanln(&nama)
    idx := cariIndex(nama)
    if idx == -1 {
        fmt.Println("❌ Data tidak ditemukan.")
        return
    }
    for i := idx; i < jumlahData-1; i++ {
        data[i] = data[i+1]
    }
    jumlahData--
    fmt.Println("🗑️  Data berhasil dihapus!\n")
}

func cariData() {
    var nama string
    fmt.Print("Masukkan nama yang dicari: ")
    fmt.Scanln(&nama)
    idx := cariIndex(nama)
    if idx == -1 {
        fmt.Println("❌ Data tidak ditemukan.")
    } else {
        d := data[idx]
        fmt.Printf("Ditemukan: %s | PC #%d | %s - %s (%s) | Rp%d\n",
            d.Nama, d.Komputer, d.Mulai, d.Selesai, d.Durasi, d.Biaya)
    }
}

func urutkanDataBiaya(naik bool) {
    for i := 0; i < jumlahData-1; i++ {
        idx := i
        for j := i + 1; j < jumlahData; j++ {
            if (naik && data[j].Biaya < data[idx].Biaya) || (!naik && data[j].Biaya > data[idx].Biaya) {
                idx = j
            }
        }
        data[i], data[idx] = data[idx], data[i]
    }
    fmt.Println("✅ Data berhasil diurutkan berdasarkan biaya.")
}

func resetData() {
    jumlahData = 0
    fmt.Println("🔁 Semua data telah direset.\n")
}