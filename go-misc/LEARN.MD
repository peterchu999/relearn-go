# Misc Learning

commonly used topic when building go application.

## Defer & Exit

`defer` is used to move the line of code to be execute after the scope (function scope) finished. Whereas, `exit` function is used to stop the application directly.

the example code could be seen too in `defer-exit.go` file

```go
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

func main() {
	withDefer()
	defer fmt.Println("won't be called")
	os.Exit(0)
	fmt.Println("program has exit, this wont be called")
}
```

OUTPUT

```sh
this would execute first
Even after sleep, i would be called first before defer
hello, currently im the  2 defer
hello, currently im the  1 defer
```

explanation:

- `defer` keyword before function execution would postpone the function execution and put it into some kind of `stack`, where we could see the `sayHello` with index 2 execute first before the index 1
- `os.Exit` function would forcibly end the application life cycle, even the defer wont be executed after exit.

## Error Handling

golang error handling is quite unique and quite deviate from usual `throw - catch - finally` approch in any other language.

### Basic Error Handling

error itself is a data type in golang, usually instead of throwing error golang function would return the error. let see the example:

```go
func division(a, b int) (int, error) {
    if b == 0 {
		return 0, errors.New("error cannot divide by 0")
	}
	return a / b, nil
}

func runErrorCode() {
	a, b := 10, 0
	val, err := division(a, b)
	if err != nil{
		fmt.Println(err.Error())
	} else {
		fmt.Println("division result :", val)
	}
}
```

OUTPUT

```sh
error cannot divide by 0
```

**explanation:**
since golang could return multiple value, this approach were often used for error handling, by checking if there is any error return from the function. `errors.New` function is used to create a new error object. and `.Error` function is a method to get the error message from Error object.

### Panic

`panic` is a function to stop the goroutine **_(process)_** immidiately and print the stacktrace. `main` function was considered as a goroutine itself, so if panic were triggered in main goroutine, the application would stop immidiately. if 1 goroutine are panic, it would affect the program as a whole unless `recover` were introduced.

example:

```go
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

func runPanic() {
	defer fmt.Println("would still be called before panic")
	if val, err := division(1, 0); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("division result were ", val)
	}
}

func main() {
	go runPanic()
	runErrorCode()
	select{}
}
```

OUTPUT

```sh
division result : 10
would still be called before panic
panic: cannot divide by 0

goroutine 18 [running]:
main.runPanic()
        /Users/peterandrew/Documents/Go Lang/relearn-go/go-misc/error.go:28 +0x9d
created by main.Error in goroutine 1
        /Users/peterandrew/Documents/Go Lang/relearn-go/go-misc/error.go:35 +0x1a
exit status 2
```

explanation:

when a goroutine went panic, the process are being stop immidiately, all of the defer function within the goroutine scope would be execute first before the panic function print the stacktrace.

### Recover

recover were a funtion to save the panic goroutine from shutting down the whole application, recover works like `catch` function in any other language to prevent error from breaking the application. However, if the `recover` function were triggered. here is the updated example from the panic case above.

```go
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

func main() {
	// defer catch()
	runErrorCode()
	runPanic()
	fmt.Println("called even after runPanic, panic")

}

```

explanation:

the `catch` function which later be defer in `runPanic` function were execute right before the panic stop the program and print the stacktrace, it recover the program immidiately.

## Random

learn about `random` package provided by golang

### random integer

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    randomizer := rand.New(rand.NewSource(10))
    fmt.Println("random 1:", randomizer.Int()) // 5221277731205826435
    fmt.Println("random 2:", randomizer.Int()) // 3852159813000522384
    fmt.Println("random 3:", randomizer.Int()) // 8532807521486154107
}
```

### random with range

```go
randomizer := rand.New(rand.NewSource(10))
randomizer := rand.New(rand.NewSource(10))
fmt.Println("random 1:", randomizer.Intn(100)) // 54
fmt.Println("random 2:", randomizer.Intn(100)) // 48
fmt.Println("random 3:", randomizer.Intn(100)) // 67
```

## Time

learn about `time` package

### basic usage

```go
var mytime = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
fmt.Printf("time %v\n", mytime)
// time 2011-12-24 10:20:00 +0000 UTC
```

### Parsing From String

here is the example to parse string data type to time data type

```go
layoutFormat := "2006-01-02 15:04:05"
value := "2015-09-02 08:04:00"
date, err := time.Parse(layoutFormat, value)

if err != nil {
    fmt.Println("error while parsing")
} else {
    fmt.Println(value, "\t->", date.String()) // // 2015-09-02 08:04:00 +0000 UTC
}
```

explanation:

- golang layout format are unique more could be seen in https://yourbasic.org/golang/format-parse-string-time-date-example/

### Formatting From Time to String

```go
var date, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")

var dateS1 = date.Format("Monday 02, January 2006 15:04 MST")
fmt.Println("dateS1", dateS1)
// Wednesday 02, September 2015 08:00 WIB

var dateS2 = date.Format(time.RFC3339)
fmt.Println("dateS2", dateS2) // 2015-09-02T08:00:00+07:00
```

## Type Casting

### strconv

strconv have several function to help for type casting

#### Atoi

`strconv.Atoi` is used to convert string to integer. here is the example:

```go
var str = "124"
var num, err = strconv.Atoi(str)

if err == nil {
    fmt.Println(num) // 124
}
```

#### Itoa

`strconv.Itoa` is used to convert integer to string. here is the example

```go
var num = 124
var str = strconv.Itoa(num)

fmt.Println(str) // "124"
```

#### Parse Bool

```go
var bul = true
var str = strconv.FormatBool(bul)

