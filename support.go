// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "strings"

// test if the input string is blank
// return <trim value, if blank>
func testBlank(value string) (string, bool) {
	newVal := strings.TrimSpace(value)
	return newVal, len(newVal) == 0
}
