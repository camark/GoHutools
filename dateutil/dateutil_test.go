package dateutil

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	before := time.Now()
	result := Now()
	after := time.Now()
	if result.Before(before) || result.After(after) {
		t.Errorf("Now() = %v, expected between %v and %v", result, before, after)
	}
}

func TestToday(t *testing.T) {
	result := Today()
	now := time.Now()
	if result.Year() != now.Year() || result.Month() != now.Month() || result.Day() != now.Day() {
		t.Errorf("Today() = %v, expected date %v", result, now.Format(FmtDate))
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 || result.Nanosecond() != 0 {
		t.Errorf("Today() should have zero time, got %v", result)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		layout string
		value  string
		want   time.Time
	}{
		{FmtDate, "2024-01-15", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)},
		{FmtDateTime, "2024-01-15 10:30:45", time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)},
	}
	for _, tt := range tests {
		got, err := Parse(tt.layout, tt.value)
		if err != nil {
			t.Errorf("Parse(%q, %q) error: %v", tt.layout, tt.value, err)
			continue
		}
		if !got.Equal(tt.want) {
			t.Errorf("Parse(%q, %q) = %v, want %v", tt.layout, tt.value, got, tt.want)
		}
	}
}

func TestParseDateTime(t *testing.T) {
	result, err := ParseDateTime("2024-01-15 10:30:45")
	if err != nil {
		t.Fatalf("ParseDateTime() error: %v", err)
	}
	if result.Year() != 2024 || result.Month() != 1 || result.Day() != 15 {
		t.Errorf("ParseDateTime() date = %v, expected 2024-01-15", result)
	}
	if result.Hour() != 10 || result.Minute() != 30 || result.Second() != 45 {
		t.Errorf("ParseDateTime() time = %v, expected 10:30:45", result)
	}

	_, err = ParseDateTime("invalid")
	if err == nil {
		t.Error("ParseDateTime('invalid') should return error")
	}
}

func TestParseDate(t *testing.T) {
	result, err := ParseDate("2024-01-15")
	if err != nil {
		t.Fatalf("ParseDate() error: %v", err)
	}
	if result.Year() != 2024 || result.Month() != 1 || result.Day() != 15 {
		t.Errorf("ParseDate() = %v, expected 2024-01-15", result)
	}
}

func TestParseISO8601(t *testing.T) {
	result, err := ParseISO8601("2024-01-15T10:30:45Z")
	if err != nil {
		t.Fatalf("ParseISO8601() error: %v", err)
	}
	if result.Year() != 2024 || result.Month() != 1 || result.Day() != 15 {
		t.Errorf("ParseISO8601() date = %v, expected 2024-01-15", result)
	}
	if result.Hour() != 10 || result.Minute() != 30 || result.Second() != 45 {
		t.Errorf("ParseISO8601() time = %v, expected 10:30:45", result)
	}
}

func TestFormat(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := Format(ti, FmtDateTime)
	if result != "2024-01-15 10:30:45" {
		t.Errorf("Format() = %q, expected %q", result, "2024-01-15 10:30:45")
	}
}

func TestFormatDateTime(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := FormatDateTime(ti)
	if result != "2024-01-15 10:30:45" {
		t.Errorf("FormatDateTime() = %q, expected %q", result, "2024-01-15 10:30:45")
	}
}

func TestFormatDate(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := FormatDate(ti)
	if result != "2024-01-15" {
		t.Errorf("FormatDate() = %q, expected %q", result, "2024-01-15")
	}
}

func TestFormatTime(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	result := FormatTime(ti)
	if result != "10:30:45" {
		t.Errorf("FormatTime() = %q, expected %q", result, "10:30:45")
	}
}

func TestYear(t *testing.T) {
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if Year(ti) != 2024 {
		t.Errorf("Year() = %d, expected 2024", Year(ti))
	}
}

func TestMonth(t *testing.T) {
	ti := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
	if Month(ti) != 3 {
		t.Errorf("Month() = %d, expected 3", Month(ti))
	}
}

func TestDay(t *testing.T) {
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if Day(ti) != 15 {
		t.Errorf("Day() = %d, expected 15", Day(ti))
	}
}

