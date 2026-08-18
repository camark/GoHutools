package dateutil

import (
	"testing"
	"time"
)

// fixed Friday 2024-03-15 14:30:45.123456789
func fixed() time.Time {
	return time.Date(2024, 3, 15, 14, 30, 45, 123456789, time.UTC)
}

func TestMinuteHourBoundaries(t *testing.T) {
	ts := fixed()
	// BeginOfMinute
	if got := BeginOfMinute(ts); got != time.Date(2024, 3, 15, 14, 30, 0, 0, time.UTC) {
		t.Errorf("BeginOfMinute = %v", got)
	}
	// EndOfMinute
	if got := EndOfMinute(ts); got != time.Date(2024, 3, 15, 14, 30, 59, 999999999, time.UTC) {
		t.Errorf("EndOfMinute = %v", got)
	}
	// BeginOfHour
	if got := BeginOfHour(ts); got != time.Date(2024, 3, 15, 14, 0, 0, 0, time.UTC) {
		t.Errorf("BeginOfHour = %v", got)
	}
	// EndOfHour
	if got := EndOfHour(ts); got != time.Date(2024, 3, 15, 14, 59, 59, 999999999, time.UTC) {
		t.Errorf("EndOfHour = %v", got)
	}
}

func TestQuarter(t *testing.T) {
	if got := QuarterOfYear(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)); got != 1 {
		t.Errorf("QuarterOfYear(Jan) = %d", got)
	}
	if got := QuarterOfYear(fixed()); got != 1 {
		t.Errorf("QuarterOfYear(Mar) = %d", got)
	}
	if got := QuarterOfYear(time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)); got != 2 {
		t.Errorf("QuarterOfYear(Jun) = %d", got)
	}
	if got := QuarterOfYear(time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC)); got != 3 {
		t.Errorf("QuarterOfYear(Sep) = %d", got)
	}
	if got := QuarterOfYear(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)); got != 4 {
		t.Errorf("QuarterOfYear(Dec) = %d", got)
	}
	// BeginOfQuarter for Feb → Jan 1, Aug → Jul 1
	if got := BeginOfQuarter(time.Date(2024, 2, 20, 12, 0, 0, 0, time.UTC)); got != time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("BeginOfQuarter(Feb) = %v", got)
	}
	if got := BeginOfQuarter(time.Date(2024, 8, 20, 12, 0, 0, 0, time.UTC)); got != time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("BeginOfQuarter(Aug) = %v", got)
	}
	// EndOfQuarter for Mar → Mar 31, Aug → Sep 30
	if got := EndOfQuarter(fixed()); got != time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.UTC) {
		t.Errorf("EndOfQuarter(Mar) = %v", got)
	}
	if got := EndOfQuarter(time.Date(2024, 8, 20, 12, 0, 0, 0, time.UTC)); got != time.Date(2024, 9, 30, 23, 59, 59, 999999999, time.UTC) {
		t.Errorf("EndOfQuarter(Aug) = %v", got)
	}
	// leap-year Feb quarter end
	if got := EndOfQuarter(time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC)); got != time.Date(2024, 3, 31, 23, 59, 59, 999999999, time.UTC) {
		t.Errorf("EndOfQuarter(Feb 2024) = %v", got)
	}
}

func TestIsSameMinuteHourSecond(t *testing.T) {
	a := time.Date(2024, 3, 15, 14, 30, 45, 0, time.UTC)
	b := time.Date(2024, 3, 15, 14, 30, 46, 0, time.UTC)
	c := time.Date(2024, 3, 15, 14, 31, 45, 0, time.UTC)
	d := time.Date(2024, 3, 15, 15, 30, 45, 0, time.UTC)
	if !IsSameMinute(a, b) || IsSameMinute(a, c) {
		t.Error("IsSameMinute")
	}
	if !IsSameHour(a, c) || IsSameHour(a, d) {
		t.Error("IsSameHour")
	}
	if IsSameSecond(a, b) {
		t.Error("IsSameSecond")
	}
	if !IsSameSecond(a, a) {
		t.Error("IsSameSecond same value")
	}
}

func TestIsWeekend(t *testing.T) {
	if IsWeekend(time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)) { // Friday
		t.Error("Friday should not be weekend")
	}
	if !IsWeekend(time.Date(2024, 3, 16, 0, 0, 0, 0, time.UTC)) { // Saturday
		t.Error("Saturday should be weekend")
	}
	if !IsWeekend(time.Date(2024, 3, 17, 0, 0, 0, 0, time.UTC)) { // Sunday
		t.Error("Sunday should be weekend")
	}
}

