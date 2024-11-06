package helpers

import (
    "strconv"
    "time"
)

func ParseDate(dateStr, layout string) (time.Time, error) {
    return time.Parse(layout, dateStr)
}

func ParseInt(intStr string) (int, error) {
    return strconv.Atoi(intStr)
}