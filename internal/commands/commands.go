package commands

import (
	"strconv"
	"time"

	"github.com/s4bb4t/malendar/internal/calendar"
)

func Info(M map[string]string) string {
	var res string
	for k, v := range M {
		res += k + " - " + v + "\n"
	}
	if res == "" {
		res = "No info yet!"
	}
	return res
}

func InfoSet(M map[string]string, key string, value string) {
	if key != "" && value != "" {
		M[key] = value
	}
}

func Week(t time.Time) string {
	w := calendar.NewWeek(t)
	var res string
	res += "-----------------------\n"

	if w.Monday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Monday.Date.Day()) + " " + w.Monday.Date.Weekday().String() + "    - " + w.Monday.Phase.String() + "\n"

	if w.Tuesday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Tuesday.Date.Day()) + " " + w.Tuesday.Date.Weekday().String() + "   - " + w.Tuesday.Phase.String() + "\n"

	if w.Wednesday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Wednesday.Date.Day()) + " " + w.Wednesday.Date.Weekday().String() + " - " + w.Wednesday.Phase.String() + "\n"

	if w.Thursday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Thursday.Date.Day()) + " " + w.Thursday.Date.Weekday().String() + "  - " + w.Thursday.Phase.String() + "\n"

	if w.Friday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Friday.Date.Day()) + " " + w.Friday.Date.Weekday().String() + "    - " + w.Friday.Phase.String() + "\n"

	if w.Saturday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Saturday.Date.Day()) + " " + w.Saturday.Date.Weekday().String() + "  - " + w.Saturday.Phase.String() + "\n"

	if w.Sunday.Date.Day() <= 9 {
		res += "0"
	}
	res += strconv.Itoa(w.Sunday.Date.Day()) + " " + w.Sunday.Date.Weekday().String() + "    - " + w.Sunday.Phase.String() + "\n"

	res += "-----------------------\n"
	return res
}

/*
	-----------------------
	01 Monday    - Месячные
	02 Tuesday   - Месячные
	03 Wednesday - Месячные
	04 Thursday  - Месячные
	05 Friday    - Месячные
	06 Saturday  -
	07 Sunday    -
	-----------------------
	Меню: (n) следующий (p) предыдущий (q) выйти
*/

func Month(t time.Time) string {
	nm := calendar.NewMonth()
	nm.SetDate(t)
	for nomer, date := nm.Date.Month(), nm.Date; nomer == date.Month(); date = date.AddDate(0, 0, 1) {
		d := calendar.NewDay()
		d.SetDate(date)
		// check date == Blocknot
		nm.AddDay(*d)
	}

	var res string
	res += nm.Date.Month().String() + " " + strconv.Itoa(nm.Date.Year()) + "\n"
	res += "-----------------------\n"
	for _, v := range nm.Days {
		var space string
		switch v.Date.Weekday() {
		case time.Monday:
			space = "    - "
		case time.Tuesday:
			space = "   - "
		case time.Wednesday:
			space = " - "
		case time.Thursday:
			space = "  - "
		case time.Friday:
			space = "    - "
		case time.Saturday:
			space = "  - "
		case time.Sunday:
			space = "    - "
		}
		if v.Date.Day() <= 9 {
			res += "0"
		}
		res += strconv.Itoa(v.Date.Day()) + " " + v.Date.Weekday().String() + space + v.Phase.String() + "\n"
	}
	res += "-----------------------\n"
	return res
}

type Note struct {
	IsPredicted bool
	Phase       calendar.Phase
}

func Set(M map[string]Note, key string, value calendar.Phase) {
	M[key] = Note{IsPredicted: false, Phase: value}
}
