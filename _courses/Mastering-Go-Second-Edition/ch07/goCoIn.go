package main

import (
	"fmt"
)

type first struct {
	name string
}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("This is shared() from first!", a.name)
}

type second struct {
	name string
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second!", a.name)
}

func (a second) F() {
	a.shared()
}

func main() {
	fmt.Println(">> First calling method F() which calls shared(), expect message from first", "\n> first{}.F()")
	first{}.F()
	fmt.Println(">> calling shared() method via second type", "\n> second{}.shared()")
	second{}.shared()
	fmt.Println(">> init var i with type second", "\n> i := second{}")
	i := second{}
	fmt.Println(">> init var j with embedded type first, so j is first", "\n> j := i.first")
	j := i.first
	fmt.Println(">> var j is type first passed via embedded type in type second", "\n> j.F()", "")
	j.F()
	fmt.Println(">> var i is type second with embedded type first, so we will get shaerd first message", "\n> (i.first).F()")
	(i.first).F()
	fmt.Println(">> should be first method called line 11", "\n> (second{}.first).F()")
	(second{}.first).F()

	fmt.Println("\n================================================\n>>> My additional code")
	f := first{"first"}
	ef := first{"embeded first"}
	s := second{"second", ef}

	fmt.Println(">>> calling first shared and F() with extra naem", "\n> f.shared()\n> f.F()")
	f.shared()
	f.F()
	fmt.Println(">>> calling embedded first directly same as above shared and F() with extra naem", "\n> ef.shared()\n> ef.F()")
	ef.shared()
	ef.F()
	fmt.Println(">>> calling secong shared and F() with extra naem", "\n> s.shared()\n> s.F()")
	s.shared()
	s.F()
	fmt.Println(">>> calling embedded first via second type same as above shared and F() with extra naem", "\n> s.first.shared()\n> s.first.F()")
	s.first.shared()
	s.first.F()
}
