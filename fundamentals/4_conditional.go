package main

import "fmt"

func main() {
	name := "Test"

	if name == "Test" {
		fmt.Println(true)
	} else if length := len(name); length < 3  {
		fmt.Println(false)
	} else if length := len(name); length > 2  {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	switch name {
	case "Test":
		fmt.Println(true)
	default:
		fmt.Println(false)
	}

	switch length := len(name); length > 2 {
	case true:
		fmt.Println(true)
	default:
		fmt.Println(false)
	}
}