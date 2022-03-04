package main

import (
	"fmt"
	"mastering_go/ch03_exercises/data_time"
	"time"
)

func main() {
	mytime := "10:49:46"
	//mydate := ""
	loc1, _ := time.LoadLocation("Europe/Dublin")
	loc2, _ := time.LoadLocation("Europe/Kiev")
	loc3, _ := time.LoadLocation("Europe/Warsaw")

	t, _ := data_time.ConvertTime("15:04:05", mytime)
	dublin := t.In(loc1)
	kiev := t.In(loc2)
	warsow := t.In(loc3)

	fmt.Printf("Default time is: [%9v]\nSame time in Dublin: [%9v]\nSame time in Warsow: [%9v]\nSame time in Kiev: [%9v]\n", mytime, dublin, warsow, kiev)

	fmt.Println()
	td := time.Now()
	formatT1 := td.Format("15:04AM Jan 1st 2006")
	fmt.Println(formatT1)
	dublinTime := td.In(loc1)
	kievTime := td.In(loc2)
	warsawTime := td.In(loc3)
	fmt.Println("Dublin:", dublinTime)
	fmt.Println("Kiev:", kievTime)
	fmt.Println("Warsaw:", warsawTime)
}
