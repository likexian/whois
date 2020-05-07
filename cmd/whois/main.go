/*
 * Copyright 2014-2020 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for domain and ip whois information query
 * https://www.likexian.com/
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/likexian/whois-go"
)

func main() {
	server := flag.String("h", "", "specify the whois server")
	version := flag.Bool("v", false, "show the whois version")
	flag.Parse()

	if *version {
		fmt.Println("whois version " + whois.Version())
		fmt.Println(whois.Author())
		os.Exit(0)
	}

	if len(flag.Args()) == 0 {
		fmt.Printf("usage:\n\t%s [-h server] domain\n", os.Args[0])
		os.Exit(1)
	}

	result, err := whois.Whois(flag.Args()[0], *server)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(result)
	os.Exit(0)
}
