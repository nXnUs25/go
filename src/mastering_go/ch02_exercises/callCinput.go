package main

//#include <stdio.h>
// void callCinput() {
//	int e;
//	char ch;
//	printf("\n[C] - Enter a character: ");
//	scanf("%c", &ch);
//	e=ch;
//	printf("\n[C] - The ASCII value of the character is : %d\n",e);
//}
import "C"
import "fmt"

func main() {
	fmt.Println("[GO] - Starting Go")
	fmt.Println("[GO] - About to call C")
	C.callCinput()
	fmt.Println("[GO] - After C been called")
	//fmt.Printf("value of c.main is : %d", v)
}
