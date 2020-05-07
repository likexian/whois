# Whois

Whois is a release tool for domain and ip whois information query.

## Download whois

Binary distributions available for Linux x86 and x86_64.

### linux x86_64

```shell
wget https://github.com/likexian/whois-go/releases/download/v1.6.0/whois.linux-amd64.tar.gz
```

OR

```shell
curl https://github.com/likexian/whois-go/releases/download/v1.6.0/whois.linux-amd64.tar.gz -OL
```

### linux x86

```shell
wget https://github.com/likexian/whois-go/releases/download/v1.6.0/whois.linux-386.tar.gz
```

OR

```shell
curl https://github.com/likexian/whois-go/releases/download/v1.6.0/whois.linux-386.tar.gz -OL
```

## Install whois

```shell
tar zxf whois.linux-*.tar.gz
chmod +x whois
mv /usr/bin/whois /usr/bin/whois.old
mv whois /usr/bin/whois
```

## Test whois

```shell
whois likexian.com
```

OR

```shell
whois 2001:dc7::1
```

OR

```shell
whois 1.1.1.1
```

## License

Copyright 2014-2020 [Li Kexian](https://www.likexian.com/)

Licensed under the Apache License 2.0

## Donation

If this project is helpful, please share it with friends.

If you want to thank me, you can [give me a cup of coffee](https://www.likexian.com/donate/).
