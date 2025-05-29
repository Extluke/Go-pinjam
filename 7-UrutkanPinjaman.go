package main

import "fmt"

func selectionSortASC(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		indexTerkecil := i
		for j := i + 1; j < n; j++ {
			if data[j].JumlahPinjaman < data[indexTerkecil].JumlahPinjaman {
				indexTerkecil = j
			}
		}
		data[i], data[indexTerkecil] = data[indexTerkecil], data[i]
	}
}

func selectionSortDESC(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		indexTerbesar := i
		for j := i + 1; j < n; j++ {
			if data[j].JumlahPinjaman > data[indexTerbesar].JumlahPinjaman {
				indexTerbesar = j
			}
		}
		data[i], data[indexTerbesar] = data[indexTerbesar], data[i]
	}
}

func insertionSortASC(data []Nasabah) {
	n := len(data)
	for i := 1; i < n; i++ {
		j := i
		temp := data[j]
		for j > 0 && temp.JumlahPinjaman < data[j-1].JumlahPinjaman {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func insertionSortDESC(data []Nasabah) {
	n := len(data)
	for i := 1; i < n; i++ {
		j := i
		temp := data[j]
		for j > 0 && temp.JumlahPinjaman > data[j-1].JumlahPinjaman {
			data[j] = data[j-1]
			j--
		}
		data[j] = temp
	}
}

func UrutkanPinjaman(dataNasabah []Nasabah) {
	fmt.Println("\n╔═══════════════════════════════════════╗")
	fmt.Println("║      📊 Pilih Metode Pengurutan       ║")
	fmt.Println("╠═══════════════════════════════════════╣")
	fmt.Println("║ 1. 🔢 Selection Sort                  ║")
	fmt.Println("║ 2. ✏️ Insertion Sort                   ║")
	fmt.Println("╚═══════════════════════════════════════╝")

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
			fmt.Println("╔═══════════════════════════════════════╗")
			fmt.Println("║ 1. 🔼 Ascending                       ║")
			fmt.Println("║ 2. 🔽 Descending                      ║")
			fmt.Println("╚═══════════════════════════════════════╝")

			var urutan int
			fmt.Print("👉 Pilihan: ")
			fmt.Scanln(&urutan)

			if urutan == 1 {
				selectionSortASC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Ascending):")
				fmt.Println("╔═══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠───────────────────────────────────────╣")
				}
				fmt.Println("╚═══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Ascending dengan Selection Sort.")
			} else if urutan == 2 {
				selectionSortDESC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Descending):")
				fmt.Println("╔═══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠───────────────────────────────────────╣")
				}
				fmt.Println("╚═══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Descending dengan Selection Sort.")
			} else {
				fmt.Println("⚠️  Pilihan tidak valid.")
			}
		} else if pilihan == 2 {
			fmt.Println("\n📥 Metode: Insertion Sort")
			fmt.Println("╔═══════════════════════════════════════╗")
			fmt.Println("║ 1. 🔼 Ascending                       ║")
			fmt.Println("║ 2. 🔽 Descending                      ║")
			fmt.Println("╚═══════════════════════════════════════╝")

			var urutan int
			fmt.Print("👉 Pilihan: ")
			fmt.Scanln(&urutan)

			if urutan == 1 {
				insertionSortASC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Ascending):")
				fmt.Println("╔═══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠───────────────────────────────────────╣")
				}
				fmt.Println("╚═══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Ascending dengan Insertion Sort.")
			} else if urutan == 2 {
				insertionSortDESC(dataNasabah)

				fmt.Println("\n📋 Hasil Pengurutan (Descending):")
				fmt.Println("╔═══════════════════════════════════════╗")
				for _, nasabah := range dataNasabah {
					fmt.Printf("║ 👤 Nama            : %s\n", nasabah.Nama)
					fmt.Printf("║ 💸 Jumlah Pinjaman : Rp.%.2f\n", nasabah.JumlahPinjaman)
					fmt.Printf("║ 🕒 Tenor           : %d Bulan\n", nasabah.Tenor)
					fmt.Println("╠───────────────────────────────────────╣")
				}
				fmt.Println("╚═══════════════════════════════════════╝")
				fmt.Println("\n✅ Data berhasil diurutkan secara Descending dengan Insertion Sort.")
			}
		}
	}
}
