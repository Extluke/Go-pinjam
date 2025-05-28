package main

import "fmt"

// Inisialisasi slice global untuk menyimpan data nasabah
var dataNasabah []Nasabah

func main() {

	var menu int
  for {
        fmt.Println("\n+-------------------------------+")
        fmt.Println("|             MENU              |")
        fmt.Println("+-------------------------------+")
        fmt.Println("| 1. Tambah Peminjam            |")
        fmt.Println("| 2. Ubah Data Peminjam         |")
        fmt.Println("| 3. Hapus Data Peminjam        |")
        fmt.Println("| 4. Simulasi Pinjaman          |")
        fmt.Println("| 5. Input Status Pembayaran    |")
        fmt.Println("| 6. Cari Peminjam              |")
        fmt.Println("| 7. Urutkan Data Peminjam      |")
        fmt.Println("| 8. Tampilkan Laporan          |")
        fmt.Println("| 9. Keluar                     |")
        fmt.Println("+-------------------------------+")
        fmt.Print("Pilihan Anda: ")
        fmt.Scan(&menu)

        if menu < 1 || menu > 9 {
					fmt.Println("Menu Tidak Valid")
        }

				if menu == 0 {
					break
				}

				if menu == 1 {
          fmt.Println("+-------------------------------+")
	        fmt.Println("|      Menu Tambah Nasabah      |")
	        fmt.Println("+-------------------------------+")

          fmt.Scanln()
					nasabahBaru := Tambah(dataNasabah)
					dataNasabah = append(dataNasabah, nasabahBaru)

          fmt.Println("\nMenu Tambah Data Nasabah Selesai")
				}

				if menu == 2 {
          fmt.Println("+-------------------------------+")
	        fmt.Println("|    Menu Ubah Data Nasabah     |")
	        fmt.Println("+-------------------------------+")

					fmt.Scanln()
          UbahData(dataNasabah)

          fmt.Println("\nMenu Ubah Data Nasabah Selesai")
				}

        if menu == 3 {
          fmt.Println("+-------------------------------+")
	        fmt.Println("|    Menu Hapus Data Nasabah    |")
	        fmt.Println("+-------------------------------+")

					fmt.Scanln()
          dataNasabah = Destroy(dataNasabah)

          fmt.Println("\nMenu Hapus Data Nasabah Selesai")
				}

        if menu == 4 {
          fmt.Println("+-------------------------------+")
	        fmt.Println("|    Menu Simulasi Pinjaman     |")
	        fmt.Println("+-------------------------------+")

					fmt.Scanln()
          SimulasiPinjaman(dataNasabah)

          fmt.Println("\nMenu Simulasi Pinjaman Selesai")
				}
        
        if menu == 6 {
          fmt.Println("+-------------------------------+")
	        fmt.Println("|         Cari Pinjaman         |")
	        fmt.Println("+-------------------------------+")

					fmt.Scanln()
          CariPeminjam(dataNasabah)

          fmt.Println("\nMenu Cari Pinjaman Selesai")
				}

        if menu == 8 {
          fmt.Println("+-------------------------------+")
          fmt.Println("|       Tampilkan Laporan       |")
          fmt.Println("+-------------------------------+")
          TampilkanLaporan(dataNasabah)
          fmt.Println("\nMenu Tampilkan Laporan Selesai")
        }

        if menu == 7 {
          fmt.Println("+-------------------------------+")
          fmt.Println("|     Urutkan Data Pinjaman     |")
          fmt.Println("+-------------------------------+")
          fmt.Scanln()
          UrutkanPinjaman(dataNasabah)
          fmt.Println("\nMenu Urutkan Data Pinjaman Selesai")
        }

        if menu == 9 {
          break
        }
  }

	fmt.Println("Terimakasih Sudah Menggunakan Program kami")

}
