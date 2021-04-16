package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

 func (d Dog) Speak(){
	 fmt.Println(d.Sound)
 }

 func (d Dog) SpeakThreeTimes() {
	 d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	 fmt.Print(d.Sound)
 }

 func (d *Dog) SpeakManyTimes() {
	 // uses pointer *Dog
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println("Pointer: ", d.Sound)
}

func main() {

	poodle := Dog{"Poodle", 37, "Woof"}

	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Afr"

	poodle.SpeakThreeTimes()
	poodle.SpeakManyTimes()
	poodle.SpeakManyTimes()

}
