# Go Medium

this part would cover the intermediate concept such as `array`, `slice`, `map`, `interface`, `struct` and `functions`

## Array

array in golang are type strict, so the element inside need to had the same data type. here is the way to declare an array

```go
var names [5]string // initialize a variable names with type of array that could hold 5 element

names := [5]string{}
```

Above way would generate an array that have and could only hold 5 element (`["", "", "", "", ""]`)

here is a way to initialize array without defining the capacity, the capacity would directly be count based on the number of element provided during initialization

```go
var numbers = [...]int{2, 3, 2, 4, 3}

fmt.Println("array data \t:", numbers) // [2, 3, 2, 4, 3]
fmt.Println("element count \t:", len(numbers)) // 5
```

### multi dimention array

just as the other language, golang accomodate multi dimension array. here is the example:

```go
var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

fmt.Println("numbers1", numbers1) // [[3 2 3] [3 4 5]]
fmt.Println("numbers2", numbers2) // [[3 2 3] [3 4 5]]
```

## Slice

slice is an reference of element inside `array`, slice could be created by referencing an value inside `array` or manipulating `array` or other `slices`. because slice is an reference to array or other slice, the value inside slice would be affected to changes of the original `array` or `slice`.

### initialize slice

the slice initialization would be similiar to array initialization, here is the example

```go
var fruits = []string{"apple", "grape", "banana", "melon"}
var fArr = [4]string{"apple", "grape", "banana", "melon"}
var fSlice = fArr[:]
```

the example above show how to initialize slice, and initilize it using array as reference

### side effect

Since slice is a reference, there would be side effect if the original data got changes. here is the example

```go
var data = [5]int{1,2,3,4,5}
var sData = data[:]

fmt.Println(data) // [1 2 3 4 5]
fmt.Println(sData) // [1 2 3 4 5]

data[0] = 6

fmt.Println(data) // [6 2 3 4 5]
fmt.Println(sData) // [6 2 3 4 5]

sData[1] = 7
fmt.Println(data) // [6 7 3 4 5]
fmt.Println(sData) // [6 7 3 4 5]

```

the side effect happen bothway, so we need to becareful when playing with array. to mitigate this side effect we could use `copy` function to copy the value instead of referencing it

```go
var data = [5]int{1,2,3,4,5}
var sData = make([]int, 5)
copy(sData, data[:])

fmt.Println(data) // [1 2 3 4 5]
fmt.Println(sData) // [1 2 3 4 5]

data[0] = 6

fmt.Println(data) // [6 2 3 4 5]
fmt.Println(sData) // [1 2 3 4 5]

sData[1] = 7
fmt.Println(data) // [6 2 3 4 5]
fmt.Println(sData) // [1 7 3 4 5]
```

### appending new data

appending new data is a crucial part of array in any other language. in golang data could only be append in slice using `append` function. here is a way to append new data

```go
var data = []int{1,2,3}
var data2 = []int{5,6,7}

data = append(data, 4) // [1, 2, 3, 4]

data = append(data, data2...) // concat : [1, 2, 3, 4, 5, 6, 7]
```

## Map

Map in golang was similiar to hashmpa in java, it's key-value pair.

### initialize map

here is the example of map initialization:

```go
var memo = map[string]int{}
memo["laundry"] = 1
memo["cook"] = 2

fmt.Println(memo) // map[cook:2 laundry:1]


var memo2 = make(map[string]int)
var memo3 = *new(map[string]int)
```

`memo3` would hold the address instead of the map value.

### check if key exists

to check if the spesific key exists in the map we could access the map this way :

```go
var memo = map[string]int{}
memo["laundry"] = 1
memo["cook"] = 2

var value, isExists = memo["laundry"]
fmt.Println(value) // 1
fmt.Println(isExists) // true
```

### delete item in array

to remove the elemen in array, so that the checking above would return false. `delete` function needed. example:

```go
var memo = map[string]int{}
memo["laundry"] = 1
memo["cook"] = 2

delete(memo, "laundry")

var value, isExists = memo["laundry"]
fmt.Println(value) // 0
fmt.Println(isExists) // false
```

## Function

### Create a function

to create a function, `func` keyword is used, with the example of below:

```go
func functionName (parameter1 int, parameter2 int) int {
    return parameter1 + parameter2
}
```

### Function with multiple return value

there is 2 ways to return multiple value in golang

- reguler

  ```go
      package main

      import "fmt"
      import "math"

      func calculate(d float64) (float64, float64) {
          var area = math.Pi * math.Pow(d / 2, 2)
          var circumference = math.Pi * d
          return area, circumference
      }
  ```

