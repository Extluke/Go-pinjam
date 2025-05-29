package main

import "fmt"

func TampilkanLaporan(dataNasabah []Nasabah) {
	fmt.Println("\n+-------------------------------+")
	fmt.Println("|         Laporan Pinjaman      |")
	fmt.Println("+-------------------------------+")
	if len(dataNasabah) == 0 {
		fmt.Println("|    Tidak ada data nasabah     |")
		return
	} else {
		for _, nasabah := range dataNasabah {
			fmt.Printf("| Nama: %s\n", nasabah.Nama)
			fmt.Printf("| Jumlah Pinjaman: %.2f\n", nasabah.JumlahPinjaman)
			fmt.Printf("| Tenor: %d bulan\n", nasabah.Tenor)
			fmt.Println("+-------------------------------+")
		}
	}
}