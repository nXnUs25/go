package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Println(str1)
	fmt.Printf("str1 value: [%v], str1 variable type: [%T]\n", str1, str1)

	var str2 string = "An explicitly typed string"
	fmt.Println(str2)
	fmt.Printf("str1 value: [%v], str1 variable type: [%T]\n", str2, str2)

	fmt.Println("ToUpper: ", strings.ToUpper(str2))
	fmt.Println("Title: ", strings.Title(str2))
	fmt.Println("ToTitle: ", strings.ToTitle(str2))

	lValue := "hellow"
	uValue := "HELLOW"

	fmt.Println("Equal case-sensitive ? :", lValue, ", ", uValue, " is ", (lValue == uValue))
	fmt.Println("Compare ? :", lValue, ", ", uValue, " is ", strings.Compare(lValue, uValue))
	// fmt.Println("Equal case-sensitive ? :", lValue, ", ", uValue, " is ", strings.Equal(lValue, uValue))
	fmt.Println("EqualFold non-case-sensitive ? :", lValue, ", ", uValue, " is ", strings.EqualFold(lValue, uValue))


	fmt.Println("Contains exp? ", str1, " ? ", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp? ", str2, " ? ", strings.Contains(str2, "exp"))
}
