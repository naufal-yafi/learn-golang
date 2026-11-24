package main

import "fmt"

func logging() {
	fmt.Println("Ini logging")
}

func runApplication(error bool) {
	defer logging();

	if error {
		// penggunaan panic akan melakukan penghentian app secara paksa
		panic("Error")
	}
}

func main() {
	runApplication(true)
	fmt.Println("ini tidak akan dieksekusi karena stop paksa appnya")
}