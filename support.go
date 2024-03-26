// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "strings"

// test if the input string is blank
// return <trim value, if not blank>
func testNotBlank(value string) (string, bool) {
	newVal := strings.TrimSpace(value)
	return newVal, len(newVal) > 0
}

func isEmpty(value string) bool {
	return len(value) == 0
}
