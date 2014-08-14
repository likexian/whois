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
    return "0.1.0"
}


func Author() string {
    return "[Li Kexian](http://www.likexian.com/)"
}


func License() string {
    return "Apache License, Version 2.0"
}


func Whois(domain string, servers ...string) (result string, err error) {
    var server string
    if len(servers) == 0 {
        domains := strings.SplitN(domain, ".", 2)
        if len(domains) != 2 {
            err = fmt.Errorf("Domain %s is invalid.", domain)
            return
        }
        server = domains[1] + "." + WHOIS_DOMAIN
    } else {
        server = servers[0]
    }

    conn, err := net.DialTimeout("tcp", net.JoinHostPort(server, WHOIS_PORT), time.Second * 30)
    if err != nil {
        return
    }

    conn.Write([]byte(domain + "\r\n"))
    var buffer []byte
    buffer, err = ioutil.ReadAll(conn)
    if err != nil {
        return
    }

    conn.Close()
    result = string(buffer[:])

    return 
}
