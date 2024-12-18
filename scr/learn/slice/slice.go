package main

import"fmt"

func main() {

// ini adalah array
var days = [...]string{
	"senin",
	"selasa",
	"rabu",
	"kamis",
	"jumat",
	"sabtu",
	"minggu",
}

// ini adalah slice
var slice = days[:6]

fmt.Println(slice)
fmt.Println(len(slice))
fmt.Println(cap(slice))
fmt.Println(days)

// cara merunah nilai array dam slice
// di array ketika mengubah nilai slice nilai array pun ikut berubah begitupun sebaliknya
// ketika mengubah nilai array maka nilai slice akan ikut berubah

slice[0] = "di ubah"
fmt.Println(slice)
fmt.Println(days)

days[3] = "jangan deh"
fmt.Println(days)
fmt.Println(slice)

// cara menambahkan nilai array dan slice dengan append dan make

// membuar slice baru dengan mengambil data di array
var slice2 = days[:]
fmt.Println(slice2)

// menambah data baru dengan append
var slice3 = append(slice2, "saya libur")
fmt.Println(slice3)
slice3[0] = "ini bukan sabtu"
fmt.Println(slice3)

fmt.Println(slice2)
fmt.Println(days)

}