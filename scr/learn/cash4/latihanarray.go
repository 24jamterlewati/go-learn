package main

import (
	"fmt"
)

var items = []string{} // Slice untuk menyimpan item

// Fungsi untuk menampilkan semua item
func showItems() {
	if len(items) == 0 {
		fmt.Println("Belum ada item yang disimpan.")
		return
	}
	fmt.Println("Daftar item:")
	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}

// Fungsi untuk menambahkan item baru
func addItem(newItem string) {
	items = append(items, newItem)
	fmt.Println("Item berhasil ditambahkan:", newItem)
}

// Fungsi untuk menghapus item berdasarkan indeks
func deleteItem(index int) {
	if index < 1 || index > len(items) {
		fmt.Println("Indeks tidak valid.")
		return
	}
	deletedItem := items[index-1]
	items = append(items[:index-1], items[index:]...)
	fmt.Println("Item berhasil dihapus:", deletedItem)
}

// Fungsi untuk menampilkan menu
func showMenu() {
	fmt.Println("\n--- Menu ATM ---")
	fmt.Println("1. Tampilkan Semua Item")
	fmt.Println("2. Tambah Item")
	fmt.Println("3. Hapus Item")
	fmt.Println("4. Keluar")
}

func main() {
	var choice int
	for {
		showMenu()
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			showItems()
		case 2:
			var newItem string
			fmt.Print("Masukkan item baru: ")
			fmt.Scanln(&newItem)
			addItem(newItem)
		case 3:
			var index int
			fmt.Print("Masukkan nomor item yang ingin dihapus: ")
			fmt.Scanln(&index)
			deleteItem(index)
		case 4:
			fmt.Println("Terima kasih telah menggunakan program ini.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
