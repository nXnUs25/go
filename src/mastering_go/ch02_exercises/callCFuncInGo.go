package main

//#include <stdio.h>
//void callC_ASCII() {
// 	  char c = "&";
//    printf("%c %d", c, c);
//}
import "C"
import "fmt"

func main() {
	fmt.Println("A Go statement!")
	C.callC_ASCII()
	fmt.Println("Another Go statement!")
}
