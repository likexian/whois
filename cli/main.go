/*
 * Go module for domain whois
 * http://www.likexian.com/
 *
 * Copyright 2014, Kexian Li
 * Released under the Apache License, Version 2.0
 *
 */

package main


import (
    "fmt"
    "os"
    ".."
)


func main() {
    if len(os.Args) < 2 {
        fmt.Println(fmt.Sprintf("usage:\n\t%s domain[ server]", os.Args[0]))
        os.Exit(1)
    }

    var server string
    if len(os.Args) > 2 {
        server = os.Args[2]
    }

    result, err := whois.Whois(os.Args[1], server)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println(result)
    os.Exit(0)
}
