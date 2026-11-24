package main

import "fmt"

func logging() {
	fmt.Println("Ini logging")
}

func runApplication() {
	/*
		digunakan jika mau dieksekusi terakhir tapi jika
		app crash ditengah jalan logging masih bisa dijalankan
		mungkin lebih jelas ada di file 2_panic.go
	*/ 
	defer logging();

	fmt.Println("Ini run app")
}

func main() {
	runApplication()
}