# GothamGo 2018 Notes

**GOROOT** - where go is set up
**GOPATH** - where projects live

[Link to Course](https://www.gopherguides.com/)

* More limitations with Go = Simpler
* Easy to read

---

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

