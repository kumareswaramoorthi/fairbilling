package validate

import (
	"regexp"
	"strings"
)

var IsAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]*$").MatchString
var IsTimestamp = regexp.MustCompile("^[:0-9]*$").MatchString

// Validate the line for 2 spaces, 8 chars in timestamp with proper regex and alphanumeric value in user name.
func ValidateLine(text string) (bool, []string) {
	spaces := strings.Count(text, " ")
	if spaces == 2 {
		parts := strings.Split(text, " ")
		lenTimestamp := len(parts[0])
		timeStampValid := IsTimestamp(parts[0])
		nameValid := IsAlphaNumeric(parts[1])
		if lenTimestamp == 8 && timeStampValid && nameValid {
			return true, parts
		}
	}
	return false, nil
}
