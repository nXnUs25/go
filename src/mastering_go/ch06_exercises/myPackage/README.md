
## To initialize package run the command below

---
```go
go mod init github.com/nXnUs25/myPackage
```
---

That should generate file called `go.mod`

```
module github.com/nXnUs25/myPackage

go 1.14
```

Add some `go` package fiels as example 
- myPackage.go

```go
package myPackage

import "fmt"

func PackageName() string {
	return "myPackage"
}

func CallPrivate(name string) {
	fmt.Println(privatePackage(name))
}

func Version() {
	fmt.Println("Version: ", version)
}
```
- samePackage.go

```go
package myPackage

const version = "1.0.0"

func OtherFileInSamePackage() string {
	return "Its other file location for this file"
}

func privatePackage(name string) string {
	return name + " " + name
}
```

Once files are created to use this package locally, run in the root directory of the package
`go install`

To test and verify the package before pushing it to the GitHub repository you will need to execute the command `go run`
and point to the package main which does import the package

Structure:
```shell
❯ ll
total 0
drwxr-xr-x  3 nonus25  staff   96  2 Apr 16:10 callmyPackage
drwxrwxr-x  6 nonus25  staff  192  2 Apr 16:07 myPackage
❯ cd myPackage
❯ ll
total 32
-rw-rw-r--  1 nonus25  staff  1293  2 Apr 16:47 README.md
-rw-rw-r--  1 nonus25  staff    45  2 Apr 15:59 go.mod
-rw-r--r--  1 nonus25  staff   208  2 Apr 16:22 myPackage.go
-rw-r--r--  1 nonus25  staff   204  2 Apr 16:21 samePackage.go
❯ ll ../callmyPackage
total 8
-rw-r--r--  1 nonus25  staff  185  2 Apr 16:18 myPackage-v1.go
❯ cat ../callmyPackage/myPackage-v1.go
package main

import (
        "fmt"

        my "github.com/nXnUs25/myPackage"
)

func main() {
        fmt.Println(my.PackageName())
        fmt.Println(my.OtherFileInSamePackage())
        my.CallPrivate("nonus25")
}
```

Execution:
```shell
❯ go run ../callmyPackage/myPackage-v1.go
myPackage
Its other file location for this file
nonus25 nonus25
```

Creating Package version 1.0.0
We need to create initial package version commit by simply sdding files and pushing it into Git repo
```shell
git add .
git commit -m "Initial commit ver 1.0.0"
git push
```

Then we will need to run the following Git commands to mark this as a package version
```shell
❯ git tag v1.0.0
❯ git push origin v1.0.0
Warning: Permanently added 'github.com' (ED25519) to the list of known hosts.
Total 0 (delta 0), reused 0 (delta 0)
remote: This repository moved. Please use the new location:
remote:   git@github.com:nXnUs25/go.git
To github.com:augustync/go.git
 * [new tag]         v1.0.0 -> v1.0.0
```
Listing the module in go
```shell
❯ go list -m all
github.com/nXnUs25/myPackage
❯ go list -m -versions
github.com/nXnUs25/myPackage
```



---
```go

```
---
