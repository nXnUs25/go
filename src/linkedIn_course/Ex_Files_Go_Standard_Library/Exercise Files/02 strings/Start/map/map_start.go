package main

import (
	"fmt"
	"strings"
)

func main() {
	// The map function returns a copy of a string with the characters modified
	// according to the mapping function
	shift := 2
	s := "The quick brown fox jumps over the lazy dog"

	// TODO: create the mapping function
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') + (int(r) - int('A') + shift)
			value = calculateValue(value, 91, 65)
			return rune(value)
		case r >= 'a' && r <= 'z':
			value := int('a') + (int(r) - int('a') + shift)
			value = calculateValue(value, 122, 97)
			return rune(value)
		}
		return r
	}

	// TODO: Encode the message
	encode := strings.Map(transform, s)
	fmt.Println(encode)

	// TODO: Decode the message
	shift = -shift
	decode := strings.Map(transform, encode)
	fmt.Println(decode)
}

func calculateValue(value, asciUpper, asciLower int) int {
	if value > asciUpper { // 91 is the limit of upper cases
		value -= 26
	} else if value < asciLower {
		value += 26
	}
	return value
}
