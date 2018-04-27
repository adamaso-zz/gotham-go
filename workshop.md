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

## Slices

Slices wrap arrays in Go, and provide a more general, powerful, and convenient interface to data sequences.

* Similar to Arrays
* Fixed type
* Dynamically sized
* Flexible

Unless you know that your list will contain a finite and fixed set of elements, you'll almost always use a slice when dealing with data.

```go
package main

import "fmt"

func main() {
	// create an array of names
	namesArray := [4]string{"John", "Paul", "George", "Ringo"}
	// create a slice of names
	namesSlice := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println(namesArray)
	fmt.Println(namesSlice)
}
```

Think of a slice as having three members:

* Length
* Capacity
* Pointer to the underlying array

```go
type slice struct {
	Length   int
	Capacity int
	Array    [10]array
}
```

### Appeding to slices

The `append` keyword allows us to add elements to a slice.

```go
package main

import "fmt"

func main() {
	names := []string{}
	names = append(names, "John")

	// Append multiple items at once
	names = append(names, "Sally", "George")

	// Append an entire slice to another slice
	moreNames := []string{"Bill", "Ginger", "Wilma"}
	names = append(names, moreNames...)

	fmt.Println(names)
}
```

## Len & Cap

* The `len` keyword tells us how many elements the slice actually has.

* The `cap` keyword tells us the capacity of the slice, or how many elements it can have.

```go
names := []string{}
fmt.Println("len:", len(names)) // 0
fmt.Println("cap:", cap(names)) // 0

names = append(names, "John")
fmt.Println("len:", len(names)) // 1
fmt.Println("cap:", cap(names)) // 1

names = append(names, "Paul")
fmt.Println("len:", len(names)) // 2
fmt.Println("cap:", cap(names)) // 2

names = append(names, "George")
fmt.Println("len:", len(names)) // 3
fmt.Println("cap:", cap(names)) // 4

names = append(names, "Ringo")
fmt.Println("len:", len(names)) // 4
fmt.Println("cap:", cap(names)) // 4

names = append(names, "Stu")
fmt.Println("len:", len(names)) // 5
fmt.Println("cap:", cap(names)) // 8
```

### Making a slice

Slices can also be created using the `make` keyword which allows us to define the starting "length" of the slice, and optionally, the starting "capacity" of the slice.


```go
package main

import "fmt"

func main() {
	a := make([]int, 1, 3)

	fmt.Println(a)      // [0]
	fmt.Println(len(a)) // 1
	fmt.Println(cap(a)) // 3
}
```

### 2-D Slices

* Go's slices are one-dimensional. To create an equivalent of a 2D slice, it is necessary to define an slice-of-slices.

* Because slices are variable-length, it is possible to have each inner slice be a different length.

```go
// a slice of byte slices
type Modules [][]byte

course := Modules{
	[]byte("Chapter One: Syntax"),
	[]byte("Chapter Two: Arrays and Slices"),
	[]byte("Chapter Three: Maps"),
}
```

### Slice Subsets

Subsets of a slice (or a slice of a slice) allow us to work with just section of a slice.

```go
package main

import "fmt"

// slice[starting_index : (starting_index + length)]

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println(names)      // [John Paul George Ringo]
	fmt.Println(names[1:3]) // [Paul George] - names[1:1+2]

	// functionally equivalent
	fmt.Println(names[2:])           // [George Ringo]
	fmt.Println(names[2:len(names)]) // [George Ringo]

	// functionally equivalent
	fmt.Println(names[:2])  // [John Paul]
	fmt.Println(names[0:2]) // [John Paul]
}
```

### Mutating Slice Subsets

```go
names := []string{"John", "Paul", "George", "Ringo"}

fmt.Println(names) // [John Paul George Ringo]

guitars := names[:3]

fmt.Println(guitars) // [John Paul George]

for i, g := range guitars {
	guitars[i] = strings.ToUpper(g)
}

fmt.Println(names) // [JOHN PAUL GEORGE Ringo]
```

Let's not do this!

### Using `copy` to not mutate data

```go
package main

import "fmt"

func main() {
	veggies := []string{"carrot", "potato", "cucumber", "onion"}
	v := veggies

	v2 := make([]string, len(veggies))
	copy(v2, veggies)

	v[0] = "zuchinni"

	fmt.Println(veggies)
	fmt.Println(v)
	fmt.Println(v2)
}
```

```go
// copy(source, destination)
```

