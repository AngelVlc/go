package main

import (
	"fmt"
	"github.com/AngelVlc/go/workingmonth"
	"log"
	"time"
)

func main() {
	var discountDays int

	fmt.Print("PTOs and holiday days: ")
	if _, err := fmt.Scan(&discountDays); err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	wm := workingmonth.WorkingMonth{
		Year:  now.Year(),
		Month: int(now.Month()),
	}

	hoursUntilToday := wm.WorkingHoursUntilToday() - (discountDays * 8)
	fmt.Printf("Working hours until today: %v\n", hoursUntilToday)
	fmt.Printf("Working hours until month end:  %v\n", wm.WorkingHours())
}
