/*
 * Copyright 2014-2023 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for domain and ip whois information query
 * https://www.likexian.com/
 */

package whois

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/likexian/gokit/assert"
	"golang.org/x/net/proxy"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestClient_SetDisableReferral(t *testing.T) {
	client := NewClient()

	resp, err := client.Whois("likexian.com")
	assert.Nil(t, err)
	assert.Equal(t, strings.Count(resp, "Domain Name: LIKEXIAN.COM"), 2)

	client.SetDisableReferral(true)

	resp, err = client.Whois("likexian.com")
	assert.Nil(t, err)
	assert.Equal(t, strings.Count(resp, "Domain Name: LIKEXIAN.COM"), 1)
}

func TestWhoisFail(t *testing.T) {
	tests := []struct {
		domain string
		err    error
	}{
		{"", ErrDomainEmpty},
		{"likexian.jp?e", ErrWhoisServerNotFound},
		{"1.1.1.1!", ErrWhoisServerNotFound},
	}

	for _, v := range tests {
		_, err := Whois(v.domain)
		assert.NotNil(t, err)
		if !errors.Is(err, v.err) {
			t.Fatalf("expect %v but got %v", v.err, err)
		}
	}

	_, err := Whois("likexian.com", "127.0.0.1")
	assert.NotNil(t, err)
}

func TestWhois(t *testing.T) {
	tests := []string{
		"com",
		"xxx",
		"cn",
		"name.com",
		"name.net",
		"name.org",
		"name.mobi",
		"name.cn",
		"name.com.cn",
		"name.in",
		"name.jp/e",
		"1.1.1.1",
		"2.1.1.1",
		"3.1.1.1",
		"4.1.1.1",
		"5.1.1.1",
		"2001:dc7::1",
		"1",
		"as2",
		"as1878",
		"as4610",
		"as27648",
		"as36865",
		"172.109.217.241",
		"144.200.46.16",
	}

	for _, v := range tests {
		b, err := Whois(v)
		assert.Nil(t, err)
		assert.NotEqual(t, b, "")
	}

	_, err := Whois("likexian.com", "com.whois-servers.net")
	assert.Nil(t, err)
}

func TestNewClient(t *testing.T) {
	c := NewClient()
	var err error

	c.SetTimeout(10 * time.Microsecond)
	_, err = c.Whois("likexian.com")
	assert.NotNil(t, err)

	c.SetTimeout(10 * time.Second)
	_, err = c.Whois("likexian.com")
	assert.Nil(t, err)

	c.SetDialer(proxy.FromEnvironment())
	_, err = c.Whois("likexian.com")
	assert.Nil(t, err)
}

func TestIsASN(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"", false},
		{"a", false},
		{"ab", false},
		{"as", false},
		{"ab1", false},
		{"as1a", false},
		{"as1", true},
		{"As1", true},
		{"AS1", true},
		{"AS123", true},
		{"1", true},
		{"123", true},
	}

	for _, v := range tests {
		assert.Equal(t, IsASN(v.in), v.out)
	}
}
