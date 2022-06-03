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

var myData1 = make(map[string]myData)
var myData2 = []myData{}
var myData3 = myData{Text: "bla1 data3", Date: "bla....", Id: 123}
var storedData1 = "myData1.gob"
var storedData2 = "myData2.gob"
var storedData3 = "myData3.gob"

func (a *myData) String() string {
	if a != nil {
		return fmt.Sprintf(myDataStrFormat, a.Text, a.date, a.id)
	}
	return ""
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
	err = encoder.Encode(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot save to file %s - %v\n", fname, err)
		return err
	}
	return nil
}

func load(fname string, data myData) error {
	fmt.Fprintf(os.Stdout, "Loading file %v\n", &data)
	loadF, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Empty key/value store.\n")
		return err
	}
	defer loadF.Close()
	dec := gob.NewDecoder(loadF)
	dec.Decode(&data)
	d := data
	fmt.Printf("Load Data -- %v - type %[1]T\n", d)
	return nil
}

func main() {
	var data myData
	//var data3 myData = myData{}
	fmt.Println("-------------------------------- load --------------------------------")
	for i, v := range []string{storedData1, storedData2, storedData3} {
		switch i {
		case 0:
			data = myData3
			//data3 = make(map[string]myData)
		case 1:
			data = myData3
			//data3 = make([]myData, 0)
		case 2:
			data = myData3
			//data3 = myData{}
		}
		fmt.Fprintf(os.Stdout, "Data before: \n- %v\n", myData3)
		fmt.Println("Load...")

		err := load(v, myData3)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot load data from %s \n%v\n", v, err)
		}
		fmt.Fprintf(os.Stdout, "Data after: \n- %v\n----\n", myData3)
		fmt.Println()
	} /*
		fmt.Println("-------------------------------- edit --------------------------------")
		for i, v := range []string{storedData1, storedData2, storedData3} {
			switch i {
			case 0:
				data = map[string]myData{
					"111": myData{Text: "bla1 data3", date: "bla....", id: 1234},
					"222": myData{Text: "bla1 data3", date: "bla....", id: 12345},
				}
			case 1:
				data = []myData{myData{Text: "bla1 data3", date: "bla....", id: 123}, myData{Text: "bla2 data2", date: "bla..33..", id: 1}}
			case 2:
				data = &myData{Text: "bla2 data3", date: "bla..22..", id: 123}
			}
			fmt.Println(v)
		}*/
	fmt.Println("-------------------------------- save --------------------------------")
	for i, v := range []string{storedData1, storedData2, storedData3} {
		switch i {
		case 0:
			//data = map[string]myData{
			//	"111": myData{Text: "bla1 data3", date: "bla....", id: 1234},
			//	"222": myData{Text: "bla1 data3", date: "bla....", id: 12345},
			//}
		case 1:
			//data = []myData{myData{Text: "bla1 data3", date: "bla....", id: 123}, myData{Text: "bla2 data2", date: "bla..33..", id: 1}}
		case 2:
			data = myData{Text: "bla2 data3", date: "bla..22..", id: 123}
		}
		err := save(v, data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot save data to %s \n%v\n", v, err)
		}
	} /*
		var data2 interface{}
		fmt.Println("-------------------------------- load2 --------------------------------")
		for i, v := range []string{storedData1, storedData2, storedData3} {
			fmt.Println("Load...")
			err := load(v, &data2)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Cannot load2 data from %s \n%v\n", v, err)
			}
			fmt.Fprintf(os.Stdout, "Data after: \n- %v\n----\n", data2)
			fmt.Println(i)
		}
	*/
	fmt.Println("-------------------------------- done --------------------------------")
}

const myDataStrFormat string = `Text: <%v>,
date: <%v>,
id: <%v>
`
