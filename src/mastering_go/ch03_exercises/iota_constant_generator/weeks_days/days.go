package main

import (
	"fmt"
	"mastering_go/ch03_exercises/iota_constant_generator"
	"time"
)

// Write a constant generator iota for the days of the week.
func main() {

	fmt.Printf("Week Days are: %s\n", iota_constant_generator.WeekDays())
	fmt.Printf("Week Days Star are: %s\n\n", iota_constant_generator.WeekDaysStar())

	fmt.Printf("Week Day and Type is: %d [%[1]T] : %s [%[2]T]\n\n", iota_constant_generator.Friday, iota_constant_generator.Friday.String())

	start := time.Now()
	for x := 0; x < len(iota_constant_generator.WeekDays()); x++ {
		fmt.Printf("%[2]d : Week Day in Polish is: %[1]s\n", iota_constant_generator.WeekDays()[x].Polish(), x)
	}
	duration := time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")
	fmt.Println("")

	start = time.Now()
	for x := 0; x < len(iota_constant_generator.WeekDays()); x++ {
		fmt.Printf("%[2]d : Week Day in String is: %[1]s\n", iota_constant_generator.WeekDays()[x].String(), x)
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")

	start = time.Now()
	for x := 0; x < len(iota_constant_generator.WeekDays()); x++ {
		fmt.Printf("%[2]d : Week Day in English is: %[1]s\n", iota_constant_generator.WeekDays()[x].English(), x)
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")
	fmt.Println("")

	start = time.Now()
	for x := 0; x < len(iota_constant_generator.WeekDaysStar()); x++ {
		fmt.Printf("%[2]d : Week Day in Star String is: %[1]s\n", iota_constant_generator.WeekDaysStar()[x].String(), x)
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")

	start = time.Now()
	for x := 0; x < len(iota_constant_generator.WeekDaysStar()); x++ {
		fmt.Printf("%[2]d : Week Day in Star English is: %[1]s\n", iota_constant_generator.WeekDaysStar()[x].English(), x)
	}
	duration = time.Since(start)
	fmt.Println("It took the for loop", duration, "to finish.")
	fmt.Println("")

	fmt.Println("\nDone.")
}
