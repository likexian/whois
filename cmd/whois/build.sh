#!/bin/bash

arch=386
for os in linux windows; do
    echo "Building whois-$os-$arch..."
    GOOS=$os GOARCH=$arch go build -trimpath -ldflags '-w -s' -o whois
    zip whois-$os-$arch.zip whois
    rm -rf whois
done

arch=amd64
for os in linux windows darwin; do
    echo "Building whois-$os-$arch..."
    GOOS=$os GOARCH=$arch go build -trimpath -ldflags '-w -s' -o whois
    zip whois-$os-$arch.zip whois
    rm -rf whois
done
