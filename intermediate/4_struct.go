package main

import "fmt"

type Customer struct {
	Name, Address string
	Age	int8
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var person Customer
	person.Name = "Test"
	person.Age = 32
	person.Address = "Yogyakarta"
	fmt.Println(person)
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Address)

	// literal
	new_person := Customer{
		Name: "Test 2",
		Age: 44,
		Address: "Jakarta",
	}
	fmt.Println(new_person)

	new_person_2 := Customer{ "Test 3", "Surabaya", 60 }
	fmt.Println(new_person_2)

	// method
	person.sayHello("comp")
	new_person.sayHello("comp")
	new_person_2.sayHello("comp")
}