package myModule

import (
	"fmt"
)

func Version() {
	fmt.Println("Version 1.1.0")
}

func Version2(version string) {
	fmt.Printf("Version 2: [%s] \n", version)
}
