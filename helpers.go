package gochan

import (
	"bytes"
	"strconv"
	"time"
)

// jsonBool allows 0/1 and "0"/"1" and "true"/"false" (strings) to also become boolean
type jsonBool bool

func (bit *jsonBool) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))
	*bit = jsonBool(txt == "1" || txt == "true")
	return nil
}

// jsonTime parses UNIX timestamps and converts them to time.Time
type jsonTime time.Time

func (t *jsonTime) UnmarshalJSON(b []byte) error {
	txt := string(bytes.Trim(b, `"`))

	// convert to int64
	i, err := strconv.ParseInt(txt, 10, 64)
	if err != nil {
		return err
	}

	// convert to time.Time
	*t = jsonTime(time.Unix(i, 0))
	return nil
}
