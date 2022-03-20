package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, " --- ", t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	fmt.Println()
	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	fmt.Println()
	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris:", londonTime)

	fmt.Println()
	td := time.Now()
	formatT1 := td.Format("15:04AM Jan 1st 2006")
	fmt.Println(formatT1)
	loc1, _ := time.LoadLocation("Europe/Dublin")
	dublinTime := td.In(loc1)
	fmt.Println("Dublin:", dublinTime)
}
