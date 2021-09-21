package dates

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	dt := GetNow()
	now := time.Now().UTC()
	assert.WithinDuration(t, dt, now, 10*time.Second)
}

func TestGetNowString(t *testing.T) {
	dt := GetNowString()
	now := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	assert.Contains(t, now, dt[0:15])
}

func TestGetNowDb(t *testing.T) {
	dt := GetNowDb()
	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	assert.Contains(t, now, dt[0:15])
}