func TestHour(t *testing.T) {
	ti := time.Date(2024, 1, 15, 14, 30, 0, 0, time.UTC)
	if Hour(ti) != 14 {
		t.Errorf("Hour() = %d, expected 14", Hour(ti))
	}
}

func TestMinute(t *testing.T) {
	ti := time.Date(2024, 1, 15, 14, 45, 0, 0, time.UTC)
	if Minute(ti) != 45 {
		t.Errorf("Minute() = %d, expected 45", Minute(ti))
	}
}

func TestSecond(t *testing.T) {
	ti := time.Date(2024, 1, 15, 14, 45, 30, 0, time.UTC)
	if Second(ti) != 30 {
		t.Errorf("Second() = %d, expected 30", Second(ti))
	}
}

func TestDayOfWeek(t *testing.T) {
	// 2024-01-15 is Monday (1)
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if DayOfWeek(ti) != 1 {
		t.Errorf("DayOfWeek() = %d, expected 1 (Monday)", DayOfWeek(ti))
	}
	// 2024-01-14 is Sunday (0)
	ti = time.Date(2024, 1, 14, 0, 0, 0, 0, time.UTC)
	if DayOfWeek(ti) != 0 {
		t.Errorf("DayOfWeek() = %d, expected 0 (Sunday)", DayOfWeek(ti))
	}
}

func TestDayOfYear(t *testing.T) {
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if DayOfYear(ti) != 15 {
		t.Errorf("DayOfYear() = %d, expected 15", DayOfYear(ti))
	}
	ti = time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	if DayOfYear(ti) != 366 { // 2024 is a leap year
		t.Errorf("DayOfYear() = %d, expected 366", DayOfYear(ti))
	}
}

func TestWeekOfYear(t *testing.T) {
	ti := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	week := WeekOfYear(ti)
	if week < 1 || week > 53 {
		t.Errorf("WeekOfYear() = %d, expected 1-53", week)
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		want bool
	}{
		{2000, true},
		{2024, true},
		{1900, false},
		{2023, false},
		{2100, false},
		{2400, true},
	}
	for _, tt := range tests {
		if got := IsLeapYear(tt.year); got != tt.want {
			t.Errorf("IsLeapYear(%d) = %v, want %v", tt.year, got, tt.want)
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		year  int
		month time.Month
		want  int
	}{
		{2024, 1, 31},
		{2024, 2, 29}, // leap year
		{2023, 2, 28},
		{2024, 4, 30},
		{2024, 12, 31},
	}
	for _, tt := range tests {
		if got := DaysInMonth(tt.year, tt.month); got != tt.want {
			t.Errorf("DaysInMonth(%d, %d) = %d, want %d", tt.year, tt.month, got, tt.want)
		}
	}
}

func TestAddYears(t *testing.T) {
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	result := AddYears(ti, 2)
	if result.Year() != 2026 {
		t.Errorf("AddYears() year = %d, expected 2026", result.Year())
	}
	result = AddYears(ti, -1)
	if result.Year() != 2023 {
		t.Errorf("AddYears() year = %d, expected 2023", result.Year())
	}
}

func TestAddMonths(t *testing.T) {
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	result := AddMonths(ti, 3)
	if result.Month() != 4 || result.Year() != 2024 {
		t.Errorf("AddMonths() = %v, expected 2024-04-15", result)
	}
	result = AddMonths(ti, 13)
	if result.Month() != 2 || result.Year() != 2025 {
		t.Errorf("AddMonths() = %v, expected 2025-02-15", result)
	}
}

func TestAddDays(t *testing.T) {
	ti := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	result := AddDays(ti, 10)
	if result.Day() != 25 {
		t.Errorf("AddDays() day = %d, expected 25", result.Day())
	}
	result = AddDays(ti, -10)
	if result.Day() != 5 {
		t.Errorf("AddDays() day = %d, expected 5", result.Day())
	}
}

func TestAddHours(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	result := AddHours(ti, 5)
	if result.Hour() != 15 {
		t.Errorf("AddHours() hour = %d, expected 15", result.Hour())
	}
	result = AddHours(ti, -5)
	if result.Hour() != 5 {
		t.Errorf("AddHours() hour = %d, expected 5", result.Hour())
	}
}

func TestAddMinutes(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	result := AddMinutes(ti, 45)
	if result.Hour() != 11 || result.Minute() != 15 {
		t.Errorf("AddMinutes() = %v, expected 11:15", result)
	}
}

func TestAddSeconds(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
	result := AddSeconds(ti, 90)
	if result.Minute() != 31 || result.Second() != 30 {
		t.Errorf("AddSeconds() = %v, expected 10:31:30", result)
	}
}

func TestBetweenMs(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 15, 10, 0, 1, 0, time.UTC)
	if BetweenMs(t1, t2) != 1000 {
		t.Errorf("BetweenMs() = %d, expected 1000", BetweenMs(t1, t2))
	}
}

