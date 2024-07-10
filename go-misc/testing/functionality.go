package main

func Add(a, b int) int { 
	return a + b
}

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}