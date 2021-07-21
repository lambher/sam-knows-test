package models

import (
	"testing"
	"time"
)

func TestNewDTime(t *testing.T) {
	testValues := []struct {
		Str      string
		Expected time.Time
	}{
		{
			Str:      "2006-05-11",
			Expected: time.Date(2006, 05, 11, 0, 0, 0, 0, time.UTC),
		},
		{
			Str:      "1993-06-20",
			Expected: time.Date(1993, 06, 20, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, testValue := range testValues {
		dtime := NewDTime(testValue.Str)
		if !dtime.ToTime().Equal(testValue.Expected) {
			t.Errorf("dtime is incorrect, got: %s, want: %s.", dtime, testValue.Expected)
		}
	}
}
