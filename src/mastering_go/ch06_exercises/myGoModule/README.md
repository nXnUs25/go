
## To initialize package run the command below

---
```go
go mod init github.com/nXnUs25/myGoModule
```
---
That should generate file called `go.mod`

---
```markdown
module github.com/nXnUs25/go-mypackage

go 1.14
```
---

Add some `go` package fiels as example 
- myPackage.go

---
```go
package myGoModule

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
---
- samePackage.go

---
```go
package myGoModule

const version = "1.0.0"

func OtherFileInSamePackage() string {
	return "Its other file location for this file"
}

func privatePackage(name string) string {
	return name + " " + name
}
```
---

Once files are created to use this package locally, run in the root directory of the package
`go install`

To test and verify the package before pushing it to the GitHub repository you will need to execute the command `go run`
and point to the package main which does import the package

Structure:
---
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

        my "github.com/nXnUs25/myGoModule"
)

func main() {
        fmt.Println(my.PackageName())
        fmt.Println(my.OtherFileInSamePackage())
        my.CallPrivate("nonus25")
}
```
---

Execution:
---
```shell
❯ go run ../callmyPackage/myPackage-v1.go
myPackage
Its other file location for this file
nonus25 nonus25
```
---

Creating Package version 1.0.0
We need to create initial package version commit by simply sdding files and pushing it into Git repo
---
```shell
git add .
git commit -m "Initial commit ver 1.0.0"
git push
```
---

Then we will need to run the following Git commands to mark this as a package version
---
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
---
Listing the module in go
---
```shell
❯ go list -m all
github.com/nXnUs25/myPackage
❯ go list -m -versions
github.com/nXnUs25/myPackage
```
---
adding extra function and changing the version to v1.1.0
---
```shell
❯ git add .
❯ git commit -m "added extra func"
[main ee22917] added extra func
 3 files changed, 128 insertions(+)
❯ git push
Warning: Permanently added 'github.com' (ED25519) to the list of known hosts.
Enumerating objects: 17, done.
Counting objects: 100% (17/17), done.
Delta compression using up to 8 threads
Compressing objects: 100% (8/8), done.
Writing objects: 100% (9/9), 1.85 KiB | 1.85 MiB/s, done.
Total 9 (delta 5), reused 0 (delta 0)
remote: Resolving deltas: 100% (5/5), completed with 5 local objects.
remote: This repository moved. Please use the new location:
remote:   git@github.com:nXnUs25/go.git
To github.com:augustync/go.git
   6703408..ee22917  main -> main
❯ git tag v1.1.0
❯ git push origin v1.1.0
Warning: Permanently added 'github.com' (ED25519) to the list of known hosts.
Total 0 (delta 0), reused 0 (delta 0)
remote: This repository moved. Please use the new location:
remote:   git@github.com:nXnUs25/go.git
To github.com:augustync/go.git
 * [new tag]         v1.1.0 -> v1.1.0
```
---


to create new release of majer version create new branch, and for the new release import 
and cleanup `go.sum` file generated before by go run if present
in the `go.mod` file edit the version of requirements to point to right version in the program which is going to use this module.
also when changing release version in your project is good to clear the old version cache by running :
`go clean -modcache`
```go
module callmyModule-v2

go 1.14

require github.com/nXnUs25/myGoModule/v2 v2.0.2 // indirect
```

so one all done from above lets create a new branch:
```shell
git checkout -b v2
git push --set-upstream origin v2
```

that will create a new branch and will switch as into 
now we should push changes to out module new release by changing whatever is needed
and then:
```shell
git commit -am "new changes for release v2"
git tag v2.0.0
git push --tags origin v2
git --no-pager branch -a
```

git tree:
```
> glog
* 4ab2f2f (HEAD -> v2, tag: v2.0.2, origin/v2) fixinig small bug with version numbers
* b2cc9e6 (tag: v2.0.1) fixinig small bug
* c69f511 (tag: v2.0.0) Initial ver 2.0.0
* a25eaa0 (tag: v1.1.0, origin/main, origin/HEAD, main) Initial ver 1.1.0
* 1f650e4 (tag: v1.0.0) Initial ver 1.0.0
* d2e28fa Initial commit

❯ git --no-pager branch -a
  main
* v2
  remotes/origin/HEAD -> origin/main
  remotes/origin/main
  remotes/origin/v2
```




---
```go

```
---
