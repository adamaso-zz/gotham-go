# Impact of Community

### Maximizing Impact

* Contribute
    * Continue ot use Go
    * Start issues
    * Create experience reports
* Become a leader
    * Start meetup
    * WWG or GoBridge
    * Start a hack or pair programming night
* Support others
    * Join Gopher SLack, Reddit, Quora, and help answer questions
    * Create a tutorial, blog, etc
    * Teach Go at your company
    * Ask "How can I help?"

>"Very great change starts from very small conversations, helped among people who care" - Margaret J Wheatley

---

# Building Gophercises

Software gets overly complex
> Your web app should...
    * JSON API w/ Microservices
    * Docker and Kubernetes
    * Never use frameworks!
     ...list goes on and on

* Content Management System
* Non traditional authenticaltion
    * Email token to login
    * i.e. http://gophercises.com/login?token=some-token

### Rendering HTML Pages

Server side rendering via `html/template` package

Code organization:
* Smalller, component like templates
* Decorator pattern
* Service objects

Keep things simple and isolate the view

### Decorator Pattern

```go
package html

type Exercise sruct {
    app.Exercise
}

func (e Exercise) Duration() string {
    //code here
}
```

^ html template allows you to use logic in the view

### Key Takeaway

You can build simpler software in Go without getting bogged down by the thigns you are "supposed" to do

calhoun.io
gophercises.com

---

# From JS to Go

**What is Go useful for?**
* Static Binaries
* Backend Services
* gRPC
* SQL DBs
* Concurrency (goroutines)
* Machine Code (no VM)

**What does Go leave out?**
* No Classes
* No Inheritance
* No Contructors
* Exceptions

