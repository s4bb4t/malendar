package calendar

import "time"

type Week struct {
	Monday    Day
	Tuesday   Day
	Wednesday Day
	Thursday  Day
	Friday    Day
	Saturday  Day
	Sunday    Day
}

func NewWeek(t time.Time) *Week {
	var monday time.Time
	switch t.Weekday() {
	case time.Monday:
		monday = t
	case time.Tuesday:
		monday = t.AddDate(0, 0, -1)
	case time.Wednesday:
		monday = t.AddDate(0, 0, -2)
	case time.Thursday:
		monday = t.AddDate(0, 0, -3)
	case time.Friday:
		monday = t.AddDate(0, 0, -4)
	case time.Saturday:
		monday = t.AddDate(0, 0, -5)
	case time.Sunday:
		monday = t.AddDate(0, 0, -6)
	}

	return &Week{
		Monday:    Day{Date: monday},
		Tuesday:   Day{Date: monday.AddDate(0, 0, 1)},
		Wednesday: Day{Date: monday.AddDate(0, 0, 2)},
		Thursday:  Day{Date: monday.AddDate(0, 0, 3)},
		Friday:    Day{Date: monday.AddDate(0, 0, 4)},
		Saturday:  Day{Date: monday.AddDate(0, 0, 5)},
		Sunday:    Day{Date: monday.AddDate(0, 0, 6)},
	}
}
