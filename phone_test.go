// @Since 2024-03-24.
// @Author Fury, All rights Reserved.

package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPhone(t *testing.T) {
	// test empty
	pts := []PhoneType{Basic, Virtual, NetOnly, IotOnly, AllType}
	for _, pt := range pts {
		assert.True(t, IsValidPhone("", true, []PhoneType{pt}))
		assert.False(t, IsValidPhone("", false, []PhoneType{pt}))
	}

	// test phones
	caseMap := map[string]map[PhoneType]bool{
		"18157610011": { // valid basic
			Basic:   true,
			Virtual: false,
			NetOnly: false,
			IotOnly: false,
			AllType: true,
		},
		"16577810902": { // valid virtual
			Basic:   false,
			Virtual: true,
			NetOnly: false,
			IotOnly: false,
			AllType: true,
		},
		"14577812563": { // valid net
			Basic:   false,
			Virtual: false,
			NetOnly: true,
			IotOnly: false,
			AllType: true,
		},
		"1440775619012": { // valid iot
			Basic:   false,
			Virtual: false,
			NetOnly: false,
			IotOnly: true,
			AllType: true,
		},
		"123456": { // all invalid
			Basic:   false,
			Virtual: false,
			NetOnly: false,
			IotOnly: false,
			AllType: false,
		},
		"181778189012": { // all invalid
			Basic:   false,
			Virtual: false,
			NetOnly: false,
			IotOnly: false,
			AllType: false,
		},
		"1817781890123": { // all invalid
			Basic:   false,
			Virtual: false,
			NetOnly: false,
			IotOnly: false,
			AllType: false,
		},
		"181ab7781cd": { // all invalid
			Basic:   false,
			Virtual: false,
			NetOnly: false,
			IotOnly: false,
			AllType: false,
		},
	}
	for phone, kv := range caseMap {
		for pt, expect := range kv {
			assert.Equal(t, expect, IsValidPhone(phone, false, []PhoneType{pt}))
			assert.Equal(t, expect, IsValidPhone(phone, true, []PhoneType{pt}))
		}
	}

}
