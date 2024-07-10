package main

import (
	"fmt"
	"os"
	"time"
)

func sayHello(idx int) {
    fmt.Println("hello, currently im the ",idx, "defer")
}

func withDefer() {
    defer sayHello(1)
	defer sayHello(2)
    fmt.Println("this would execute first")
	time.Sleep(1 * time.Second)
	fmt.Println("Even after sleep, i would be called first before defer")
}

func DeferExit() {
	withDefer()
	defer fmt.Println("won't be called")
	os.Exit(0) 
	fmt.Println("program has exit, this wont be called")
}