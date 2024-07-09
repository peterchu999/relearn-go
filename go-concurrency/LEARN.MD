# Concurrency

this part would cover the concurrency concept of golang. for more access https://go.dev/tour/concurrency/11

## Goroutine

Goroutine concept are quite similiar with concept of threading, however Goroutine is not a thread, it's very light weight and more fit to be called **mini thread**. goroutine execution is asynchronous, so the goroutine didn't wait for other goroutine execution.

it's nice to learn more concept behind goroutine [medium](https://medium.com/@hatronix/go-multitasking-comparing-concurrency-and-parallelism-in-go-812dc15d4245#:~:text=Go%20is%20designed%20to%20efficiently,by%20modern%20multi%2Dcore%20processors)

### basic goroutine implementation

goroutine process need to be warp inside a function, then add the keyword `go` before the function execution. Here is the example:

```go
package main

import "fmt"
import "runtime"

func print(till int, message string) {
    for i := 0; i < till; i++ {
        fmt.Println((i + 1), message)
    }
}

func main() {
    runtime.GOMAXPROCS(2)

    go print(5, "Goroutine Hello")
    print(5, "Normal Hello")

    var input string
    fmt.Scanln(&input) // we need this to hald the main function from exiting after the usual function execution is done to wait for goroutine function
}
```

OUTPUT:

```sh
1 Normal Hello
1 Goroutine Hello
2 Goroutine Hello
3 Goroutine Hello
4 Goroutine Hello
5 Goroutine Hello
2 Normal Hello
3 Normal Hello
4 Normal Hello
5 Normal Hello
```

explanation:

- `runtime.GOMAXPROCS(n)` => used to decide the how many core need to be activated for function execution
- `go` => keyword to make the function to run in it's own goroutine

## Channel

Channel is a mechanism to make connect 2 goroutine so that the value could be passed across goroutine. The channel mechanism execution is **_blocking_** or **_synchronous_**, let say we have 2 goroutine A and B, and B goroutine need a value from A goroutine that could be passed via channel. B goroutine wouldn't be execute until A goroutine passed the value to the channel.

### Channel Basic Implementation

channel could be make via `make` function and `chan` keyword, let see the example:

```go
package main

import (
	"fmt"
	"runtime"
)

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
```

output:

```sh
main func 1
hello john wick
hello jason bourne
main func 2
hello ethan hunt
main func 3
```

explanation:

- since goroutine execution is asynchronous so which goroutine run and finish first is decided by the core.
- the finised goroutine would put the value into the channel via `<-` keyword. `messages <- data`.
- eachtime time the goroutine is filled with value, the `var message1 = <-messages` block would be executed

### Channel as Parameter

we could pass channel as a function parameter, here is what it looks like:

```go
func printMessage(what chan string) {
    fmt.Println(<-what)
}
func main() {
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
```

### Buffer Channel

By default we are using `Unbuffered` Channel means that the channel would hold 1 value, and if the value isn't took by other goroutine, the new value couldn't be added to the channel.

However, using the `Buffered` Channel, we coudl defined the number of value the channel could hold and making it non blocking (async) until the channel value hold max amount of value. here is the example:

```go
package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    runtime.GOMAXPROCS(2)

    messages := make(chan int, 1)

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
```

OUTPUT:

```sh
send data 0
send data 1
send data 2
receive data 0
receive data 1
receive data 2
send data 3
send data 4
receive data 3
receive data 4
```

explantion:

- our channel could hold maximum 2 value so we could send over 3 data (1 was received immediatly) before it could send more value
