package main

import "fmt"

func Destroy(dataNasabah []Nasabah) []Nasabah {
	var target, konfirmasi string
	for {
		fmt.Print("Masukkan Nama Nasabah Yang Datanya Ingin Dihapus : ")
		fmt.Scan(&target)
		hasil := sequentialSearch(dataNasabah, target)
		
		if hasil == -1 {
			continue
		}
			
		
		for {
				fmt.Printf("Apa kamu yakin untuk menghapus %s [y/n] : ", dataNasabah[hasil].Nama)
				fmt.Scan(&konfirmasi)
				if konfirmasi != "y" && konfirmasi != "n" {
					fmt.Println("\ninput tidak valid!")
					continue
				} 
				if konfirmasi == "y" {
					dataNasabah = append(dataNasabah[:hasil], dataNasabah[hasil+1:]...)
				} else {
					fmt.Println("\nPenghapusan Dibatalkan")
				}	
			break
		}

		break
	}

	return dataNasabah
}