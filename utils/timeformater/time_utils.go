package timeformater

import (
	"fmt"
	"time"
)

type Time time.Time

func (t *Time) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
}
