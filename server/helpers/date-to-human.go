package helpers

import (
	"strconv"
	"time"
)

// DateToHuman converts a date to human readable time
func DateToHuman(d time.Time) string {
	year, month, day, hour, min, _ := diff(d, time.Now())
	dHumand := "few seconds ago"
	if year == 1 {
		dHumand = "A year ago"
	} else if year >= 1 {
		dHumand = strconv.Itoa(year) + " years ago"
	} else if month == 1 {
		dHumand = "A months ago"
	} else if month > 1 {
		dHumand = strconv.Itoa(month) + " months ago"
	} else if day == 1 {
		dHumand = "A day ago"
	} else if day > 1 {
		dHumand = strconv.Itoa(day) + " days ago"
	} else if hour == 1 {
		dHumand = "An hour ago"
	} else if hour > 1 {
		dHumand = strconv.Itoa(hour) + " hours ago"
	} else if min == 1 {
		dHumand = "A minute ago"
	} else if min >= 1 {
		dHumand = strconv.Itoa(min) + " minutes ago"
	}
	return dHumand
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