func TestBetweenDays(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	if BetweenDays(t1, t2) != 5 {
		t.Errorf("BetweenDays() = %d, expected 5", BetweenDays(t1, t2))
	}
}

func TestBetweenMonths(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 4, 15, 0, 0, 0, 0, time.UTC)
	if BetweenMonths(t1, t2) != 3 {
		t.Errorf("BetweenMonths() = %d, expected 3", BetweenMonths(t1, t2))
	}
	t2 = time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	if BetweenMonths(t1, t2) != 12 {
		t.Errorf("BetweenMonths() = %d, expected 12", BetweenMonths(t1, t2))
	}
}

func TestBetweenYears(t *testing.T) {
	t1 := time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if BetweenYears(t1, t2) != 4 {
		t.Errorf("BetweenYears() = %d, expected 4", BetweenYears(t1, t2))
	}
}

func TestIsSameDay(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 15, 20, 0, 0, 0, time.UTC)
	t3 := time.Date(2024, 1, 16, 10, 0, 0, 0, time.UTC)
	if !IsSameDay(t1, t2) {
		t.Error("IsSameDay() should return true for same day")
	}
	if IsSameDay(t1, t3) {
		t.Error("IsSameDay() should return false for different day")
	}
}

func TestIsSameMonth(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
	if !IsSameMonth(t1, t2) {
		t.Error("IsSameMonth() should return true for same month")
	}
	if IsSameMonth(t1, t3) {
		t.Error("IsSameMonth() should return false for different month")
	}
}

func TestIsSameYear(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	t3 := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	if !IsSameYear(t1, t2) {
		t.Error("IsSameYear() should return true for same year")
	}
	if IsSameYear(t1, t3) {
		t.Error("IsSameYear() should return false for different year")
	}
}

func TestIsBefore(t *testing.T) {
	t1 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	if !IsBefore(t1, t2) {
		t.Error("IsBefore() should return true")
	}
	if IsBefore(t2, t1) {
		t.Error("IsBefore() should return false")
	}
}

func TestIsAfter(t *testing.T) {
	t1 := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	if !IsAfter(t1, t2) {
		t.Error("IsAfter() should return true")
	}
	if IsAfter(t2, t1) {
		t.Error("IsAfter() should return false")
	}
}

func TestIsBetween(t *testing.T) {
	start := time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 1, 20, 0, 0, 0, 0, time.UTC)
	mid := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	before := time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC)
	after := time.Date(2024, 1, 25, 0, 0, 0, 0, time.UTC)
	if !IsBetween(mid, start, end) {
		t.Error("IsBetween() should return true for middle date")
	}
	if !IsBetween(start, start, end) {
		t.Error("IsBetween() should return true for start date (inclusive)")
	}
	if !IsBetween(end, start, end) {
		t.Error("IsBetween() should return true for end date (inclusive)")
	}
	if IsBetween(before, start, end) {
		t.Error("IsBetween() should return false for date before range")
	}
	if IsBetween(after, start, end) {
		t.Error("IsBetween() should return false for date after range")
	}
}

func TestBeginOfDay(t *testing.T) {
	ti := time.Date(2024, 1, 15, 14, 30, 45, 123456789, time.UTC)
	result := BeginOfDay(ti)
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 || result.Nanosecond() != 0 {
		t.Errorf("BeginOfDay() = %v, expected zero time", result)
	}
	if result.Year() != 2024 || result.Month() != 1 || result.Day() != 15 {
		t.Errorf("BeginOfDay() date changed: %v", result)
	}
}

