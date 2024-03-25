// @Since 2024-03-22.
// @Author Leo, All rights Reserved.

package validator

import "regexp"

var moneyRegexp = regexp.MustCompile("^-?([0-9]+|[1-9][0-9]{0,2}(,[0-9]{3})*)(.[0-9]{1,4})?$")

// IsValidMoney 金钱
// allowEmpty: return true if id is nil or ""
// Validate is the money a valid money
func IsValidMoney(money string, allowEmpty bool) bool {
	if value, ok := testNotBlank(money); ok {
		return moneyRegexp.MatchString(value)
	}

	return allowEmpty && isEmpty(money)
}
