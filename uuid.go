package gostrutils

import "regexp"

var (
	uuidRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
)

// IsValidUUID validate if a given string is a uuid.
// On error it will return false.
//
// The code is not mine, and was taken from: https://stackoverflow.com/questions/25051675/how-to-validate-uuid-v4-in-go#25051754
func IsValidUUID(uuid string) bool {
	return uuidRegex.MatchString(uuid)
}