- predefined return
  ```go
      func calculate(d float64) (area float64, circumference float64) {
          area = math.Pi * math.Pow(d / 2, 2)
          circumference = math.Pi * d
          return
      }
  ```

### Vardiac Function

vardiac function is used to have unlimited amount of parameter.

Example of vardiac function:

```go
func sum(numbers ...int) int {
    res := 0
    for num := range numbers {
        res += num
    }
    return res
}

// called by
acc := sum(1, 2, 3, 4, 5, 6) // acc = 21
numbers := []int {1, 2, 3, 4, 5, 6}
acc2 := sum(numbers...) // acc2 = 21
```

we could combine the usual function parameter with vardiac parameter, example:

```go
func operation(operand string, numbers ...int) int {
    switch operand {
        case "+":
            fmt.Println(operand, numbers)
        case "-":
            fmt.Println(operand, numbers)
        default:
            fmt.Println("invalid operand")
    }
}

operation("+", 1, 2, 3)
//or
operation("+", []int{1, 2, 3}...)
```

### Closure / Anoymus function

we could store a function inside a variable, or have it as return value, here is the example

```go
package main

import "fmt"

func main() {
    var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

    var newNumbers = func(min int) []int {
        var r []int
        for _, e := range numbers {
            if e < min {
                continue
            }
            r = append(r, e)
        }
        return r
    }(3)

    fmt.Println("original number :", numbers)
    fmt.Println("filtered number :", newNumbers)
}
```

and

```go
func findMax(numbers []int, max int) (int, func() []int) {
    var res []int
    for _, e := range numbers {
        if e <= max {
            res = append(res, e)
        }
    }
    return len(res), func() []int {
        return res
    }
}
```

### Function as parameter

we could have a function as another function parameter, here is what it looks like:

```go
package main

import "fmt"
import "strings"

func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}
```

We could make an alias to gave our function, example:

```go
package main

import "fmt"
import "strings"

type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}
```

## Pointer

pointer is an reference of a value in memory, usually pointer hold the **address** of the value. pointer in golang is very similiar with c pointer.

### Example of pointer

```go
var numberA int = 4
var numberB *int = &numberA

fmt.Println("numberA (value)   :", numberA)  // 4
fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

fmt.Println("numberB (value)   :", *numberB) // 4
fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
```

since pointer is a reference, same with slice, when the original data got change the reference value also changed

```go
var numberA int = 4
var numberB *int = &numberA

fmt.Println("numberA (value)   :", numberA)  // 4
fmt.Println("numberA (address) :", &numberA) // 0xc20800a220

numberA = 5

fmt.Println("numberB (value)   :", *numberB) // 5
fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
```

### Pass pointer

when we are using basic data type as parameter, usually it is pass by value scenario, if we want to make side effect we could pass the pointer to pass it by reference. example:

```go
func main() {
	var a int = 10
	unchanged(a, 11)
	fmt.Println(a) // 10
	change(&a, 12)
	fmt.Println(a) // 12
}

func unchanged(original int, value int) {
	original = value
}

func change(original *int, value int) {
	*original = value
}
```

## Struct

go doens't have class like any other oop language, so it have struct instead like c language.

struct is a bunch of variable and, or method that are group together to become a new data type.

### define and initialize struct

here is the example of how to define a new struct and initilize it's _object_ to be store in variable

```go
type StructName struct {
    name string
    age int
}

//initialize it's object

var s StructName // initialize struct with each property zero value
var s2 = StructName{} // initialize struct with each property zero value
var s3 = StructName{"name", 11} // initialize struct with property
s4 := StructName{"name", 11} // initialize struct with property


// set value
s.age = 3
s.name = "golang"

fmt.Println(s.age, s.name) // 3 golang
```

### struct pointer

struct object could have pointer too like any other data type, and we could use it to pass by reference on a function. Example

```go
type Student struct {
	name string
	age int
}

func main() {
	var s = Student {
		"golang",
		10,
	}
	unchanged(s, 11)
	fmt.Println(s) // {golang 10}
	change(&s, 12)
	fmt.Println(s) // {golang 12}
}

func unchanged(original Student, value int) {
	original.age = value
}

func change(original *Student, value int) {
	original.age = value
}
```

**PS:** notice that we didn't use `*` in our `change` function ?. in struct we could access the value of our struct directly without using the `*`.

### Nested Struct

struct could be the data type of other struct too, here is the example:

```go
type person struct {
    name string
    age  int
}

type student struct {
    person
    age   int
    grade int
}

// or

type student struct {
    person struct {
        name string
        age  int
    }
    age   int
    grade int
}

```

### Anonymus struct

we could directly use struct in variable instead of making it's type, or we could initialize struct object without given the name of the struct.

```go
var s1 = struct {
    person
    grade int
}{}
```

### tag in struct

