package main

import "fmt"

func main() {
	// array
	var values = [3]int8{ 90, 90, 90, }
	fmt.Println(len(values), values)

	// array tanpa batas
	var unlimited = [...]int8{ 90, 90, 90, 90, 90, 90, }
	fmt.Println(len(unlimited), unlimited)

	// slice
	buah := []string{ "nangka", "jambu" }

	// cap(slice): mendapatkan kapasitas
	fmt.Println(cap(buah))
	
	// append(slice, data): menambahkan data paling akhir
	buah = append(buah, "semangka")
	fmt.Println(buah)

	// make([]TypeData, length, capacity): membuat slice baru
	rakBuah := make([]string, len(buah))

	// copy(destination, source)
	copy(rakBuah, buah)
	fmt.Println(rakBuah)
}
