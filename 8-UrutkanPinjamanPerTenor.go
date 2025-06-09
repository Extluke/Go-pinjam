package main

import "fmt"

func TenorSelectionSortASC(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		indexTerkecil := i
		for j := i + 1; j < n; j++ {
			if float64(data[j].Tenor) < float64(data[indexTerkecil].Tenor) {
				indexTerkecil = j
			}
		}
		data[i], data[indexTerkecil] = data[indexTerkecil], data[i]
	}
}

func TenorSelectionSortDESC(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		indexTerbesar := i
		for j := i + 1; j < n; j++ {
			if float64(data[j].Tenor) > float64(data[indexTerbesar].Tenor) {
				indexTerbesar = j
			}
		}
		data[i], data[indexTerbesar] = data[indexTerbesar], data[i]
	}
}

func TenorInsertionSortASC(data []Nasabah) {
	n := len(data)
	for i := 1; i < n; i++ {
		j := i
		temp := data[j]
		for j > 0 && float64(temp.Tenor) < float64(data[j-1].Tenor) {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func TenorInsertionSortDESC(data []Nasabah) {
	n := len(data)
	for i := 1; i < n; i++ {
		j := i
		temp := data[j]
		for j > 0 && float64(temp.Tenor) > float64(data[j-1].Tenor) {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func TenorUrutkanPinjaman(dataNasabah []Nasabah) {
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║     📊 Pilih Metode Pengurutan       ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║ 1. 🔢 Selection Sort                 ║")
	fmt.Println("║ 2. ✏️ Insertion Sort                 ║")
	fmt.Println("╚══════════════════════════════════════╝")

	var pilihan int

	for {
		fmt.Print("👉 Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan >= 1 && pilihan <= 2 {
			break
		}
		fmt.Println("⚠️  Input tidak valid, silakan pilih ulang.")
	}

	if len(dataNasabah) == 0 {
		fmt.Println("\n🚫 Tidak ada data nasabah.")
		return
	} else {
		if pilihan == 1 {
			fmt.Println("\n📥 Metode: Selection Sort")
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║ 1. 🔼 Ascending                      ║")
			fmt.Println("║ 2. 🔽 Descending                     ║")
			fmt.Println("╚══════════════════════════════════════╝")

			var urutan int
			fmt.Print("👉 Pilihan: ")
			fmt.Scanln(&urutan)

			if urutan == 1 {
				TenorSelectionSortASC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Ascending):")
				fmt.Println("╔══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠──────────────────────────────────────╣")
				}
				fmt.Println("╚══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Ascending dengan Selection Sort.")
			} else if urutan == 2 {
				TenorSelectionSortDESC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Descending):")
				fmt.Println("╔══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠──────────────────────────────────────╣")
				}
				fmt.Println("╚══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Descending dengan Selection Sort.")
			} else {
				fmt.Println("⚠️  Pilihan tidak valid.")
			}
		} else if pilihan == 2 {
			fmt.Println("\n📥 Metode: Insertion Sort")
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║ 1. 🔼 Ascending                      ║")
			fmt.Println("║ 2. 🔽 Descending                     ║")
			fmt.Println("╚══════════════════════════════════════╝")

			var urutan int
			fmt.Print("👉 Pilihan: ")
			fmt.Scanln(&urutan)

			if urutan == 1 {
				TenorInsertionSortASC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Ascending):")
				fmt.Println("╔══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠──────────────────────────────────────╣")
				}
				fmt.Println("╚══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Ascending dengan Insertion Sort.")
			} else if urutan == 2 {
				TenorInsertionSortDESC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Descending):")
				fmt.Println("╔══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠──────────────────────────────────────╣")
				}
				fmt.Println("╚══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Descending dengan Insertion Sort.")
			}
		}
	}
}
