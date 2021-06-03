package main

import (
	"fmt"
	"unicode"
)

// import (
// 	"fmt"
// 	"unicode"
// )

func main() {
	// Declare a sample string
	fmt.Println("// Declare a sample string")
	const s = "The 'quick' brown fox, jumped over the *LAZY* dog!"
	fmt.Println("Constant s:", s)

	// Init some count variables
	t := `// Init some count variables
	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount := 0`
	fmt.Println("\n", t)
	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount := 0

	// iterate over the characters and call unicode function tests
	fmt.Println("\n", "// iterate over the characters and call unicode function tests")
	for _, ch := range s {
		if unicode.IsPunct(ch) {
			punctCount++
		}
		if unicode.IsLower(ch) {
			lowerCount++
		}
		if unicode.IsUpper(ch) {
			upperCount++
		}
		if unicode.IsSpace(ch) {
			spaceCount++
		}
		if unicode.Is(unicode.Hex_Digit, ch) {
			hexdigitCount++
		}
	}

	t = `// Print the results
	fmt.Println("Punctuation:", punctCount)
	fmt.Println("Lowercase:", lowerCount)
	fmt.Println("Uppercase:", upperCount)
	fmt.Println("Whitespace:", spaceCount)
	fmt.Println("Hex Digits:", hexdigitCount)`
	// Print the results
	fmt.Println("\n", t)
	fmt.Println("Punctuation:", punctCount)
	fmt.Println("Lowercase:", lowerCount)
	fmt.Println("Uppercase:", upperCount)
	fmt.Println("Whitespace:", spaceCount)
	fmt.Println("Hex Digits:", hexdigitCount)
}
