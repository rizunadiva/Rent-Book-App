package main

import (
	"fmt"
	"group-project-kel5/config"
	_user "group-project-kel5/controller/user"
	_entity "group-project-kel5/entity"
)

func main() {
	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesUser := _user.AksesUsers{DB: conn}
	// aksesBuku := _entity.AksesBuku{DB: conn}
	var input int = 0
	for input != 99 {
		fmt.Println("=============================================================")
		fmt.Println("|SELAMAT DATANG DI PERPUSTAKAAN UNIVERSITAS LANGSUNG SARJANA|")
		fmt.Println("=============================================================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Lihat Daftar Buku")
		fmt.Println("0. Kembali")
		fmt.Println("Masukkan Pilihan menu: ")
		fmt.Scanln(&input)
		switch input {
		case 1:
			// AksesLogin := _entity.AksesLogin{DB: conn}
			// username, result, err := _entity.Login(DBConn)
			// if err != nil {
			// 	fmt.Println("Login Gagal, username tidak terdaftar")
			// }
			// if result {
			// 	fmt.Println("\n=Login /berhasil=")
			// 	fmt.Println("\n=================")
			// 	fmt.Println("Selamat Datang\nPilih menu dibawah ini")
			// 	fmt.Println("1. Sewa Buku\n2. Tambah Buku")
			// }

			// fmt.Print("Masukkan username: ")
			// fmt.Scanln(&)
			// fmt.Print("Masukkan password: ")
			// fmt.Scanln(&)
		case 2:
			var newUsers _entity.Users
			result := aksesUser.TambahUser(newUsers)
			if result.ID == 0 {
				fmt.Println("Tidak bisa input user")
				break
			}
			fmt.Println("Berhasil input user")
		case 3:
			AksesBuku := _entity.AksesBuku{DB: conn}
			for _, val := range AksesBuku.GetAllData() {
				fmt.Println(val)
			}
		default:
			continue
		}
	}
	fmt.Println("Terima Kasih")
}
