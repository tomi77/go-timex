package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDuration(t *testing.T) {
	d, _ := time.ParseDuration("1h30m")
	assert.Equal(t, "1:30:00", FormatDuration(d))
	d, _ = time.ParseDuration("2h31m15s")
	assert.Equal(t, "2:31:15", FormatDuration(d))
	d, _ = time.ParseDuration("25h")
	assert.Equal(t, "25:00:00", FormatDuration(d))
}

func ExampleFormatDuration() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")

	fmt.Println(FormatDuration(hours))
	fmt.Println(FormatDuration(complex))
	// Output:
	// 10:00:00
	// 1:10:10
}
