// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

// (?i) 定义了不区分大小写
var httpUrlRegexp = regexp.MustCompile("(?i)^https?://(?:www\\.)?[\\w.-]+\\.[a-zA-Z]{2,}(?:/\\S*)?$")

// IsValidHttpURL HTTP(S) URL
// Validate is the value a valid phone number
func IsValidHttpURL(url string, allowBlank bool) bool {
	if value, ok := testBlank(url); ok {
		return httpUrlRegexp.MatchString(value)
	}

	return allowBlank
}
