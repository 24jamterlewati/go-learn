package main

import "fmt"

func showMenu() {
	fmt.Println("== Pilih Opsi ==")
	fmt.Println("1.Tambah")
	fmt.Println("2.Kali")
	fmt.Println("3.Kurang")
	fmt.Println("4.Bagi")
}

func main() {
	var num1, num2 float64
	var choice int

	fmt.Println("== Kalkulator Sederhana ==")

	fmt.Print("masukkan nilai pertama: ")
	fmt.Scan(&num1)

	fmt.Print("masukkan nilai ke dua: ")
	fmt.Scan(&num2)

	showMenu()

	fmt.Print("Masukkan Opsi: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Hasil Dari %.f + %.f = %.f", num1, num2, num1+num2)
	case 2:
		fmt.Printf("Hasil Dari %.f * %.f = %.f", num1, num2, num1*num2)
	case 3:
		fmt.Printf("Hasil Dari %.f - %.f = %.f", num1, num2, num1-num2)
	case 4:
		if num2 != 0 {
			fmt.Printf("Hasil Dari %.f / %.f = %.f", num1, num2, num1/num2)
		} else {
			fmt.Printf("Hasil Dari %.f / %.f = 0", num1, num2)
		}
	default:
		fmt.Println("Pilihan Tidak Valid")
	}

}