func TestIsOverlap(t *testing.T) {
	sec := func(v int) time.Time { return time.Unix(int64(v), 0) } // small epoch for readability
	// touching endpoints count as overlap
	if !IsOverlap(sec(1), sec(3), sec(3), sec(5)) {
		t.Error("touching should overlap")
	}
	// partial overlap
	if !IsOverlap(sec(2), sec(4), sec(3), sec(5)) {
		t.Error("partial should overlap")
	}
	// containment
	if !IsOverlap(sec(1), sec(5), sec(2), sec(3)) {
		t.Error("contained should overlap")
	}
	// disjoint
	if IsOverlap(sec(1), sec(2), sec(3), sec(4)) {
		t.Error("disjoint should not overlap")
	}
}

func TestBetweenMore(t *testing.T) {
	base := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	if got := BetweenSeconds(base, base.Add(20*time.Second)); got != 20 {
		t.Errorf("BetweenSeconds = %d", got)
	}
	if got := BetweenMinutes(base, base.Add(90*time.Minute)); got != 90 {
		t.Errorf("BetweenMinutes = %d", got)
	}
	if got := BetweenHours(base, base.Add(5*time.Hour)); got != 5 {
		t.Errorf("BetweenHours = %d", got)
	}
	// Monday Jan 1 → Jan 8 = 1 week; → Jan 15 = 2 weeks
	if got := BetweenWeeks(base, time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)); got != 1 {
		t.Errorf("BetweenWeeks = %d", got)
	}
	if got := BetweenWeeks(base, time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)); got != 2 {
		t.Errorf("BetweenWeeks(2) = %d", got)
	}
}

func TestDateOf(t *testing.T) {
	d := DateOf(2024, 3, 15)
	if d.Year() != 2024 || d.Month() != 3 || d.Day() != 15 {
		t.Errorf("DateOf = %v", d)
	}
	if d.Hour() != 0 || d.Minute() != 0 {
		t.Errorf("DateOf time part = %v", d)
	}
	dt := DateTimeOf(2024, 12, 31, 23, 59, 58)
	if dt.Year() != 2024 || dt.Second() != 58 {
		t.Errorf("DateTimeOf = %v", dt)
	}
}

func TestTruncateTo(t *testing.T) {
	ts := fixed()
	sec := time.Date(2024, 3, 15, 14, 30, 45, 123456789, time.UTC)

	if got := TruncateTo(sec, TruncSecond); got != time.Date(2024, 3, 15, 14, 30, 45, 0, time.UTC) {
		t.Errorf("TruncateTo(Second) = %v", got)
	}
	if got := TruncateTo(sec, TruncMinute); got != time.Date(2024, 3, 15, 14, 30, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Minute) = %v", got)
	}
	if got := TruncateTo(sec, TruncHour); got != time.Date(2024, 3, 15, 14, 0, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Hour) = %v", got)
	}
	if got := TruncateTo(sec, TruncDay); got != time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Day) = %v", got)
	}
	// week truncation goes back to Monday (2024-03-15 is a Friday)
	if got := TruncateTo(sec, TruncWeek); got != time.Date(2024, 3, 11, 0, 0, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Week) = %v", got)
	}
	if got := TruncateTo(sec, TruncMonth); got != time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Month) = %v", got)
	}
	if got := TruncateTo(sec, TruncYear); got != time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Year) = %v", got)
	}
	if got := TruncateTo(sec, TruncQuarter); got != time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("TruncateTo(Quarter) = %v", got)
	}
	if got := TruncateTo(ts, TruncSecond); got.Nanosecond() != 0 {
		t.Errorf("TruncateTo(exact) should drop nanos: %v", got)
	}
}

func TestFormatChinese(t *testing.T) {
	d := time.Date(2024, 5, 6, 9, 8, 7, 0, time.UTC)
	if got := FormatChineseDate(d); got != "二〇二四年五月六日" {
		t.Errorf("FormatChineseDate = %q", got)
	}
	if got := FormatChineseDateTime(d); got != "二〇二四年五月六日 9时8分7秒" {
		t.Errorf("FormatChineseDateTime = %q", got)
	}
	// teen/tens handling
	if FormatChineseNumber(10) != "十" || FormatChineseNumber(15) != "十五" ||
		FormatChineseNumber(21) != "二十一" || FormatChineseNumber(30) != "三十" {
		t.Error("FormatChineseNumber tens/teens")
	}
	if FormatChineseNumber(100) != "一百" {
		t.Error("FormatChineseNumber(100)")
	}
}

func TestCurrentMs(t *testing.T) {
	now := time.Now()
	if got := Current(); got < now.UnixMilli()-1 || got > now.UnixMilli()+50 {
		t.Errorf("Current() = %d, now=%d", got, now.UnixMilli())
	}
}