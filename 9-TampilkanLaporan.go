package main

import "fmt"

func TampilkanLaporan(dataNasabah []Nasabah) {
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║         📄 LAPORAN PINJAMAN          ║")
	fmt.Println("╚══════════════════════════════════════╝")
	if len(dataNasabah) == 0 {
		fmt.Println("⚠️ Tidak ada data nasabah.")
		return
	} else {
		fmt.Println("╔══════════════════════════════════════╗")
		for _, nasabah := range dataNasabah {
			progres := 100 - (float64(nasabah.StatusPembayaran) / float64(nasabah.Tenor) * 100)
			fmt.Printf("║ 👤 Nama               : %s\n", nasabah.Nama)
			fmt.Printf("║ 💰 Jumlah Pinjaman    : Rp.%.2f\n", nasabah.JumlahPinjaman)
			fmt.Printf("║ 🕒 Tenor              : %d Bulan\n", nasabah.Tenor)
			fmt.Printf("║ 💵 Status Pembayaran  : Kurang %d Bulan\n", nasabah.StatusPembayaran)
			fmt.Printf("║ 📊 Progres Pembayaran : %.f%%\n", progres)
			fmt.Println("╠──────────────────────────────────────╣")
		}
		fmt.Println("╚══════════════════════════════════════╝")
	}
}
