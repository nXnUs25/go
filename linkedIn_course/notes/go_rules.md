
# Golang

## Essential Characteristics of Go

* GO is a compiled, statically typed language.
* The Go tool can run a file without pre-compiling.
* Compiled executables are OS specific
* Applications have a statically linked runtime.
* No external virtual machine is needed.

## Is Go Object Objected ?

### Go has some object oriented features

* you can defain custom interfaces
* Custom types implement one or more interfaces
* Custom types can have member methods
* Custom structs (data structures) can have member fields

### What Go doesn't support

* Type inheritance (no classes)
* Method or operator overloading
* Structured exceptions handling
* Implicit numeric conversions

## Essential Syntax Rules

* Go is case sensitive
* Variable and package names are lower and mixed case
* Exported functions and fields have an initial upper*case character

---

### Go Shell commands

```shell
godoc *http=:6060 *play *index
```

#### Go documentation command

`godoc fmt`

#### Formatting source code

`gofmt badformatting.go` - to display the bad formatting
`gofmt -w badformatting2.go` to correct the formatting of the file

#### Go command

`go run script.go` - execute the script
`go build script.go` - build the executable for that script
`go install` - command to build more complex executables
`go env` - to list Go environment variables

#### To build more advanced application folder structure is needed

```none
- project_name
|- src
 |- application_name
|- bin
|- pkg
```

```shell
mkdir ~/project_name
cd project_name
mkdir {src,bin,pkg}
export GOPATH=$(pwd)
cd src
mkdir application_name
cd application_name
touch code.go
go install
cd ../../bin
./application_name
```

#### Creating packages as function libraries

```none
- project_name
|- src
 |- application_name
  |- scripts.go
 |- pkg_name
  |- pkg_name.go
|- bin
|- pkg
```

Under `pkg_name` create file with the same name, and declare package and functions as example:

```go
package pkg_name

func function1(){}
func function2() {}
```

In the application scripts import that package, example `script.go`

```go
package main
import (
  "fmt"
  "pkg_name"
)

func main() {
  pkg_name.function1()
  pkg_name.function2()
}

```

#### Set variable to generate executable for different OS/Platform

```shell
export GOOS=linux
export GOOS=darwin
export GOOS=windows

go install
```

---

## Golang Declarations

### Explicitly Typed Declaration

* Use `var` keyword and `=` assignment operator
* Setting internal value its optional:
  * var variableName type = something (= something is optional)
  * `var anInteger int = 42`
  * if value not set the default value for variable above would be 0
  * `var aString string = "This is String"`

### Implicitly Typed Declaration

* Use `:=` assignment operation without `var` keyword
* `anInteger := 42`
* Type is inferred from initial value
* `aString := "This is a String"`
* Type is still static, and can't be changed.

### Constants

* A constants is a simple, unchanging value
* you can use Explicit or Implicit typing
* Explicit typing
  * `const anInteger int = 34`
* Implicit typing
  * `const anInteger = 34`

### Pre-declared Boolean and String Types

* Boolean values `bool`
* String type `string`
* all identifiers are spelt with lower case
* fixed integer types declare unsigned or signed values
  * `uint8 uint16 uint32 uint32 uint64`
  * `int8 int16 int32 int64`
* Aliases
  * `byte uint int uintptr`

### Other Pre-declared Numeric Types

* Floating values
  * `float32 float64`
* Complex numbers
  * `complex64 complex128`

### Pre-declared Complex Types

* Data collections
  * `Arrays Slices Maps Structs`
* Language organization
  * `Functions Interfaces Channels`
* Data management
  * `Pointers`

### Arithmetic Operations

* Go supports all math operators used in C:
  * `+` Sum
  * `-` Difference
  * `*` Product
  * `/` Quotient
  * `%` Remainder
  * `&` Bitwise *AND*
  * `|` Bitwise *OR*
  * `^` Bitwise *XOR*
  * `&^` Bit clear
  * `<<` Left shift
  * `>>` Right shift

---

### Memory Is Managed by the Runtime

* The Go runtime is statically linked into the application
* Memory is allocated and deallocated automatically
* Use `make()` or `new()` to initialize
  * Maps, slices, channels

#### Memory Allocation

* the `new()` function
  * Allocates but does not initialize memory
  * Results in zeroed storage but returns a memory address
* the `make()` function
  * Allocates and initializes memory
  * Allocates non-zeroed memory and returns a memory address

#### Creating nil Objects

* You must initialize complex objects before adding values
* Declarations without `make()` can cause a panic
* Correct Memory initialization
  * Use `make()` to allocate and initialize memory

```go
m := make(map[string]int)
m["key"] = 42
fmt.Println(m)
```

#### Memory Deallocation

* Memory is deallocated by garbage collector (**GC**)
* Objects out of scope or set to nil are eligible
* GC was rebuilt in Go 1.5 for very low latency
  * <https://golang.org/pkg/runtime/>
  * <https://talks.golang.org/2015/go-gc.pdf>

---

### Loops

* `for` loop
* `for` loop is used as `while` and for this u can drop semicolon `;`
* Example of `while` loop in Go

```go
 sum := 1
 for sum < 100 {
  sum += sum
 }
 ```

 ---

## Conclusion of what next

### Concurrency in Go

* **Goroutine**
  * A lightweight synchronized thread managed by the runtime
* **Channel**
  * A typed conduit for messages between goroutines
* **Select**
  * Lets a goroutine wait for multiple communication operations

### REST API

* Beego: <http://beego.me>
* Martini: <http://martini.codegangsta.io/>
* Gorilla: <http://www.gorillatoolkit.org/>
* Gocraft: <https://github.com/gocraft/web>
* Revel: <https://revel.github.io/>

### Databases

* Standard Database functions: `database/sql`
* Driver interfaces: `database/sql/driver`
* Available for relational and NoSQL databases
* Many community driven
  * <https://github.com/golang/go/wiki/SQLDrivers>
