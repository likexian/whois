# whois.go

whois-go is a simple Go module for domain whois.

[![Build Status](https://secure.travis-ci.org/likexian/whois-go.png)](https://secure.travis-ci.org/likexian/whois-go)

## Overview

whois.go: A golang module for domain whois query.

cli/main.go: A golang cli command for domain whois query.

cli/build.sh: Build the cli command and test it.

Work for most domain extensions and most of the time.

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
    if err != nil {
        fmt.Println(result)
    }

## LICENSE

Copyright 2014, Kexian Li

Apache License, Version 2.0

## Whois info parser in Go

Please refer to [whois-parser-go](https://github.com/likexian/whois-parser-go)

## About

- [Kexian Li](http://github.com/likexian)
- [http://www.likexian.com/](http://www.likexian.com/)
