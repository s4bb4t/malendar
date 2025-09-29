package calendar

import "time"

type Day struct {
	Date  time.Time
	Phase Phase
}

func NewDay() *Day {
	return &Day{
		Date:  time.Now(),
		Phase: NonPhase,
	}
}

func (d *Day) SetDate(date time.Time) {
	d.Date = date
}

func (d *Day) SetPhase(phase Phase) {
	d.Phase = phase
}

func (d *Day) Validate() bool {
	if d.Date.Unix() != 0 {
		return true
	}
	if d.Phase == NonPhase || d.Phase == OvulationPhase || d.Phase == MenstrualPhase {
		return true
	}
	return false
}