func TestEndOfDay(t *testing.T) {
	ti := time.Date(2024, 1, 15, 14, 30, 45, 0, time.UTC)
	result := EndOfDay(ti)
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfDay() = %v, expected 23:59:59", result)
	}
	if result.Nanosecond() != 999999999 {
		t.Errorf("EndOfDay() nanoseconds = %d, expected 999999999", result.Nanosecond())
	}
}

func TestBeginOfWeek(t *testing.T) {
	// 2024-01-17 is Wednesday
	ti := time.Date(2024, 1, 17, 14, 30, 0, 0, time.UTC)
	result := BeginOfWeek(ti)
	// Should be Monday 2024-01-15
	if result.Weekday() != time.Monday {
		t.Errorf("BeginOfWeek() weekday = %v, expected Monday", result.Weekday())
	}
	if result.Day() != 15 {
		t.Errorf("BeginOfWeek() day = %d, expected 15", result.Day())
	}
}

func TestEndOfWeek(t *testing.T) {
	// 2024-01-17 is Wednesday
	ti := time.Date(2024, 1, 17, 14, 30, 0, 0, time.UTC)
	result := EndOfWeek(ti)
	// Should be Sunday 2024-01-21
	if result.Weekday() != time.Sunday {
		t.Errorf("EndOfWeek() weekday = %v, expected Sunday", result.Weekday())
	}
	if result.Day() != 21 {
		t.Errorf("EndOfWeek() day = %d, expected 21", result.Day())
	}
}

func TestBeginOfMonth(t *testing.T) {
	ti := time.Date(2024, 3, 15, 14, 30, 0, 0, time.UTC)
	result := BeginOfMonth(ti)
	if result.Day() != 1 || result.Month() != 3 || result.Year() != 2024 {
		t.Errorf("BeginOfMonth() = %v, expected 2024-03-01", result)
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("BeginOfMonth() should have zero time, got %v", result)
	}
}

func TestEndOfMonth(t *testing.T) {
	ti := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC)
	result := EndOfMonth(ti)
	if result.Day() != 29 { // 2024 is leap year
		t.Errorf("EndOfMonth() day = %d, expected 29", result.Day())
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfMonth() time = %v, expected 23:59:59", result)
	}
}

func TestBeginOfYear(t *testing.T) {
	ti := time.Date(2024, 6, 15, 14, 30, 0, 0, time.UTC)
	result := BeginOfYear(ti)
	if result.Year() != 2024 || result.Month() != 1 || result.Day() != 1 {
		t.Errorf("BeginOfYear() = %v, expected 2024-01-01", result)
	}
	if result.Hour() != 0 || result.Minute() != 0 || result.Second() != 0 {
		t.Errorf("BeginOfYear() should have zero time, got %v", result)
	}
}

func TestEndOfYear(t *testing.T) {
	ti := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	result := EndOfYear(ti)
	if result.Year() != 2024 || result.Month() != 12 || result.Day() != 31 {
		t.Errorf("EndOfYear() = %v, expected 2024-12-31", result)
	}
	if result.Hour() != 23 || result.Minute() != 59 || result.Second() != 59 {
		t.Errorf("EndOfYear() time = %v, expected 23:59:59", result)
	}
}

func TestAge(t *testing.T) {
	birthday := time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
	age := Age(birthday)
	now := time.Now()
	expected := now.Year() - 1990
	if now.Month() < 6 || (now.Month() == 6 && now.Day() < 15) {
		expected--
	}
	if age != expected {
		t.Errorf("Age() = %d, expected %d", age, expected)
	}
}

