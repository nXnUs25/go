package main

import (
	"fmt"
	"strings"
)

// import (
// 	"fmt"
// 	"strings"
// )

func main() {
	// TODO: Declare an empty strings.Builder
	var sb strings.Builder

	// TODO: Write some content
	fmt.Println("\nTODO: Write some content")
	sb.WriteString("This is string 1\n")
	sb.WriteString("This is string 2\n")
	sb.WriteString("This is string 3\n")

	// TODO: Output the concatenated string
	fmt.Println("\nTODO: Output the concatenated string")
	fmt.Println(sb.String())

	// TODO: Examine the builder's capacity
	fmt.Println("\nTODO: Examine the builder's capacity")
	fmt.Println("Capacity:", sb.Cap())
	sb.Grow(1024)
	fmt.Println("Capacity:", sb.Cap())

	// TODO:
	// Grow the capacity - use this if you know in advance how much data
	// you're going to be writing into the buffer to minimize copies
	fmt.Println("\nTODO:\nEGrow the capacity - use this if you know in advance how much data\nyou're going to be writing into the buffer to minimize copies")
	for i := 0; i <= 10; i++ {
		fmt.Fprintf(&sb, "String %d -- ", i)
	}
	fmt.Println(sb.String())

	// TODO: we can get the length of what the final string will be
	fmt.Println("\nTODO: we can get the length of what the final string will be")
	fmt.Println("Builder size is", sb.Len())

	// TODO: The Reset function will reset the builder to original state
	fmt.Println("\nTODO: The Reset function will reset the builder to original state")
	sb.Reset()
	fmt.Println("After Reset:")
	fmt.Println("Capacity:", sb.Cap())
	fmt.Println("Builder size is", sb.Len())
}
