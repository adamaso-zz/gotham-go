# GothamGo 2018 Notes

[Link to Course](https://www.gopherguides.com/)

* More limitations with Go = Simpler
* Easy to read

## Syntax and Types

* string
    * immutable and can not be modified
    * `a := "Say hello to Go!"`
* bool

### Declaring Variables

**Without Initialization**
`var a int`

**With initialization, explicit type**
`var a int = 1`

**With initialization, implicit type**
`a := 1`

All builtin types have a zero value. Any allocated variable is usable even if it never has a value assigned.

### Iota
Go's `iota` identifier is used in const declarations to simplify definitions of incrementing numbers. Because it can be used in expressions, it provides a generality beyond that of simple enumerations.

**Input**
```go
package main

import "fmt"

const (
	Apple int = iota
	Orange
	Banana
)

func main() {
	fmt.Printf("The value of Apple is %v\n", Apple)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Banana is %v\n", Banana)
}
```

**Output**
```go
The value of Apple is 0
The value of Orange is 1
The value of Banana is 2
```

### Structs

A struct is a collection of fields, often called members (or attributes).

Structs are used so we can create our own custom complex types in Go.

When trying to understand structs, it helps to think of them as a blueprint for what your new type will do.

A struct definition, does NOT contain any data.

**Structs may have zero or more fields**
```go
type User struct {
	Name  string
	Email string
}
```

**Fields can be referenced using a period and the field name.**
```go
func main() {
	u := User{}
	u.Name
	u.Email
}
```

#### Initializing Structs

```go
u := User{
	Name:  "Homer Simpson",
	Email: "homer@example.com",
}
```
**You need a comma after values ^ ^**

You can set as many of the field values on a struct at initialization time as you want.
```go
u := User{Email: "marge@example.com"}
u.Name = "Marge Simpson"
```

## Arrays and Iteration

* Fixed length
* Fixed type
* Zero based

The capacity of an array is defined at creation time. Once an array is allocated it's size can no longer be changed.

### Defining an array

```go
names := [4]string{}
names[0] = "John"
names[1] = "Paul"
names[2] = "George"
names[3] = "Ringo"
```

### Initializing an array

```go
package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}

	fmt.Println(names)
}
```

## Array Type
The length is actually part of the type that is defined for arrays.

```go
package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}

	a2 = a1

	fmt.Println(a2)
	a3 := [3]string{}

	// This can't be done, as it is not of the same type
	//a3 = a2

	fmt.Println(a3)
}
```
## 2-D Arrays (Matrix)

Go's arrays are one-dimensional. To create an equivalent of a 2D array, it is necessary to define an array-of-arrays.



```go
type Matrix [3][3]int

func main() {
	m := Matrix{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 3},
    }
```

### The `for` Loop
Go only has one loop! Use it for `for`, `while`, `do while`, `do until`, etc...

```go
for i := 0; i < N; i++ {
  // do work until i equals N
}
```

```go
for {
  // loop forever
}
```

### Iterating over arrays with `len`

```go
package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
```

**NOTE:** Explicit array length MUST be declared, it cannot infer any type data.

### `range` Keyword

Looping over arrays, and other collection types, is so common that Go created the range tool to simplify this code.

```go
package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}

	for i, n := range names {
		fmt.Printf("%d - %s\n", i, n)
	}
}
```

Range returns the the index and the value of the array.

