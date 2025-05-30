package main

import "fmt"

func sequentialSearch(data []Nasabah, target string) int {
	for i, nasabah := range data {
		if nasabah.Nama == target {
			return i
		}
	}

	fmt.Println("❌ Data tidak ditemukan")
	return -1 
}

func CariPeminjam(dataNasabah []Nasabah) {
	var target string

	fmt.Println("╔═══════════════════════════════════════")
	fmt.Print("║ 🔍 Masukkan nama yang ingin dicari : ")
	fmt.Scanln(&target)
	fmt.Println("╚═══════════════════════════════════════")

	index := sequentialSearch(dataNasabah, target)
	if index == -1 {
		return
	}

	nasabah := dataNasabah[index]
	fmt.Println("\n📄 Data Peminjam Ditemukan:")
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Printf("║ 👤 Nama             : %s\n", nasabah.Nama)
	fmt.Printf("║ 💸 Jumlah Pinjaman  : Rp.%.2f\n", nasabah.JumlahPinjaman)
	fmt.Printf("║ 🕒 Tenor            : %d bulan\n", nasabah.Tenor)
	fmt.Println("╚══════════════════════════════════════╝")
}
