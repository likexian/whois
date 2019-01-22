# whois.go

whois-go is a simple Go module for domain and ip whois info query.

[![Build Status](https://secure.travis-ci.org/likexian/whois-go.png)](https://secure.travis-ci.org/likexian/whois-go)

## Overview

whois.go: A golang module for domain and ip whois query.

whois: A golang cli command for domain and ip whois query.

*Works for most domain extensions and most ip most of the time.*

## Installation

    go get github.com/likexian/whois-go

## Importing

    import (
        "github.com/likexian/whois-go"
    )

## Documentation

    func Whois(domain string, servers ...string) (result string, err error)

## Example

    result, err := whois.Whois("example.com")
    if err == nil {
        fmt.Println(result)
    }

    result, err := whois.Whois("1.1.1.1")
    if err == nil {
        fmt.Println(result)
    }

## Whois info parser in Go

Please refer to [whois-parser-go](https://github.com/likexian/whois-parser-go)

## LICENSE

Copyright 2014-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
