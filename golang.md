# `Golang 101`

Martin J. Schreiner

---

## `Overview`

0. What is Go about?
1. Installing Go
2. Text editors and IDEs
3. Structure of Go code
4. Hello, world!
5. Typing system
6. Structs and methods
7. Interfaces and polymorphism
8. Closures

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

## `Typing system`

##### Type inference

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

## `Typing system`

##### Arrays

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

## `Typing system`

##### Slices

Unlike arrays, slices are dynamically allocated blocks of memory having three main attributes: a **type**, a **length** and a **capacity**.

Slices grow as necessary, and are supported by a bunch of standard library functions.

Of course, the casual `realloc` will be called under the bonnet when it increases beyond its capacity.

---

## `Typing system`

##### Slices

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

## `Structs and methods`
