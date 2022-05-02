/*
Explore reflection using your own example. How does reflection work on Go maps?

(Page 378).

also good to read
https://golangbot.com/reflection/

How does reflection work on Go maps?

(Page 378).
*/

package main

import (
	"fmt"
	"reflect"
)

type deepEqual struct {
	v1 string
	v2 int
	v3 float64
	v4 bool
}

func main() {
	testDeepEqual()
	fmt.Println("======")

	typeValue(&deepEqual{v1: "testme", v2: 2, v3: 34.34, v4: false})
	fmt.Println("=====")

	var m map[string]int
	fmt.Println(reflect.TypeOf(m).Key())
	fmt.Println(reflect.TypeOf(m).Elem())
}

func testDeepEqual() {
	a := []string{""}
	b := []string{""}

	fmt.Printf("a: %v is deepEqual [%v] to b: %v\n", a, reflect.DeepEqual(a, b), b)
	b = []string{"a", "b", "c"}
	fmt.Printf("a: %v is deepEqual [%v] to b: %v\n", a, reflect.DeepEqual(a, b), b)
	a = []string{"a", "d", "c"}
	fmt.Printf("a: %v is deepEqual [%v] to b: %v\n", a, reflect.DeepEqual(a, b), b)
	a = []string{"a", "b", "c"}
	fmt.Printf("a: %v is deepEqual [%v] to b: %v\n", a, reflect.DeepEqual(a, b), b)

	d1 := &deepEqual{
		v1: "test",
		v2: 12,
		v3: 2.456,
		v4: true,
	}
	d2 := &deepEqual{
		v1: "test",
		v2: 12,
		v3: 4.252,
		v4: true,
	}
	fmt.Printf("d1: %v is deepEqual [%v] to d2: %v\n", d1, reflect.DeepEqual(d1, d2), d2)
	d1 = &deepEqual{
		v1: "test",
		v2: 12,
		v3: 2.456,
		v4: true,
	}
	d2 = &deepEqual{
		v1: "test",
		v2: 12,
		v3: 2.456,
		v4: true,
	}
	fmt.Printf("d1: %v is deepEqual [%v] to d2: %v\n", d1, reflect.DeepEqual(d1, d2), d2)

	m1 := map[string]int{
		"m1key1": 1,
		"m1key2": 2,
	}
	m2 := map[string]int{
		"m2key1": 1,
		"m2key2": 2,
	}
	fmt.Printf("m1: %v is deepEqual [%v] to m2: %v\n", m1, reflect.DeepEqual(m1, m2), m2)

	m1 = map[string]int{
		"key1": 1,
		"key2": 2,
	}
	m2 = map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Printf("m1: %v is deepEqual [%v] to m2: %v\n", m1, reflect.DeepEqual(m1, m2), m2)

	fmt.Println("method deepequal is good fro testing code and comparing things")
}

func typeValue(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	for _, v := range []interface{}{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s\n", v.Kind())
		}
	}
}
