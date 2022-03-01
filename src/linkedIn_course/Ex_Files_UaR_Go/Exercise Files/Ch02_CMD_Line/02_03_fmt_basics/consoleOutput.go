package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42
	isTrue := true

	//fmt.Println(str1, str2, str3)

	strLen, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String Length: ", strLen)
	}

	strLen2, _ := fmt.Println(str1, str2)

	fmt.Println("String Length: ", strLen2)

	fmt.Printf("Value of aNumber: %v\n", aNumber)
	fmt.Printf("Value of isTrue: %v\n", isTrue)
	fmt.Printf("Value of aNumber as float: %.3f\n", float64(aNumber))

	fmt.Printf("Data types: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)

	myStr := fmt.Sprintf("Data types as var: %T, %T, %T, %T, and %T\n", str1, str2, str3, aNumber, isTrue)
	fmt.Print(myStr)
}
