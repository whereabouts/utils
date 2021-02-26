package timer

import "time"

const (
	defaultFormatLayout       = "2006-01-02 15:04:05"
	defaultFormat2MilliLayout = "2006-01-02 15:04:05.00"
)

func Format(time time.Time) string {
	return time.Format(defaultFormatLayout)
}

func Format2Milli(time time.Time) string {
	return time.Format(defaultFormat2MilliLayout)
}

func Parse(value string) (time.Time, error) {
	return time.ParseInLocation(defaultFormatLayout, value, time.Local)
}

func ParseMilli(value string) (time.Time, error) {
	return time.ParseInLocation(defaultFormat2MilliLayout, value, time.Local)
}

func Unix2Time(unix int64) time.Time {
	return time.Unix(unix, 0)
}
