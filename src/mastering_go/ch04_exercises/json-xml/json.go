package main

import (
	"bytes"
	"encoding/json"
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

--- json ---
*/

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
	file   = pwd + utils + "/data.json"
)

type WorkPlace struct {
	Employees []struct {
		Name    string `json:"name"`
		Salary  int    `json:"salary"`
		Married bool   `json:"married"`
	} `json:"employees"`
}

func GetOutputJsonLogger() *bytes.Buffer {
	return &buf
}

func main() {
	l.SetOutput(GetOutputJsonLogger())
	data, err := ReadJSON(file)
	if err != nil {
		lln(E+"Failed to read ", file)
	}
	var myWorkPlace WorkPlace
	ParseJSON(data, &myWorkPlace)
	Pln(myWorkPlace)

	for i := 0; i < len(myWorkPlace.Employees); i++ {
		rise := rand.Intn(100) + i
		myWorkPlace.Employees[i].Salary += rise
	}
	Pln(myWorkPlace)
	saveToJSON(os.Stdout, myWorkPlace)
	f, err := os.Create(file + ".Save")
	defer f.Close()
	if err != nil {
		lf(E+"Failed to create [%V] = func() [%v] \n", file, GetFunc())
	}
	saveToJSON(f, myWorkPlace)
	pjson, _ := myWorkPlace.Pretty()
	Pln(pjson)
}

func ReadJSON(file string) ([]byte, error) {
	lf(I+"Reading a file - [%v] - func: [%v]\n", file, GetFunc())
	data, err := ioutil.ReadFile(file)
	if err != nil {
		lf(E+"Could not read file [%s] - %v\n", file, err)
		return nil, err
	}
	lln(D+"Data: ", data, "eding func: ", GetFunc())
	return data, nil
}

func ParseJSON(data []byte, structure interface{}) error {
	lf(I+"Parsing a data - [%v] - func: [%v]\n", file, GetFunc())
	dec := json.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(structure)
	if err != nil {
		lf(E+"Could not parse data: [%v]\n", data)
		return err
	}
	lln(D+"Data: ", structure, "eding func: <", GetFunc(), ">")
	return nil
}

func (wp *WorkPlace) Pretty() (string, error) {
	j, err := json.MarshalIndent(&wp, " ", "   ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func saveToJSON(file *os.File, structure interface{}) bool {
	enc := json.NewEncoder(file)
	enc.SetIndent("", "   ")
	err := enc.Encode(structure)
	if err != nil {
		lf(E + "Could not encode structure: [%v]\n")
		return false
	}
	return true
}
