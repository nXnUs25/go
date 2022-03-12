package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type WorkPlace struct {
	Employees []Employees `json:"employees"`
}
type Employees struct {
	Name    string `json:"name"`
	Salary  int    `json:"salary"`
	Married bool   `json:"married"`
}

var (
	pln = fmt.Println
	pf  = fmt.Printf
)

func main() {
	myWorkPlace := WorkPlace{
		[]Employees{
			{
				Name:    "sonoo1",
				Salary:  49000,
				Married: true,
			},
			{
				Name:    "sonoo2",
				Salary:  59000,
				Married: true,
			},
		},
	}

	j, _ := myWorkPlace.Pretty()
	pf("Before Update: \n%v\n", j)

	for i, v := range myWorkPlace.Employees {
		rise := rand.Intn(100) + i
		v.Salary += rise
	}
	for i, _ := range myWorkPlace.Employees {
		rise := rand.Intn(100) + i
		myWorkPlace.Employees[i].Salary += rise
	}
	j1, _ := myWorkPlace.Pretty()
	pf("After Update using range: \n%v\n", j1)

	for i := 0; i < len(myWorkPlace.Employees); i++ {
		rise := rand.Intn(100) + i
		myWorkPlace.Employees[i].Salary += rise
	}
	j2, _ := myWorkPlace.Pretty()
	pf("After Update without of range: \n%v\n", j2)

	pln("================================================================")
	pln("========================== xml =================================")

}

func (wp *WorkPlace) Pretty() (string, error) {
	j, err := json.MarshalIndent(&wp, " ", "   ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}
