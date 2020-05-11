# Whois

Whois is a release tool for domain and ip whois information query.

## Download whois

The latest version of whois can be downloaded using the links below. Please download the proper package for your operating system and architecture.

Whois is released as a single binary file. Install whois by unzipping it and moving it to a directory included in your system's PATH.

### macOS

- [64-bit](https://github.com/likexian/whois-go/releases/download/v1.7.1/whois-darwin-amd64.zip)

### Linux

- [32-bit](https://github.com/likexian/whois-go/releases/download/v1.7.1/whois-linux-386.zip)
- [64-bit](https://github.com/likexian/whois-go/releases/download/v1.7.1/whois-linux-amd64.zip)

### Windows

- [32-bit](https://github.com/likexian/whois-go/releases/download/v1.7.1/whois-windows-386.zip)
- [64-bit](https://github.com/likexian/whois-go/releases/download/v1.7.1/whois-windows-amd64.zip)

## Usage

### whois query for domain

```shell
whois likexian.com
```

### whois query for ipv6

```shell
whois 2001:dc7::1
```

### whois query for ipv4

```shell
whois 1.1.1.1
```

### whois query output as json

```shell
whois -j likexian.com
```

## License

Copyright 2014-2020 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
