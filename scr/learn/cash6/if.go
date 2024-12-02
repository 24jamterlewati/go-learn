package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("masukkan dua angka dan pisahkan dengan spasi(ex: 10 6): ")
	fmt.Scanf("%d %d", &a, &b)
	
		if b == 1 {
			fmt.Println(a + a)
		} else if b == 2 {
			fmt.Println(a * a)
		} else if b == 3 {
			fmt.Println("error tidak dikenali")
		} else {
			fmt.Println("keluar Program")
		}

}
