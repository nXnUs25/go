# Questions 2

---
## Questions
>
1. What do the SOLID acronym initials stand for?
2. Why does the following piece of code violate the SRP? How would you refactor it to make sure it doesn't violate it?
```go 
 import (
     "crypto/ecdsa"
 )
 
 type Document struct { //... }
 
 // Append adds a line to the end of the document.
 func (d *Document) Append(line string) { //...  }
 
 // Content returns the document contents as a string.
 func (d *Document) Content() string { //... }
 
 // Sign calculates a hash for the document contents, signs it with the
 // provided private key and returns back the result.
 func (d *Document) Sign(pk *ecdsa.PrivateKey) (string, error) { //... }
```
3. What is the main concept behind the ISP? Discuss how would you apply it to improve the following function signature:
```go 
 // write a set of lines to a file. 
 func write(lines []string, f *os.File) error {
    //...
 }
```
4. Explain why `util` is considered to be a less-than-ideal name for a Go package.
5. Why are import cycles an issue for Go programs?
6. Name some of the advantages of using the zero value when designing new Go types.”
 
> Excerpt From
> Hands-On Software Engineering with Golang
> Achilleas Anagnostopoulos
> This material may be protected by copyright.
---

## Answers 2

- ### 2.1
  **SOLID**: 
  The SOLID principles are essentially a set of rules for helping you write clean and maintainable object-oriented code.
- **`S`** - **Single responsibility principle** _(SRP)_
_In any well-designed system, objects should only have a single responsibility_
- **`O`** - **Open/closed principle**
_A software module should be open for extension but closed for modification._
- **`L`** - **Liskov substitution principle** _(LSP)_
_If, for each object, O1 of type S there is an object O2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when O1 is substituted for O2, then S is a subtype of T._
- **`I`** - **Interface segregation principle** _(ISP)_
_The bigger the interface, the weaker the abstraction._
- **`D`** - **Dependency inversion principle** _(DIP)_ 
_High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions._

---
- ### 2.2
**SRP** - In a nutshell, object implementations should focus on doing one thing well, and in an efficient way.
On this example we can see that `Sign()` method could be used to not only *Sign* this type of document, but also any other kind of file.
```go
 // Sign calculates a hash for the document contents, signs it with the
 // provided private key and returns back the result.
 func (d *Document) Sign(pk *ecdsa.PrivateKey) (string, error) { //... }
```
and return the hash of the file. As the case may be, this is a valid, working solution. However, the extra coupling that is introduced makes the implementation harder to maintain and extend. For instance, what if we want to evaluate different Document/Text types for object-recognition purposes? What if we want to use the same object hash generation code for different Document/Text files types? We could move this into a separate standalone object such as `Security` `Hashes` and this way we could use this `Sign` method on any other kind of Document/Text object.

---
- ### 2.3
**ISP** - states that no code should be forced to depend on methods it does not use. ISP splits interfaces that are very large into smaller and more specific ones so that clients will only have to know about the methods that are of interest to them. The method which takes `os.File` as a parameter we can change to `io.Writer` type, so it will gave us more flexibility to write data, not only to file but also other types such `sockets`
```go 
 // write a set of lines to a file. 
 func write(lines []string, w io.Writer) error {
    //...
 }
```

---
- ### 2.4
_Go package names should be short and concise and provide a clear indication of their purpose to the intended users of the package._

Package names should be kept short, we should avoid coming up with package names that can potentially clash with variable names that are commonly used by the code importing the package. Otherwise, package users would have to import the package using an alias 
example: `import blah "path-to-package"`. In such cases, it is usually better to abbreviate the package name (if possible). Typical examples from the Go standard library include the `fmt` and `bufio` packages. 
More specifically, the `bufio` package is named as such to avoid name clashes with `buf`, a variable name you are very likely to encounter when dealing with a piece of code that uses buffers.

In other programming languages whose standard libraries usually come with utility libraries or packages with generic-sounding names such as `common` or `util`, Go is quite opinionated against this practice. This is actually justified from the SOLID principles' point of view as those packages are more likely to be violating the SRP versus aptly named packages whose name enforces a logical boundary for their contents. To add to this, as the number of published Go packages grows over time, searching for and locating packages with generic-sounding names will become more and more difficult.

---
- ### 2.5

For a Go program to be well formed, its import graph must be acyclic; in other words, it must not contain any loops. Any violation of this predicate will cause the Go compiler to emit an error. As the systems you are building grow in complexity, so does the probability of eventually hitting the dreaded import cycle detected error.

Usually, import cycles are an indication of a fault in the design of a software solution. Fortunately, in many cases, we can refactor our code and work around most import cycles.

#### don't repeat yourself (DRY)
Example of breaking DRY code rule to fix an issue with import cycle 
```go
package y

import "z"

func IsPrime(v uint64) bool {
    // ... 
}
// Other functions referencing types exported from package z
```

```go
package z

import "x"

// functions referencing types exported from package x”
```

