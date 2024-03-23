// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import (
	"fmt"
	"regexp"
	"unicode"
)

var id15Regexp = regexp.MustCompile(
	"^" +
		"\\d{6}" + // 6位地区码
		"\\d{2}" + // 2位年份
		"((0[1-9])|(10|11|12))" + // 2位月份
		"(([0-2][1-9])|10|20|30|31)" + // 2位日期
		"\\d{3}" + // 3位顺序码
		"$")
var id18Regexp = regexp.MustCompile(
	"^" +
		"\\d{6}" + // 6位地区码
		"(18|19|([23]\\d))\\d{2}" + // 4位年份
		"((0[1-9])|(10|11|12))" + // 2位月份
		"(([0-2][1-9])|10|20|30|31)" + // 2位日期
		"\\d{3}" + // 3位顺序码
		"[0-9Xx]" + // 1位校验码
		"$")

// IsValidIdCard 身份证
// Validate is the value a valid id card number
func IsValidIdCard(idCardNo string, allowBlank bool) bool {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	if value, ok := testBlank(idCardNo); ok {
		if len(value) == 15 {
			return verify(upgrade(value))
		}

		return verify(value)
	}

	return allowBlank
}

func verify(value string) bool {
	return id18Regexp.MatchString(value) && rune(value[17]) == idCvc(value)
}

func upgrade(value string) string {
	if id15Regexp.MatchString(value) {
		nv := value[:6] + "19" + value[6:]
		return nv + string(idCvc(nv))
	}

	panic("Illegal argument of value: " + value)
}

func idCvc(value string) rune {
	ida := []rune(value)
	cca := []rune{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	wa := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}

	sum := 0
	for i := 0; i < 17; i++ {
		if !unicode.IsDigit(ida[i]) {
			panic("Illegal argument of value: " + value)
		}
		sum = sum + (int(ida[i])-int('0'))*wa[i]
	}

	return cca[sum%11]
}