fmt.Println(str) // true
```

## Typecasting Technqiue

instead of using `strconv` golang prepare other way of typecasting. here is the example

```go
// string to array of byte
var text1 = "halo"
var b = []byte(text1)

fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3]) // 104 97 108 111
```

```go
// byte to string
var byte1 = []byte{104, 97, 108, 111}
var s = string(byte1)

fmt.Printf("%s \n", s) // halo
```

```go
// char to int
var c int64 = int64('h')
fmt.Println(c) // 104

// int to char to string
var d string = string(104)
fmt.Println(d) // h
```

### Type Assertion

if we are using any / interface{} data type we could implement type assertion to cast this variable into the spesific data type we wanted

```go
var data = map[string]interface{}{
    "name":    "john wick",
    "grade":   2,
    "height":  156.5,
    "isMale":  true,
    "hobbies": []string{"eating", "sleeping"},
}

fmt.Println(data["name"].(string))
fmt.Println(data["grade"].(int))
fmt.Println(data["height"].(float64))
fmt.Println(data["isMale"].(bool))
fmt.Println(data["hobbies"].([]string))
```

## File System

learn how to create, manipulate and delete a file within a system

### create a new file

```go
package main

import "fmt"
import "os"

var path = "./test.txt"

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func createFile() {
    // detect if file exists
    var _, err = os.Stat(path)

    // create new file if it's not exists
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) { return }
        defer file.Close()
    }

    fmt.Println("==> file created", path) // ==> file created ./test.txt
}

func main() {
    createFile()
}
```

### Writing to file

we need to open the file with write access beforehand, then update the file accordingly.

```go
func writeFile() {
    // open filw with READ & WRITE access
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) { return }
    defer file.Close()

    // write string to the file
    _, err = file.WriteString("Hello World\n")
    if isError(err) { return }
    _, err = file.WriteString("write from the code\n")
    if isError(err) { return }

    // save the changes with Sync function
    err = file.Sync()
    if isError(err) { return }

    fmt.Println("==> File Edited")
}

func main() {
    writeFile()
}
```

### Reading from file

we need to use `io` package to read data stream from a file. here is the example:

```go
func readFile() {
    // Open file with Read only access
    var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
    if isError(err) { return }
    defer file.Close()

    // Read the string and store it into our array of byte
    var text = make([]byte, 1024)
    for {
        n, err := file.Read(text)
        if err != io.EOF {
            if isError(err) { break }
        }
        if n == 0 {
            break
        }
    }
    if isError(err) { return }

    fmt.Println("==> File Successfully read")
    fmt.Println(string(text))
}
```

### Delete File

we could remove a file using `os.Remove`

```go
func deleteFile() {
    var err = os.Remove(path)
    if isError(err) { return }

    fmt.Println("==> file deleted")
}

func main() {
    deleteFile()
}
```

## Unit Testing

golang provide `testing` package that consist of a lot of testing helper function. the code example would be under `testing/` folder

### Unit Testing Example

we would make a simple calculator function `add` and `diff`. then we would test the code via unit test

```go
// functionality.go
package main

func Add(a, b int) int {
	return a + b
}

func Diff(a, b int) int {
	return a - b
}
```

```go
// functionality_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1,2)
	t.Logf("Case A: 1, B: 2, expected: 3, result: %d",result)

	if result != 3 {
		t.Errorf("result is wrong")
	}
}

func TestDiff(t *testing.T) {
	result := Diff(1,2)
	t.Logf("Case A: 1, B: 2, expected: 1, result: %d",result)

	if result != 1 {
		t.Errorf("result is wrong")
	}
}
```

to run the test use `go test *.go -v` command, `-v` was used to also print which case falied and success.

### BenchMarking Example

golang test also provide benchmarking capabilities. here is the example:

```go
func BenchmarkAdd(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		Add(i, i+1)
	}
}
```

output:

```sh
BenchmarkAdd-16         1000000000               0.2533 ns/op
PASS
ok      command-line-arguments  0.787s
```

### Other Unit Testing Package

there is popular 3rd party api `github.com/stretchr/testify` that have more similiarity with other testing api out there.

run `go get github.com/stretchr/testify` to get this package. this package contains serveral function namely

| subpackage | usage                                                                   |
| ---------- | ----------------------------------------------------------------------- |
| `assert`   | Standart tool for unit testing                                          |
| `http`     | Tools for http testing                                                  |
| `mock`     | Tools to mock an object while testing                                   |
| `require`  | similiar with assert, however it stop the testing immidiately when fail |
| `suite`    | Tools for testing for `struct` and `method`                             |

## Generics

Generic is a way to define data type when the function are being called rather than on the function definition itself.

### Generic function example

```go
func Sum(numbers []int) int {
    var total int
    for _, e := range numbers {
        total += e
    }
    return total
}

func SumGeneric[V int | float32 | float64](numbers []V) V {
    var total V
    for _, e := range numbers {
        total += e
    }
    return total
}

// when calling
SumGeneric[int]([]int{1, 2, 3, 4, 5})
SumGeneric[float32]([]float32{2.5, 7.2})
SumGeneric[float64]([]float64{1.23, 6.33, 12.6})
// or
SumGeneric([]int{1, 2, 3, 4, 5})
SumGeneric([]float32{2.5, 7.2})
SumGeneric([]float64{1.23, 6.33, 12.6})
```

### `comparable`

`comaprable` keyword, is used to make the data type to be any, example

```go
func MyPrint[K comparable](para K) {
    fmt.Println(para)
}
```

### Generic Type Constraint

```go
type Number interface {
    int64 | float64
}

func Add[V Number](a,b V) V {
    return a + b
}
```

### Generic Struct

```go
type UserModel[T int | float64] struct {
    Name string
    Scores []T
}

var m1 UserModel[int]
m1.Scores = []int{1,2,3}
```
