package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s. [%v]\n", xType, reflect.ValueOf(&x))
	fmt.Printf("The type of x is %s.\n", reflect.TypeOf(x))

	A := a{100, 200.12, "Struct a"}
	B := b{1, 2, "Struct b", -1.2}
	var r reflect.Value
	//var t interface{}
	arguments := os.Args
	if len(arguments) == 1 {
		r = reflect.ValueOf(&A).Elem()
		//t = A
	} else {
		r = reflect.ValueOf(&B).Elem()
		//t = B
	}

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s - %T [%v]\n", iType.Field(i).Name, i, i)
		fmt.Printf("with type: %s - %T [%v]", r.Field(i).Type(), r.Field(i), i)
		fmt.Printf(" and value %v - %T\n", r.Field(i).Interface(), r.Field(i).Interface())
	}
}
