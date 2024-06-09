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

<!-- ## Basic Data type -->
