---
title: dateutil - 日期工具
---

# dateutil - 日期工具

日期时间工具包，提供了丰富的日期时间操作函数，包括解析、格式化、计算、判断等功能。

## 导入

```go
import "github.com/camark/GoHutools/dateutil"
```

## 常量

```go
const (
    FmtDateTime     = "2006-01-02 15:04:05"
    FmtDate         = "2006-01-02"
    FmtTime         = "15:04:05"
    FmtDateTimeZone = "2006-01-02 15:04:05 -0700"
    FmtISO8601      = "2006-01-02T15:04:05Z07:00"
    FmtCNDate       = "2006年01月02日"
    FmtCNDateTime   = "2006年01月02日 15时04分05秒"
)
```

## 函数列表

### Now

```go
func Now() time.Time
```

返回当前时间。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.FormatDateTime(now))  // "2026-08-14 12:00:00"
```

### Today

```go
func Today() time.Time
```

返回今天的日期（时间部分为零）。

**示例:**

```go
today := dateutil.Today()
fmt.Println(dateutil.FormatDate(today))  // "2026-08-14"
```

### Parse

```go
func Parse(layout, value string) (time.Time, error)
```

解析字符串为 time.Time。

**示例:**

```go
t, err := dateutil.Parse("2006-01-02", "2026-08-14")
fmt.Println(t, err)  // 2026-08-14 00:00:00 +0000 UTC <nil>
```

### ParseDateTime

```go
func ParseDateTime(s string) (time.Time, error)
```

解析 "2006-01-02 15:04:05" 格式的字符串。

**示例:**

```go
t, err := dateutil.ParseDateTime("2026-08-14 12:30:00")
fmt.Println(t, err)
```

### ParseDate

```go
func ParseDate(s string) (time.Time, error)
```

解析 "2006-01-02" 格式的字符串。

**示例:**

```go
t, err := dateutil.ParseDate("2026-08-14")
fmt.Println(t, err)
```

### ParseISO8601

```go
func ParseISO8601(s string) (time.Time, error)
```

解析 ISO8601 格式的字符串。

**示例:**

```go
t, err := dateutil.ParseISO8601("2026-08-14T12:30:00Z")
fmt.Println(t, err)
```

### Format

```go
func Format(t time.Time, layout string) string
```

格式化时间。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Format(now, "2006/01/02 15:04"))  // "2026/08/14 12:30"
```

### FormatDateTime

```go
func FormatDateTime(t time.Time) string
```

格式化为 "2006-01-02 15:04:05" 格式。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.FormatDateTime(now))  // "2026-08-14 12:30:00"
```

### FormatDate

```go
func FormatDate(t time.Time) string
```

格式化为 "2006-01-02" 格式。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.FormatDate(now))  // "2026-08-14"
```

### FormatTime

```go
func FormatTime(t time.Time) string
```

格式化为 "15:04:05" 格式。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.FormatTime(now))  // "12:30:00"
```

### Year

```go
func Year(t time.Time) int
```

返回年份。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Year(now))  // 2026
```

### Month

```go
func Month(t time.Time) int
```

返回月份（1-12）。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Month(now))  // 8
```

### Day

```go
func Day(t time.Time) int
```

返回日期。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Day(now))  // 14
```

### Hour

```go
func Hour(t time.Time) int
```

返回小时。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Hour(now))  // 12
```

### Minute

```go
func Minute(t time.Time) int
```

返回分钟。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Minute(now))  // 30
```

### Second

```go
func Second(t time.Time) int
```

返回秒。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.Second(now))  // 0
```

### DayOfWeek

```go
func DayOfWeek(t time.Time) int
```

返回星期几（0=周日, 1=周一, ..., 6=周六）。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.DayOfWeek(now))  // 4 (周四)
```

### DayOfYear

```go
func DayOfYear(t time.Time) int
```

