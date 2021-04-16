package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
   var s string
   fmt.Print("Type some words: ")
   fmt.Scanln(&s)
   fmt.Println("You typed first word: ", s)

   reader := bufio.NewReader(os.Stdin)
   fmt.Print("Enter text: ")

   str, _ := reader.ReadString('\n')
   fmt.Println(str)

   fmt.Print("Enter number: ")
   str, _ = reader.ReadString('\n')
   f, err :=  strconv.ParseFloat(strings.TrimSpace(str), 64)

   if err != nil {
       fmt.Println(err)
   } else {
       fmt.Println(f)
   }

}