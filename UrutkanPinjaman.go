package main

import "fmt"

func selectionSortASC(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if data[j].JumlahPinjaman < data[idxMin].JumlahPinjaman {
				idxMin = j
			}
		}
		data[i], data[idxMin] = data[idxMin], data[i]
	}
}

func selectionSortDESC(data []Nasabah) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		maxIndex := i
		for j := i + 1; j < n; j++ {
			if data[j].JumlahPinjaman > data[maxIndex].JumlahPinjaman {
				maxIndex = j
			}
		}
		data[i], data[maxIndex] = data[maxIndex], data[i]
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
	fmt.Println("\n+-------------------------------+")
	fmt.Println("|       Pilih Metode Urut       |")
	fmt.Println("+-------------------------------+")
	fmt.Println("| 1. Selection Sort             |")
	fmt.Println("| 2. Insertion Sort             |")
	fmt.Println("+-------------------------------+")

	var pilihan int

	for {
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		if pilihan >= 1 && pilihan <= 2 {
			break
		}
		fmt.Println("\nInput tidak valid, silakan pilih ulang.")
	}

	if len(dataNasabah) == 0 {
		fmt.Println("\n+-------------------------------+")
		fmt.Println("|    Tidak ada data nasabah     |")
		fmt.Println("+-------------------------------+")
		return
	} else {
		if pilihan == 1 {
			fmt.Println("\n+-------------------------------+")
			fmt.Println("|  Pilih Metode Selection Sort  |")
			fmt.Println("+-------------------------------+")
			fmt.Println("| 1. Ascending                  |")
			fmt.Println("| 2. Descending                 |")
			fmt.Println("+-------------------------------+")

			var urutan int
			fmt.Print("Pilihan: ")
			fmt.Scanln(&urutan)

			if urutan == 1 {
				selectionSortASC(dataNasabah)

				fmt.Println("\n+-------------------------------+")
				fmt.Println("|     Data Setelah Diurutkan    |")
				fmt.Println("+-------------------------------+")

				for _, nasabah := range dataNasabah {
					fmt.Printf("Nama: %s, Jumlah Pinjaman: %.2f, Tenor: %d\n", nasabah.Nama, nasabah.JumlahPinjaman, nasabah.Tenor)
				}

				fmt.Println("\nData nasabah telah diurutkan menggunakan Selection Sort secara Ascending berdasarkan jumlah pinjaman.")
			} else if urutan == 2 {
				selectionSortDESC(dataNasabah)

				fmt.Println("\n+-------------------------------+")
				fmt.Println("|     Data Setelah Diurutkan    |")
				fmt.Println("+-------------------------------+")

				for _, nasabah := range dataNasabah {
					fmt.Printf("Nama: %s, Jumlah Pinjaman: %.2f, Tenor: %d\n", nasabah.Nama, nasabah.JumlahPinjaman, nasabah.Tenor)
				}

				fmt.Println("\nData nasabah telah diurutkan menggunakan Selection Sort secara Descending berdasarkan jumlah pinjaman.")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else if pilihan == 2 {
			fmt.Println("\n+-------------------------------+")
			fmt.Println("|  Pilih Metode Insertion Sort  |")
			fmt.Println("+-------------------------------+")
			fmt.Println("| 1. Ascending                  |")
			fmt.Println("| 2. Descending                 |")
			fmt.Println("+-------------------------------+")

			var urutan int
			fmt.Print("Pilihan: ")
			fmt.Scanln(&urutan)

			if urutan == 1 {
				insertionSortASC(dataNasabah)

				fmt.Println("\n+-------------------------------+")
				fmt.Println("|     Data Setelah Diurutkan    |")
				fmt.Println("+-------------------------------+")

				for _, nasabah := range dataNasabah {
					fmt.Printf("Nama: %s, Jumlah Pinjaman: %.2f, Tenor: %d\n", nasabah.Nama, nasabah.JumlahPinjaman, nasabah.Tenor)
				}

				fmt.Println("\nData nasabah telah diurutkan menggunakan Insertion Sort secara Ascending berdasarkan jumlah pinjaman.")
			} else if urutan == 2 {
				insertionSortDESC(dataNasabah)

				fmt.Println("\n+-------------------------------+")
				fmt.Println("|     Data Setelah Diurutkan    |")
				fmt.Println("+-------------------------------+")

				for _, nasabah := range dataNasabah {
					fmt.Printf("Nama: %s, Jumlah Pinjaman: %.2f, Tenor: %d\n", nasabah.Nama, nasabah.JumlahPinjaman, nasabah.Tenor)
				}

				fmt.Println("\nData nasabah telah diurutkan menggunakan Insertion Sort secara Descending berdasarkan jumlah pinjaman.")
			}
		}
	}
}