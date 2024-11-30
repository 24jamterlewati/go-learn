package main

import (
	"fmt"
)

func main() {
	var a, b int

	// input dari pengguna
	fmt.Print("masukkan dua angka (pisahkan dengan spasi): ")

	// untuk menyimpan nilai yg di input oleh pengguna
	fmt.Scanf("%d %d", &a, &b)

	// output yg di hasilkan
	fmt.Printf("hasil perkalian adalah %d\n", a*b)
	fmt.Printf("hasil penjumlahan adalah %d\n", a+b)
	fmt.Printf("hasil pembagian adalah %d\n", a/b)
	fmt.Printf("hasil perkurangan adalah %d\n", a-b)
	fmt.Printf("hasil modulus adalah %d\n", a%b)
}
