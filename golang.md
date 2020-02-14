# `Golang 101`

Martin J. Schreiner

---

## `Overview`

0. My experience with Go
1. What is Go about?
2. Installing Go
3. Text editors and IDEs
4. Structure of Go code
5. Hello, world!
6. Typing system
7. Structs and methods
8. Interfaces and polymorphism
9. Closures and higher-order functions
10. Goroutines and channels

---

## `What is Go about?`

Go is a fairly small, multi-platform, statically typed, compiled, garbage collected programming language.

It's ABI-compatible with C and provides neat concurrency primitives.

Mostly meant for service development, but works well for mostly all desktop or server software.

---

## `Installing Go`

On Linux, you may use your package manager.

`sudo apt install golang`

`sudo snap install go --classic`

`sudo dnf install golang`

`sudo zypper in go1.13`

On Windows or macOS, visit https://golang.org/dl

---

## `Text editors and IDEs`

I love and use Emacs for Go. It kicks arse.

VSCode works really well too.

Haven't tried Vim, but should work beautifully with some plugins.

There's also Goland: https://www.jetbrains.com/go

---

## `Structure of Go code`

> Go programmes are organised into _packages_. A package is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.

Source: https://golang.org/doc/code.html

---

## `Structure of Go code`

```text
src/foo/
src/foo/const.go
src/foo/lib.go
src/foo/types.go
src/mayhem/
src/mayhem/lib.go
src/mayhem/main.go
src/mayhem/types.go
src/mayhem/oauth/
src/mayhem/oauth/oauth.go
```

`mayhem` is our main project (from Fight Club).

`foo` is a library we pulled from GitHub.

---

## `Structure of Go code` - `src/foo/const.go`

Package is named after the directory.

```go
package foo

// untyped constant, public
const Version = 1

// typed constant, private
const baseUrl string = "http://www.foo.org/api/v1"
```

---

## `Structure of Go code` - `src/foo/lib.go`

Package is named after the directory.

```go
package foo

import (
    "fmt"
    "net.http"
)

// notice that baseUrl belongs to the same namespace
// this function is private
func url(path string) string {
    return fmt.Sprintf("%s/%s", baseUrl, path)
}

// this function is public
func Status() Json, error {
    if response, err := http.Get(url("status")); err != nil {
        return response.Json(), nil
    } else {
        return nil, err
    }
}
```

---

## `Structure of Go code` - `src/mayhem/types.go`

Directory changed, and so did the package's name.

```go
package mayhem

type Character struct {
    Name      string
    Inventory []Item
}

type Item struct {
    Name   string
    Weight int
}
```

---

## `Structure of Go code` - `src/mayhem/oauth/oauth.go`

Subpackages are named after their own subdirectories.

```go
package oauth

func Authenticate(token string) Session, err {
    // ...
}
```

---

## `Structure of Go code` - `src/mayhem/main.go`

Notice this is the `main` package.

```go
package main

import (
    "foo"
    "mayhem"
    "mayhem/oauth"
)

func main() {
    if status, err := foo.Status(); err != nil {
        // ...
    }
    if session, err := oauth.Authenticate(oauth.Token()); err != nil {
        // ...
    }

    character := mayhem.Character{Name: "John Smith", Inventory: make([]Item, 0)}
    // ...
}
```

---

## `Hello, world!`

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello, world!")
}
```

---

## `Typing system`

Go supports the following types:

- `string`.
- `bool`.

Numeric types:
- `byte`, `int16`, `int64`, `int8 (uint8)`, `int`, `rune`, `uint16 (int32)`, `uint32`, `uint64`, `uint`, `uintptr`.
- `float32`, `float64`.
- `complex64`, `complex128`.

Note that unsigned types aren't defined via a keyword, but through the type's identifier itself.

--- 

## `Typing system`

##### Variable declaration and assignment

Variables may be declared and then have their values assigned at a later time.

```go
package main

func main() {
    var name string
    
    name = "John Smith"
    
    fmt.Printf("My name is: %s.\n", name)
}
```

---

## `Typing system` - `Type inference`

Types may also be inferred.

```go
package main

import "fmt"

func main() {
    name := "John Smith"
    
    fmt.Printf("My name is: %s.\n", name)

    // invalid assignment, won't compile, type mismatch
    name = 3
}
```

---

## `Typing system` - `Arrays`

Unlike in C, arrays are values, and we can't pass them around by leveraging their starting memory addresses as pointers.

They are continuous areas of memory the compiler knows the boundaries and type of.

```go
package main

import "fmt"

func main() {
    // size inference syntax: [...]
    var names [3]string = [...]string{"John Smith", "Paul Murray", "Linus Pauling"}
    
    for _, name := range(names) {
        fmt.Printf("There once lived a fellow, his name was: %s.\n", name)
    }
}
```

---

## `Typing system` - `Slices`

Unlike arrays, slices are dynamically allocated blocks of memory having three main attributes: a **type**, a **length** and a **capacity**.

Slices grow as necessary, and are supported by a bunch of standard library functions.

---

## `Typing system` - `Slices`

```go
package main

