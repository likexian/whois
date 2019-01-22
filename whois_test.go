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
    "testing"
)


func TestWhois(t *testing.T) {
    _, err := Whois("likexian")
    if err == nil {
        t.Error("Not a domain shall got error")
    }

    _, err = Whois("8.8.8.888")
    if err == nil {
        t.Error("Not an ip shall got error")
    }

    test_domains := []string{
        "likexian.com",
        "likexian.net",
        "likexian.org",
        "likexian.cn",
        "likexian.com.cn",
        "1.1.1.1",
        "2.1.1.1",
        "3.1.1.1",
        "4.1.1.1",
        "5.1.1.1",
    }

    for _, v := range test_domains {
        _, err = Whois(v)
        if err != nil {
            t.Errorf("Domain %s shall got result but got an error: %s", v, err.Error())
        }
    }

    _, err = Whois("likexian.com", "127.0.0.1")
    if err == nil {
        t.Error("Invalid server shall got error")
    }

    _, err = Whois("likexian.com", "com.whois-servers.net")
    if err != nil {
        t.Errorf("Domain shall got result but got and error: %s", err.Error())
    }
}
