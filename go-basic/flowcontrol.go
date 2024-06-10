package main

import "fmt"

func FlowControl() {
	var score = 60

	if score > 70 {
		fmt.Println("great")
	} else if score > 50 {
		fmt.Println("well done")
	} else {
		fmt.Println("could be improve")
	}

	var grade = "A"

	switch grade {
		case "A":
			fmt.Println("excellent")
		case "B":
			fmt.Println("great")
		case "C":
			fmt.Println("well done")
		case "D":
			fmt.Println("need improvement")
	}

	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	var points = 6

	switch  {
	case points == 8:
		fmt.Println("perfect")
	case (points < 8) && (points > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value=", v)
	}
	var arr = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range arr {
		fmt.Println("Value=", v)
	}

	var sli = arr[0:2] // slice
	for _, v := range sli {
		fmt.Println("Value=", v)
	}
	
	var kv = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
	for k, v := range kv {
		fmt.Println("Key=", k, "Value=", v)
	}

}