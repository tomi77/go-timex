package timex

import (
	"fmt"
	"strings"
	"time"
)

// ParseDuration parses a duration string in format "h:mm:ss".
func ParseDuration(s string) (time.Duration, error) {
	parts := strings.Split(s, ":")
	s = fmt.Sprintf("%sh%sm%ss", parts[0], parts[1], parts[2])
	d, _ := time.ParseDuration(s)
	return d, nil
}

// FormatDuration returns a string representing the duration in the form "h:mm:ss".
func FormatDuration(d time.Duration) string {
	hours := int64(d.Hours())
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}