[more slice tricks + tips](https://github.com/golang/go/wiki/SliceTricks)

## Maps

An unordered set of values indexed by a unique key

```go
package main

import "fmt"

func main() {
	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	fmt.Println(beatles)
}

// Prints:
// map[Ringo:drums John:guitar Paul:bass George:guitar]
```

### Initializing Maps

```go
package main

import "fmt"

func main() {
	beatles := map[string]string{
		"John":   "guitar",
		"Paul":   "bass",
		"George": "guitar",
		"Ringo":  "drums",
	}

	fmt.Println(beatles)
}
```

The len function can be used to find the length (the number of keys) in the map.

`fmt.Println(len(beatles))`

### Deleting Map Keys

```go
package main

import "fmt"

func main() {
	beatles := map[string]string{
		"John":   "guitar",
		"Paul":   "bass",
		"George": "guitar",
		"Ringo":  "drums",
	}

	delete(beatles, "John")

	fmt.Println(beatles)
	// map[Paul:bass George:guitar Ringo:drums]
}
```

### Map Values

Map values can be set and retrieved using the the [] syntax.

```go
package main

import "fmt"

func main() {
	beatles := map[string]string{}

	beatles["Paul"] = "bass"

	paul := beatles["Paul"]
	fmt.Println(paul) // bass
}
```

Maps in Go return an optional second argument that will tell you if the key exists in the map.

```go
if paul, ok := beatles["Paul"]; ok {
	fmt.Println(paul) // bass
}

// If the value is paul, and ok equals Paul, do this
```

### Getting Map Keys

Go does not provide a way to get just a list of keys, or values, from a map. To do this you must build that list yourself.


```go
package main

import "fmt"

func main() {
	m := map[string]string{}
	m["foo"] = "bar"
	m["bin"] = "baz"
	m["boo"] = "woo"

	keys := make([]string, len(m))

	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	fmt.Printf("%+v", keys)
}
```

## Pointers

A pointer is a type that holds the address that points to a variables value.

### How does Go pass data (default)
```go
package main

import "fmt"

type Beatle struct {
	Name string
}

func main() {
	b := Beatle{Name: "Ringo"}
	changeBeatleName(b)
	fmt.Println(b.Name) // Ringo
}

func changeBeatleName(b Beatle) {
	b.Name = "George"
	fmt.Println(b.Name) // George
}
```
The changeBeatleName function is receiving a copy of the Beatle value.

```go
func changeBeatleName(b Beatle) {
	b.Name = "George"
	fmt.Println(b.Name) // George
}
```
 
It can not modify the original value.

### Pointer Syntax

Pointers allow us to point at the original memory space of a value.

To indicate we are expecting a pointer, we use the * modifier:

`func changeBeatleName(b *Beatle) {}`

To get the address of a value we use the & modifier:

`&Beatle{}`

And we can store pointers too:

`b := &Beatle{}`

### When to use pointers?

Use pointers if:

* You want others to be able to modify your values or
* You have a large (memory) value that you don't want to keep copying


## Interfaces

* Interfaces allow us to specify behavior. They are about doing, not being.

* Unlike other languages interfaces in Go are implicitly not explicitly defined.

* Interfaces are found all around the standard library.

### Concrete Types Vs. Interfaces

Imagine you own a venue, such as a night club. You want to have music at the venue. Do you require that the entertainer be a Beatle or would any musician do?

That is the difference between a concrete type and an interface.

A Beatle is a concrete type. You are either a Beatle or you are not. There are only two people alive that can fulfill that requirement.

A musician, however, is anyone who can play an instrument.

### Defining an Interface

Interfaces are created by defining an interface type that has a list of N methods that must be satisfied to implement that interface.

```go
type Writer interface {
	Write(p []byte) (int, error)
}
```

It is important to note that interfaces are a collection of methods, not fields.

```go
// good
type Writer interface {
	Write(p []byte) (int, error)
}

// bad
type Emailer interface {
	Email string
}
```

> The io.Writer interface in the standard library requires the implementation of a Write method that matches the signature of Write(p []byte) (n int, err error).

## Errors

### Gracefully Handle Errors

```go
var f io.Reader
var err error

// try to read a file
f, err = os.Open("/path/to/some/content.file")
if err != nil {
	// create a fall back io.Reader so our program works
	f = bytes.NewBufferString("some fall back content")
}

b, err := ioutil.ReadAll(f)
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(b))
```

### Creating Errors

There are two "built-in" ways to create errors in the standard libary.

```go
// in the `errors` package
errors.New("a message")

// in the `fmt` package
fmt.Errorf("a %s message", "formatted")
```

### Returning Errors

Errors are returned from functions just like any other type.

```go
func boom() error {
	return errors.New("boom!")
}

func greetOrBoom() (string, error) {
	return "hello", errors.New("boom!")
}
```

## Concurrency

### Goroutines

- A goroutine is a lightweight "thread" managed by the Go runtime.
- A goroutine is a function capable of running concurrently with other functions.
- Goroutines are not threads. They are small, lightweight resources that require only a few bytes of memory.
- It's not uncommon to have hundreds, if not thousands, of goroutines.
- In contrast to system threads, which have a fixed stack size and therefore a finite limit to the number of threads you can run, you can run millions of goroutines

---

> Concurrency is about dealing with a lot of things at once.
Parallelism is about doing a lot of things at once.

```go
package main

import "fmt"

func main() {
	go sayHello()
}

func sayHello() {
	fmt.Println("hello")
}
```

Goroutines execute concurrently. In this case, main() and sayHello() were running at the same time. main() finished and sayHello() never got a chance to run!

To solve this problem you need to use some sort of communication, or synchronization device, to signal that a goroutine is complete.

We typically do that with channels, or waitgroups, depending on whether you need to return a value from the function you're running.

### Waitgroups

WaitGroups are defined in the `sync` package in the standard library. They're a simple, and powerful, way to ensure that all of the goroutines you launch have completed before you move on to do other things.