```go
package x

import "y" // circular dependency: x imports y, y imports z and z imports x

func IsValid(v uint64) bool {
    return v != 0 && y.IsPrime(v)
}
```
In cases like this, and assuming the code we need from the included package is small enough, we can just duplicate it (along with its tests) and avoid the extra import that triggers a circular dependency. 

As a popular Go proverb goes:
"_A little copying is better than a little dependency."_
Other way of fixing circular dependencies its when we do breaking circular dependencies via implicit interfaces.

```go
package warehouse

import "context"

// Robot navigates the warehouse floor and fetches items for packing.
type Robot struct {
    // various fields
}

// AcquireRobot blocks until a Robot becomes available or until the 
// context expires.
func AcquireRobot(ctx context.Context) *Robot { //...  }

// Pack instructs the robot to pick up an item from its shelf and place 
// it into a box that will be shipped to the customer.
func (r *Robot) Pack(item *entity.Item, to *entity.Box) error { //...  }
```
```go
package entity

// Box contains a list of items that are shipped to the customer.
type Box struct {
    // various fields
}

// Pack qty items of type i into the box.
func (b *Box) Pack(i *Item, qty int) error {
    robot := warehouse.Acquire() // compile error: import cycle detected
    // ...
}
```
Work around this problem using Go's support for implicit interfaces
```go
package entity
import "context"

// Packer is implemented by objects that can pack an Item into a Box.
type Packer interface {
    Pack(*Item, *Box) error
}

// AcquirePacker returns a Packer instance.
var AcquirePacker func(context.Context) Packer

// Pack qty items of type i into the box.
func (b *Box) Pack(i *Item, qty int) error {
    p := AcquirePacker(context.Background())
    for j := 0; j < qty; j++ {
        if err := p.Pack(i, b); err != nil {
            return err 
        }
    }
    return nil
}

```
And by litter code modifications we can now add 3rd package and via the 3rd package we can do implicit import _warehouse_ and _entity_ packages
```go
package main

import "github.com/achilleasa/logistics/entity"
import "github.com/achilleasa/logistics/warehouse"

func wireComponents() {
    entity.AcquirePacker = func(ctx context.Context) entity.Packer {
        return warehouse.AcquireRobot(ctx)
    }
}
```

---
- ### 2.6 

One great feature that Go offers is that each type is automatically assigned its zero value when it gets instantiated. Some interesting examples from Go and its standard library are as follows:
 
  - `Go channels`; nil channels indefinitely block go-routines attempting to read off them
  - The zero value for a Go `slice`; this is an empty slice that things can be appended to
  - The `sync.Mutex`type, whose zero value indicates that the `mutex` is unlocked
  - The`bytes.Buffer`type, whose zero value indicates an empty buffer

By relying on zero values when designing new types, we can provide implementations that work out of the box without the need to explicitly invoke a constructor or any other initializer method.

---

> Excerpt From
> Hands-On Software Engineering with Golang
> Achilleas Anagnostopoulos
> This material may be protected by copyright.

---

# Answers to Chapter 2 ... by author of the book 

### Chapter 2 
#### 1. 
The following is what the SOLID acronym initials stand for: 
-   Single responsibility 
-   Open/Closed 
-   Liskov substitution 
-   Interface segregation 
-   Dependency inversion 
#### 2. 
The code conflates two responsibilities: retrieving/mutating the state of a document and creating a signature for the document's content. Furthermore, the proposed implementation is inflexible as it forces the use of a specific signing algorithm. To address this problem, we can remove the Sign method from the Document type and provide an external helper that can sign not only instances of Document but also any type that can export its content as a string:
 ```go 
 type ContentProvider interface { 
     Content() string 
 } 
 type ECDADocumentSigner struct {//...} 
 
 func (s ECDADocumentSigner) Sign(pk *ecdsa.PrivateKey, contentProvider ContentProvider) (string, error) { 
     //... 
 } 
 ```
 #### 3. 
 The idea behind the interface segregation principle is to provide clients with the smallest possible interface that satisfies their needs and thus avoid depending on interfaces that will not actually be used. In the provided example, the write method receives an `*os.File` instance as an argument. However, as the function implementation probably only needs to be able to write data to the file, we could achieve the same result by passing an `io.Writer` in the place of the `*os.File` instance. Apart from breaking the dependency to the `*os.File` concrete type, this change will also allow us to reuse the implementation for any type that implements `io.Writer` (for example, sockets, loggers, or others). 
 
 #### 4. 
 The use of `util` as a package name is not a recommended practice due to the following reasons: 
 - It provides little context as to the package's purpose and contents. 
 - It can end up as the home for miscellaneous, possibly unrelated types and/or methods that would undoubtedly violate the single responsibility principle.
#### 5. 
Import cycles cause the Go compiler to emit compile-time errors when you attempt to compile and/or run your code. 
#### 6. 
Some of the advantages of using zero values when defining new Go types are as follows: 
- An explicit constructor is not required as Go will automatically assign the zero value to the fields of an object when it is allocated. 
- The types can be embedded into other types and used out-of-the-box without any further initialization (for example, embedding a `sync.Mutex` into a `struct`).
  
---
