package main

import (
	"fmt"
	"log"
	"github.com/AngelVlc/go/workingmonth"
)

func main() {
	var year, month int

	fmt.Print("Year: ")
	if _, err:= fmt.Scan(&year); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Month: ")
	if _, err:= fmt.Scan(&month); err != nil {
		log.Fatal(err)
	}

	wm := workingmonth.WorkingMonth{
		Year: year, 
		Month: month,
	}
	fmt.Printf("Working hours for %v/%02d: %v\n", wm.Year, wm.Month, wm.WorkingHours())
	fmt.Printf("Working hours until today: %v\n", wm.WorkingHoursUntilToday())
}
