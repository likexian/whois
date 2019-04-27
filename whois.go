/*
 * Copyright 2014-2019 Li Kexian
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
 * Go module for domain and ip whois info query
 * https://www.likexian.com/
 */

package whois

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

// Query server const
const (
	IP_WHOIS_SERVER     = "whois.iana.org"
	DOMAIN_WHOIS_SERVER = "whois-servers.net"
	WHOIS_PORT          = "43"
)

// Version returns package version
func Version() string {
	return "1.0.3"
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
		err = fmt.Errorf("Domain is empty")
		return
	}

	result, err = query(domain, servers...)
	if err != nil {
		return
	}

	token := "Registrar WHOIS Server:"
	if IsIpv4(domain) {
		token = "whois:"
	}

	start := strings.Index(result, token)
	if start == -1 {
		return
	}

	start += len(token)
	end := strings.Index(result[start:], "\n")
	server := strings.TrimSpace(result[start : start+end])
	if server == "" {
		return
	}

	tmpResult, err := query(domain, server)
	if err != nil {
		return
	}

	result += tmpResult

	return
}

// query do the query
func query(domain string, servers ...string) (result string, err error) {
	var server string
	if len(servers) == 0 || servers[0] == "" {
		if IsIpv4(domain) {
			server = IP_WHOIS_SERVER
		} else {
			domains := strings.Split(domain, ".")
			if len(domains) < 2 {
				err = fmt.Errorf("Domain %s is invalid", domain)
				return
			}
			server = domains[len(domains)-1] + "." + DOMAIN_WHOIS_SERVER
		}
	} else {
		server = strings.ToLower(servers[0])
		if server == "whois.arin.net" {
			domain = "n + " + domain
		}
	}

	conn, e := net.DialTimeout("tcp", net.JoinHostPort(server, WHOIS_PORT), time.Second*30)
	if e != nil {
		err = e
		return
	}

	defer conn.Close()
	conn.Write([]byte(domain + "\r\n"))
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))

	buffer, e := ioutil.ReadAll(conn)
	if e != nil {
		err = e
		return
	}

	result = string(buffer)

	return
}

// IsIpv4 returns string is an ipv4 ip
func IsIpv4(ip string) bool {
	i := net.ParseIP(ip)
	return i.To4() != nil
}
