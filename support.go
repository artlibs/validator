// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "strings"

// check if the input string is blank
func isBlank(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}
