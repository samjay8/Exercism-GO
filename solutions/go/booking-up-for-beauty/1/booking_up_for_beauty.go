package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{
		"1/2/2006 15:04:05",
		"January 2, 2006 15:04:05",
		"Monday, January 2, 2006 15:04:05", // Included for common Exercism test variations
	}

	for _, layout := range layouts {
		parsedate, err := time.Parse(layout, date)
		if err == nil {
			return parsedate // Return immediately when a layout successfully parses
		}
	}

	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Directly returning the condition cleans up redundant if-else branches
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	hour := Schedule(date).Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedate := Schedule(date)

	return parsedate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}
