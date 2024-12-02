package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan dua angka pisahkan dengan spasi:")
	fmt.Scanf("%d %d", &a, &b)

	switch {
	case b == 1:
		fmt.Println(a + a)
	case b == 2:
		fmt.Println(a * a)
	case b == 3:
		fmt.Println("error")
	default:

	}
}
