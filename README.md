# whois.go

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/whois-go?status.svg)](https://godoc.org/github.com/likexian/whois-go)
[![Build Status](https://travis-ci.org/likexian/whois-go.svg?branch=master)](https://travis-ci.org/likexian/whois-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/whois-go)](https://goreportcard.com/report/github.com/likexian/whois-go)
[![Code Cover](https://codecov.io/gh/likexian/whois-go/graph/badge.svg)](https://codecov.io/gh/likexian/whois-go)

whois-go is a simple Go module for domain and ip whois info query.

## Overview

You can directly using the binary distributions whois, follow [whois release tool](whois).

Or you can do development by using the golang module as below.

*Works for most domain extensions and most ip most of the time.*

## Installation

    go get -u github.com/likexian/whois-go

## Importing

    import (
        "github.com/likexian/whois-go"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/whois-go)

## Example

### whois query for domain

```go
result, err := whois.Whois("example.com")
if err == nil {
    fmt.Println(result)
}
```

### whois query for ip

```go
result, err := whois.Whois("1.1.1.1")
if err == nil {
    fmt.Println(result)
}
```

## Whois info parser in Go

Please refer to [whois-parser-go](https://github.com/likexian/whois-parser-go)

## LICENSE

Copyright 2014-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
