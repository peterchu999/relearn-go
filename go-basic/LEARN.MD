# Basic Go

this part would cover the basic such as `variable`, `data type`, `loop`, `conditional`. and basic way to run the golang

## Run Apps

`run.go` would have basic hello world code and the way to run the script was using

```
go run main.go
```

if you have dependent from other file (not package), then you need to include it in the go run, ex for this case `main.go` depend on `lvariable.go`

then we need to run like this:

```sh
go run main.go lvariable.go
```

or run this script for simplicity

```sh
go run *.go
```

### package

`package` keyword in the top of the `run.go` is mean to packaged a code (based on folder), so that other file (from other folder) could import the code from the current file.

however each project need to have `package main` as a way, so that go could know which file to run first.

### main function

`func main()` was also needed in the `package main` so that the golang could know which function would need to be execute first.

### comment

- `//` was used to comment a line.
- `/* */` was used for multiline comment

## fmt

`fmt` was a basic built-in library to access I/O kernel, such as scanning input, or printing anything. all of the function could be access from https://pkg.go.dev/fmt

## variable

there is several way to declare the variable

### using `var`

`var` keyword was use to declare a new variable that could be changed at the later point of the code.

var declaration could be seperate into several part, see the code below for example:

```go
var myString string = "value"
```

- `var` **keyword** was used to initialize var
- `myString` was the **variable name**
- `string` was the **data type**
- `"value"` was the **value** of the variable

if we directly assign the `value` to the variable, it's optional to declare the `data type`. Otherwise, if we declare the type but not the value go would directly assign an empty value on the variable such as `""` to `string` or `0` to `int` or `[]` to `array`.

### using `:=`

the other way to declare a new variable is using the `:=` syntax. but this way we need to directly assign the value. let see the example below:

```go
myString := "value"
```

- `myString` : was the **variable name**
- `:=` : was the **syntax to declare + assign**
- `"value"`: was **value** of the variable

### using `const`

exactly the same with var, and the multiline decalration using var but the changes was on the the keyword difference `var` to `const`.

`var` and `const` have difference use case, because `const` variable couldn't be reassign with a new value. `const` or constant usually use to store a constant value which we didn't want to change throughout our application life cycle such as value of `pi (22/7)` or number of second in hour `3600`.

### multiline declaration

we could declare and assing multi variable at the same time, example:

```go
var a, b, c string = "a", "b", "c"
```

or

```go
a, b, c := "a", "b", "c"
```

using `:=` multi declare-assign operator we could declare variable with multiple datatype

```go
a, b, c := "a", 2, true
```

### using `new()`

unlike in other language, `new` keyword is used to create a new pointer [pointer part]("")

example:

```go
name := new(string)

fmt.Println(name)   // 0x20818a220
fmt.Println(*name)  // ""
```

### using `make()`

make keyword was used to declare and create a special data type such as:

- channel [link]()
- slice [link]()
- map [link]()

## Basic Data type

example of basic data type code could be found in `ldatatype.go`

### integer

real number (not decimal) data type with variant of:
| Int Type | Range of number |
| :------: | :---------------------------------------------------: |
| `uint8` | 0 ↔ 255 |
| `uint16` | 0 ↔ 65535 |
| `uint32` | 0 ↔ 4294967295 |
| `uint64` | 0 ↔ 18446744073709551615 |
| `uint` | equals to `uint32` or `uint64` (depends on the value) |
| `byte` | equals to `uint8` |
| `int8` | -128 ↔ 127 |
| `int16` | -32768 ↔ 32767 |
| `int32` | -2147483648 ↔ 2147483647 |
| `int64` | -9223372036854775808 ↔ 9223372036854775807 |
| `int` | equals to `int32` or `int64` (depends on value) |
| `rune` | equals to `int32` |

### Float

