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
)

const (
	// IANA_WHOIS_SERVER is iana whois server
	IANA_WHOIS_SERVER = "whois.iana.org"
	// DEFAULT_WHOIS_PORT is default whois port
	DEFAULT_WHOIS_PORT = "43"
)

// Version returns package version
func Version() string {
	return "1.9.0"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Licensed under the Apache License 2.0"
}

// Whois do the whois query and returns whois info
func Whois(domain string, servers ...string) (result string, err error) {
	domain = strings.Trim(strings.TrimSpace(domain), ".")
	if domain == "" {
		return "", fmt.Errorf("whois: domain is empty")
	}

	isASN := IsASN(domain)
	if isASN {
		if !strings.HasPrefix(strings.ToUpper(domain), "AS") {
			domain = "AS" + domain
		}
	}

	if !strings.Contains(domain, ".") && !strings.Contains(domain, ":") && !isASN {
		return query(domain, IANA_WHOIS_SERVER)
	}

	var server string
	if len(servers) == 0 || servers[0] == "" {
		ext := getExtension(domain)
		result, err := query(ext, IANA_WHOIS_SERVER)
		if err != nil {
			return "", fmt.Errorf("whois: query for whois server failed: %v", err)
		}
		server = getServer(result)
		if server == "" {
			return "", fmt.Errorf("whois: no whois server found for domain: %s", domain)
		}
	} else {
		server = strings.ToLower(servers[0])
	}

	result, err = query(domain, server)
	if err != nil {
		return
	}

	refServer := getServer(result)
	if refServer == "" || refServer == server {
		return
	}

	data, err := query(domain, refServer)
	if err == nil {
		result += data
	}

	return
}

// query send query to server
func query(domain, server string) (string, error) {
	if server == "whois.arin.net" {
		if IsASN(domain) {
			domain = "a + " + domain
		} else {
			domain = "n + " + domain
		}
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, DEFAULT_WHOIS_PORT), time.Second*30)
	if err != nil {
		return "", fmt.Errorf("whois: connect to whois server failed: %v", err)
	}

	defer conn.Close()
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 30))
	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return "", fmt.Errorf("whois: send to whois server failed: %v", err)
	}

	_ = conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	buffer, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", fmt.Errorf("whois: read from whois server failed: %v", err)
	}

	return string(buffer), nil
}

// getExtension returns extension of domain
func getExtension(domain string) string {
	ext := domain

	if net.ParseIP(domain) == nil {
		if net.ParseIP(strings.Split(domain, "/")[0]) != nil {
			ext = strings.Split(domain, "/")[0]
		} else {
			domains := strings.Split(domain, ".")
			ext = domains[len(domains)-1]
		}
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
			return strings.TrimSpace(data[start : start+end])
		}
	}

	return ""
}

// IsASN returns if s is ASN
func IsASN(s string) bool {
	s = strings.ToUpper(s)

	if strings.HasPrefix(s, "AS") {
		s = strings.TrimPrefix(s, "AS")
	}

	_, err := strconv.Atoi(s)

	return err == nil
}
