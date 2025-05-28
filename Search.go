package main

import "fmt"

// untuk func sequential search
func sequentialSearch(data []Nasabah, target string) int {
	for i, n := range data {
		if n.Nama == target {
      return i
		}
	}

  fmt.Println("❌ Data tidak ditemukan")
	return -1 
}

func CariPeminjam(dataNasabah []Nasabah) {
	var target string

	fmt.Print("Masukkan nama yang ingin dicari: ")
	fmt.Scanln(&target)

	index := sequentialSearch(dataNasabah, target)
	if index == -1 {
		return
	}

	n := dataNasabah[index]
	fmt.Println("\n+-------------------------------+")
	fmt.Println("|         Data Peminjam         |")
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Nama: %s\n", n.Nama)
	fmt.Printf("| Jumlah Pinjaman: %.2f\n", n.JumlahPinjaman)
	fmt.Printf("| Tenor: %d bulan\n", n.Tenor)
	fmt.Println("+-------------------------------+")
}
