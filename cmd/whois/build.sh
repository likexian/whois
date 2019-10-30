#!/bin/bash

for i in 386 amd64; do
    echo "Building whois.linux-$i..."
    GOOS=linux GOARCH=$i go build -v -ldflags '-w -s' -o whois
    tar zcf whois.linux-$i.tar.gz whois
    rm -rf whois
done
