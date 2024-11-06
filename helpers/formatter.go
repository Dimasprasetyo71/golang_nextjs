// helpers/formatter.go
package helpers

import (
    "fmt"
    "time"
)

func FormatDate(t time.Time, layout string) string {
    return t.Format(layout)
}

func FormatCurrency(amount int) string {
    return fmt.Sprintf("Rp %d", amount) 
}