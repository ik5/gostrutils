package gostrutils

import "regexp"

// IsIsraeliPhoneNumber check to see if a pettern that looks like Israeli phone number
// is actually an Israeli phone number
func IsIsraeliPhoneNumber(number string) bool {
	re := regexp.MustCompile("^(0[23489]2|05[69]|0[1-9]{1,6}$|0[1-9]{11,}$|[1-9])")
	return re.MatchString(number) == false
}

// IsE164 validates if a given number is on e164 standard
func IsE164(number string) bool {
	re := regexp.MustCompile("^\\+?[1-9]\\d{1,14}$")
	return re.MatchString(number)
}

// CleanPhoneNumber Remove non numeric values from a phone number
func CleanPhoneNumber(number string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(number, "")
}
