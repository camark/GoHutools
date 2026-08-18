package dateutil

import (
	"strconv"
	"strings"
	"time"
)

// TruncateUnit identifies a time unit for TruncateTo.
type TruncateUnit int

const (
	// TruncSecond truncates to the whole second (drops nanoseconds)
	TruncSecond TruncateUnit = iota
	// TruncMinute truncates to the whole minute
	TruncMinute
	// TruncHour truncates to the whole hour
	TruncHour
	// TruncDay truncates to midnight
	TruncDay
	// TruncWeek truncates to the Monday of the week
	TruncWeek
	// TruncMonth truncates to the 1st of the month
	TruncMonth
	// TruncQuarter truncates to the 1st of the quarter
	TruncQuarter
	// TruncYear truncates to Jan 1st
	TruncYear
)

// BeginOfMinute returns the start of the minute (second 0, zero nanos).
func BeginOfMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute returns the last instant of the minute.
func EndOfMinute(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, 999999999, t.Location())
}

// BeginOfHour returns the start of the hour.
func BeginOfHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour returns the last instant of the hour.
func EndOfHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 59, 59, 999999999, t.Location())
}

// QuarterOfYear returns the quarter (1-4) of t's month.
func QuarterOfYear(t time.Time) int {
	return (int(t.Month())-1)/3 + 1
}

// BeginOfQuarter returns the first day of the quarter containing t.
func BeginOfQuarter(t time.Time) time.Time {
	q := QuarterOfYear(t)
	firstMonth := time.Month((q - 1) * 3)
	return time.Date(t.Year(), firstMonth+1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfQuarter returns the last instant of the quarter containing t.
func EndOfQuarter(t time.Time) time.Time {
	q := QuarterOfYear(t)
	lastMonth := time.Month(q*3)
	return EndOfMonth(time.Date(t.Year(), lastMonth, 1, 0, 0, 0, 0, t.Location()))
}

// IsSameSecond reports whether a and b agree down to the second.
func IsSameSecond(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay() &&
		a.Hour() == b.Hour() && a.Minute() == b.Minute() && a.Second() == b.Second()
}

// IsSameMinute reports whether a and b fall in the same minute.
func IsSameMinute(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay() &&
		a.Hour() == b.Hour() && a.Minute() == b.Minute()
}

// IsSameHour reports whether a and b fall in the same hour.
func IsSameHour(a, b time.Time) bool {
	return a.Year() == b.Year() && a.YearDay() == b.YearDay() && a.Hour() == b.Hour()
}

// IsWeekend reports whether t falls on Saturday or Sunday.
func IsWeekend(t time.Time) bool {
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		return true
	}
	return false
}

// IsOverlap reports whether the two inclusive time ranges overlap.
// Ranges that merely touch at an endpoint are considered overlapping.
func IsOverlap(s1, e1, s2, e2 time.Time) bool {
	return !e1.Before(s2) && !e2.Before(s1)
}

// BetweenSeconds returns the signed difference t2-t1 in seconds.
func BetweenSeconds(t1, t2 time.Time) int {
	return int(t2.Sub(t1) / time.Second)
}

// BetweenMinutes returns the signed difference t2-t1 in minutes.
func BetweenMinutes(t1, t2 time.Time) int {
	return int(t2.Sub(t1) / time.Minute)
}

// BetweenHours returns the signed difference t2-t1 in hours.
func BetweenHours(t1, t2 time.Time) int {
	return int(t2.Sub(t1) / time.Hour)
}

// BetweenWeeks returns the signed difference t2-t1 in whole weeks.
func BetweenWeeks(t1, t2 time.Time) int {
	return BetweenDays(t1, t2) / 7
}

// DateOf builds a date at midnight in the local time zone.
func DateOf(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// DateTimeOf builds a timestamp at the given wall-clock fields in local time.
func DateTimeOf(year int, month time.Month, day, hour, minute, second int) time.Time {
	return time.Date(year, month, day, hour, minute, second, 0, time.Local)
}

// TruncateTo drops the parts of t finer than the given unit.
func TruncateTo(t time.Time, unit TruncateUnit) time.Time {
	switch unit {
	case TruncSecond:
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
	case TruncMinute:
		return BeginOfMinute(t)
	case TruncHour:
		return BeginOfHour(t)
	case TruncDay:
		return BeginOfDay(t)
	case TruncWeek:
		return BeginOfWeek(t)
	case TruncMonth:
		return BeginOfMonth(t)
	case TruncQuarter:
		return BeginOfQuarter(t)
	case TruncYear:
		return BeginOfYear(t)
	}
	return t
}

// Current returns the current time in Unix milliseconds.
func Current() int64 {
	return time.Now().UnixMilli()
}

var cnDigits = []string{"〇", "一", "二", "三", "四", "五", "六", "七", "八", "九"}

// FormatChineseNumber renders n (0-999) in Chinese numerals
// (e.g. 15 → 十五, 21 → 二十一, 105 → 一百零五).
func FormatChineseNumber(n int) string {
	if n < 0 {
		return "-" + FormatChineseNumber(-n)
	}
	if n < 10 {
		return cnDigits[n]
	}
	var b strings.Builder
	if n >= 100 {
		b.WriteString(cnDigits[n/100])
		b.WriteString("百")
		n %= 100
		if n == 0 {
			return b.String()
		}
		if n < 10 {
			b.WriteString("零")
		}
	}
	switch {
	case n < 10:
		b.WriteString(cnDigits[n])
	case n >= 20:
		b.WriteString(cnDigits[n/10])
		b.WriteString("十")
		if n%10 != 0 {
			b.WriteString(cnDigits[n%10])
		}
	case n >= 10:
		b.WriteString("十")
		if n%10 != 0 {
			b.WriteString(cnDigits[n%10])
		}
	}
	return b.String()
}

// FormatChineseDate renders a date like 二〇二四年五月六日.
func FormatChineseDate(t time.Time) string {
	var b strings.Builder
	y := t.Year()
	b.WriteString(cnDigits[(y/1000)%10])
	b.WriteString(cnDigits[(y/100)%10])
	b.WriteString(cnDigits[(y/10)%10])
	b.WriteString(cnDigits[y%10])
	b.WriteString("年")
	b.WriteString(FormatChineseNumber(int(t.Month())))
	b.WriteString("月")
	b.WriteString(FormatChineseNumber(t.Day()))
	b.WriteString("日")
	return b.String()
}

// FormatChineseDateTime renders date and clock like
// 二〇二四年五月六日 9时8分7秒 (following Hutool: arabic numerals
// for the clock part, Chinese numerals for the date).
func FormatChineseDateTime(t time.Time) string {
	return FormatChineseDate(t) + " " +
		strconv.Itoa(t.Hour()) + "时" +
		strconv.Itoa(t.Minute()) + "分" +
		strconv.Itoa(t.Second()) + "秒"
}