/*
 * Copyright 2014-2021 Li Kexian
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
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

const (
	// defaultWhoisServer is iana whois server
	defaultWhoisServer = "whois.iana.org"
	// defaultWhoisPort is default whois port
	defaultWhoisPort = "43"
	// defaultTimeout is query default timeout
	defaultTimeout = 30 * time.Second
	// asnPrefix is asn prefix string
	asnPrefix = "AS"
)

// DefaultClient is default whois client
var DefaultClient = NewClient()

// Client is whois client
type Client struct {
	dialer  proxy.Dialer
	timeout time.Duration
	elapsed time.Duration
}

// Version returns package version
func Version() string {
	return "1.12.4"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Whois do the whois query and returns whois information
func Whois(domain string, servers ...string) (result string, err error) {
	return DefaultClient.Whois(domain, servers...)
}

// NewClient returns new whois client
func NewClient() *Client {
	return &Client{
		dialer: &net.Dialer{
			Timeout: defaultTimeout,
		},
		timeout: defaultTimeout,
	}
}

// SetDialer set query net dialer
func (c *Client) SetDialer(dialer proxy.Dialer) {
	c.dialer = dialer
}

// SetTimeout set query timeout
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Whois do the whois query and returns whois information
func (c *Client) Whois(domain string, servers ...string) (result string, err error) {
	start := time.Now()
	defer func() {
		result = fmt.Sprintf("%s\n\n;; Query time: %d msec\n;; WHEN: %s\n",
			strings.TrimRight(result, "\n"),
			time.Since(start).Milliseconds(),
			start.Format("Mon Jan 02 15:04:05 MST 2006"),
		)
	}()

	domain = strings.Trim(strings.TrimSpace(domain), ".")
	if domain == "" {
		return "", ErrDomainEmpty
	}

	isASN := IsASN(domain)
	if isASN {
		if !strings.HasPrefix(strings.ToUpper(domain), asnPrefix) {
			domain = asnPrefix + domain
		}
	}

	if !strings.Contains(domain, ".") && !strings.Contains(domain, ":") && !isASN {
		return c.rawQuery(domain, defaultWhoisServer)
	}

	var server string
	if len(servers) > 0 && servers[0] != "" {
		server = strings.ToLower(servers[0])
	} else {
		ext := getExtension(domain)
		result, err := c.rawQuery(ext, defaultWhoisServer)
		if err != nil {
			return "", fmt.Errorf("whois: query for whois server failed: %w", err)
		}
		server = getServer(result)
		if server == "" {
			return "", fmt.Errorf("%w: %s", ErrWhoisServerNotFound, domain)
		}
	}

	result, err = c.rawQuery(domain, server)
	if err != nil {
		return
	}

	refServer := getServer(result)
	if refServer == "" || refServer == server {
		return
	}

	data, err := c.rawQuery(domain, refServer)
	if err == nil {
		result += data
	}

	return
}

// rawQuery do raw query to the server
func (c *Client) rawQuery(domain, server string) (string, error) {
	c.elapsed = 0
	start := time.Now()

	if server == "whois.arin.net" {
		if IsASN(domain) {
			domain = "a + " + domain
		} else {
			domain = "n + " + domain
		}
	}

	// See: https://github.com/likexian/whois/issues/17
	if server == "whois.godaddy" {
		server = "whois.godaddy.com"
	}

	conn, err := c.dialer.Dial("tcp", net.JoinHostPort(server, defaultWhoisPort))
	if err != nil {
		return "", fmt.Errorf("whois: connect to whois server failed: %w", err)
	}

	defer conn.Close()
	c.elapsed = time.Since(start)

	_ = conn.SetWriteDeadline(time.Now().Add(c.timeout - c.elapsed))
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return "", fmt.Errorf("whois: send to whois server failed: %w", err)
	}

	c.elapsed = time.Since(start)

	_ = conn.SetReadDeadline(time.Now().Add(c.timeout - c.elapsed))
	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", fmt.Errorf("whois: read from whois server failed: %w", err)
	}

	c.elapsed = time.Since(start)

	return string(buffer), nil
}

// getExtension returns extension of domain
func getExtension(domain string) string {
	ext := domain

	if net.ParseIP(domain) == nil {
		domains := strings.Split(domain, ".")
		ext = domains[len(domains)-1]
	}

	if strings.Contains(ext, "/") {
		ext = strings.Split(ext, "/")[0]
	}

	return ext
}

// getServer returns server from whois data
func getServer(data string) string {
	tokens := []string{
		"Registrar WHOIS Server: ",
		"whois: ",
	}

	for _, token := range tokens {
		start := strings.Index(data, token)
		if start != -1 {
			start += len(token)
			end := strings.Index(data[start:], "\n")
			server := strings.TrimSpace(data[start : start+end])
			server = strings.Trim(server, "/")
			return server
		}
	}

	return ""
}

// IsASN returns if s is ASN
func IsASN(s string) bool {
	s = strings.ToUpper(s)

	s = strings.TrimPrefix(s, asnPrefix)
	_, err := strconv.Atoi(s)

	return err == nil
}
