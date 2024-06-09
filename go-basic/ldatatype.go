package main

import "fmt"

func LDataType() {
	var integer int = 1
	var floater	float32 = 0.1
	var boolean bool = true
	var stringer string = "string"
	var zeroVal string
	var nilVal interface{} = nil
	fmt.Println(integer, floater, boolean, stringer, zeroVal, nilVal)
}