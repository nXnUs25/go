package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"runtime"
)

/*
Read a JSON file with 10 integer values, store it in a struct variable, increment each integer value by one,
and write the updated JSON entry on disk. Now, do the same for an XML file.

--- xml.go ----
*/
type WorkPlace struct {
	XMLName   xml.Name `xml:"WorkPlace"`
	Text      string   `xml:",chardata"`
	Employees []struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name"`
		Salary  int    `xml:"salary"`
		Married string `xml:"married"`
	} `xml:"employees"`
}

var (
	Pf      = fmt.Printf
	Pln     = fmt.Println
	buf     bytes.Buffer
	l       = log.New(&buf, "json ", log.LstdFlags|log.Lshortfile)
	I       = "[INFO] "
	D       = "[DEBUG] "
	E       = "[ERROR] "
	GetFunc = func() string {
		pc, _, _, _ := runtime.Caller(1)
		return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
	}
	lf     = log.Printf
	lln    = log.Println
	utils  = "/src/mastering_go/ch04_exercises/json-xml"
	pwd, _ = os.Getwd()
	file   = pwd + utils + "/data.xml"
)

func GetOutputJsonLogger() *bytes.Buffer {
	return &buf
}

func main() {
	log.SetOutput(&buf)
	data, err := ReadXML(file)
	if err != nil {
		lln(E+"Failed to read XML file", file)
	}
	var myWorkPlace WorkPlace
	ParseXML(data, &myWorkPlace)
	Pln("Before Salary update:\n", myWorkPlace.Pretty())
	for i := 0; i < len(myWorkPlace.Employees); i++ {
		rise := rand.Intn(100) + i
		myWorkPlace.Employees[i].Salary += rise
	}
	Pln("After Salary update:\n", myWorkPlace.Pretty())
	Pln("Sending same as above but to the file:", file+".Save")
	SaveToXML(file+".Save", myWorkPlace)

	f, _ := ReadXML2(file)
	in := xml.NewDecoder(f)
	in.Decode(&myWorkPlace)
	Pln("Other method reading file:\n", myWorkPlace.Pretty())
}

func ReadXML(file string) ([]byte, error) {
	lf(I+"Reading file [%v] in func() [%v]\n", file, GetFunc())
	data, err := ioutil.ReadFile(file)
	if err != nil {
		lf(E+"Cannot read file [%v] in func() [%v] - err message: %v\n", file, GetFunc(), err)
		return nil, err
	}
	return data, nil
}

func ReadXML2(file string) (*os.File, error) {
	lf(I+"Reading file [%v] in func() [%v]\n", file, GetFunc())
	f, err := os.Open(file)
	if err != nil {
		lf(E+"Cannot read file [%v] in func() [%v] - err message: %v\n", file, GetFunc(), err)
		return nil, err
	}
	return f, nil
}

func ParseXML(data []byte, structure interface{}) error {
	lf(I+"Parsing data [%v] in func() [%v]\n", file, GetFunc())
	dec := xml.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(structure)
	if err != nil {
		lf(I+"Cannot parse data [%v] in func() [%v] - err message: %v\n", file, GetFunc(), err)
		return err
	}
	lln(D+"Data: ", structure, "eding func: <", GetFunc(), ">\n")
	return nil
}

func (wp *WorkPlace) Pretty() string {
	lf(I+"Indent data [%v] in func() [%v]\n", file, GetFunc())
	x, err := xml.MarshalIndent(&wp, " ", "   ")
	if err != nil {
		return ""
	}
	return string(x)
}

func SaveToXML(file string, structure interface{}) bool {
	lf(I+"Saving data to file [%v] in func() [%v]\n", file, GetFunc())
	data, err := xml.MarshalIndent(structure, "", "   ")
	if err != nil {
		lf(E+"Problem with Saving data to file [%v] in func() [%v]\n", file, GetFunc())
	}
	ioutil.WriteFile(file, []byte(xml.Header+string(data)), 0644)
	return true
}
