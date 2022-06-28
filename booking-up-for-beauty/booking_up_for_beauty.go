package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)

	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	d, _ := time.Parse("January 2, 2006 15:04:05", date)

	return d.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	d, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return d.Hour() >= 12 && d.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	d := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", d.Weekday(), d.Month(), d.Day(), d.Year(), d.Hour(), d.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(2022, time.September, 15, 00, 00, 00, 00, time.UTC)
}
