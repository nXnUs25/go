package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// declare sample strings
	s := "the quick brown fox jumped over the lazy dog"
	s2 := []string{"one", "two", "three", "four"}
	s3 := "This is a string. With some punctionation, for a demo! Yep."

	// The Split function splits a string into substrings
	fmt.Println("\nThe Split function splits a string into substrings")
	sub1 := strings.Split(s, " ")
	fmt.Println(sub1)
	fmt.Printf("%q\n", sub1)
	sub2 := strings.Split(s, "the")
	fmt.Println(sub2)
	fmt.Printf("%q\n", sub2)

	// Join concatenates substrings, with the separator between each
	fmt.Println("\nJoin concatenates substrings, with the separator between each")
	result := strings.Join(s2, " - ")
	fmt.Println(result)

	// The Fields function splits a string around white space
	fmt.Println("\nThe Fields function splits a string around white space")
	subF := strings.Fields(s)
	fmt.Println(subF)
	fmt.Printf("%q\n", subF)

	// FieldsFunc is a customizable version of fields that uses a callback
	fmt.Println("\nFieldsFunc is a customizable version of fields that uses a callback")
	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	result3 := strings.FieldsFunc(s3, f)
	fmt.Println(result3)

	// Replacer can be used for multiple replacement operations
	fmt.Println("\nReplacer can be used for multiple replacement operations")
	rep := strings.NewReplacer(".", "|", ",", "|", "!", "|")
	result4 := rep.Replace(s3)
	fmt.Println(result4)
}
