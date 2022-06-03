package main

import (
	//"bufio"
	"encoding/gob"
	"fmt"
	"os"
	//"strings"
)

type myData struct {
	Text string
	Date string
	Id   int
}

var myData3 = myData{Text: "bla1 new file", Date: "bla....", Id: 123}
var decideMyData myData
var storedData3 = "myData4.gob"

func main() {
	fmt.Fprintf(os.Stdout, "Data before: \n- %v\n", myData3)

	err := save(storedData3, myData3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot save data to %s \n%v\n", storedData3, err)
	}

	any, err := load(storedData3, decideMyData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot load file %s\n%v\n", storedData3, err)
	}
	convert(any)
	fmt.Fprintf(os.Stdout, "Data after: \n- %v\n", decideMyData)
}

func convert(data interface{}) {
	switch d := data.(type) {
	case myData:
		fmt.Printf("%T, %[1]v\n", d)
	default:
		fmt.Println("unknown")
	}
}

func save(fname string, data myData) error {
	fmt.Fprintf(os.Stdout, "%v\n", data)
	if _, err := os.Stat(fname); err == nil {
		err := os.Remove(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot remove file %s - %v\n", fname, err)
			return err
		}
	}

	saveTo, err := os.Create(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file: %s, %v\n", fname, err)
		return err
	}
	defer saveTo.Close()

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(&data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot save to file %s - %v\n", fname, err)
		return err
	}
	return nil
}

func load(fname string, data interface{}) (interface{}, error) {
	fmt.Fprintf(os.Stdout, "Loading file %v\n", &data)
	loadF, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Empty key/value store.\n")
		return nil, err
	}
	defer loadF.Close()
	dec := gob.NewDecoder(loadF)
	dec.Decode(&data)
	fmt.Printf("Load Data -- %v - type %[1]T\n", data)
	return data, nil
}

func (a *myData) String() string {
	if a != nil {
		return fmt.Sprintf(myDataStrFormat, a.Text, a.Date, a.Id)
	}
	return ""
}

const myDataStrFormat string = `
Text: <%v>,
date: <%v>,
id: <%v>
`
