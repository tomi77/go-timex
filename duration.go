package timex

import (
	"fmt"
	"time"
)

// FormatDuration returns a string representing the duration in the form "72:03:00".
func FormatDuration(d time.Duration) string {
	hours := int64(d.Hours())
	minutes := int64(d.Minutes()) % 60
	seconds := int64(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds)
}
