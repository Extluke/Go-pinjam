package main

import "fmt"

func SimulasiPinjaman(dataNasabah []Nasabah) {
	var target string
	var bunga int

	hasil := -1

	for {
		fmt.Println("╔═══════════════════════════════════════")
		fmt.Print("║ 👤 Nama Nasabah Yang Akan Disimulasi : ")
		fmt.Scan(&target)
		fmt.Println("╚═══════════════════════════════════════")

		hasil = sequentialSearch(dataNasabah, target)
		if hasil == -1 {
			fmt.Println("⚠️  Nama tidak ditemukan. Coba lagi.")
			continue
		}
		break
	}

	fmt.Scanln()

	// Pilihan Bunga yang akan diambil
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║        📊 Pilih Jenis Bunga          ║")
	for {
		fmt.Println("╠══════════════════════════════════════╣")
		fmt.Println("║ 1. 💰 Bunga Tetap                    ║")
		fmt.Println("║ 2. 📉 Bunga Variabel                 ║")
		fmt.Println("╚══════════════════════════════════════╝")
		fmt.Print("👉 Pilihan: ")
		_, err := fmt.Scanln(&bunga)
		if err != nil || bunga < 1 || bunga > 2 {
			fmt.Println("⚠️  Input tidak valid, silakan pilih ulang.")
			continue
		}
		break
	}

	if bunga == 1 {
		Tetap(dataNasabah[hasil])
	} else if bunga == 2 {
		Variabel(dataNasabah[hasil])
	}
}

func Tetap(hasil Nasabah) {
	sukuBunga := 0.01 // 1% per bulan
	bungaPerBulan := hasil.JumlahPinjaman * sukuBunga
	bungaTotal := bungaPerBulan * float64(hasil.Tenor)
	totalBayar := hasil.JumlahPinjaman + bungaTotal
	cicilan := totalBayar / float64(hasil.Tenor)

	fmt.Println("\n📈 Simulasi Cicilan dengan Bunga Tetap")
	fmt.Println("╔══════════════════════════════════════")
	fmt.Printf("║ Bunga/Bulan   : %.2f × 1%%         : %.2f\n", hasil.JumlahPinjaman, bungaPerBulan)
	fmt.Printf("║ Bunga Total   : %.2f   × %d Bulan     : %.2f\n", bungaPerBulan, hasil.Tenor, bungaTotal)
	fmt.Printf("║ Total Bayar   : %.2f + %.2f : %.2f\n", hasil.JumlahPinjaman, bungaTotal, totalBayar)
	fmt.Printf("║ Cicilan/Bulan : %.2f / %d Bulan   : %.2f\n", totalBayar, hasil.Tenor, cicilan)
	fmt.Println("╚══════════════════════════════════════")
}

func Variabel(hasil Nasabah) {
	sukuBunga := 0.01 // 1% per bulan
	pokok := hasil.JumlahPinjaman
	tenor := hasil.Tenor
	angsuranPokok := pokok / float64(tenor)

	fmt.Println("\n📊 Simulasi Cicilan dengan Bunga Variabel")
	fmt.Println("╔══════════════════════════════════════╗")

	for bulan := 1; bulan <= tenor; bulan++ {
		bungaBulanIni := pokok * sukuBunga
		cicilanBulanIni := angsuranPokok + bungaBulanIni

		fmt.Printf("║ 📆 Bulan ke-%d\n", bulan)
		fmt.Printf("║    Sisa Pokok     : %.2f\n", pokok)
		fmt.Printf("║    Bunga          : %.2f (%.2f%% dari %.2f)\n", bungaBulanIni, sukuBunga*100, pokok)
		fmt.Printf("║    Angsuran Pokok : %.2f\n", angsuranPokok)
		fmt.Printf("║    Cicilan Bulan  : %.2f\n", cicilanBulanIni)
		fmt.Println("╠──────────────────────────────────────╣")

		// Kurangi pokoknya setelah pembayaran bulan ini
		pokok -= angsuranPokok
	}

	fmt.Println("╚══════════════════════════════════════╝")
}
