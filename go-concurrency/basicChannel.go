package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMessage(what chan string) {
    fmt.Println(<-what)
}
func BasicChannel2() {
    runtime.GOMAXPROCS(2)

    var messages = make(chan string)

    for _, each := range []string{"wick", "hunt", "bourne"} {
        go func(who string) {
            var data = fmt.Sprintf("hello %s", who)
            messages <- data
        }(each)
    }

    for i := 0; i < 3; i++ {
        printMessage(messages)
    }
}

func BasicChannel() {
    runtime.GOMAXPROCS(3)

    var messages = make(chan string)

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("hello %s", who)
        messages <- data
    }

    go sayHelloTo("john wick")
    go sayHelloTo("ethan hunt")
    go sayHelloTo("jason bourne")
	fmt.Println("main func 1")
	
	
    var message1 = <-messages
    fmt.Println(message1)
	
    var message2 = <-messages
    fmt.Println(message2)
	fmt.Println("main func 2")
	
    var message3 = <-messages
    fmt.Println(message3)

	fmt.Println("main func 3")
}

func BasicChannel3() {
    runtime.GOMAXPROCS(3)

    var messages = make(chan string)

    var sayHelloTo = func(who string) {
        var data = fmt.Sprintf("hello %s", who)
        messages <- data
		fmt.Println("insert who to channel: ", messages)
    }

    go sayHelloTo("john wick")
    go sayHelloTo("ethan hunt")
    go sayHelloTo("jason bourne")
	
	
	fmt.Println("main func 0")
    var message1 = <-messages
    fmt.Println(message1)
	fmt.Println("main func 1")
    var message2 = <-messages
    fmt.Println(message2)
	fmt.Println("main func 2")
    var message3 = <-messages
    fmt.Println(message3)
	fmt.Println("main func 3")
}


func BasicBufferChannel() {
    runtime.GOMAXPROCS(2)

    messages := make(chan int, 0)

    go func() {
        for {
            i := <-messages
            fmt.Println("receive data", i)
        }
    }()

    for i := 0; i < 5; i++ {
        fmt.Println("send data", i)
        messages <- i
    }

    time.Sleep(1 * time.Second)
}