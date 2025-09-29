package calendar

import "time"

type Month struct {
	Days []Day
	Date time.Time
}

func NewMonth() *Month {
	return &Month{
		Days: make([]Day, 0),
		Date: time.Now(),
	}
}
func (m *Month) AddDay(day Day) {
	m.Days = append(m.Days, day)
}
func (m *Month) SetDate(date time.Time) {
	m.Date = date
}
