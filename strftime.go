package timex

import (
	"strings"
	"time"
)

var fmt2layout = map[string]string{
	"%a": StdWeekDay,
	"%A": StdLongWeekDay,
	// "%w":
	"%d":  StdZeroDay,
	"%-d": StdDay,
	"%b":  StdMonth,
	"%B":  StdLongMonth,
	"%m":  StdZeroMonth,
	"%-m": StdNumMonth,
	"%y":  StdYear,
	"%Y":  StdLongYear,
	"%H":  StdHour,
	// "%-H":
	"%I":  StdZeroHour12,
	"%-I": StdHour12,
	"%p":  StdPM,
	"%M":  StdZeroMinute,
	"%-M": StdMinute,
	"%S":  StdZeroSecond,
	"%-S": StdSecond,
	// "%f":
	"%z": StdNumTZ,
	"%Z": StdTZ,
	// "%j":
	// "%-j":
	// "%U":
	// "%W":
	// "%c":
	// "%x":
	// "%X":
	"%%": "%",
}

// Strftime formats "time.Time" in the same way as the strftime Python function
func Strftime(fmt string, t time.Time) string {
	for _fmt, _layout := range fmt2layout {
		fmt = strings.Replace(fmt, _fmt, _layout, -1)
	}
	return t.Format(fmt)
}
