#!/bin/bash

go build -o whois main.go

if [ ! -x ./whois ]; then
    echo "build fail."
    exit 1
fi

echo
echo "build success."
echo "you can use [./whois domain] to query whois now."
echo
echo "for example: ./whois likexian.com"
echo
echo "press any key to test the above example."
echo
read -n1
./whois likexian.com
