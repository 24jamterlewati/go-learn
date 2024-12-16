package main

import (
	"fmt"
)

func main() {

	var Member = []string{}

	for {
		fmt.Println("1.Tampil Semua")
		fmt.Println("2.Tambah")
		fmt.Println("3.Hapus")
		fmt.Println("4.Keluar")

		var choice int
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Data Tersimpan\n", Member)
			for i, value := range Member {
				fmt.Printf("%d %s/n", i+1, value)
			}
		case 2:
			fmt.Println("Masukkan Data Baru: ")
			var newData string
			fmt.Scan(&newData)
			Member = append(Member, newData)
			fmt.Println("Member Baru Berhasil di Tambahkan")
		case 3:
			fmt.Println("Hapus Data")
			Member = []string{}
		case 4:
			fmt.Println("Terima Kasih")
			return
		default:
			fmt.Println("Masukkan Nomor yg Benar")
		}
	}
}
