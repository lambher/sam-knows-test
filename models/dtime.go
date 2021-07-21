package models

import (
	"fmt"
	"strings"
	"time"
)

type DTime time.Time

const DTLayout = "2006-01-02"

func NewDTime(date string) DTime {
	nt, err := time.Parse(DTLayout, date)
	if err != nil {
		return DTime{}
	}
	return DTime(nt)
}

func (dt *DTime) UnmarshalJSON(b []byte) (err error) {
	*dt = NewDTime(strings.Trim(string(b), `"`))
	return
}

func (dt DTime) MarshalJSON() ([]byte, error) {
	return []byte(dt.String()), nil
}

func (dt DTime) String() string {
	t := time.Time(dt)
	return strings.Trim(fmt.Sprintf("%q", t.Format(DTLayout)), `"`)
}

func (dt *DTime) ToTime() time.Time {
	return time.Time(*dt)
}
