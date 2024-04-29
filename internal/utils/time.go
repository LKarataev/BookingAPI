package utils

import (
	"time"
)

func DaysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DatesToStr(dates []time.Time) string {
	res := ""
	for _, date := range dates {
		res += date.Format("2006/01/02") + ", "
	}
	if len(res) >= 2 {
		res = res[:len(res)-2]
	}
	return res
}
