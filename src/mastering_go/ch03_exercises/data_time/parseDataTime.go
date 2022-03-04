package data_time

import (
	"fmt"
	"time"
)

/*
how to parse time
15 h
04 min
05 sec
PM PM
pm pm
Jan month 3 letters
January January
2006 year
02 day of the month
*/

/*
The Go constants for working with times are 15 for parsing the hour, 04 for parsing the minutes, and 05 for parsing the seconds.
You can easily guess that all these numeric values must be unique.
You can use PM for parsing the PM string in uppercase and pm for parsing the lowercase version.

The Go constants for working with dates are Jan for parsing the three-letter abbreviation used for describing a month, 2006 for parsing the year,
and 02 for parsing the day of the month. If you use January instead of Jan, you will get the long name of the month instead of its three-letter abbreviation,
which makes perfect sense.
*/

const (
	FormatData1 string = "2006/02/01"
	FormatFull2 string = "2006-01-02T15:04:05"
	FormatTime  string = "15-04.05"
)

func ConvertDate(format, current_date string) (time.Time, error) {
	d, err := time.Parse(format, current_date)
	if err != nil {
		fmt.Printf("[Error] Problem with converting [current_date]::[%50v] with [format]::[%20v]\n", current_date, format)
		fmt.Println(err)
		return d, err
	}
	return d, nil
}

func ConvertTime(format, current_time string) (time.Time, error) {
	d, err := time.Parse(format, current_time)
	if err != nil {
		fmt.Printf("[Error] Problem with converting [current_time]::[%50v] with [format]::[%20v]\n", current_time, format)
		fmt.Println(err)
		return d, err
	}
	return d, nil
}
