/*
 * Copyright 2014-2024 Li Kexian
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
	"context"
	"fmt"
	"io"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// Dialer is a means to establish a connection.
//
// It uses the same function definition as `net/Dialer.DialContext`
type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

// Client is whois client
type Client struct {
	dialer          Dialer
	timeout         time.Duration
	disableStats    bool
	disableReferral bool
}

type hasTimeout struct {
	Timeout time.Duration
	Dialer
}

// Version returns package version
func Version() string {
	return "1.15.5"
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
func (c *Client) SetDialer(dialer Dialer) *Client {
	c.dialer = dialer
	return c
}

// SetTimeout set query timeout
func (c *Client) SetTimeout(timeout time.Duration) *Client {
	if d, ok := c.dialer.(*hasTimeout); ok {
		d.Timeout = timeout
	}
	c.timeout = timeout
	return c
}

// SetDisableStats set disable stats
func (c *Client) SetDisableStats(disabled bool) *Client {
	c.disableStats = disabled
	return c
}

// SetDisableReferral if set to true, will not query the referral server.
func (c *Client) SetDisableReferral(disabled bool) *Client {
	c.disableReferral = disabled
	return c
}

// Whois do the whois query and returns whois information
func (c *Client) Whois(domain string, servers ...string) (result string, err error) {
	return c.WhoisContext(context.Background(), domain, servers...)
}

// WhoisContext do the whois query and returns whois information
func (c *Client) WhoisContext(ctx context.Context, domain string, servers ...string) (result string, err error) {
	start := time.Now()
	defer func() {
		result = strings.TrimSpace(result)
		if result != "" && !c.disableStats {
			result = fmt.Sprintf("%s\n\n%% Query time: %d msec\n%% WHEN: %s\n",
				result, time.Since(start).Milliseconds(), start.Format("Mon Jan 02 15:04:05 MST 2006"),
			)
		}
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
		return c.rawQuery(ctx, domain, defaultWhoisServer, defaultWhoisPort)
	}

	var server, port string
	if len(servers) > 0 && servers[0] != "" {
		server = strings.ToLower(servers[0])
		port = defaultWhoisPort
		reHasNumericPort := regexp.MustCompile(`:(\d+)$`)
		if matches := reHasNumericPort.FindStringSubmatch(server); matches != nil {
			server = server[:len(server)-len(matches[0])]
			port = matches[1]
		}
	} else {
		ext := getExtension(domain)
		result, err := c.rawQuery(ctx, ext, defaultWhoisServer, defaultWhoisPort)
		if err != nil {
			return "", fmt.Errorf("whois: query for whois server failed: %w", err)
		}
		server, port = getServer(result)
		if server == "" {
			return "", fmt.Errorf("%w: %s", ErrWhoisServerNotFound, domain)
		}
	}

	result, err = c.rawQuery(ctx, domain, server, port)
	if err != nil {
		return
	}

	if c.disableReferral {
		return
	}

	refServer, refPort := getServer(result)
	if refServer == "" || refServer == server {
		return
	}

	data, err := c.rawQuery(ctx, domain, refServer, refPort)
	if err == nil {
		result += data
	}

	return
}

// rawQuery do raw query to the server
func (c *Client) rawQuery(ctx context.Context, domain, server, port string) (string, error) {
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

	// See: https://github.com/likexian/whois/pull/30
	if server == "porkbun.com/whois" {
		server = "whois.porkbun.com"
	}

	conn, err := c.dialer.DialContext(ctx, "tcp", net.JoinHostPort(server, port))
	if err != nil {
		return "", fmt.Errorf("whois: connect to whois server failed: %w", err)
	}

	defer conn.Close()
	elapsed := time.Since(start)

	_ = conn.SetWriteDeadline(time.Now().Add(c.timeout - elapsed))
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		// Some servers may refuse a request with a reason, immediately closing the connection after sending.
		// For example, GoDaddy returns "Number of allowed queries exceeded.\r\n", and immediately closes the connection.
		//
		// We return both the response _and_ the error, to allow callers to try to parse the response, while
		// still letting them know an error occurred. In particular, this helps catch rate limit errors.
		buffer, _ := io.ReadAll(conn)
		if len(buffer) > 0 {
			return string(buffer), err
		}

		return "", fmt.Errorf("whois: send to whois server failed: %w", err)
	}

	elapsed = time.Since(start)

	_ = conn.SetReadDeadline(time.Now().Add(c.timeout - elapsed))
	buffer, err := io.ReadAll(conn)
	if err != nil {
		if len(buffer) > 0 {
			// Some servers may refuse a request with a reason, immediately closing the connection after sending.
			// For example, GoDaddy returns "Number of allowed queries exceeded.\r\n", and immediately closes the connection.
			//
			// We return both the response _and_ the error, to allow callers to try to parse the response, while
			// still letting them know an error occurred (potentially short reads). In particular, this helps
			// catch rate limit errors.
			return string(buffer), err
		}

		return "", fmt.Errorf("whois: read from whois server failed: %w", err)
	}

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

// getServer returns the first referral server from whois data (if any)
func getServer(data string) (string, string) {
	tokens := []string{
		"Registrar WHOIS Server: ",
		"whois: ",
		"ReferralServer: ",
		"refer: ",
		"%referral ", // e.g. %referral rwhois://root.rwhois.net:4321/auth-area=.
	}

	for _, token := range tokens {
		start := strings.Index(data, token)
		if start != -1 {
			start += len(token)
			end := strings.Index(data[start:], "\n")
			server := strings.TrimSpace(data[start : start+end])
			server = strings.TrimPrefix(server, "http:")
			server = strings.TrimPrefix(server, "https:")
			server = strings.TrimPrefix(server, "whois:")
			server = strings.TrimPrefix(server, "rwhois:")
			server = strings.Trim(server, "/")
			port := defaultWhoisPort
			if strings.Contains(server, ":") {
				v := strings.Split(server, ":")
				server, port = v[0], v[1]
				// Strip trailing non-numeric characters from port
				reNumericTrailing := regexp.MustCompile(`^(\d+)\D.*$`)
				matches := reNumericTrailing.FindStringSubmatch(port)
				if matches != nil {
					port = matches[1]
				}
			}
			return server, port
		}
	}

	return "", ""
}

// IsASN returns if s is ASN
func IsASN(s string) bool {
	s = strings.ToUpper(s)

	s = strings.TrimPrefix(s, asnPrefix)
	_, err := strconv.Atoi(s)

	return err == nil
}
