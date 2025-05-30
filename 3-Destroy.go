package main

import "fmt"

func Destroy(dataNasabah []Nasabah) []Nasabah {
	var target, konfirmasi string
	for {
		fmt.Println("╔═══════════════════════════════════════")
		fmt.Print("║ 🗑️ Nama Nasabah Untuk Dihapus : ")
		fmt.Scan(&target)
		fmt.Println("╚═══════════════════════════════════════")

		hasil := sequentialSearch(dataNasabah, target)
		
		if hasil == -1 {
			fmt.Println("⚠️  Nama tidak ditemukan. Coba lagi.")
			continue
		}

		for {
			fmt.Printf("\n❓ Apakah yakin ingin menghapus \"%s\"? [y/n]: ", dataNasabah[hasil].Nama)
			fmt.Scan(&konfirmasi)

			if konfirmasi != "y" && konfirmasi != "n" {
				fmt.Println("⚠️  Input tidak valid! Pilih 'y' atau 'n'.")
				continue
			}
			
			if konfirmasi == "y" {
				dataNasabah = append(dataNasabah[:hasil], dataNasabah[hasil+1:]...)
				fmt.Println("\n✅ Data berhasil dihapus.")
			} else {
				fmt.Println("\n❎ Penghapusan dibatalkan.")
			}
			break
		}

		break
	}

	return dataNasabah
}
