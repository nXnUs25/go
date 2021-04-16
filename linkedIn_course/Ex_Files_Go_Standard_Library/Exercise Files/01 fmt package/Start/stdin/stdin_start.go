package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Insert some text: ")
	s, _ := reader.ReadString('\n')
	fmt.Println(s)
}
