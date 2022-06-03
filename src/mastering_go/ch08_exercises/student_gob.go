package main

import (
	//"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

type Student struct {
	Name string
	Age  int32
	id   int // u need capital names to allow serialization otherwise it will initialize to default values ... in int example will be 0
}

func main() {

	fmt.Println("Gob Example")

	studentEncode := Student{Name: "Ketan", Age: 30, id: 3}
	fname := "student.gob"
	if _, err := os.Stat(fname); err == nil {
		err := os.Remove(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot remove file %s - %v\n", fname, err)
		}
	}
	saveTo, err := os.Create(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file: %s, %v\n", fname, err)
	}
	defer saveTo.Close()

	e := gob.NewEncoder(saveTo)
	if err := e.Encode(studentEncode); err != nil {
		panic(err)
	}
	fmt.Println("Encoded Struct ", saveTo)

	loadF, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Empty key/value store.\n")
	}
	defer loadF.Close()
	var studentDecode Student
	d := gob.NewDecoder(loadF)
	if err := d.Decode(&studentDecode); err != nil {
		panic(err)
	}

	fmt.Println("Decoded Struct ", studentDecode.Name, "\t", studentDecode.Age)
	fmt.Println(studentDecode)

}
