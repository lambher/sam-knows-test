package models

import (
	"fmt"
	"strings"
	"time"
)

type DTime time.Time

const DTLayout = "2006-01-02"

func (dt *DTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(DTLayout, s)
	*dt = DTime(nt)
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
