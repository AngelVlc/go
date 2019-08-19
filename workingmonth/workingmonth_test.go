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
		testName := testNameForTestMonth(&testMonth, t)
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
		testName := testNameForTestMonth(&testMonth, t)
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

func TestWorkingHoursUntilToday(t *testing.T) {
	for _, testDay := range testDays() {
		testName := testNameForTestDay(&testDay, t)
		t.Run(testName, func(t *testing.T) {
			expected := testDay.hoursUntilDay

			wm := WorkingMonth{
				Year:  testDay.year,
				Month: testDay.month,
				nowFunc: func() time.Time {
					return time.Date(testDay.year, time.Month(testDay.month), testDay.day, 0, 0, 0, 0, time.UTC)
				},
			}
			got := wm.WorkingHoursUntilToday()
			if expected != got {
				t.Errorf("Expected to be %v but got %v", expected, got)
			}
		})
	}
}

func testNameForTestMonth(tm *testMonth, t *testing.T) string {
	t.Helper()

	return fmt.Sprintf("For %v/%02d", tm.year, tm.month)
}

func testNameForTestDay(td *testDay, t *testing.T) string {
	t.Helper()

	return fmt.Sprintf("For %v/%02d/%02d", td.year, td.month, td.day)
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
	wm := WorkingMonth{
		Year: 2019, 
		Month: 8,
	}

	result := wm.WorkingHours()

	fmt.Println(result)

	// Output: 176
}

func ExampleWorkingMonth_WorkingHoursUntilToday() {
	// if today is 2019/08/18
	wm := WorkingMonth{
		Year: 2019, 
		Month: 8,
		nowFunc: func() time.Time {
			return time.Date(2019, time.Month(8), 18, 0, 0, 0, 0, time.UTC)
		},
	}

	result := wm.WorkingHoursUntilToday()

	fmt.Println(result)

	// Output: 96
}

func ExampleWorkingMonth_WorkingDays() {
	wm := WorkingMonth{
		Year: 2019, 
		Month: 8,
	}

	result := wm.WorkingDays()

	fmt.Println(result)

	// Output: 22
}

func ExampleWorkingMonth_WorkingDaysUntilToday() {
	// if today is 2019/08/18
	wm := WorkingMonth{
		Year: 2019, 
		Month: 8,
		nowFunc: func() time.Time {
			return time.Date(2019, time.Month(8), 18, 0, 0, 0, 0, time.UTC)
		},
	}

	result := wm.WorkingDaysUntilToday()

	fmt.Println(result)

	// Output: 12
}

func ExampleWorkingMonth_TotalDays() {
	wm := WorkingMonth{
		Year: 2019, 
		Month: 8,
	}

	result := wm.TotalDays()

	fmt.Println(result)

	// Output: 31
}
