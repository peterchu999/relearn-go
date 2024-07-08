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

## initialize map

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

## check if key exists

to check if the spesific key exists in the map we could access the map this way :

```go
var memo = map[string]int{}
memo["laundry"] = 1
memo["cook"] = 2

var value, isExists = memo["laundry"]
fmt.Println(value) // 1
fmt.Println(isExists) // true
```

## delete item in array

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
