/*
 * Go module for domain whois
 * https://www.likexian.com/
 *
 * Copyright 2014-2018, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package whois

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestWhois(t *testing.T) {
	result, err := Whois("likexian", true)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "", result)

	result, err = Whois("likexian.com", true)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", result)

	result, err = Whois("likexian.com", true, "127.0.0.1")
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "", result)

	result, err = Whois("likexian.com", true, "com.whois-servers.net")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", result)

	result, err = Whois("likexian.com.cn", false)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", result)
}
