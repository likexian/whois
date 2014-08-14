/*
 * Go module for domain whois
 * http://www.likexian.com/
 *
 * Copyright 2014, Kexian Li
 * Released under the Apache License, Version 2.0
 *
 */

package whois


import (
    "testing"
    "github.com/bmizerany/assert"
)


func TestWhois(t *testing.T) {
    result, err := Whois("likexian")
    assert.NotEqual(t, nil, err)
    assert.Equal(t, "", result)

    result, err = Whois("likexian.com")
    assert.Equal(t, nil, err)
    assert.NotEqual(t, "", result)

    result, err = Whois("likexian.com", "127.0.0.1")
    assert.NotEqual(t, nil, err)
    assert.Equal(t, "", result)

    result, err = Whois("likexian.com", "com.whois-servers.net")
    assert.Equal(t, nil, err)
    assert.NotEqual(t, "", result)
}
