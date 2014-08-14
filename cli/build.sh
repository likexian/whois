#!/bin/bash

go build main.go
mv main whois

echo
echo "build success."
echo "you can use [./whois domain] to query whois now."
echo
echo "for example:"
echo "./whois likexian.com"
echo
echo "press any key to test the above example."
echo
read -n1
./whois likexian.com
