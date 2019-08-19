package workingmonth

type testDay struct {
	year          int
	month         int
	day           int
	isWeekend     bool
	hoursUntilDay int
}

func testDays() []testDay {
	r := make([]testDay, 7)

	r[0] = testDay{
		year:          2019,
		month:         8,
		day:           12,
		isWeekend:     false,
		hoursUntilDay: 64,
	}

	r[1] = testDay{
		year:          2019,
		month:         8,
		day:           13,
		isWeekend:     false,
		hoursUntilDay: 72,
	}

	r[2] = testDay{
		year:          2019,
		month:         8,
		day:           14,
		isWeekend:     false,
		hoursUntilDay: 80,
	}

	r[3] = testDay{
		year:          2019,
		month:         8,
		day:           15,
		isWeekend:     false,
		hoursUntilDay: 88,
	}

	r[4] = testDay{
		year:          2019,
		month:         8,
		day:           16,
		isWeekend:     false,
		hoursUntilDay: 96,
	}

	r[5] = testDay{
		year:          2019,
		month:         8,
		day:           17,
		isWeekend:     true,
		hoursUntilDay: 96,
	}

	r[6] = testDay{
		year:          2019,
		month:         8,
		day:           18,
		isWeekend:     true,
		hoursUntilDay: 96,
	}

	return r
}

type testMonth struct {
	year         int
	month        int
	lastDay      int
	workingHours int
}

func testMonths() []testMonth {
	r := make([]testMonth, 13)

	r[0] = testMonth{
		year:         2019,
		month:        1,
		lastDay:      31,
		workingHours: 23 * 8,
	}

	r[1] = testMonth{
		year:         2019,
		month:        2,
		lastDay:      28,
		workingHours: 20 * 8,
	}

	r[2] = testMonth{
		year:         2019,
		month:        3,
		lastDay:      31,
		workingHours: 21 * 8,
	}

	r[3] = testMonth{
		year:         2019,
		month:        4,
		lastDay:      30,
		workingHours: 22 * 8,
	}

	r[4] = testMonth{
		year:         2019,
		month:        5,
		lastDay:      31,
		workingHours: 23 * 8,
	}

	r[5] = testMonth{
		year:         2019,
		month:        6,
		lastDay:      30,
		workingHours: 20 * 8,
	}

	r[6] = testMonth{
		year:         2019,
		month:        7,
		lastDay:      31,
		workingHours: 23 * 8,
	}

	r[7] = testMonth{
		year:         2019,
		month:        8,
		lastDay:      31,
		workingHours: 22 * 8,
	}

	r[8] = testMonth{
		year:         2019,
		month:        9,
		lastDay:      30,
		workingHours: 21 * 8,
	}

	r[9] = testMonth{
		year:         2019,
		month:        10,
		lastDay:      31,
		workingHours: 23 * 8,
	}

	r[10] = testMonth{
		year:         2019,
		month:        11,
		lastDay:      30,
		workingHours: 21 * 8,
	}

	r[11] = testMonth{
		year:         2019,
		month:        12,
		lastDay:      31,
		workingHours: 22 * 8,
	}

	r[12] = testMonth{
		year:         2021,
		month:        2,
		lastDay:      28,
		workingHours: 20 * 8,
	}

	return r
}
