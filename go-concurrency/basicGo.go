package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func BasicRun() {
    runtime.GOMAXPROCS(3)

    go print(5, "Goroutine Hello")
    print(5, "Normal Hello")

    var input string
    fmt.Scanln(&input) // we need this to hald the main function from exiting after the usual function execution is done to wait for goroutine function
}