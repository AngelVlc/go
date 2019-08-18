package workingmonth

import (
	"fmt"
	"testing"
	"time"
)

func TestDayIsWeekend(t *testing.T) {
	for _, testDay := range testDays() {
		d := time.Date(testDay.year, time.Month(testDay.month), testDay.day, 0, 0, 0, 0, time.UTC)
		testName := fmt.Sprintf("For %q", d.Weekday())
		t.Run(testName, func(t *testing.T) {
			expected := testDay.isWeekend

			got := dayIsWeekend(&d)

			if expected != got {
				t.Errorf("Expected to be %v but got %v", expected, got)
			}
		})
	}
}

func TestTotalDays(t *testing.T) {
	for _, testMonth := range testMonths() {
		testName := fmt.Sprintf("For %v_%v", testMonth.year, testMonth.month)
		t.Run(testName, func(t *testing.T) {
			expected := testMonth.lastDay

			wm := WorkingMonth{
				Year:  testMonth.year,
				Month: testMonth.month,
			}
			got := wm.TotalDays()

			if expected != got {
				t.Errorf("Expected to be %v but got %v", expected, got)
			}
		})
	}
}

func TestWorkingHours(t *testing.T) {
	for _, testMonth := range testMonths() {
		testName := fmt.Sprintf("For %v_%v", testMonth.year, testMonth.month)
		t.Run(testName, func(t *testing.T) {
			expected := testMonth.workingHours

			wm := WorkingMonth{
				Year:  testMonth.year,
				Month: testMonth.month,
			}
			got := wm.WorkingHours()

			if expected != got {
				t.Errorf("Expected to be %v but got %v", expected, got)
			}
		})
	}
}

func BenchmarkWorkingHours(b *testing.B) {
	wm := WorkingMonth{
		Year:  2019,
		Month: 9,
	}

	for i := 0; i < b.N; i++ {
		wm.WorkingHours()
	}
}

func ExampleWorkingMonth_WorkingHours() {
	wm := WorkingMonth{2019, 8}

	result := wm.WorkingHours()

	fmt.Println(result)

	// Output: 176
}
