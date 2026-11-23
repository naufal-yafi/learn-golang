package main

import (
	"fmt"
)

func main() {
	// seperti halnya di javascript
	var name string = "Cak Husen"          // bisa diubah
	const birth_date string = "10/12/1999" // tidak bisa diubah / constant
	fmt.Println(name, birth_date)

	// versi lebih singkatnya
	nama := "Cak Husen"
	fmt.Println(nama)

	// multiple variable
	var (
		test_var   = 1
		test_var_1 = 2
	)
	fmt.Println(test_var, test_var_1)

	// integer bisa minus
	const min_small_number int8 = -128
	const max_small_number int8 = 127
	fmt.Println("small number: ", min_small_number, max_small_number)

	const min_number int16 = -32768
	const max_number int16 = 32767
	fmt.Println("number: ", min_number, max_number)

	const min_large_number int32 = -2147483648
	const max_large_number int32 = 2147483647
	fmt.Println("large number: ", min_large_number, max_large_number)

	const min_very_large_number int64 = -9223372036854775808
	const max_very_large_number int64 = 9223372036854775807
	fmt.Println("very large number: ", min_very_large_number, max_very_large_number)

	// integer ga bisa minus
	const small_number uint8 = 255
	const number uint16 = 65535
	const large_number uint32 = 4294967295
	const very_large_number uint64 = 18446744073709551615
	fmt.Println(small_number, number, large_number, very_large_number)

	// floating point
	const min_float float32 = -3.4028234663852886e+38
	const max_float float32 = 3.4028234663852886e+38
	const min_float64 float64 = -1.7976931348623157e+308
	const max_float64 float64 = 1.7976931348623157e+308
	fmt.Println(min_float, max_float, min_float64, max_float64)

	// string
	fmt.Println(len("Test"))
	fmt.Println("Test Sajah"[0], ": ", string("Test Sajah"[0]))
	fmt.Println("Test Sajah"[1], ": ", string("Test Sajah"[1]))

	// type (sama kek typescript) memberikan tipe data alias
	type TempatLahir string
	var tempat_lahir TempatLahir = "Yogyakarta"
	fmt.Println(tempat_lahir)
}