float data type was used to store **decimal value**, there is 2 type of float which is `float32` and `float64` and the differences could be found on [stackoverflow](https://stackoverflow.com/questions/22337418/golang-floating-point-precision-float32-vs-float64)

### Boolean

boolean data type was used to store `true` or `false` data, usually used for conditional.

### string

string data type was used to store words in the variable such as `("this is string")`

### Nil and Zero value

`nil` is equivalent to `null` in other programming language, however it could only be assign to non-basic data type such as:

- pointer
- function
- slice
- map
- channel
- interface

`zero value` on the other hand was kind of default value of basic data type if we declare the variable without assignning any value, ex:

- `string` &#8594; `""`
- `int` &#8594; `0`
- `float` &#8594; `0.0`
- `boolean` &#8594; `false`

## Flow Control

code example of control flow could be seen in `flowcontrol.go`

### Basic if - else if - else

as usual programming language go also cover conditional using `if` `else if` and `else`. example:

```go
var score = 60

if score > 70 {
    fmt.Println("great")
} else if condition > 50 {
    fmt.Println("well done")
} else {
    fmt.Println("could be improve")
}
```

### Temporary If - else

unique in golang, we could use a temporary variable which only exists on the if - else if - else scope. ex:

```go
var point = 8840.0

if percent := point / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
} else if percent >= 70 {
    fmt.Printf("%.1f%s good\n", percent, "%")
} else {
    fmt.Printf("%.1f%s not bad\n", percent, "%")
}
```

### basic switch case

as the other programming language go also support `switch` case coditional scenario. example :

```go
var grade = "A"

switch grade {
    case "A":
        fmt.Println("excellent")
    case "B":
        fmt.Println("great")
    case "C":
        fmt.Println("well done")
    case "D", "F":
        fmt.Println("need improvement")
    default:
        {
            fmt.Println("default value")
        }
}
```

the `{}` was optional on switch case.

unlike other language, Golang switch case also support the if - else if - else style of condition, ex:

```go
var point = 6

switch {
case point == 8:
    fmt.Println("perfect")
case (point < 8) && (point > 3):
    fmt.Println("awesome")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}
```

unique in golang, the switch case didn't need a `break` keyword to prevent other case to be checked and run. if we want the other case to be checked and run, we need to use `fallthrough` keyword, example:

```go
var point = 6

switch {
case point == 8:
    fmt.Println("perfect")
case (point < 8) && (point > 3):
    fmt.Println("awesome")
    fallthrough
case point < 5:
    fmt.Println("you need to learn more")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}

/*
Output:
awsome
you need to learn more
*/
```

### for loop

looping in golang use for loop with the structure of:

```go

for i := 0; i < 5; i++ {
    fmt.Println(i)
}

```

golang didn't have `while` or `do-while` statement, instead the `for` loop have the capabilities of `while` loop. example:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

even more unique in golang `for`, we could make an infinite loop, which in other language we need to use `while true {}` with this structure:

```go
for {
    fmt.Println("infinite loop")
    break // to ensure it didn't really infinite and overload our machine
}
```

golang also used the `continue` and `break` keyword to control the flow inside the loop. example:

```go
for i := 1; i <= 10; i++ {
    if i % 2 == 1 {
        continue
    }
    if i == 7 {
        break
    }

    fmt.Println("odd number", i)
}
```

In golang we could use _label_ to label our loop, so when we use break, we could break through the label we had set. here is the example:

```go
outerLoopLabel:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i == 3 {
            break outerLoopLabel
        }
        fmt.Print("matrix [", i, "][", j, "]", "\n")
    }
}
```

golang have other method to loop through array, slice and map, etc. Using the `for` - `range` method.

- loop through `string` (_rune[]_)

```go
var s = "123" // string
for i, v := range s {
    fmt.Println("Index=", i, "Value=", v)
}

/*
Output:
Index= 0 Value= 49
Index= 1 Value= 50
Index= 2 Value= 51
*/
```

- loop through `array`

```go
var arr = [5]int{10, 20, 30, 40, 50} // array
for _, v := range arr {
    fmt.Println("Value=", v)
}
/*
Output:
Value= 10
Value= 20
Value= 30
Value= 40
Value= 50
*/
```

- loop through `slice`

```go
var sli = arr[0:2] // slice
for _, v := range sli {
    fmt.Println("Value=", v)
}
/*
Output:
Value= 10
Value= 20
*/
```

- loop through `map`

```go
var kv = map[byte]int{'a': 0, 'b': 1, 'c': 2} // map
for k, v := range kv {
    fmt.Println("Key=", k, "Value=", v)
}

/*
Output:
Key= 97 Value= 0
Key= 98 Value= 1
Key= 99 Value= 2
*/
```

- loop with range (like python)

```go
for i := range 5 {
    fmt.Print(i) // 01234
}
/*
Output:
01234
*/
```
