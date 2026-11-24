package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Test",
		"address": "Yogyakarta",
		"test": "test",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	delete(person, "test")
	fmt.Println(person)

	// inisialisasi map dengan cara lain
	newPerson := make(map[string]string)
	newPerson["name"] = "Test"
	newPerson["address"] = "Jakarta"
	fmt.Println(newPerson)
}