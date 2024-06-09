package main

import "fmt"

func VariableMain() {
	var firstName = "john"
    var lastName string
	fmt.Printf("empty string: %s\n", lastName)
    lastName = "wick"
	middleName := "terry"

    fmt.Printf("halo %s %s %s!\n", firstName,middleName ,lastName)
}
