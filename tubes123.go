package main

import "fmt"

const NMAX int = 50

type Jadwal struct {
	id         string
	lapangan   string
	jamMulai   int
	jamSelesai int
	status     string
	penyewa    string
}

type tJadwal [NMAX]Jadwal

func initData(J *tJadwal, n *int) {
	*n = 6
	J[0] = Jadwal{"01", "Lapangan A", 14, 15, "Kosong", "-"}
	J[1] = Jadwal{"02", "Lapangan B", 19, 20, "Kosong", "-"}
	J[2] = Jadwal{"03", "Lapangan A", 9, 10, "Kosong", "-"}
	J[3] = Jadwal{"04", "Lapangan A", 21, 22, "Dipesan", "kesmoy FC"}
	J[4] = Jadwal{"05", "Lapangan B", 16, 17, "Kosong", "-"}
	J[5] = Jadwal{"06", "Lapangan B", 10, 11, "Dipesan", "yessir"}
}

func cetakTabel(J tJadwal, n int) {
	fmt.Println("\n-------------------------------------------------------------------------")
	fmt.Printf("| %-6s | %-12s | %-11s | %-9s | %-15s |\n", "ID", "Lapangan", "Jam Main", "Status", "Penyewa")
	fmt.Println("-------------------------------------------------------------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-6s | %-12s | %02d:00-%02d:00 | %-9s | %-15s |\n",
			J[i].id, J[i].lapangan, J[i].jamMulai, J[i].jamSelesai, J[i].status, J[i].penyewa)
	}
	fmt.Println("-------------------------------------------------------------------------")
}

func urutJadwal(J *tJadwal, n int) {
	var temp tJadwal
	var m int = 0

	for i := 0; i < n; i++ {
		if J[i].status == "Kosong" {
			temp[m] = J[i]
			m++
		}
	}

	if m == 0 {
		fmt.Println("\nTidak ada jadwal kosong saat ini.")
		return
	}

	for i := 0; i < m-1; i++ {
		minIdx := i
		for j := i + 1; j < m; j++ {
			if temp[j].jamMulai < temp[minIdx].jamMulai {
				minIdx = j
			}
		}
		tukar := temp[i]
		temp[i] = temp[minIdx]
		temp[minIdx] = tukar
	}

	fmt.Println("\n>>> DAFTAR JADWAL KOSONG <<<")
	cetakTabel(temp, m)
}

func cariJadwalIdx(J tJadwal, n int, idCari string) int {
	for i := 0; i < n; i++ {
		if J[i].id == idCari {
			return i
		}
	}
	return -1
}

func bookingLapangan(J *tJadwal, n int) {
	var idTarget, namaTim string
	cetakTabel(*J, n)

	fmt.Print("Pilih ID Jadwal yang mau dipesan: ")
	fmt.Scan(&idTarget)

	idx := cariJadwalIdx(*J, n, idTarget)

	if idx != -1 {
		if J[idx].status == "Kosong" {
			fmt.Print("Masukkan Nama Tim Anda: ")
			fmt.Scan(&namaTim)

			J[idx].status = "Dipesan"
			J[idx].penyewa = namaTim
			fmt.Printf("Booking sukses untuk tim %s!\n", namaTim)
		} else {
			fmt.Println("Gagal! Lapangan sudah penuh.")
		}
	} else {
		fmt.Println("ID Jadwal tidak ditemukan.")
	}
}

func main() {
	var dataJadwal tJadwal
	var nJadwal int
	var menuUtama int

	initData(&dataJadwal, &nJadwal)

	for menuUtama != 4 {
		fmt.Println("\n=====================================")
		fmt.Println("       LAPANGAN FUTSAL       ")
		fmt.Println("=====================================")
		fmt.Println("1. Lihat Semua Jadwal")
		fmt.Println("2. Jadwal kosong")
		fmt.Println("3. Booking Lapangan")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menuUtama)

		switch menuUtama {
		case 1:
			cetakTabel(dataJadwal, nJadwal)
		case 2:
			urutJadwal(&dataJadwal, nJadwal)
		case 3:
			bookingLapangan(&dataJadwal, nJadwal)
		case 4:
			fmt.Println("Program selesai.")
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}
