# whois.go

whois-go is a simple Go module for domain whois.

[![Build Status](https://secure.travis-ci.org/likexian/whois-go.png)](https://secure.travis-ci.org/likexian/whois-go)

## Overview

A golang module for domain whois query.

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

    result, err := Whois("example.com")
    if err != nil {
        fmt.Println(result)
    }

## LICENSE

Copyright 2014, Kexian Li

Apache License, Version 2.0

## About

- [Kexian Li](http://github.com/likexian)
- [http://www.likexian.com/](http://www.likexian.com/)
