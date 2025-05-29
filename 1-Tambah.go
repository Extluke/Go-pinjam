package main

import (
	"fmt"
)

func Tambah(dataNasabah []Nasabah) Nasabah {
	var nama string
	var jumlah float64
	var tenor, StatusPembayaran int
	
for {
	fmt.Print("Masukkan Nama : ")
	fmt.Scanln(&nama)
	
	if nama == "" {
		fmt.Println("Nama tidak boleh kosong.")
		continue
	}

		// Validasi duplikat
		duplikat := false
		for _, n := range dataNasabah {
			if n.Nama == nama {
				fmt.Println("⚠️  Maaf, nama sudah digunakan. Mohon masukkan nama yang lain.")
				duplikat = true
				break
			}
		}
		if !duplikat {
			break
		}
	}

	// Input jumlah pinjaman
	for {
		fmt.Print("║ Masukkan Nominal Yang Diajukan : ")
		_, err := fmt.Scanln(&jumlah)
		fmt.Println("╚═══════════════════════════════════════")
		
		if err != nil {
			fmt.Println("⚠️  Input tidak valid. Masukkan angka.")
			continue
		}
		break
	}

	// Pilihan tenor
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║            🕒 PILIH TENOR            ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║ 1. Tenor 3 Bulan                     ║")
	fmt.Println("║ 2. Tenor 6 Bulan                     ║")
	fmt.Println("║ 3. Tenor 1 Tahun                     ║")
	fmt.Println("╚══════════════════════════════════════╝")
	for {
		fmt.Print("👉 Pilihan: ")
		_, err := fmt.Scanln(&tenor)
		if err != nil || tenor < 1 || tenor > 3 {
			fmt.Println("⚠️  Pilihan tenor tidak valid, silakan pilih ulang.")
			continue
		}
		break
	}

	// Konversi pilihan ke bulan
	var tenorBulan int
	switch tenor {
	case 1:
		tenorBulan = 3
	case 2:
		tenorBulan = 6
	case 3:
		tenorBulan = 12
	}

	if StatusPembayaran > 0 {
		
	}

	fmt.Println("\n🎉 ✅ Berhasil Ditambahkan:")
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Printf("║ Nama            : %s\n", nama)
	fmt.Printf("║ Jumlah Pinjaman : Rp.%.0f\n", jumlah)
	fmt.Printf("║ Tenor           : %d bulan\n", tenorBulan)
	fmt.Println("╚══════════════════════════════════════╝")

	return Nasabah{
		Nama:           nama,
		JumlahPinjaman: jumlah,
		Tenor:          tenorBulan,
	}
}
