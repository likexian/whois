/*
 * Go module for domain whois
 * https://www.likexian.com/
 *
 * Copyright 2014-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
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
	return "1.0.1"
}

// Author returns package author
func Author() string {
	return "[Li Kexian](https://www.likexian.com/)"
}

// License returns package license
func License() string {
	return "Apache License, Version 2.0"
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
