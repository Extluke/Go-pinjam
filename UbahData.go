package main

import (
	"fmt"
)

func UbahData(dataNasabah []Nasabah) {
	var target string

	for {
		fmt.Println("\n╔══════════════════════════════════════╗")
		fmt.Println("║     ✏️ Nasabah Yang Ingin Diubah     ║")
		fmt.Println("╚══════════════════════════════════════╝")
		fmt.Print("👉 Nama: ")
		fmt.Scanln(&target)

		hasil := sequentialSearch(dataNasabah, target)
		if hasil == -1 {
			fmt.Println("⚠️ Nama tidak ditemukan. Silakan coba lagi.")
			continue
		} else {
<<<<<<< HEAD
			fmt.Println("✅Nama Ditemukan")
=======
			fmt.Println("\n🔄 Silakan masukkan data baru untuk nasabah ini:")
>>>>>>> 8effdf20a5437c5de2d16a1b0971e242d396901f
			nasabahBaru := Tambah(dataNasabah)

			dataNasabah[hasil].Nama = nasabahBaru.Nama
			dataNasabah[hasil].JumlahPinjaman = nasabahBaru.JumlahPinjaman
			dataNasabah[hasil].Tenor = nasabahBaru.Tenor

			fmt.Println("\n✅ Data Nasabah Sukses Diperbarui!")
			break
		}
	}
}
