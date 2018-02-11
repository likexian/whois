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
    "fmt"
    "net"
    "time"
    "strings"
    "io/ioutil"
)


const (
    WHOIS_DOMAIN = "whois-servers.net"
    WHOIS_PORT = "43"
)


func Version() string {
    return "0.5.0"
}


func Author() string {
    return "[Li Kexian](https://www.likexian.com/)"
}


func License() string {
    return "Apache License, Version 2.0"
}


func Whois(domain string, servers ...string) (result string, err error) {
    domain = strings.Trim(strings.Trim(domain, " "), ".")
    if domain == "" {
        err = fmt.Errorf("Domain is empty")
        return
    }

    result, err = query(domain, servers...)
    if err != nil {
        return
    }

    start := strings.Index(result, "Registrar WHOIS Server:")
    if start == -1 {
        return
    }

    start += 23
    end := strings.Index(result[start:], "\n")
    server := strings.Trim(strings.Replace(result[start:start + end], "\r", "", -1), " ")
    if server == "" {
        return
    }

    tmp_result, err := query(domain, server)
    if err != nil {
        return
    }

    result += tmp_result

    return
}


func query(domain string, servers ...string) (result string, err error) {
    var server string
    if len(servers) == 0 || servers[0] == "" {
        domains := strings.Split(domain, ".")
        if len(domains) < 2 {
            err = fmt.Errorf("Domain %s is invalid", domain)
            return
        }
        server = domains[len(domains) - 1] + "." + WHOIS_DOMAIN
    } else {
        server = servers[0]
    }

    conn, e := net.DialTimeout("tcp", net.JoinHostPort(server, WHOIS_PORT), time.Second * 30)
    if e != nil {
        err = e
        return
    }

    defer conn.Close()
    conn.Write([]byte(domain + "\r\n"))
    buffer, e := ioutil.ReadAll(conn)
    if e != nil {
        err = e
        return
    }

    result = string(buffer)

    return
}
