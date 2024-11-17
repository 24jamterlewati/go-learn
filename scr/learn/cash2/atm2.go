package cash2

import (
	"fmt"
)

// Saldo awal
var saldo float64 = 100000.0

// Fungsi untuk cek saldo
func checkBalance() {
	fmt.Printf("Saldo Anda saat ini adalah: Rp%.2f\n", saldo)
}

// Fungsi untuk tarik tunai
func withdraw(amount float64) {
	if amount <= saldo {
		saldo -= amount
		fmt.Printf("Berhasil menarik: Rp%.2f\n", amount)
		checkBalance()
	} else {
		fmt.Println("Saldo tidak mencukupi untuk melakukan penarikan.")
	}
}

// Fungsi untuk setor tunai
func deposit(amount float64) {
	saldo += amount
	fmt.Printf("Berhasil menyetor: Rp%.2f\n", amount)
	checkBalance()
}

// Fungsi untuk menampilkan menu ATM
func showMenu() {
	fmt.Println("\n--- Menu ATM ---")
	fmt.Println("1. Cek Saldo")
	fmt.Println("2. Tarik Tunai")
	fmt.Println("3. Setor Tunai")
	fmt.Println("4. Keluar")
}

func main() {
	var choice int
	var amount float64

	for {
		showMenu()
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			checkBalance()
		case 2:
			fmt.Print("Masukkan jumlah yang ingin ditarik: Rp")
			fmt.Scanln(&amount)
			withdraw(amount)
		case 3:
			fmt.Print("Masukkan jumlah yang ingin disetor: Rp")
			fmt.Scanln(&amount)
			deposit(amount)
		case 4:
			fmt.Println("Terima kasih telah menggunakan ATM.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
