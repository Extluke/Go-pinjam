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

func binarySearch(data []Nasabah, target string) int {
	left, right := 0, len(data)-1
	for left <= right {
		mid := (left + right) / 2
		if data[mid].Nama == target {
			return mid
		} else if data[mid].Nama < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("❌ Data tidak ditemukan")
	return -1
}

func binarySelectionSort(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		indexTerkecil := i
		for j := i + 1; j < n; j++ {
			if data[j].Nama < data[indexTerkecil].Nama {
				indexTerkecil = j
			}
		}
		data[i], data[indexTerkecil] = data[indexTerkecil], data[i]
	}
}

func CariPeminjam(dataNasabah []Nasabah) {
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║       🔍 Pilih Metode Pencarian      ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║ 1. 🎰 Sequential Search              ║")
	fmt.Println("║ 2. 🎲 Binary Search                  ║")
	fmt.Println("╚══════════════════════════════════════╝")

	var pilihan int
	for {
		fmt.Print("👉 Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan == 1 || pilihan == 2 {
			break
		}
		fmt.Println("⚠️  Input tidak valid, silakan pilih ulang.")
	}

	var nama string
	fmt.Print("Masukkan Nama Peminjam: ")
	fmt.Scanln(&nama)

	var index int
	if pilihan == 1 {
		index = sequentialSearch(dataNasabah, nama)
	} else {
		binarySelectionSort(dataNasabah)
		index = binarySearch(dataNasabah, nama)
	}

	if index != -1 {
		nasabah := dataNasabah[index]
		fmt.Println("\n📄 Data Peminjam Ditemukan:")
		fmt.Println("╔══════════════════════════════════════╗")
		fmt.Printf("║ 👤 Nama             : %s\n", nasabah.Nama)
		fmt.Printf("║ 💸 Jumlah Pinjaman  : Rp.%.2f\n", nasabah.JumlahPinjaman)
		fmt.Printf("║ 🕒 Tenor            : %d bulan\n", nasabah.Tenor)
		fmt.Println("╚══════════════════════════════════════╝")
	}
}
