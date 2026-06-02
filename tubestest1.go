package main

import "fmt"

const MAX_JADWAL int = 50
const MAX_USER int = 50

// =========================================================================
// 1. Struktur Data (Struct/Rekor)
// =========================================================================

type JadwalFutsal struct {
	IDJadwal   string
	NamaLap    string
	JamMulai   int
	JamSelesai int
	Status     string // "Kosong" atau "Dipesan"
	Penyewa    string // Nama penyewa/team jika status "Dipesan"
}

type AkunUser struct {
	Username string
	Password string
	NamaTeam string
}

// Tipe data alias untuk Array
type ArrJadwal [MAX_JADWAL]JadwalFutsal
type ArrUser [MAX_USER]AkunUser

// =========================================================================
// 2. Subprogram - Manajemen Akun (Modular & Terstruktur)
// =========================================================================

// Fungsi untuk mengecek apakah username sudah terdaftar (Sequential Search)
func cariUser(U ArrUser, nUser int, usernameCari string) int {
	var ketemuIdx int = -1
	var i int = 0
	for i < nUser && ketemuIdx == -1 {
		if U[i].Username == usernameCari {
			ketemuIdx = i
		}
		i++
	}
	return ketemuIdx
}

// Prosedur pendaftaran akun User baru
func registrasiUser(U *ArrUser, nUser *int) {
	if *nUser >= MAX_USER {
		fmt.Println("\n[Gagal] Kapasitas pendaftaran user penuh.")
		return
	}

	var uname, pass, team string
	fmt.Println("\n--- MENU REGISTRASI USER ---")
	fmt.Print("Masukkan Username Baru  : ")
	fmt.Scan(&uname)

	// Validasi username unik
	if cariUser(*U, *nUser, uname) != -1 || uname == "admin" {
		fmt.Println("[Gagal] Username sudah digunakan, cari nama lain!")
		return
	}

	fmt.Print("Masukkan Password       : ")
	fmt.Scan(&pass)
	fmt.Print("Masukkan Nama Team Futsal: ")
	fmt.Scan(&team)

	// Simpan ke array user
	U[*nUser] = AkunUser{Username: uname, Password: pass, NamaTeam: team}
	*nUser++
	fmt.Println("\n[Sukses] Akun berhasil dibuat! Silakan gunakan menu login.")
}

// Fungsi Autentikasi Login (Mengembalikan role: 1 untuk Admin, 2 untuk User, 0 jika Gagal)
func loginSistem(U ArrUser, nUser int, currentTeam *string) int {
	var uname, pass string
	fmt.Println("\n--- MENU LOGIN ---")
	fmt.Print("Username: ")
	fmt.Scan(&uname)
	fmt.Print("Password: ")
	fmt.Scan(&pass)

	// 1. Cek Kredensial Admin (Hardcoded sesuai permintaan)
	if uname == "admin" && pass == "admin123" {
		*currentTeam = "Administrator"
		return 1 // Role Admin
	}

	// 2. Cek Kredensial User di Array (Searching)
	idx := cariUser(U, nUser, uname)
	if idx != -1 && U[idx].Password == pass {
		*currentTeam = U[idx].NamaTeam // Simpan nama team user yang aktif login
		return 2                       // Role User
	}

	fmt.Println("\n[Gagal] Username atau Password salah!")
	return 0 // Gagal Login
}

// =========================================================================
// 3. Subprogram - Fitur Inti Aplikasi (Jadwal, Sorting, Searching)
// =========================================================================

func inisialisasiData(A *ArrJadwal, n *int) {
	*n = 6
	A[0] = JadwalFutsal{"J001", "Lapangan A", 14, 15, "Kosong", "-"}
	A[1] = JadwalFutsal{"J002", "Lapangan B", 19, 20, "Kosong", "-"}
	A[2] = JadwalFutsal{"J003", "Lapangan A", 9, 10, "Kosong", "-"}
	A[3] = JadwalFutsal{"J004", "Lapangan A", 21, 22, "Dipesan", "AAAteam"}
	A[4] = JadwalFutsal{"J005", "Lapangan B", 16, 17, "Kosong", "-"}
	A[5] = JadwalFutsal{"J006", "Lapangan B", 10, 11, "Dipesan", "Yebs FC"}
}

