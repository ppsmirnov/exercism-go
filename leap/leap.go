package leap

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// IsLeapYear is a function, that checks weather year is leap
func IsLeapYear(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}
