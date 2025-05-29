package main

import "fmt"

// Inisialisasi slice global untuk menyimpan data nasabah
var dataNasabah []Nasabah

func main() {
	var DataDummy string
	var menu int

	for {
		fmt.Print("Apakah Anda Ingin Menggunakan Data Dummy? [y/n] : ")
		fmt.Scan(&DataDummy)
		if DataDummy == "y" {
			dataNasabah = []Nasabah{
				{"Andi Setiawan", 15000000, 12, 6},
				{"Budi Hartono", 20000000, 12, 12},
				{"Citra Ayu", 10000000, 6, 6},
				{"Dewi Lestari", 25000000, 12, 3},
				{"Eka Pratama", 30000000, 3, 1},
				{"Fajar Nugroho", 12000000, 12, 3},
				{"Gina Marissa", 18000000, 12, 6},
				{"Hadi Santoso", 22000000, 6, 4},
				{"Ika Putri", 9000000, 6, 3},
				{"Joko Susanto", 27000000, 12, 12},
			}
			break

		} else if DataDummy == "n" {
			break
		} else {
			fmt.Println("⚠️  Inputan Tidak Valid!")
			continue
		}
	}
	
	for {
		fmt.Println("\n╔══════════════════════════════════════╗")
		fmt.Println("║              📋  MENU                ║")
		fmt.Println("╠══════════════════════════════════════╣")
		fmt.Println("║  1. ➕ Tambah Peminjam               ║")
		fmt.Println("║  2. ✏️ Ubah Data Peminjam             ║")
		fmt.Println("║  3. ❌ Hapus Data Peminjam           ║")
		fmt.Println("║  4. 💰 Simulasi Pinjaman             ║")
		fmt.Println("║  5. 💵 Input Status Pembayaran       ║")
		fmt.Println("║  6. 🔍 Cari Peminjam                 ║")
		fmt.Println("║  7. 📊 Urutkan Data Peminjam         ║")
		fmt.Println("║  8. 📑 Tampilkan Laporan             ║")
		fmt.Println("║  9. 🚪 Keluar                        ║")
		fmt.Println("╚══════════════════════════════════════╝")
		fmt.Print("👉 Pilihan Anda: ")
		fmt.Scan(&menu)

		if menu < 1 || menu > 9 {
			fmt.Println("⚠️  Menu Tidak Valid!")
		}

		if menu == 0 {
			break
		}

		if menu == 1 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║        ➕ Menu Tambah Nasabah        ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			nasabahBaru := Tambah(dataNasabah)
			dataNasabah = append(dataNasabah, nasabahBaru)

			fmt.Println("✅ Menu Tambah Data Nasabah Selesai")
		}

		if menu == 2 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║      📝  Menu Ubah Data Nasabah      ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			UbahData(dataNasabah)

			fmt.Println("✅ Menu Ubah Data Nasabah Selesai")
		}

		if menu == 3 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║        ❌ Menu Hapus Nasabah         ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			dataNasabah = Destroy(dataNasabah)

			fmt.Println("✅ Menu Hapus Data Nasabah Selesai")
		}

		if menu == 4 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║      💰 Menu Simulasi Pinjaman       ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			SimulasiPinjaman(dataNasabah)

			fmt.Println("✅ Menu Simulasi Pinjaman Selesai")
		}

		if menu == 5 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║      💵 Menu Status Pembayaran       ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			StatusPembayaran(dataNasabah)

			fmt.Println("✅ Menu Simulasi Pinjaman Selesai")
		}

		if menu == 6 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║         🔍 Cari Pinjaman             ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			CariPeminjam(dataNasabah)

			fmt.Println("✅ Menu Cari Pinjaman Selesai")
		}

		if menu == 7 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║      📊 Urutkan Data Pinjaman        ║")
			fmt.Println("╚══════════════════════════════════════╝")

			fmt.Scanln()
			UrutkanPinjaman(dataNasabah)

			fmt.Println("✅ Menu Urutkan Data Pinjaman Selesai")
		}

		if menu == 8 {
			fmt.Println("╔══════════════════════════════════════╗")
			fmt.Println("║        📑 Tampilkan Laporan          ║")
			fmt.Println("╚══════════════════════════════════════╝")

			TampilkanLaporan(dataNasabah)

			fmt.Println("✅ Menu Tampilkan Laporan Selesai")
		}

		if menu == 9 {
			break
		}
	}

	fmt.Println("🙏 Terimakasih Sudah Menggunakan Program Kami")
}