tag is an optional property in struct, it could be used for encode / decode data process, usually used in decode and encoding JSON.

```go
type person struct {
    name string `tag1`
    age  int    `tag2`
}
```

### Method in struct

we could attached a method / function into a struct. and there is two way of attaching it, either by reference or value

```go
type student struct {
    name  string
    grade int
}

func (s student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
    fmt.Println("---> on changeName1, name changed to", name)
    // the object name won't change
    s.name = name
}

func (s *student) changeName2(name string) {
    fmt.Println("---> on changeName2, name changed to", name)
    // object name would change
    s.name = name
}
```

## Interface

Interface is type of contract that need to be fullfilled by the object. the interface concept in golang is similiar with any other language, the difference is only it need to be fullfill by struct

here is the example:

```go
type Animal interface {
	move(x int)
	getCoordinate() int
}

type Bird struct {
	coordinate int
}

func (bird *Bird) move(x int) {
	fmt.Println("Fly", x)
	bird.coordinate += x
}

func (bird *Bird) getCoordinate() int {
	return bird.coordinate
}

type Fish struct {
	coordinate int
}

func (fish *Fish) move(x int) {
	fmt.Println("Swin", x)
	fish.coordinate -= x
}

func (fish *Fish) getCoordinate() int {
	return fish.coordinate
}

func (fish *Fish) sayHello() {
	fmt.Println("hello")
}

func main() {
	var bird Animal = &Bird{}
	var fish Animal = &Fish{}
	bird.move(1)
	fmt.Println(bird.getCoordinate())
	fish.move(1)
	fmt.Println(fish.getCoordinate())
	// used only fish method by casting to fish
	fish.(*Fish).sayHello()
}
```

PS: notice that we are using `&` and `*` to access our struct, this happen because we want to create a struct with pointer refence on it.

### Embedded Interface

we could have 2 or more interface merged into 1 interface by embedding it. here is the example.

```go
type Human interface {
    getAge() int
    getName() string
    sayHello()
}

type Worker interface {
    work()
}

type SalaryMan interface {
    Human
    Worker
}

type Sales struct {
	age int
	name string
}

func (sales Sales) getAge() int {
	return sales.age
}

func (sales Sales) getName() string {
	return sales.name
}

func (sales Sales) sayHello() {
	fmt.Println("Hello, my name is", sales.name)
}

func (sales Sales) work() {
	fmt.Println("Gotta work well")
}
func main() {
	var salaryMan SalaryMan = &Sales{}
	fmt.Println(salaryMan)
}
```

### any / empty interface

even thought interface is a contract, what would happen if it's empty ? it become any. here is some example of any and empty interface:

```go
var json map[string]interface{} = make(map[string]interface{})
var jsonAny map[string]any = make(map[string]any)
json["key1"] = 1
json["key2"] = "value"

jsonAny["key1"] = 12
jsonAny["key2"] = "string"
fmt.Println(json, jsonAny) // map[key1:1 key2:value] map[key1:12 key2:string]
```

## Reflect

reflect is a technique to get the variable metadata to advancedly manipulate the variable. reflect in golang itself is a library that have a lot of function for reflect technique usage. more of reflect https://pkg.go.dev/reflect

### Common / Basic Reflect Usage

- Checking data type

  ```go
  var number = 23
  var reflectValue = reflect.ValueOf(number)

  fmt.Println(reflect.TypeOf(number)) // int
  fmt.Println("variable type :", reflectValue.Type()) // variable type : int
  fmt.Println("variable value :", reflectValue.Interface()) // variable value : 23
  ```

- accessing object property

  ```go
  type student struct {
      Name  string
      Grade int
  }
  func (s *student) getPropertyInfo() {
      var reflectValue = reflect.ValueOf(s)

      if reflectValue.Kind() == reflect.Ptr {
          reflectValue = reflectValue.Elem()
      }

      var reflectType = reflectValue.Type()

      for i := 0; i < reflectValue.NumField(); i++ {
          fmt.Println("property name:", reflectType.Field(i).Name)
          fmt.Println("data type :", reflectType.Field(i).Type)
          fmt.Println("value     :", reflectValue.Field(i).Interface())
          fmt.Println("")
      }
  }
  ```

- accessing method from object

  ```go
  func (s *student) SetName(name string) {
      s.Name = name
  }

  func main() {
      var s1 = &student{Name: "john wick", Grade: 2}
      fmt.Println("name :", s1.Name) // john wick

      var reflectValue = reflect.ValueOf(s1)
      var method = reflectValue.MethodByName("SetName")
      method.Call([]reflect.Value{
          reflect.ValueOf("wick"),
      })

      fmt.Println("name :", s1.Name) // wick
  }
  ```
