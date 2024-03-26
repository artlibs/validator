// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

// (?i) 定义了不区分大小写
var httpUrlRegexp = regexp.MustCompile("^(?i)https?://(www\\.)?[-A-Z0-9@:%._+~#=]{1,256}\\.[A-Z0-9()]{1,6}\\b([-A-Z0-9()@:%_+.~#?&/=]*)")

// IsValidHttpURL HTTP(S) URL
// allowEmpty: return true if url is nil or ""
// Validate is the url a valid http url
func IsValidHttpURL(url string, allowEmpty bool) bool {
	if value, ok := testNotBlank(url); ok {
		return httpUrlRegexp.MatchString(value)
	}

	return allowEmpty && isEmpty(url)
}
