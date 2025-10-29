package main

import (
	"fmt"
	"strings"
)

const NMAX int = 100

// Struct Tiket
type Tiket struct {
	IdPembeli  int
	Nama       string
	JenisTiket string
	Harga      float64
}

// Struct User
type User struct {
	Username string
	Password string
	Role     string // "admin" atau "user"
}

type tabTiket [NMAX]Tiket
type tabUser [NMAX]User

// Data global
var dataTiket tabTiket
var nTiket int = 0

var dataUser = tabUser{
	{"admin", "admin123", "admin"},
	{"user", "user123", "user"},
}
var nUser int = 2

func main() {
	var pilihan int
	for pilihan != 3 {
		fmt.Println("\n====== MENU UTAMA ======")
		fmt.Println("1. Login")
		fmt.Println("2. Daftar Akun Baru")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var username, password string
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)
			user, loginSuccess := login(username, password)
			if !loginSuccess {
				fmt.Println("Login gagal.")
				continue
			}
			fmt.Printf("Login berhasil sebagai %s\n", user.Role)
			menuUtama(user)
		case 2:
			daftarAkunBaru()
		case 3:
			fmt.Println("Terima kasih telah menggunakan program.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func menuUtama(user User) {
	var pilihan int
	for pilihan != 8 {
		fmt.Println("\n======== MENU ========")
		fmt.Println("1. Tambah Tiket")
		fmt.Println("2. Tampilkan Semua Tiket")
		fmt.Println("3. Update Tiket")
		fmt.Println("4. Hapus Tiket")
		fmt.Println("5. Cari Tiket")
		fmt.Println("6. Urutkan Tiket (Harga)")
		fmt.Println("7. Total Pendapatan")
		fmt.Println("8. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if user.Role != "admin" && (pilihan == 1 || pilihan == 3 || pilihan == 4 || pilihan == 6 || pilihan == 7) {
			fmt.Println("Akses ditolak. Hanya admin yang dapat melakukan aksi ini.")
			continue
		}

		switch pilihan {
		case 1:
			tambahTiket(&dataTiket, &nTiket)
		case 2:
			tampilkanTiket(dataTiket, nTiket)
		case 3:
			updateTiket(&dataTiket, nTiket)
		case 4:
			hapusTiket(&dataTiket, &nTiket)
		case 5:
			cariTiket(dataTiket, nTiket)
		case 6:
			urutTiket(&dataTiket, nTiket)
		case 7:
			fmt.Printf("Total Pendapatan: Rp%.0f\n", totalPendapatan(dataTiket, nTiket))
		case 8:
			fmt.Println("Logout dari akun.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func login(username, password string) (User, bool) {
	for i := 0; i < nUser; i++ {
		u := dataUser[i]
		if u.Username == username && u.Password == password {
			return u, true
		}
	}
	return User{}, false
}

func daftarAkunBaru() {
	if nUser >= NMAX {
		fmt.Println("Jumlah akun penuh.")
		return
	}
	var username, password, role string
	fmt.Print("Masukkan username baru: ")
	fmt.Scan(&username)
	for i := 0; i < nUser; i++ {
		if dataUser[i].Username == username {
			fmt.Println("Username sudah digunakan.")
			return
		}
	}
	fmt.Print("Masukkan password (minimal 6 karakter): ")
	fmt.Scan(&password)
	if len(password) < 6 {
		fmt.Println("Password terlalu pendek.")
		return
	}
	fmt.Print("Masukkan role (admin/user): ")
	fmt.Scan(&role)
	role = strings.ToLower(role)
	if role != "admin" && role != "user" {
		fmt.Println("Role tidak valid. Hanya 'admin' atau 'user' yang diperbolehkan.")
		return
	}
	dataUser[nUser] = User{username, password, role}
	nUser++
	fmt.Println("Akun berhasil didaftarkan! Silakan login.")
}

func tambahTiket(data *tabTiket, n *int) {
	var id int
	var nama, jenis string
	var harga float64
	fmt.Print("Masukkan ID Pembeli: ")
	fmt.Scan(&id)
	fmt.Print("Masukkan Nama Pembeli: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Jenis Tiket: ")
	fmt.Scan(&jenis)
	fmt.Print("Masukkan Harga Tiket: ")
	fmt.Scan(&harga)
	data[*n] = Tiket{id, nama, jenis, harga}
	*n++
	fmt.Println("Tiket berhasil ditambahkan!")
}

func tampilkanTiket(data tabTiket, n int) {
	for i := 0; i < n; i++ {
		t := data[i]
		fmt.Printf("ID: %d | Nama: %s | Jenis: %s | Harga: Rp%.0f\n", t.IdPembeli, t.Nama, t.JenisTiket, t.Harga)
	}
}

func updateTiket(data *tabTiket, n int) {
	var id int
	fmt.Print("Masukkan ID yang ingin diupdate: ")
	fmt.Scan(&id)
	for i := 0; i < n; i++ {
		if data[i].IdPembeli == id {
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&data[i].Nama)
			fmt.Print("Masukkan Jenis Tiket Baru: ")
			fmt.Scan(&data[i].JenisTiket)
			fmt.Print("Masukkan Harga Tiket Baru: ")
			fmt.Scan(&data[i].Harga)
			fmt.Println("Tiket berhasil diupdate!")
			return
		}
	}
	fmt.Println("Tiket tidak ditemukan.")
}

func hapusTiket(data *tabTiket, n *int) {
	var id int
	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scan(&id)
	for i := 0; i < *n; i++ {
		if data[i].IdPembeli == id {
			for j := i; j < *n-1; j++ {
				data[j] = data[j+1]
			}
			*n--
			fmt.Println("Tiket berhasil dihapus!")
			return
		}
	}
	fmt.Println("Tiket tidak ditemukan.")
}

func cariTiket(data tabTiket, n int) {
	var id int
	fmt.Print("Masukkan ID yang dicari: ")
	fmt.Scan(&id)
	for i := 0; i < n; i++ {
		if data[i].IdPembeli == id {
			t := data[i]
			fmt.Printf("Tiket ditemukan: ID: %d | Nama: %s | Jenis: %s | Harga: Rp%.0f\n", t.IdPembeli, t.Nama, t.JenisTiket, t.Harga)
			return
		}
	}
	fmt.Println("Tiket tidak ditemukan.")
}

func urutTiket(data *tabTiket, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if data[j].Harga < data[minIdx].Harga {
				minIdx = j
			}
		}
		data[i], data[minIdx] = data[minIdx], data[i]
	}
	fmt.Println("Tiket berhasil diurutkan berdasarkan harga.")
}

func totalPendapatan(data tabTiket, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += data[i].Harga
	}
	return total
}
