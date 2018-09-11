package util

import "time"

// TimeFormatDms3Fs is the format dms3fs uses to represent time in string form.
var TimeFormatDms3Fs = time.RFC3339Nano

// ParseRFC3339 parses an RFC3339Nano-formatted time stamp and
// returns the UTC time.
func ParseRFC3339(s string) (time.Time, error) {
	t, err := time.Parse(TimeFormatDms3Fs, s)
	if err != nil {
		return time.Time{}, err
	}
	return t.UTC(), nil
}

// FormatRFC3339 returns the string representation of the
// UTC value of the given time in RFC3339Nano format.
func FormatRFC3339(t time.Time) string {
	return t.UTC().Format(TimeFormatDms3Fs)
}
