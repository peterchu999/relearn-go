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

### Select Channel

what happen when there is 2 channel and we want to process it simultaneously ? if we use usual way our channel is receiving code would got block. so here is the example

we got two function to calculate maximum and average:

```go
func getAverage(numbers []int, ch chan float64) {
    var sum = 0
    for _, e := range numbers {
        sum += e
    }
    ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
    var max = numbers[0]
    for _, e := range numbers {
        if max < e {
            max = e
        }
    }
    ch <- max
}
```

here is what we know about processing channel:

```go
runtime.GOMAXPROCS(2)
var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
fmt.Println("numbers :", numbers)

var ch1 = make(chan float64)
go getAverage(numbers, ch1)

var ch2 = make(chan int)
go getMax(numbers, ch2)

var avg = <-ch1
fmt.Printf("Avg \t: %.2f \n", avg)
var max = <-ch2
fmt.Printf("Max \t: %d \n", max)
```

using this code, no matter how fast the `getMax` function calculated, it need to wait for `getAverage` function to put the value to channel `ch1` before the code below `<-ch1` got execute.

here is where the `select` come to help. by using select we could execute the line of code which the got value first.

```go
runtime.GOMAXPROCS(2)

var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
fmt.Println("numbers :", numbers)

var ch1 = make(chan float64)
go getAverage(numbers, ch1)

var ch2 = make(chan int)
go getMax(numbers, ch2)

for i := 0; i < 2; i++ {
    select {
    case avg := <-ch1:
        fmt.Printf("Avg \t: %.2f \n", avg)
    case max := <-ch2:
        fmt.Printf("Max \t: %d \n", max)
    }
}
```

explanation:

- the select part make sure, whoever channel send data first got process.
- the for loop make sure we execute the seelct 2 times since we got 2 channel and 2 goroutine execution to that channel

**PS:** note that channel in goroutine were made for communication and data sharing and not for execution control. the `select` use case would better shown in later part of concurrency concept

### Channel Range and Close

`for - range` syntax could be also be used for channel loop execution. let say we had a goroutine that produce 20 data. we would code it this way for now:

```go
func sendMessage(ch chan<- string) {
    for i := 0; i < 20; i++ {
        ch <- fmt.Sprintf("data %d", i)
    }
    close(ch)
}

func printMessage(ch <-chan string) {
    for i := 0; i < 20; i++ {
        message := <-ch
        fmt.Println(message)
    }
}
```

we used constant 20 value in both part, however what happen if we forgot to update either part if the code changes ? we could use `for - range` loop for `printMessage` function.

```go
func printMessage(ch <-chan string) {
    for message := range ch {
        fmt.Println(message)
    }
}
```

now the printMessage would print all of the message untill the channel are closed by `close(ch)` function

### Channel Timeout

since we are usually just setting the boundaries of value retriving from channel via our loop condition, or channel closed. we could actually set a timeout condition to our loop that wait for the channel value. here is what it looks like

```go
func sendData(ch chan<- int) {
    for i := 0; true; i++ {
        ch <- i
        time.Sleep(time.Duration(i) * time.Second)
    }
}

func retreiveData(ch <-chan int) {
    loop:
    for {
        select {
        case data := <-ch:
            fmt.Print(`receive data "`, data, `"`, "\n")
        case <-time.After(time.Second * 5):
            fmt.Println("timeout. no activities under 5 seconds")
            break loop
        }
    }
}

func Timeout() {
 	runtime.GOMAXPROCS(2)

    var messages = make(chan int)

    go sendData(messages)
    retreiveData(messages)
}
```

OUTPUT:

```sh
receive data "0"
receive data "1"
receive data "2"
receive data "3"
receive data "4"
receive data "5"
timeout. no activities under 5 seconds
```

since we are setting out sleep period based on our loop index, when the loop index are more than 5 the system would wait more than 5 second, therefore triggering the timeout scope block `case <-time.After(time.Second * 5)` and break the loop.

## Sync

golang provide `sync` package to manage goroutine

### Wait Group

`sync.WaitGroup` is use to synchronize goroutine flow, it is crafted to manage goroutine easier and effective compared to our mangement using channel

#### Implementation

```go
package main

import "sync"
import "runtime"
import "fmt"

func doPrint(wg *sync.WaitGroup, message string) {
    defer wg.Done()
    fmt.Println(message)
}

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup

    for i := 0; i < 5; i++ {
        var data = fmt.Sprintf("data %d", i)

        wg.Add(1)
        go doPrint(&wg, data)
    }

    wg.Wait()
}
```

in this implemenetation the main process would wait untill the waitgroup empty before ended the process, it different with channel that are waiting for spesific channel. this waitgroup wait for all goroutine.

### Mutex

Mutex "mutual exclusion" is used to handle a **race condition**. a race condition is a condition where more than 1 goroutine access the same piece of data at the same time, therefore the data would be messy (example when a process are reading a data and the other are updating it, we dont know the get result are updated or not).

Mutex change the data access level to be exclusive so it could only be access by 1 goroutine at time. If a race condition occur, only the first gorutine that access the data could access it at time, other goroutine need to wait untill the data finished to be access by this goroutine before accesing it.

#### race condition code

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
    val int
}

func (c *counter) Add(int) {
    c.val++
}

func (c *counter) Value() (int) {
    return c.val
}

func RaceCondition() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    var meter counter

    for i := 0; i < 1000; i++ {
        wg.Add(1)

        go func() {
            for j := 0; j < 1000; j++ {
                meter.Add(1)
            }

            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(meter.Value())
}
```

OUTPUT:

```
785426
783606
786793
```

explanation:

- inconsistent output throughout testing, and it suppose to be 1000 \* 1000 = 1000000
- the race condition occur could be detected via this command `go run -race *.go`

```sh
981929
Found 2 data race(s)
exit status 66
```

#### mutex help

to prevent the race condition like above, we need to implements mutex to the data source. here is what the code looks like

```go
type counter struct {
    sync.Mutex
    val int
}

func (c *counter) Add(int) {
    c.Lock()
    c.val++
    c.Unlock()
}

func (c *counter) Value() (int) {
    return c.val
}
```

with this code we are locking the counter object so that it could perform the add operation before other access it.

we could also attached the locking logic in the goroutine code, however, attaching it to object were the advised approach.

```go
func (c *counter) Add(int) {
    c.val++
}

func (c *counter) Value() (int) {
    return c.val
}

func main() {
    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    var mtx sync.Mutex
    var meter counter

    for i := 0; i < 1000; i++ {
        wg.Add(1)

        go func() {
            for j := 0; j < 1000; j++ {
                mtx.Lock()
                meter.Add(1)
                mtx.Unlock()
            }

            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println(meter.Value())
}
```
