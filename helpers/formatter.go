// helpers/formatter.go
package helpers

import (
    "fmt"
    "time"
)

// FormatDate formats a time.Time object to a string with the specified layout
func FormatDate(t time.Time, layout string) string {
    return t.Format(layout)
}

// FormatCurrency formats an integer to a string with currency notation
func FormatCurrency(amount int) string {
    return fmt.Sprintf("Rp %d", amount) // Format untuk Indonesia Rupiah
}