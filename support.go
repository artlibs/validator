// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "strings"

// check if the input string is blank
func testBlank(value string) (string, bool) {
	newVal := strings.TrimSpace(value)
	return newVal, len(newVal) == 0
}
