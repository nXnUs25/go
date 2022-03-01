package main

import "fmt"

func main() {
	n1, l1 := fullName("Adam", "Chmiel")
	fmt.Printf("Fullname: %-20v, number of chars: %-3v\n", n1, l1)

	n2, l2 := fullNameNakedReturn("Augustyn", "Chmiel")
	fmt.Printf("Fullname: %-20v, number of chars: %-3v\n", n2, l2)

}

func fullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full) - 1
	return full, length
}

func fullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}