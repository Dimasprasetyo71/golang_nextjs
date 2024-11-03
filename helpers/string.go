package helpers

import "strings"

// ToUpperCase converts a string to uppercase
func ToUpperCase(s string) string {
    return strings.ToUpper(s)
}

// ToLowerCase converts a string to lowercase
func ToLowerCase(s string) string {
    return strings.ToLower(s)
}

// TrimSpaces removes leading and trailing spaces from a string
func TrimSpaces(s string) string {
    return strings.TrimSpace(s)
}