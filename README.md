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

