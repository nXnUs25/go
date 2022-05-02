/*
Write your own interface and use it in another Go program.
Then state why your interface is useful.

(Page 378).
*/

package main

import (
	"fmt"
)

type tweeter interface {
	Tweet()
}

type sayer interface {
	Say()
}

type talker interface {
	tweeter
	sayer
}

type person struct {
	name string
	age  int
}

func (p *person) String() string {
	return fmt.Sprintf("%s is %v years old", p.name, p.age)
}

func (p *person) Tweet() {
	fmt.Printf("Tweet from person %s\n", p.name)
}

func (p *person) Say() {
	fmt.Printf("Say from person %s\n", p.name)
}

func main() {
	var t talker = &person{
		name: "Kevin",
		age:  1,
	}
	t.Tweet()
	t.Say()
	fmt.Println(t)

	t = &person{
		name: "Adam",
		age:  2,
	}
	t.Tweet()
	t.Say()
	fmt.Println(t)
}
