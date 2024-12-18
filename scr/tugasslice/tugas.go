// tugas
// cobalah membuat slice, dan tampilkan data slice, lalu tambahkan data slice, dan hapus data slice

package main

import "fmt"

func main() {

	// membuat slice
	var slice = []int{1, 2, 3, 4, 5}

	// menampilkan data slice
	fmt.Println(slice)

	// menambah data sli,ce
	var tambahData = append(slice, 6, 7, 8, 9)
	fmt.Println(tambahData)

	// menghapus semua data slice
	slice = []int{}
	fmt.Println("data slice di hapus", slice)

}
