package helpers

import (
    "strconv"
    "time"
)

// ParseDate parses a string to time.Time with the specified layout
func ParseDate(dateStr, layout string) (time.Time, error) {
    return time.Parse(layout, dateStr)
}

// ParseInt parses a string to an integer
func ParseInt(intStr string) (int, error) {
    return strconv.Atoi(intStr)
}