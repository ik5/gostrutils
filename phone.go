package gostrutils

import "regexp"

var (
	israeliPhoneRegex = regexp.MustCompile("^(0[23489]2|05[69]|0[1-9]{1,6}$|0[1-9]{11,}$|[1-9])")
	e160Regex         = regexp.MustCompile(`\+?[1-9]\d{1,14}$`)
	nonDigits         = regexp.MustCompile("[^0-9]")
)

// IsIsraeliPhoneNumber check to see if a pettern that looks like Israeli phone number
// is actually an Israeli phone number
func IsIsraeliPhoneNumber(number string) bool {
	return !israeliPhoneRegex.MatchString(number)
}

// IsE164 validates if a given number is on e164 standard
func IsE164(number string) bool {
	return e160Regex.MatchString(number)
}

// CleanPhoneNumber Remove non numeric values from a phone number
func CleanPhoneNumber(number string) string {
	return nonDigits.ReplaceAllString(number, "")
}
