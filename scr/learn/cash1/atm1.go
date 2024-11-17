package main

import (
	"fmt"
)

func main() {

	// ini adalah saldo awal
	saldo := 100000.0

	for {
		fmt.Println("=== ATM MENU ===")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Setor Uang")
		fmt.Println("3. Tarik Uang")
		fmt.Println("4. Keluar ")

		var choice int
		fmt.Scan(&choice)

		// pengkondisian
		if choice == 1 {
			fmt.Printf("saldo anda adalah: Rp %.2f\n", saldo)
		} else if choice == 2 {
			var jumlah float64
			fmt.Print("masukkan jumlah setoran: Rp ")
			fmt.Scan(&jumlah)
			if jumlah > 0 {
				saldo += jumlah
				fmt.Printf("Rp %.2f saldo berhasil di tambahkan, saldo anda sekarang berjumlah: Rp %.2f\n", jumlah, saldo)
			}
		} else if choice == 3 {
			var jumlah float64
			fmt.Print("masukkan jumlah penarikan: Rp ")
			fmt.Scan(&jumlah)
			if jumlah > 0 && jumlah <= saldo {
				saldo -= jumlah
				fmt.Printf("Rp %.2f berhasil di tarik, jumlah saldo anda sekarang adalah %.2f\n", jumlah, saldo)
			} else {
				fmt.Println("Jumlah nominal yang anda masukkan salah")
			}
		} else if choice == 4 {
			fmt.Println("Terima Kasih Sudah Mengunakan Aplikasi Ini")
			break
		} else {
			fmt.Println("input tidak benar, silahkan pilih angka 1 sampai 4")
		}

	}

}
