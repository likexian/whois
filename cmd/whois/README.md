# Whois

Whois is a release tool for domain and ip whois information query.

## Overview

All of domain, IP include IPv4 and IPv6, ASN are supported.

## Download whois

The latest version of whois can be downloaded using the links below. Please download the proper package for your operating system and architecture.

Whois is released as a single binary file. Install whois by unzipping it and moving it to a directory included in your system's PATH.

### macOS

- [64-bit](https://github.com/likexian/whois/releases/latest/download/whois-darwin-amd64.tar.gz)

### Linux

- [64-bit](https://github.com/likexian/whois/releases/latest/download/whois-linux-amd64.tar.gz)
- [32-bit](https://github.com/likexian/whois/releases/latest/download/whois-linux-386.tar.gz)

### Windows

- [64-bit](https://github.com/likexian/whois/releases/latest/download/whois-windows-amd64.zip)
- [32-bit](https://github.com/likexian/whois/releases/latest/download/whois-windows-386.zip)

## Usage

### whois query for domain

```shell
whois likexian.com
```

### whois query for IPv6

```shell
whois 2001:dc7::1
```

### whois query for IPv4

```shell
whois 1.1.1.1
```

### whois query for ASN

```shell
# or whois as60614
whois 60614
```

### whois query output as json

```shell
whois -j likexian.com
```

## License

Copyright 2014-2023 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
