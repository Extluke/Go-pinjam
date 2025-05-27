package main

import (
	"fmt"
)

func UbahData(dataNasabah []Nasabah) {
	var target string

	for {
		fmt.Println("\nMasukkan Nama Nasabah Yang Ingin Diubah : ")
		fmt.Scanln(&target)

		hasil := sequentialSearch(dataNasabah, target)
		if hasil == -1 {
			continue
		} else {
			nasabahBaru := Tambah(dataNasabah)

			dataNasabah[hasil].Nama = nasabahBaru.Nama
			dataNasabah[hasil].JumlahPinjaman = nasabahBaru.JumlahPinjaman
			dataNasabah[hasil].Tenor = nasabahBaru.Tenor
			
			println("\n✅Data Nasabah Sukses Diperbarui!")

			break
		}
	}
}
