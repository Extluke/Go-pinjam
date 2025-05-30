package main

import (
	"fmt"
)

func UbahData(dataNasabah []Nasabah) {
	var target string

	for {
		fmt.Println("\n╔══════════════════════════════════════╗")
		fmt.Println("║     ✏️ Nasabah Yang Ingin Diubah      ║")
		fmt.Println("╚══════════════════════════════════════╝")
		fmt.Print("👉 Nama: ")
		fmt.Scanln(&target)

		hasil := sequentialSearch(dataNasabah, target)
		if hasil == -1 {
			fmt.Println("⚠️ Nama tidak ditemukan. Silakan coba lagi.")
			continue
		} else {
			fmt.Println("\n🔄 Silakan masukkan data baru untuk nasabah ini:")
			nasabahBaru := Tambah(dataNasabah)

			dataNasabah[hasil].Nama = nasabahBaru.Nama
			dataNasabah[hasil].JumlahPinjaman = nasabahBaru.JumlahPinjaman
			dataNasabah[hasil].Tenor = nasabahBaru.Tenor

			if dataNasabah[hasil].StatusPembayaran > 0 {
				bulan := TambahkanStatusPembayaran(dataNasabah, hasil)
				dataNasabah[hasil].StatusPembayaran = bulan
			}

			fmt.Println("\n✅ Data Nasabah Sukses Diperbarui!")
			break
		}
	}
}