func TestAgeAt(t *testing.T) {
	birthday := time.Date(1990, 6, 15, 0, 0, 0, 0, time.UTC)
	at := time.Date(2024, 6, 14, 0, 0, 0, 0, time.UTC)
	if AgeAt(birthday, at) != 33 {
		t.Errorf("AgeAt() = %d, expected 33", AgeAt(birthday, at))
	}
	at = time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)
	if AgeAt(birthday, at) != 34 {
		t.Errorf("AgeAt() = %d, expected 34", AgeAt(birthday, at))
	}
	at = time.Date(2024, 6, 16, 0, 0, 0, 0, time.UTC)
	if AgeAt(birthday, at) != 34 {
		t.Errorf("AgeAt() = %d, expected 34", AgeAt(birthday, at))
	}
}

func TestZodiac(t *testing.T) {
	tests := []struct {
		year int
		want string
	}{
		{2024, "龙"},
		{2023, "兔"},
		{2022, "虎"},
		{2021, "牛"},
		{2020, "鼠"},
		{2019, "猪"},
	}
	for _, tt := range tests {
		if got := Zodiac(tt.year); got != tt.want {
			t.Errorf("Zodiac(%d) = %q, want %q", tt.year, got, tt.want)
		}
	}
}

func TestConstellation(t *testing.T) {
	tests := []struct {
		month int
		day   int
		want  string
	}{
		{1, 1, "摩羯座"},
		{1, 20, "水瓶座"},
		{2, 19, "双鱼座"},
		{3, 21, "白羊座"},
		{4, 20, "金牛座"},
		{5, 21, "双子座"},
		{6, 22, "巨蟹座"},
		{7, 23, "狮子座"},
		{8, 23, "处女座"},
		{9, 23, "天秤座"},
		{10, 24, "天蝎座"},
		{11, 22, "射手座"},
		{12, 22, "摩羯座"},
	}
	for _, tt := range tests {
		if got := Constellation(tt.month, tt.day); got != tt.want {
			t.Errorf("Constellation(%d, %d) = %q, want %q", tt.month, tt.day, got, tt.want)
		}
	}
}

func TestToMillisecond(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	ms := ToMillisecond(ti)
	if ms != 1705314645000 {
		t.Errorf("ToMillisecond() = %d, expected 1705314645000", ms)
	}
}

func TestFromMillisecond(t *testing.T) {
	ms := int64(1705314645000)
	ti := FromMillisecond(ms)
	if ti.Year() != 2024 || ti.Month() != 1 || ti.Day() != 15 {
		t.Errorf("FromMillisecond() date = %v, expected 2024-01-15", ti)
	}
}

func TestToUnix(t *testing.T) {
	ti := time.Date(2024, 1, 15, 10, 30, 45, 0, time.UTC)
	unix := ToUnix(ti)
	if unix != 1705314645 {
		t.Errorf("ToUnix() = %d, expected 1705314645", unix)
	}
}

func TestFromUnix(t *testing.T) {
	sec := int64(1705314645)
	ti := FromUnix(sec)
	if ti.Year() != 2024 || ti.Month() != 1 || ti.Day() != 15 {
		t.Errorf("FromUnix() date = %v, expected 2024-01-15", ti)
	}
}

func TestElapsed(t *testing.T) {
	ti := time.Now().Add(-5 * time.Second)
	elapsed := Elapsed(ti)
	if elapsed < 4*time.Second || elapsed > 6*time.Second {
		t.Errorf("Elapsed() = %v, expected around 5s", elapsed)
	}
}

func TestFriendlyTime(t *testing.T) {
	tests := []struct {
		name   string
		locale string
		diff   time.Duration
		want   string
	}{
		{"just now zh", "zh", 30 * time.Second, "刚刚"},
		{"1 min zh", "zh", 2 * time.Minute, "2分钟前"},
		{"1 hour zh", "zh", 3 * time.Hour, "3小时前"},
		{"1 day zh", "zh", 5 * 24 * time.Hour, "5天前"},
		{"just now en", "en", 30 * time.Second, "just now"},
		{"2 min en", "en", 2 * time.Minute, "2 minutes ago"},
		{"3 hours en", "en", 3 * time.Hour, "3 hours ago"},
		{"5 days en", "en", 5 * 24 * time.Hour, "5 days ago"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ti := time.Now().Add(-tt.diff)
			if got := FriendlyTime(ti, tt.locale); got != tt.want {
				t.Errorf("FriendlyTime() = %q, want %q", got, tt.want)
			}
		})
	}
}
