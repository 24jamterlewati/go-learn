package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan dua angka pisahkan dengan spasi:")
	fmt.Scanf("%d %d", &a, &b)

	// switch tanpa expresion
	switch {
	case b == 1:
		fmt.Println(a + a)
	case b == 2:
		fmt.Println(a * a)
	case b == 3:
		fmt.Println("error kode salah")
	default:
		fmt.Println("Keluar Program")
	}

	// switch dengan expresion
	switch b {
	case 1:
		fmt.Println(a + a)
	case 2:
		fmt.Println(a * a)
	case 3:
		fmt.Println("error kode salah")
	default:
		fmt.Println("Keluar Program")
	}
}

// catatan pada switch expresion:
// inputan dgn tipe data int menggunakan case 1:
// inputan dgn tipe data char atau string menggunakan case "1":
