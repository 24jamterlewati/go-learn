package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("masukkan dua angka (pisahkan dengan spasi): ")
	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("hasilnya pertama adalah %d\n", a)
	fmt.Printf("hasil ke dua adalah %d", b)
}
