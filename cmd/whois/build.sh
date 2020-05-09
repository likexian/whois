#!/bin/bash

for os in linux windows; do
    for arch in 386 amd64; do
        echo "Building whois-$os-$arch..."
        GOOS=$os GOARCH=$arch go build -v -trimpath -ldflags '-w -s' -o whois
        zip whois-$os-$arch.zip whois
        rm -rf whois
    done
done
