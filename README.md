# Whois

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/whois-go?status.svg)](https://godoc.org/github.com/likexian/whois-go)
[![Build Status](https://travis-ci.org/likexian/whois-go.svg?branch=master)](https://travis-ci.org/likexian/whois-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/whois-go)](https://goreportcard.com/report/github.com/likexian/whois-go)
[![Code Cover](https://codecov.io/gh/likexian/whois-go/graph/badge.svg)](https://codecov.io/gh/likexian/whois-go)

Whois is a simple Go module for domain and ip whois information query.

## Overview

All of domain, ipv4 and ipv6 are support.

You can directly using the binary distributions whois, follow [whois release tool](cmd/whois).

Or you can do development by using this golang module as below.

## Installation

```shell
go get -u github.com/likexian/whois-go
```

## Importing

```go
import (
    "github.com/likexian/whois-go"
)
```

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/whois-go)

## Example

### whois query for domain

```go
result, err := whois.Whois("likexian.com")
if err == nil {
    fmt.Println(result)
}
```

### whois query for ipv6

```go
result, err := whois.Whois("2001:dc7::1")
if err == nil {
    fmt.Println(result)
}
```

### whois query for ipv4

```go
result, err := whois.Whois("1.1.1.1")
if err == nil {
    fmt.Println(result)
}
```

## Whois information parsing

Please refer to [whois-parser-go](https://github.com/likexian/whois-parser-go)

## License

Copyright 2014-2020 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
