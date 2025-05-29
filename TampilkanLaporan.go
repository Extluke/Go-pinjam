package main

import "fmt"

func TampilkanLaporan(dataNasabah []Nasabah) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║            📄 LAPORAN PINJAMAN           ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	if len(dataNasabah) == 0 {
		fmt.Println("⚠️ Tidak ada data nasabah.")
		return
	} else {
		for _, nasabah := range dataNasabah {
			fmt.Println("╔══════════════════════════════════════════╗")
			fmt.Printf("║ 👤 Nama             : %s\n", nasabah.Nama)
			fmt.Printf("║ 💰 Jumlah Pinjaman  : Rp.%.2f\n", nasabah.JumlahPinjaman)
			fmt.Printf("║ 🕒 Tenor            : %d Bulan\n", nasabah.Tenor)
			fmt.Println("╚══════════════════════════════════════════╝")
		}
	}
}
