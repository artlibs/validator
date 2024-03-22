package validator

import "regexp"

var moneyRegexp = regexp.MustCompile("^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$")

func IsValidMoney(value string, allowBlank bool) bool {
	if isBlank(value) {
		return allowBlank
	}

	return moneyRegexp.MatchString(value)
}
