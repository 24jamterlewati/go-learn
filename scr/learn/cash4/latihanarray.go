package main

import "fmt"

func showMenu() {
	var Menu [3]string
	Menu[0] = "1.Tampil Semua"
	Menu[1] = "2.Tambah"
	Menu[2] = "3.Hapus"

	for _, value := range Menu {
		fmt.Println(value)
	}
}

func main() {
	showMenu()
}
