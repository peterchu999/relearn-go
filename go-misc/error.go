package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
    if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

func runErrorCode() {
	a, b := 10, 1
	val, err := division(a, b)
	if err != nil{
		fmt.Println(err.Error())
	} else {
		fmt.Println("division result :", val)
	}
}

func catch() {
	if rec := recover(); rec != nil {
		fmt.Println("Error occured ", rec)
	} else {
		fmt.Println("Everything works as expected")
	}
}

func runPanic() {
	defer catch()
	defer fmt.Println("would still be called before panic")
	if val, err := division(1, 0); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("division result were ", val)
	}
	fmt.Println("still called")
}

func Error() {
	// defer catch()
	runErrorCode()
	runPanic()
	fmt.Println("called even after runPanic, panic")
	
}