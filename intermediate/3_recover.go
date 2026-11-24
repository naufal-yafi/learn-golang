package main

import "fmt"

func logging() {
	fmt.Println("Ini logging")

	// recover itu dikenal dengan try catch
	message := recover()
	fmt.Println("Error: ", message)
}

func runApplication(error bool) {
	defer logging();

	if error {
		// penggunaan panic akan melakukan penghentian app secara paksa
		panic("Stop sementara")
	}
}

func main() {
	runApplication(true)
	fmt.Println("ini akan dieksekusi karena sudah ada recover")
}