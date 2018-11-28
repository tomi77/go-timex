package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStrftime(t *testing.T) {
	d, _ := time.Parse(time.RFC3339, "2013-09-30T07:06:05Z")

	assert.Equal(t, "Mon", Strftime("%a", d))
	assert.Equal(t, "Monday", Strftime("%A", d))
	// assert.Equal(t, "1", Strftime("%w", d))
	assert.Equal(t, "30", Strftime("%d", d))
	assert.Equal(t, "30", Strftime("%-d", d))
	assert.Equal(t, "Sep", Strftime("%b", d))
	assert.Equal(t, "September", Strftime("%B", d))
	assert.Equal(t, "09", Strftime("%m", d))
	assert.Equal(t, "9", Strftime("%-m", d))
	assert.Equal(t, "13", Strftime("%y", d))
	assert.Equal(t, "2013", Strftime("%Y", d))
	assert.Equal(t, "07", Strftime("%H", d))
	// assert.Equal(t, "7", Strftime("%-H", d))
	assert.Equal(t, "07", Strftime("%I", d))
	assert.Equal(t, "7", Strftime("%-I", d))
	assert.Equal(t, "AM", Strftime("%p", d))
	assert.Equal(t, "06", Strftime("%M", d))
	assert.Equal(t, "6", Strftime("%-M", d))
	assert.Equal(t, "05", Strftime("%S", d))
	assert.Equal(t, "5", Strftime("%-S", d))
	// assert.Equal(t, "000000", Strftime("%f", d))
	assert.Equal(t, "+0000", Strftime("%z", d))
	assert.Equal(t, "UTC", Strftime("%Z", d))
	// assert.Equal(t, "273", Strftime("%j", d))
	// assert.Equal(t, "273", Strftime("%-j", d))
	// assert.Equal(t, "39", Strftime("%U", d))
	// assert.Equal(t, "39", Strftime("%W", d))
	// assert.Equal(t, "Mon Sep 30 07:06:05 2013", Strftime("%c", d))
	// assert.Equal(t, "09/30/13", Strftime("%x", d))
	// assert.Equal(t, "07:06:05", Strftime("%X", d))
	assert.Equal(t, "%", Strftime("%%", d))
}
