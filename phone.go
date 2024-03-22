// @Since 2024-03-21.
// @Author Fury, All rights Reserved.

package validator

import "regexp"

// https://github.com/VincentSit/ChinaMobilePhoneNumberRegex/blob/master/README-CN.md
const basicRegexp string = "^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[235-8]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[0-35-9]\\d{2}|66\\d{2})\\d{6}$"
const virtualRegexp string = "^(?:\\+?86)?1(?:7[01]|6[257])\\d{8}$"
const netOnlyRegexp string = "^(?:\\+?86)?14[579]\\d{8}$"
const iotOnlyRegexp string = "^(?:\\+?86)?14(?:[14]0|41|[68]\\d)\\d{9}$"
const allTypeRegexp string = "^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[0-35-9]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[0-35-9]\\d{2}|6[2567]\\d{2}|4(?:(?:10|4[01])\\d{3}|[68]\\d{4}|[579]\\d{2}))\\d{6}$"

type PhoneType int

const (
	// Basic 11位手机卡-基础运营商,支持语音通话/短信/数据流量
	Basic PhoneType = iota
	// Virtual 11位手机卡-虚拟运营商,支持语音通话/短信/数据流量
	Virtual
	// NetOnly 11位上网卡,支持语音通话(部分)/短信/数据流量
	NetOnly
	// IotOnly 13位物联网数据卡,支持数据流量
	IotOnly
)

func getRegexp(phoneType PhoneType) *regexp.Regexp {
	switch phoneType {
	case Basic:
		return regexp.MustCompile(basicRegexp)
	case Virtual:
		return regexp.MustCompile(virtualRegexp)
	case NetOnly:
		return regexp.MustCompile(netOnlyRegexp)
	case IotOnly:
		return regexp.MustCompile(iotOnlyRegexp)
	default: // all type, 11或13位,支持所有号码
		return regexp.MustCompile(allTypeRegexp)
	}
}

// IsValidPhone 手机号码
// Validate is the value a valid phone number
func IsValidPhone(value string, allowBlank bool, types []PhoneType) bool {
	if isBlank(value) {
		return allowBlank
	}

	if len(types) == 0 {
		types = []PhoneType{Basic, Virtual, NetOnly}
	}

	for _, phoneType := range types {
		if getRegexp(phoneType).MatchString(value) {
			return true
		}
	}

	return false
}
