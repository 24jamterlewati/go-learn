// Buatlah program yang menggunakan for loop untuk menghitung jumlah dari semua angka genap antara 1 dan 100.

package main

import "fmt"

func main() {
	// Inisialisasi variabel untuk menyimpan total
	total := 0

	// Loop dari 1 hingga 100
	for i := 1; i <= 10; i++ {
		// Periksa apakah angka genap
		if i%2 == 0 {
			total += i // Tambahkan angka genap ke total
		}
	}

	// Cetak hasil
	fmt.Printf("Jumlah total angka genap dari 1 hingga 10 adalah: %d\n", total)
}
