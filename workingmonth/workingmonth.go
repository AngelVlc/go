// Package workingmonth contains methods to ge the number of
// working hours in a given month
package workingmonth

import (
	"time"
)

// WorkingMonth contains methods to get the working hours
// in a given month
type WorkingMonth struct {
	Year    int
	Month   int
	nowFunc func() time.Time
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
	totalDays := m.TotalDays()
	return m.workingDaysUntilDay(totalDays)
}

// WorkingDaysUntilToday returns the number of working days until today
func (m WorkingMonth) WorkingDaysUntilToday() int {
	return m.workingDaysUntilDay(m.now().Day())
}

// WorkingHours returns the number of working hours in the month
func (m WorkingMonth) WorkingHours() int {
	return m.WorkingDays() * 8
}

// WorkingHoursUntilToday returns the number of working hours at the
// end of today
func (m WorkingMonth) WorkingHoursUntilToday() int {
	return m.WorkingDaysUntilToday() * 8
}

func (m WorkingMonth) workingDaysUntilDay(day int) int {
	month := m.toTimeMonth()
	date := time.Date(m.Year, month, 1, 0, 0, 0, 0, time.UTC)
	var result int
	for date.Month() == month && date.Day() <= day {
		if !dayIsWeekend(&date) {
			result++
		}
		date = date.AddDate(0, 0, 1)
	}

	return result
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

func (m WorkingMonth) now() time.Time {
	if m.nowFunc != nil {
		return m.nowFunc()
	}
	return time.Now()
}