func tampilkanSemuaTabel(A ArrJadwal, n int) {
	fmt.Println("\n=========================================================================")
	fmt.Printf("| %-6s | %-12s | %-11s | %-9s | %-15s |\n", "ID", "Lapangan", "Jam Main", "Status", "Penyewa")
	fmt.Println("=========================================================================")
	for i := 0; i < n; i++ {
		jamFormat := fmt.Sprintf("%02d:00-%02d:00", A[i].JamMulai, A[i].JamSelesai)
		fmt.Printf("| %-6s | %-12s | %-11s | %-9s | %-15s |\n",
			A[i].IDJadwal, A[i].NamaLap, jamFormat, A[i].Status, A[i].Penyewa)
	}
	fmt.Println("=========================================================================")
}

// Sorting: Selection Sort (Mengurutkan Jadwal Kosong Berdasarkan Jam Mulai)
func urutkanJadwalKosong(A *ArrJadwal, n int) {
	var temp ArrJadwal
	var m int = 0

	for i := 0; i < n; i++ {
		if A[i].Status == "Kosong" {
			temp[m] = A[i]
			m++
		}
	}

	if m == 0 {
		fmt.Println("\n[!] Maaf, tidak ada jadwal kosong yang tersedia.")
		return
	}

	for i := 0; i < m-1; i++ {
		idxMin := i
		for j := i + 1; j < m; j++ {
			if temp[j].JamMulai < temp[idxMin].JamMulai {
				idxMin = j
			}
		}
		tukar := temp[i]
		temp[i] = temp[idxMin]
		temp[idxMin] = tukar
	}

	fmt.Println("\n>>> JADWAL KOSONG TERSEDIA (Diurutkan Jam Paling Awal) <<<")
	tampilkanSemuaTabel(temp, m)
}

// Searching: Sequential Search berdasarkan ID Jadwal
func cariJadwalByID(A ArrJadwal, n int, idCari string) int {
	var ketemuIdx int = -1
	var i int = 0
	for i < n && ketemuIdx == -1 {
		if A[i].IDJadwal == idCari {
			ketemuIdx = i
		}
		i++
	}
	return ketemuIdx
}

// Fitur User: Booking Lapangan secara otomatis memakai nama team pendaftar
func pesanLapangan(A *ArrJadwal, n int, namaTeamPenyewa string) {
	var idTarget string
	tampilkanSemuaTabel(*A, n)
	fmt.Print("Masukkan ID Jadwal yang ingin dibooking: ")
	fmt.Scan(&idTarget)

	idx := cariJadwalByID(*A, n, idTarget)

	if idx != -1 {
		if A[idx].Status == "Kosong" {
			A[idx].Status = "Dipesan"
			A[idx].Penyewa = namaTeamPenyewa
			fmt.Printf("\n[Sukses] Lapangan berhasil dibooking untuk team: %s!\n", namaTeamPenyewa)
		} else {
			fmt.Println("\n[Gagal] Lapangan sudah dibooking team lain.")
		}
	} else {
		fmt.Println("\n[Gagal] ID Jadwal tidak valid.")
	}
}

