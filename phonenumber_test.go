// @Since 2024-03-24.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPhone(t *testing.T) {
	// test empty
	pts := []PhoneNumberType{BasicPhoneNumber, VirtualPhoneNumber, NetOnlyPhoneNumber, IotOnlyPhoneNumber, AllTypePhoneNumber}
	for _, pt := range pts {
		assert.True(t, IsValidPhoneNumber("", true, pt))
		assert.False(t, IsValidPhoneNumber("", false, pt))
	}

	// test phones
	caseMap := map[string]map[PhoneNumberType]bool{
		"18157610011": { // valid basic
			BasicPhoneNumber:   true,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: true,
		},
		"16577810902": { // valid virtual
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: true,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: true,
		},
		"14577812563": { // valid net
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: true,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: true,
		},
		"1440775619012": { // valid iot
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: true,
			AllTypePhoneNumber: true,
		},
		"123456": { // all invalid
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: false,
		},
		"181778189012": { // all invalid
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: false,
		},
		"1817781890123": { // all invalid
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: false,
		},
		"181ab7781cd": { // all invalid
			BasicPhoneNumber:   false,
			VirtualPhoneNumber: false,
			NetOnlyPhoneNumber: false,
			IotOnlyPhoneNumber: false,
			AllTypePhoneNumber: false,
		},
	}
	for phone, kv := range caseMap {
		for pt, expect := range kv {
			assert.Equal(t, expect, IsValidPhoneNumber(phone, false, pt))
			assert.Equal(t, expect, IsValidPhoneNumber(phone, true, pt))
		}
	}

}
