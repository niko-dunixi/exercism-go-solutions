// Leap package contains leap year related functions
package leap

// IsLeapYear will return true or false if a given year is a leap-year.
func IsLeapYear(year int) bool {
	return (year%4 == 0) && (year%100 != 0 || year%400 == 0)
}