// Fitur Khusus Admin: Menambah slot jadwal baru
func tambahSlotJadwal(A *ArrJadwal, n *int) {
	if *n >= MAX_JADWAL {
		fmt.Println("\n[Gagal] Database jadwal penuh.")
		return
	}

	var id, lap string
	var mulai, selesai int

	fmt.Println("\n--- FORMAT TAMBAH SLOT JADWAL (ADMIN) ---")
	fmt.Print("Masukkan ID Jadwal Baru (ex: J007): ")
	fmt.Scan(&id)

	if cariJadwalByID(*A, *n, id) != -1 {
		fmt.Println("[Gagal] ID Jadwal tersebut sudah ada!")
		return
	}

	fmt.Print("Nama Lapangan (Lapangan A / Lapangan B): ")
	fmt.Scan(&lap)
	fmt.Print("Jam Mulai (0-23): ")
	fmt.Scan(&mulai)
	fmt.Print("Jam Selesai (0-23): ")
	fmt.Scan(&selesai)

	A[*n] = JadwalFutsal{
		IDJadwal:   id,
		NamaLap:    lap,
		JamMulai:   mulai,
		JamSelesai: selesai,
		Status:     "Kosong",
		Penyewa:    "-",
	}
	*n++
	fmt.Println("\n[Sukses] Slot jadwal baru berhasil ditambahkan oleh Admin!")
}

// =========================================================================
// 4. Program Utama & Alur Menu (Main Driver)
// =========================================================================
func main() {
	var masterJadwal ArrJadwal
	var nJadwal int
	var masterUser ArrUser
	var nUser int

	var pilihanUtama, pilihanSub int
	var roleAktif int
	var teamAktif string

	inisialisasiData(&masterJadwal, &nJadwal)

	for pilihanUtama != 4 {
		fmt.Println("\n=========================================")
		fmt.Println("       SISTEM FUTSAL SENTRAL (CMD)       ")
		fmt.Println("=========================================")
		fmt.Println("1. Registrasi Akun User Baru")
		fmt.Println("2. Login Aplikasi (Admin / User)")
		fmt.Println("3. Lihat Jadwal Terbuka (Tanpa Login)")
		fmt.Println("4. Keluar Program")
		fmt.Print("Pilih Opsi (1-4): ")
		fmt.Scan(&pilihanUtama)

		switch pilihanUtama {
		case 1:
			registrasiUser(&masterUser, &nUser)
		case 2:
			roleAktif = loginSistem(masterUser, nUser, &teamAktif)

			// JIKA LOGIN SEBAGAI ADMIN (Role = 1)
			if roleAktif == 1 {
				pilihanSub = 0
				for pilihanSub != 3 {
					fmt.Println("\n--- DASHBOARD UTAMA: ADMINISTRATOR ---")
					fmt.Println("1. Lihat Seluruh Log Jadwal")
					fmt.Println("2. Tambah Slot Jadwal Baru")
					fmt.Println("3. Logout Kembali")
					fmt.Print("Pilih Opsi: ")
					fmt.Scan(&pilihanSub)

					if pilihanSub == 1 {
						tampilkanSemuaTabel(masterJadwal, nJadwal)
					} else if pilihanSub == 2 {
						tambahSlotJadwal(&masterJadwal, &nJadwal)
					}
				}
				fmt.Println("\n[Info] Sesi Admin berakhir.")
			}

			// JIKA LOGIN SEBAGAI USER (Role = 2)
			if roleAktif == 2 {
				pilihanSub = 0
				for pilihanSub != 3 {
					fmt.Printf("\n--- DASHBOARD USER (Team: %s) ---\n", teamAktif)
					fmt.Println("1. Urutkan & Cari Jadwal Kosong (Sorting)")
					fmt.Println("2. Lakukan Booking Lapangan (Searching)")
					fmt.Println("3. Logout Kembali")
					fmt.Print("Pilih Opsi: ")
					fmt.Scan(&pilihanSub)

					if pilihanSub == 1 {
						urutkanJadwalKosong(&masterJadwal, nJadwal)
					} else if pilihanSub == 2 {
						pesanLapangan(&masterJadwal, nJadwal, teamAktif)
					}
				}
				fmt.Println("\n[Info] Sesi User berakhir.")
			}

		case 3:
			fmt.Println("\n>>> JADWAL SAAT INI (PENGUNJUNG UMUM) <<<")
			tampilkanSemuaTabel(masterJadwal, nJadwal)
		case 4:
			fmt.Println("\nAplikasi ditutup. Selamat berolahraga!")
		default:
			fmt.Println("\n[!] Opsi menu tidak valid.")
		}
	}
}
