/*
 * Copyright 2014-2020 Li Kexian
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
	"testing"

	"github.com/likexian/gokit/assert"
)

func TestVersion(t *testing.T) {
	assert.Contains(t, Version(), ".")
	assert.Contains(t, Author(), "likexian")
	assert.Contains(t, License(), "Apache License")
}

func TestWhoisFail(t *testing.T) {
	tests := []string{
		"",
		"likexian.jp?e",
		"1.1.1.1!",
	}

	for _, v := range tests {
		_, err := Whois(v)
		assert.NotNil(t, err)
	}

	_, err := Whois("likexian.com", "127.0.0.1")
	assert.NotNil(t, err)
}

func TestWhois(t *testing.T) {
	tests := []string{
		"com",
		"xxx",
		"cn",
		"google.com",
		"google.net",
		"google.org",
		"google.mobi",
		"google.cn",
		"google.com.cn",
		"google.in",
		"google.jp/e",
		"1.1.1.1",
		"2.1.1.1",
		"3.1.1.1",
		"4.1.1.1",
		"5.1.1.1",
		"2001:dc7::1",
	}

	for _, v := range tests {
		b, err := Whois(v)
		assert.Nil(t, err)
		assert.NotEqual(t, b, "")
	}

	_, err := Whois("likexian.com", "com.whois-servers.net")
	assert.Nil(t, err)
}
