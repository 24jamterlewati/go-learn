package main

import (
	"fmt"
)

// array adalah kumpulan data data yg tipe datanya sama

func main() {
	var Member [5]string
	Member[0] = "messi"
	Member[1] = "neymar"
	Member[2] = "suares"
	Member[3] = "ronaldo"
	Member[4] = "benzema"

	// mengambil semua data array
	fmt.Println(Member)

	// mengambil beberapa data array
	fmt.Println(Member[1])
	fmt.Println(Member[3])

	// cara mudah mengambil data array bisa menggunakan perulangan
	for i := 0; i < len(Member); i++ {
		fmt.Println(Member[i])
	}

	// cara paling simple untuk mengambil semua data array
	for index, value := range Member {
		fmt.Println("index", index, "value", value)
	}

	// jika hanya ingin mengambil value dari array kita bisa menghilangkan var index dengan menggantinya menjadi _,
	for _, value := range Member {
		fmt.Println(value)
	}

}
