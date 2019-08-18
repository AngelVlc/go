// Package workingmonth contains methods to ge the number of
// working hours in a given month
package workingmonth

import (
	"time"
)

// WorkingMonth contains methods to get the working hours
// in a given month
type WorkingMonth struct {
	Year  int
	Month int
}

func (m WorkingMonth) toTimeMonth() time.Month {
	return time.Month(m.Month)
}

// TotalDays returns the number of days in the month
func (m WorkingMonth) TotalDays() int {
	firstDayOfMonth := time.Date(m.Year, m.toTimeMonth(), 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	return lastDayOfMonth.Day()
}

// WorkingDays returns the number of working days in the month
func (m WorkingMonth) WorkingDays() int {
	month := m.toTimeMonth()
	totalDays := m.TotalDays()
	date := time.Date(m.Year, month, 1, 0, 0, 0, 0, time.UTC)
	var result int
	for date.Month() == month && date.Day() <= totalDays {
		if !dayIsWeekend(&date) {
			result++
		}
		date = date.AddDate(0, 0, 1)
	}

	return result
}

// WorkingHours returns the number of working hours in the month
func (m WorkingMonth) WorkingHours() int {
	return m.WorkingDays() * 8
}

// dayIsWeekend returns true if the week day of a given time
// is saturday or sunday
func dayIsWeekend(t *time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return true
	default:
		return false
	}
}