import "fmt"

func main() {
    // array
    var names [3]string = [...]string{"John Smith", "Paul Murray", "Linus Pauling"}

    var sliceNames []string
    sliceNames = make([]string, 0, len(names))

    for _, name := range(names) {
        sliceNames = append(sliceNames, name)
    }

    for _, name := range(sliceNames) {
        fmt.Printf("There one lived a fellow, he lived in a slice, his name was: %s.\n", name)
    }
}
```

---

## `Typing system` - `Maps`

```go
package main

import "fmt"

type Vertex struct { Lat, Long float64 }

func main() {
    var m map[string]Vertex
    m = make(map[string]Vertex)

    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }

    fmt.Println(m["Bell Labs"])
}
```

--- 

## `Typing system` - `Pointers`

```go
package main

import "fmt"

func main() {
    // array
    var names [3]string = [...]string{"John Smith", "Paul Murray", "Linus Pauling"}

    var sliceNames *[]string = new([]string)

    for _, name := range(names) {
        *sliceNames = append(*sliceNames, name)
    }

    for _, name := range(*sliceNames) {
        fmt.Printf("There one lived a fellow, he lived in a slice, his name was: %s.\n", name)
    }
}
```

--- 

## `Typing system` - `Pointers`

```go
package main

import "fmt"

func newSlice() *[]string {
    var s []string = make([]string, 0)
    return &s
}

func main() {
    var names [3]string = [...]string{"John Smith", "Paul Murray", "Linus Pauling"}

    var sliceNames *[]string = newSlice()

    for _, name := range(names) {
        *sliceNames = append(*sliceNames, name)
    }

    for _, name := range(*sliceNames) {
        fmt.Printf("There one lived a fellow, he lived in a slice, his name was: %s.\n", name)
    }
}
```

--- 

## `Structs and methods`

```go
package main

import "fmt"

type Point struct {
    X, Y uint
}

func (p Point) GetX() uint {
    return p.X
}

func (p Point) Display() {
    fmt.Printf("X: %d Y: %d\n", p.X, p.Y)
}

func main() {
    point := Point{X: 12, Y: 27}
    point.Display()
}
```

---

## `Structs and methods` - `Type embeddings`

```go
package main

import "fmt"

type Animal struct {
    Name string
}

type Dog struct {
    Class Animal
    Breed string
}

func (a Animal) SayName() {
    fmt.Println("%s", a.Name)
}

func (d Dog) SayBreed() {
    fmt.Println("%s", d.Breed)
}

func main() {
    d := Dog{Class: Animal{Name: "Billy"}, Breed: "Labrador"}
    d.SayName()
}
```

---

## `Structs and methods` - `Immutability`

```go
package main

import "fmt"

type Point struct {
    X, Y uint
}

// actually modifying copy
func (p Point) SetX(x uint) {
    p.X = x
}

func (p Point) Display() {
    fmt.Printf("X: %d Y: %d\n", p.X, p.Y)
}

func main() {
    point := Point{X: 12, Y: 27}
    point.Display()
    point.SetX(18)
    point.Display()
}
```

---

## `Structs and methods` - `Immutability`

```go
package main

import "fmt"

type Point struct {
    X, Y uint
}

// pointer to struct allows us to modify it
func (p *Point) SetX(x uint) {
    // implicit dereference, same as: (*p).X = x
    p.X = x
}

// still works, implicit dereference
func (p Point) Display() {
    fmt.Printf("X: %d Y: %d\n", p.X, p.Y)
}

func main() {
    point := &Point{X: 12, Y: 27}
    point.Display()
    point.SetX(18)
    point.Display()
}
```

---

## `Interfaces and polymorphism`

```go
package main

import "fmt"

type Greetable interface {
    Greet()
    AskName() string
}

type Person struct {Name string}
type Crow struct {Name string}

func (p Person) Greet() {
    fmt.Printf("Hello! I am %s!\n", p.Name)
}

func (p Person) AskName() string {
    return p.Name
}

func (c Crow) Greet() {
    fmt.Printf("Coo coo! Greetings, I'm %s!\n", c.Name)
}

func (c Crow) AskName() string {
    return c.Name
}

func main() {
    var g Greetable

    g = Person{Name: "Martin"}
    g.Greet()

    g = Crow{Name: "Barry"}
    g.Greet()
}
```

---

## `Closures and higher-order functions`

```go
package main

import "fmt"

func Counter() func() uint {
    current := uint(0)
    return func() uint {
        current++
        return current
    }
}

func main() {
    counter := Counter()
    fmt.Println(counter())
    fmt.Println(counter())
    fmt.Println(counter())
}
```

---

## `Closures and higher-order functions`

```go
package main

import "fmt"

func Counter(callback func(uint)) func() uint {
    current := uint(0)
    return func() uint {
        callback(current)
        current++
        return current
    }
}

func main() {
    f := func(c uint) {
        fmt.Printf("Current state: %d.\n", c)
    }
    counter := Counter(f)
    fmt.Println(counter())
    fmt.Println(counter())
    fmt.Println(counter())
}
```
