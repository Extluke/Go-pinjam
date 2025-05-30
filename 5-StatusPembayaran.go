package main

import "fmt"

func StatusPembayaran (data []Nasabah) {
 var target string
 hasil := -1

	for {
		fmt.Println("╔═══════════════════════════════════════")
		fmt.Print("║ 👤 Nama Nasabah Yang Akan Diinput Status Pembayarannya  : ")
		fmt.Scan(&target)
		fmt.Println("╚═══════════════════════════════════════")

		hasil = sequentialSearch(dataNasabah, target)
		if hasil == -1 {
			fmt.Println("⚠️  Nama tidak ditemukan. Coba lagi.")
			continue
		}
		break
	}

	fmt.Printf("\n🎉 ✅Data Nasabah Behasil ditemukan\n")

	bulan := TambahkanStatusPembayaran(dataNasabah, hasil)

	dataNasabah[hasil].StatusPembayaran = bulan

	TampilkanStatusPembayaran(dataNasabah, hasil)

}

func TampilkanStatusPembayaran(dataNasabah []Nasabah, hasil int) {
	SisaCicilan := dataNasabah[hasil].Tenor - dataNasabah[hasil].StatusPembayaran

	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Printf("║ 👤 Nama              : %s\n", dataNasabah[hasil].Nama)
	fmt.Printf("║ 💰 Jumlah Pinjaman   : Rp.%.2f\n", dataNasabah[hasil].JumlahPinjaman)
	fmt.Printf("║ 🕒 Tenor             : %d Bulan\n", dataNasabah[hasil].Tenor)
	if dataNasabah[hasil].StatusPembayaran < dataNasabah[hasil].Tenor {
			fmt.Printf("║ 💵 Status Pembayaran : Belum Lunas, Tersisa %d dari %d cicilan tersisa\n", SisaCicilan, dataNasabah[hasil].Tenor )
		} else if dataNasabah[hasil].StatusPembayaran == dataNasabah[hasil].Tenor {
			fmt.Printf("║ 💵 Status Pembayaran : Lunas\n")
		}
	fmt.Println("╚══════════════════════════════════════════╝")
}

func TambahkanStatusPembayaran(dataNasabah []Nasabah, hasil int) int {
	var bulan int
	for {
		fmt.Println("╔═══════════════════════════════════════")
		fmt.Print("║ 💲Masukkan Cicilan Yang sudah Dibayar : ")
		fmt.Scan(&bulan)
		fmt.Println("╚═══════════════════════════════════════")
		
		if bulan < 0 {
			fmt.Printf("\n ⚠️ Mohon Masukkan Inputan yang tidak lebih kecil dari 0 \n")
			continue
		} 

		if bulan > dataNasabah[hasil].Tenor {
			fmt.Printf("\n ⚠️ Mohon Masukkan Inputan yang tidak lebih besar dari Tenor \n")
			continue
		}

		break
	}
	return bulan
}