返回一年中的第几天。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.DayOfYear(now))  // 226
```

### WeekOfYear

```go
func WeekOfYear(t time.Time) int
```

返回 ISO 周数。

**示例:**

```go
now := dateutil.Now()
fmt.Println(dateutil.WeekOfYear(now))  // 33
```

### IsLeapYear

```go
func IsLeapYear(year int) bool
```

检查是否为闰年。

**示例:**

```go
fmt.Println(dateutil.IsLeapYear(2024))  // true
fmt.Println(dateutil.IsLeapYear(2025))  // false
```

### DaysInMonth

```go
func DaysInMonth(year int, month time.Month) int
```

返回指定月份的天数。

**示例:**

```go
fmt.Println(dateutil.DaysInMonth(2026, time.February))  // 28
fmt.Println(dateutil.DaysInMonth(2024, time.February))  // 29
```

### AddYears

```go
func AddYears(t time.Time, years int) time.Time
```

增加年份。

**示例:**

```go
now := dateutil.Now()
future := dateutil.AddYears(now, 1)
fmt.Println(dateutil.FormatDate(future))  // "2027-08-14"
```

### AddMonths

```go
func AddMonths(t time.Time, months int) time.Time
```

增加月份。

**示例:**

```go
now := dateutil.Now()
future := dateutil.AddMonths(now, 3)
fmt.Println(dateutil.FormatDate(future))  // "2026-11-14"
```

### AddDays

```go
func AddDays(t time.Time, days int) time.Time
```

增加天数。

**示例:**

```go
now := dateutil.Now()
future := dateutil.AddDays(now, 7)
fmt.Println(dateutil.FormatDate(future))  // "2026-08-21"
```

### AddHours

```go
func AddHours(t time.Time, hours int) time.Time
```

增加小时。

**示例:**

```go
now := dateutil.Now()
future := dateutil.AddHours(now, 2)
fmt.Println(dateutil.FormatDateTime(future))
```

### AddMinutes

```go
func AddMinutes(t time.Time, minutes int) time.Time
```

增加分钟。

**示例:**

```go
now := dateutil.Now()
future := dateutil.AddMinutes(now, 30)
fmt.Println(dateutil.FormatDateTime(future))
```

### AddSeconds

```go
func AddSeconds(t time.Time, seconds int) time.Time
```

增加秒数。

**示例:**

```go
now := dateutil.Now()
future := dateutil.AddSeconds(now, 60)
fmt.Println(dateutil.FormatDateTime(future))
```

### BetweenMs

```go
func BetweenMs(t1, t2 time.Time) int64
```

返回两个时间之间的毫秒差。

**示例:**

```go
t1, _ := dateutil.ParseDateTime("2026-08-14 10:00:00")
t2, _ := dateutil.ParseDateTime("2026-08-14 10:00:01")
fmt.Println(dateutil.BetweenMs(t1, t2))  // 1000
```

### BetweenDays

```go
func BetweenDays(t1, t2 time.Time) int
```

返回两个时间之间的天数差。

**示例:**

```go
t1, _ := dateutil.ParseDate("2026-08-01")
t2, _ := dateutil.ParseDate("2026-08-14")
fmt.Println(dateutil.BetweenDays(t1, t2))  // 13
```

### BetweenMonths

```go
func BetweenMonths(t1, t2 time.Time) int
```

返回两个时间之间的月数差。

**示例:**

```go
t1, _ := dateutil.ParseDate("2026-01-01")
t2, _ := dateutil.ParseDate("2026-08-01")
fmt.Println(dateutil.BetweenMonths(t1, t2))  // 7
```

### BetweenYears

```go
func BetweenYears(t1, t2 time.Time) int
```

返回两个时间之间的年数差。

**示例:**

```go
t1, _ := dateutil.ParseDate("2020-01-01")
t2, _ := dateutil.ParseDate("2026-01-01")
fmt.Println(dateutil.BetweenYears(t1, t2))  // 6
```

### IsSameDay

```go
func IsSameDay(t1, t2 time.Time) bool
```

检查两个时间是否在同一天。

**示例:**

```go
t1, _ := dateutil.ParseDateTime("2026-08-14 10:00:00")
t2, _ := dateutil.ParseDateTime("2026-08-14 15:00:00")
fmt.Println(dateutil.IsSameDay(t1, t2))  // true
```

### IsSameMonth

```go
func IsSameMonth(t1, t2 time.Time) bool
```

检查两个时间是否在同一月。

**示例:**

```go
t1, _ := dateutil.ParseDate("2026-08-01")
t2, _ := dateutil.ParseDate("2026-08-31")
fmt.Println(dateutil.IsSameMonth(t1, t2))  // true
```

### IsSameYear

```go
func IsSameYear(t1, t2 time.Time) bool
```

检查两个时间是否在同一年。

**示例:**

```go
t1, _ := dateutil.ParseDate("2026-01-01")
t2, _ := dateutil.ParseDate("2026-12-31")
fmt.Println(dateutil.IsSameYear(t1, t2))  // true
```

### IsBefore

```go
func IsBefore(t1, t2 time.Time) bool
```

检查 t1 是否在 t2 之前。

**示例:**

```go
t1, _ := dateutil.ParseDate("2026-08-01")
t2, _ := dateutil.ParseDate("2026-08-14")
fmt.Println(dateutil.IsBefore(t1, t2))  // true
```

### IsAfter

```go
func IsAfter(t1, t2 time.Time) bool
```

检查 t1 是否在 t2 之后。

**示例:**

```go
t1, _ := dateutil.ParseDate("2026-08-14")
t2, _ := dateutil.ParseDate("2026-08-01")
fmt.Println(dateutil.IsAfter(t1, t2))  // true
```

### IsBetween

```go
func IsBetween(t, start, end time.Time) bool
```

检查时间是否在范围内。

**示例:**

```go
start, _ := dateutil.ParseDate("2026-08-01")
end, _ := dateutil.ParseDate("2026-08-31")
t, _ := dateutil.ParseDate("2026-08-14")
fmt.Println(dateutil.IsBetween(t, start, end))  // true
```

### BeginOfDay

```go
func BeginOfDay(t time.Time) time.Time
```

返回当天的开始时间（00:00:00）。

**示例:**

```go
now := dateutil.Now()
begin := dateutil.BeginOfDay(now)
fmt.Println(dateutil.FormatDateTime(begin))  // "2026-08-14 00:00:00"
```

### EndOfDay

```go
func EndOfDay(t time.Time) time.Time
```

返回当天的结束时间（23:59:59.999999999）。

**示例:**

```go
now := dateutil.Now()
end := dateutil.EndOfDay(now)
fmt.Println(dateutil.FormatDateTime(end))  // "2026-08-14 23:59:59"
```

### BeginOfWeek

```go
func BeginOfWeek(t time.Time) time.Time
```

返回本周的开始时间（周一 00:00:00）。

**示例:**

```go
now := dateutil.Now()
begin := dateutil.BeginOfWeek(now)
fmt.Println(dateutil.FormatDate(begin))
```

### EndOfWeek

```go
func EndOfWeek(t time.Time) time.Time
```

返回本周的结束时间（周日 23:59:59.999999999）。

**示例:**

```go
now := dateutil.Now()
end := dateutil.EndOfWeek(now)
fmt.Println(dateutil.FormatDate(end))
```

### BeginOfMonth

```go
func BeginOfMonth(t time.Time) time.Time
```

返回本月的开始时间。

**示例:**

```go
now := dateutil.Now()
begin := dateutil.BeginOfMonth(now)
fmt.Println(dateutil.FormatDate(begin))  // "2026-08-01"
```

### EndOfMonth

```go
func EndOfMonth(t time.Time) time.Time
```

返回本月的结束时间。

**示例:**

```go
now := dateutil.Now()
end := dateutil.EndOfMonth(now)
fmt.Println(dateutil.FormatDate(end))  // "2026-08-31"
```

### BeginOfYear

```go
func BeginOfYear(t time.Time) time.Time
```

返回本年的开始时间。

**示例:**

```go
now := dateutil.Now()
begin := dateutil.BeginOfYear(now)
fmt.Println(dateutil.FormatDate(begin))  // "2026-01-01"
```

### EndOfYear

```go
func EndOfYear(t time.Time) time.Time
```

返回本年的结束时间。

**示例:**

```go
now := dateutil.Now()
end := dateutil.EndOfYear(now)
fmt.Println(dateutil.FormatDate(end))  // "2026-12-31"
```

### Age

```go
func Age(birthday time.Time) int
```

根据生日计算年龄。

**示例:**

```go
birthday, _ := dateutil.ParseDate("1990-05-15")
fmt.Println(dateutil.Age(birthday))  // 36
```

### AgeAt

```go
func AgeAt(birthday, at time.Time) int
```

计算在指定时间点的年龄。

**示例:**

```go
birthday, _ := dateutil.ParseDate("1990-05-15")
at, _ := dateutil.ParseDate("2026-08-14")
fmt.Println(dateutil.AgeAt(birthday, at))  // 36
```

### Zodiac

```go
func Zodiac(year int) string
```

返回生肖。

**示例:**

```go
fmt.Println(dateutil.Zodiac(2026))  // "马"
fmt.Println(dateutil.Zodiac(2024))  // "龙"
```

### Constellation

```go
func Constellation(month int, day int) string
```

返回星座。

**示例:**

```go
fmt.Println(dateutil.Constellation(8, 14))  // "狮子座"
fmt.Println(dateutil.Constellation(12, 25)) // "摩羯座"
```

### ToMillisecond

```go
func ToMillisecond(t time.Time) int64
```

转换为毫秒时间戳。

**示例:**

```go
now := dateutil.Now()
ms := dateutil.ToMillisecond(now)
fmt.Println(ms)
```

### FromMillisecond

```go
func FromMillisecond(ms int64) time.Time
```

从毫秒时间戳创建时间。

**示例:**

```go
ms := int64(1692000000000)
t := dateutil.FromMillisecond(ms)
fmt.Println(dateutil.FormatDateTime(t))
```

### ToUnix

```go
func ToUnix(t time.Time) int64
```

转换为 Unix 时间戳。

**示例:**

```go
now := dateutil.Now()
ts := dateutil.ToUnix(now)
fmt.Println(ts)
```

### FromUnix

```go
func FromUnix(sec int64) time.Time
```

从 Unix 时间戳创建时间。

**示例:**

```go
ts := int64(1692000000)
t := dateutil.FromUnix(ts)
fmt.Println(dateutil.FormatDateTime(t))
```

### Elapsed

```go
func Elapsed(t time.Time) time.Duration
```

返回从 t 到现在的时间间隔。

**示例:**

```go
past, _ := dateutil.ParseDateTime("2026-08-14 10:00:00")
duration := dateutil.Elapsed(past)
fmt.Println(duration)  // 如 "2h30m0s"
```

### FriendlyTime

```go
func FriendlyTime(t time.Time, locale string) string
```

返回友好的时间描述。

**示例:**

```go
past, _ := dateutil.ParseDateTime("2026-08-14 10:00:00")
fmt.Println(dateutil.FriendlyTime(past, "zh"))    // "2小时前"
fmt.Println(dateutil.FriendlyTime(past, "en"))     // "2 hours ago"
```
