package dateutil

import (
	"strconv"
	"time"
)

// Common date formats
const (
	FmtDateTime     = "2006-01-02 15:04:05"
	FmtDate         = "2006-01-02"
	FmtTime         = "15:04:05"
	FmtDateTimeZone = "2006-01-02 15:04:05 -0700"
	FmtISO8601      = "2006-01-02T15:04:05Z07:00"
	FmtCNDate       = "2006年01月02日"
	FmtCNDateTime   = "2006年01月02日 15时04分05秒"
)

// Now returns current time
func Now() time.Time {
	return time.Now()
}

// Today returns today's date (zero time)
func Today() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// Parse parses string to time.Time
func Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

// ParseDateTime parses "2006-01-02 15:04:05" format
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse(FmtDateTime, s)
}

// ParseDate parses "2006-01-02" format
func ParseDate(s string) (time.Time, error) {
	return time.Parse(FmtDate, s)
}

// ParseISO8601 parses ISO8601 format
func ParseISO8601(s string) (time.Time, error) {
	return time.Parse(FmtISO8601, s)
}

// Format formats time to string
func Format(t time.Time, layout string) string {
	return t.Format(layout)
}

// FormatDateTime formats as "2006-01-02 15:04:05"
func FormatDateTime(t time.Time) string {
	return t.Format(FmtDateTime)
}

// FormatDate formats as "2006-01-02"
func FormatDate(t time.Time) string {
	return t.Format(FmtDate)
}

// FormatTime formats as "15:04:05"
func FormatTime(t time.Time) string {
	return t.Format(FmtTime)
}

// Year returns year
func Year(t time.Time) int {
	return t.Year()
}

// Month returns month (1-12)
func Month(t time.Time) int {
	return int(t.Month())
}

// Day returns day of month
func Day(t time.Time) int {
	return t.Day()
}

// Hour returns hour
func Hour(t time.Time) int {
	return t.Hour()
}

// Minute returns minute
func Minute(t time.Time) int {
	return t.Minute()
}

// Second returns second
func Second(t time.Time) int {
	return t.Second()
}

// DayOfWeek returns day of week (1=Sunday, 7=Saturday)
func DayOfWeek(t time.Time) int {
	return int(t.Weekday())
}

// DayOfYear returns day of year
func DayOfYear(t time.Time) int {
	return t.YearDay()
}

// WeekOfYear returns ISO week number
func WeekOfYear(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// IsLeapYear checks if year is leap year
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysInMonth returns days in month
func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// AddYears adds years
func AddYears(t time.Time, years int) time.Time {
	return t.AddDate(years, 0, 0)
}

// AddMonths adds months
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// AddDays adds days
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// AddHours adds hours
func AddHours(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

// AddMinutes adds minutes
func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

// AddSeconds adds seconds
func AddSeconds(t time.Time, seconds int) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

// BetweenMs returns difference in milliseconds
func BetweenMs(t1, t2 time.Time) int64 {
	return t2.Sub(t1).Milliseconds()
}

// BetweenDays returns difference in days
func BetweenDays(t1, t2 time.Time) int {
	return int(t2.Sub(t1).Hours() / 24)
}

// BetweenMonths returns difference in months
func BetweenMonths(t1, t2 time.Time) int {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return (y2-y1)*12 + int(m2-m1)
}

// BetweenYears returns difference in years
func BetweenYears(t1, t2 time.Time) int {
	return t2.Year() - t1.Year()
}

// IsSameDay checks if two times are on the same day
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsSameMonth checks if two times are in the same month
func IsSameMonth(t1, t2 time.Time) bool {
	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()
	return y1 == y2 && m1 == m2
}

// IsSameYear checks if two times are in the same year
func IsSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

// IsBefore checks if t1 is before t2
func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// IsAfter checks if t1 is after t2
func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// IsBetween checks if t is between start and end
func IsBetween(t, start, end time.Time) bool {
	return (t.Equal(start) || t.After(start)) && (t.Equal(end) || t.Before(end))
}

// BeginOfDay returns beginning of day
func BeginOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns end of day (23:59:59.999999999)
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// BeginOfWeek returns beginning of week (Monday)
func BeginOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return BeginOfDay(t.AddDate(0, 0, -weekday+1))
}

// EndOfWeek returns end of week (Sunday)
func EndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return EndOfDay(t.AddDate(0, 0, 7-weekday))
}

// BeginOfMonth returns beginning of month
func BeginOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns end of month
func EndOfMonth(t time.Time) time.Time {
	return EndOfDay(BeginOfMonth(t).AddDate(0, 1, -1))
}

// BeginOfYear returns beginning of year
func BeginOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns end of year
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

// Age calculates age from birthday
func Age(birthday time.Time) int {
	return AgeAt(birthday, time.Now())
}

// AgeAt calculates age at specific time
func AgeAt(birthday, at time.Time) int {
	age := at.Year() - birthday.Year()
	if at.Month() < birthday.Month() || (at.Month() == birthday.Month() && at.Day() < birthday.Day()) {
		age--
	}
	return age
}

// Zodiac returns Chinese zodiac sign
func Zodiac(year int) string {
	zodiacs := []string{"猴", "鸡", "狗", "猪", "鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊"}
	return zodiacs[year%12]
}

// Constellation returns constellation (星座)
func Constellation(month int, day int) string {
	days := []int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 22, 22}
	constellations := []string{
		"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座",
		"巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座",
	}
	if day < days[month-1] {
		return constellations[month-1]
	}
	return constellations[month%12]
}

// ToMillisecond converts time to millisecond timestamp
func ToMillisecond(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// FromMillisecond creates time from millisecond timestamp
func FromMillisecond(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}

// ToUnix converts time to unix timestamp
func ToUnix(t time.Time) int64 {
	return t.Unix()
}

// FromUnix creates time from unix timestamp
func FromUnix(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// Elapsed returns elapsed time since t
func Elapsed(t time.Time) time.Duration {
	return time.Since(t)
}

// FriendlyTime returns friendly time description
func FriendlyTime(t time.Time, locale string) string {
	now := time.Now()
	diff := now.Sub(t)

	if diff < 0 {
		if locale == "zh" {
			return "未来"
		}
		return "in the future"
	}

	seconds := int(diff.Seconds())
	minutes := int(diff.Minutes())
	hours := int(diff.Hours())
	days := int(diff.Hours() / 24)

	if locale == "zh" {
		switch {
		case seconds < 60:
			return "刚刚"
		case minutes < 60:
			return strconv.Itoa(minutes) + "分钟前"
		case hours < 24:
			return strconv.Itoa(hours) + "小时前"
		case days < 30:
			return strconv.Itoa(days) + "天前"
		case days < 365:
			return strconv.Itoa(days/30) + "个月前"
		default:
			return strconv.Itoa(days/365) + "年前"
		}
	}

	switch {
	case seconds < 60:
		return "just now"
	case minutes < 60:
		if minutes == 1 {
			return "1 minute ago"
		}
		return strconv.Itoa(minutes) + " minutes ago"
	case hours < 24:
		if hours == 1 {
			return "1 hour ago"
		}
		return strconv.Itoa(hours) + " hours ago"
	case days < 30:
		if days == 1 {
			return "1 day ago"
		}
		return strconv.Itoa(days) + " days ago"
	case days < 365:
		months := days / 30
		if months == 1 {
			return "1 month ago"
		}
		return strconv.Itoa(months) + " months ago"
	default:
		years := days / 365
		if years == 1 {
			return "1 year ago"
		}
		return strconv.Itoa(years) + " years ago"
	}
}
